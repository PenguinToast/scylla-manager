// Copyright (C) 2022 ScyllaDB

package backup

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/scylladb/gocqlx/v2/qb"
	. "github.com/scylladb/scylla-manager/v3/pkg/service/backup/backupspec"
	"github.com/scylladb/scylla-manager/v3/pkg/util/uuid"
)

// GetRestoreTarget converts runner properties into RestoreTarget.
func (s *Service) GetRestoreTarget(ctx context.Context, clusterID uuid.UUID, properties json.RawMessage) (RestoreTarget, error) {
	s.logger.Info(ctx, "GetRestoreTarget", "cluster_id", clusterID)

	t := defaultRestoreTarget()

	if err := json.Unmarshal(properties, &t); err != nil {
		return t, err
	}

	if err := t.validateProperties(); err != nil {
		return t, err
	}

	locations := make(map[string]struct{})
	for _, l := range t.Location {
		rp := l.RemotePath("")
		if _, ok := locations[rp]; ok {
			return t, errors.Errorf("location '%s' is specified multiple times", l)
		}
		locations[rp] = struct{}{}

		if l.DC == "" {
			s.logger.Info(ctx, "No datacenter specified for location - using all nodes for this location", "location", l)
		}
	}
	t.sortLocations()

	if t.RestoreSchema {
		t.Keyspace = []string{"system_schema"}
	}
	if t.RestoreTables {
		if t.Keyspace == nil {
			t.Keyspace = []string{"*"}
		}
		// Skip restoration of those tables regardless of the '--keyspace' param
		doNotRestore := []string{
			"system",        // system.* tables are recreated on every cluster and shouldn't even be backed-up
			"system_schema", // Schema restoration is only possible with '--restore-schema' flag
			// Don't restore tables related to CDC.
			// Currently, it is forbidden to alter those tables, so SM wouldn't be able to ensure their data consistency.
			// Moreover, those tables usually contain data with small TTL value,
			// so their contents would probably expire right after restore has ended.
			"system_distributed_everywhere.cdc_generation_descriptions_v2",
			"system_distributed.cdc_streams_descriptions_v2",
			"system_distributed.cdc_generation_timestamps",
			"*.*_scylla_cdc_log", // All regular CDC tables have "_scylla_cdc_log" suffix
		}

		for _, ks := range doNotRestore {
			t.Keyspace = append(t.Keyspace, "!"+ks)
		}

		// Filter out all materialized views and secondary indexes. They are not a part of restore procedure at the moment.
		// See https://docs.scylladb.com/stable/operating-scylla/procedures/backup-restore/restore.html#repeat-the-following-steps-for-each-node-in-the-cluster.
		views, err := s.listAllViews(ctx, clusterID)
		if err != nil {
			return t, errors.Wrapf(err, "list all views of cluster %s", clusterID.String())
		}
		for _, viewName := range views {
			t.Keyspace = append(t.Keyspace, "!"+viewName)
		}
	}

	client, err := s.scyllaClient(ctx, clusterID)
	if err != nil {
		return t, errors.Wrapf(err, "get client")
	}

	status, err := client.Status(ctx)
	if err != nil {
		return t, errors.Wrap(err, "get status")
	}
	if len(status) == 0 {
		return t, errors.New("empty status")
	}
	// Check if for each location there is at least one host
	// living in location's dc with access to it.
	for _, l := range t.Location {
		var (
			remotePath     = l.RemotePath("")
			locationStatus = status
		)
		// In case location does not have specified dc, use nodes from all dcs.
		if l.DC != "" {
			locationStatus = status.Datacenter([]string{l.DC})
			if len(locationStatus) == 0 {
				return t, errors.Errorf("no nodes in location's datacenter: %s", l)
			}
		}

		if _, err = client.GetLiveNodesWithLocationAccess(ctx, locationStatus, remotePath); err != nil {
			if strings.Contains(err.Error(), "NoSuchBucket") {
				return t, errors.Errorf("specified bucket does not exist: %s", l)
			}
			return t, errors.Wrap(err, "location is not accessible")
		}
	}

	return t, nil
}

// Restore executes restore on a given target.
func (s *Service) Restore(ctx context.Context, clusterID, taskID, runID uuid.UUID, target RestoreTarget) error {
	s.logger.Info(ctx, "Restore",
		"cluster_id", clusterID,
		"task_id", taskID,
		"run_id", runID,
		"target", target,
	)

	run := &RestoreRun{
		ClusterID:   clusterID,
		TaskID:      taskID,
		ID:          runID,
		SnapshotTag: target.SnapshotTag,
		Stage:       StageRestoreInit,
	}

	// Get cluster name
	clusterName, err := s.clusterName(ctx, clusterID)
	if err != nil {
		return errors.Wrap(err, "invalid cluster")
	}
	// Get the cluster client
	client, err := s.scyllaClient(ctx, clusterID)
	if err != nil {
		return errors.Wrap(err, "get client proxy")
	}
	// Get cluster session
	clusterSession, err := s.clusterSession(ctx, clusterID)
	if err != nil {
		return errors.Wrap(err, "get CQL cluster session")
	}
	defer clusterSession.Close()

	tools := restoreWorkerTools{
		workerTools: workerTools{
			ClusterID:   clusterID,
			ClusterName: clusterName,
			TaskID:      taskID,
			RunID:       runID,
			SnapshotTag: target.SnapshotTag,
			Client:      client,
			Config:      s.config,
			Logger:      s.logger.Named("restore"),
		},
		metrics:                 s.metrics.Restore,
		managerSession:          s.session,
		clusterSession:          clusterSession,
		forEachRestoredManifest: s.forEachRestoredManifest(clusterID, target),
	}

	var w restoreWorker
	if target.RestoreTables {
		w = &tablesWorker{restoreWorkerTools: tools}
	} else {
		w = &schemaWorker{restoreWorkerTools: tools}
	}

	if target.Continue {
		if err := w.decorateWithPrevRun(ctx, run); err != nil {
			return err
		}

		w.insertRun(ctx, run)
		// Update run with previous progress.
		if run.PrevID != uuid.Nil {
			w.clonePrevProgress(ctx, run)
		}
	} else {
		w.insertRun(ctx, run)
	}
	// Check if restore is the continuation of previous run.
	if target.Continue && run.PrevID != uuid.Nil {
		w.continuePrevRun()
	} else {
		s.metrics.Restore.ResetClusterMetrics(clusterID)
	}
	// As manifests are immutable, units can be initialized only once per task
	if run.Units == nil {
		run.Units, err = w.newUnits(ctx, target)
		if err != nil {
			return errors.Wrap(err, "initialize units")
		}
	}

	run.Stage = StageRestoreData
	w.insertRun(ctx, run)

	if err = w.restore(ctx, run, target); err != nil {
		return errors.Wrapf(err, "restore data")
	}

	run.Stage = StageRestoreDone
	w.insertRun(ctx, run)

	return nil
}

// GetRestoreUnits restored units specified in restore target.
func (s *Service) GetRestoreUnits(ctx context.Context, clusterID uuid.UUID, target RestoreTarget) ([]RestoreUnit, error) {
	clusterSession, err := s.clusterSession(ctx, clusterID)
	if err != nil {
		return nil, errors.Wrap(err, "get CQL cluster session")
	}
	defer clusterSession.Close()

	w := &restoreWorkerTools{
		clusterSession:          clusterSession,
		forEachRestoredManifest: s.forEachRestoredManifest(clusterID, target),
	}

	return w.newUnits(ctx, target)
}

// GetRestoreProgress aggregates progress for the run of the task and breaks it down
// by keyspace and table.json.
func (s *Service) GetRestoreProgress(ctx context.Context, clusterID, taskID, runID uuid.UUID) (RestoreProgress, error) {
	w := &restoreWorkerTools{
		workerTools: workerTools{
			ClusterID: clusterID,
			TaskID:    taskID,
			RunID:     runID,
		},
		managerSession: s.session,
	}

	return w.getProgress(ctx)
}

// forEachRestoredManifest returns a wrapper for forEachManifest that iterates over
// manifests with specified in restore target.
func (s *Service) forEachRestoredManifest(clusterID uuid.UUID, target RestoreTarget) func(context.Context, Location, func(ManifestInfoWithContent) error) error {
	return func(ctx context.Context, location Location, f func(content ManifestInfoWithContent) error) error {
		filter := ListFilter{
			SnapshotTag: target.SnapshotTag,
			Keyspace:    target.Keyspace,
		}
		return s.forEachManifest(ctx, clusterID, []Location{location}, filter, f)
	}
}

// listAllViews is the utility function that queries system_schema.views table to get list of all views created on the cluster.
// system_schema.views contains view definitions for materialized views and for secondary indexes.
func (s *Service) listAllViews(ctx context.Context, clusterID uuid.UUID) (views []string, err error) {
	// Get cluster session
	clusterSession, err := s.clusterSession(ctx, clusterID)
	if err != nil {
		return nil, errors.Wrap(err, "get CQL cluster session")
	}

	iter := qb.Select("system_schema.views").
		Columns("keyspace_name", "view_name").
		Query(clusterSession).Iter()
	defer func() {
		err = iter.Close()
	}()

	var keyspace, view string
	for iter.Scan(&keyspace, &view) {
		views = append(views, keyspace+"."+view)
	}

	return
}
