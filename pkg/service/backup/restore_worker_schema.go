// Copyright (C) 2022 ScyllaDB

package backup

import (
	"context"
	"path"
	"strconv"

	"github.com/pkg/errors"
	"go.uber.org/atomic"

	. "github.com/scylladb/scylla-manager/v3/pkg/service/backup/backupspec"
	"github.com/scylladb/scylla-manager/v3/pkg/util/parallel"
	"github.com/scylladb/scylla-manager/v3/pkg/util/timeutc"
)

type schemaWorker struct {
	restoreWorkerTools

	hosts           []string
	generationCnt   atomic.Int64
	renamedSSTables map[string]string
	miwc            ManifestInfoWithContent // Currently restored manifest
	// Maps original SSTable name to its existing older version (with respect to currently restored snapshot tag)
	// that should be used during the restore procedure. It should be initialized per each restored table.
	versionedFiles VersionedMap
}

// restore downloads all backed-up schema files to each node in the cluster. This approach is necessary because
// it's not possible to alter gc_grace_seconds or tombstone_gc on schema tables (safety requirement for nodetool refresh).
// It introduces great data duplication, but is necessary in order to simulate schema repair on each node.
// Luckily, schema files are small, so this shouldn't be noticeable in terms of performance.
// When all files are downloaded, they are restored using nodetool refresh.
// Note that due to small schema size:
// - resuming schema restoration will always start from scratch
// - schema restoration does not use long polling for updating download progress
// Adding the ability to resume schema restoration might be added in the future.
func (w *schemaWorker) restore(ctx context.Context, run *RestoreRun, target RestoreTarget) error {
	w.AwaitSchemaAgreement(ctx, w.clusterSession)

	w.Logger.Info(ctx, "Started restoring schema")
	defer w.Logger.Info(ctx, "Restoring schema finished")

	status, err := w.Client.Status(ctx)
	if err != nil {
		return errors.Wrap(err, "get status")
	}
	if len(status) != len(status.Live()) {
		return errors.New("not all nodes are in the UN state")
	}
	// Clean upload dirs.
	// This is required as we rename SSTables during download in order to avoid name overlaps.
	for _, u := range run.Units {
		for _, t := range u.Tables {
			version, err := w.GetTableVersion(ctx, u.Keyspace, t.Table)
			if err != nil {
				return err
			}
			uploadDir := UploadTableDir(u.Keyspace, t.Table, version)

			for _, h := range status {
				if err := w.cleanUploadDir(ctx, h.Addr, uploadDir, nil); err != nil {
					return err
				}
			}
		}
	}
	// Download files
	for _, l := range target.Location {
		if err = w.locationDownloadHandler(ctx, run, target, l); err != nil {
			return err
		}
	}
	// Set restore start in all run progresses
	w.ForEachProgress(ctx, run, func(pr *RestoreRunProgress) {
		pr.setRestoreStartedAt()
		w.insertRunProgress(ctx, pr)
	})
	// Load schema SSTables on all nodes
	f := func(i int) error {
		host := status[i]
		for _, ks := range run.Units {
			for _, t := range ks.Tables {
				if _, err := w.Client.LoadSSTables(ctx, host.Addr, ks.Keyspace, t.Table, false, false); err != nil {
					return errors.Wrap(err, "restore schema")
				}
			}
		}
		return nil
	}

	notify := func(i int, err error) {
		w.Logger.Error(ctx, "Failed to load schema on host",
			"host", status[i],
			"error", err,
		)
	}

	if err = parallel.Run(len(status), parallel.NoLimit, f, notify); err != nil {
		return err
	}
	// Set restore completed in all run progresses
	w.ForEachProgress(ctx, run, func(pr *RestoreRunProgress) {
		pr.setRestoreCompletedAt()
		w.insertRunProgress(ctx, pr)
	})

	return nil
}

func (w *schemaWorker) locationDownloadHandler(ctx context.Context, run *RestoreRun, target RestoreTarget, location Location) error {
	w.Logger.Info(ctx, "Downloading schema from location", "location", location)
	defer w.Logger.Info(ctx, "Downloading schema from location finished", "location", location)

	w.location = location
	run.Location = location.String()

	if err := w.initHosts(ctx); err != nil {
		return errors.Wrap(err, "initialize hosts")
	}

	tableDownloadHandler := func(fm FilesMeta) error {
		w.Logger.Info(ctx, "Downloading schema table", "keyspace", fm.Keyspace, "table", fm.Table)
		defer w.Logger.Info(ctx, "Downloading schema table finished", "keyspace", fm.Keyspace, "table", fm.Table)

		run.Table = fm.Table
		run.Keyspace = fm.Keyspace

		return w.workFunc(ctx, run, target, fm)
	}

	manifestDownloadHandler := func(miwc ManifestInfoWithContent) error {
		w.Logger.Info(ctx, "Downloading schema from manifest", "manifest", miwc.ManifestInfo)
		defer w.Logger.Info(ctx, "Downloading schema from manifest", "manifest", miwc.ManifestInfo)

		w.miwc = miwc
		run.ManifestPath = miwc.Path()
		w.insertRun(ctx, run)

		return miwc.ForEachIndexIterWithError(target.Keyspace, tableDownloadHandler)
	}

	return w.forEachRestoredManifest(ctx, location, manifestDownloadHandler)
}

func (w *schemaWorker) workFunc(ctx context.Context, run *RestoreRun, target RestoreTarget, fm FilesMeta) error {
	version, err := w.GetTableVersion(ctx, fm.Keyspace, fm.Table)
	if err != nil {
		return err
	}

	var (
		srcDir = w.location.RemotePath(w.miwc.SSTableVersionDir(fm.Keyspace, fm.Table, fm.Version))
		dstDir = UploadTableDir(fm.Keyspace, fm.Table, version)
	)

	w.Logger.Info(ctx, "Start downloading schema files",
		"keyspace", fm.Keyspace,
		"table", fm.Table,
		"src_dir", srcDir,
		"dst_dir", dstDir,
		"files", fm.Files,
	)

	w.initRenamedID(fm.Files)
	w.versionedFiles, err = ListVersionedFiles(ctx, w.Client, w.SnapshotTag, w.hosts[0], srcDir, w.Logger)
	if err != nil {
		return errors.Wrap(err, "initialize versioned SSTables")
	}

	f := func(i int) error {
		host := w.hosts[i]

		if err := w.checkAvailableDiskSpace(ctx, hostInfo{IP: host}); err != nil {
			return errors.Wrapf(err, "validate free disk space on host: %s", host)
		}

		start := timeutc.Now()

		fHost := func(j int) error {
			file := fm.Files[j]
			// Rename SSTable in the destination in order to avoid name conflicts
			dstFile := w.renamedSSTables[file]
			// Take the correct version of restored file
			srcFile := file
			if v, ok := w.versionedFiles[file]; ok {
				srcFile = v.FullName()
			}

			srcPath := path.Join(srcDir, srcFile)
			dstPath := path.Join(dstDir, dstFile)

			return w.Client.RcloneCopyFile(ctx, host, dstPath, srcPath)
		}

		notifyHost := func(j int, err error) {
			w.Logger.Error(ctx, "Failed to download schema SSTable",
				"host", host,
				"file", fm.Files[j],
				"error", err,
			)
		}

		// Rely on rclone ability to limit number of concurrent transfers
		err := parallel.Run(len(fm.Files), parallel.NoLimit, fHost, notifyHost)
		if err != nil {
			return errors.Wrapf(err, "download renamed SSTables on host: %s", host)
		}
		end := timeutc.Now()
		// In order to ensure that the size calculated in newUnits matches the sum of restored bytes from
		// run progresses, insert only fraction of the whole downloaded size. This is caused by the data duplication.
		proportionalSize := int64((int(fm.Size) + i) / len(w.hosts))

		w.insertRunProgress(ctx, &RestoreRunProgress{
			ClusterID:           run.ClusterID,
			TaskID:              run.TaskID,
			RunID:               run.ID,
			ManifestPath:        run.ManifestPath,
			Keyspace:            run.Keyspace,
			Table:               run.Table,
			Host:                host,
			DownloadStartedAt:   &start,
			DownloadCompletedAt: &end,
			Downloaded:          proportionalSize,
		})

		return nil
	}

	notify := func(i int, err error) {
		w.Logger.Error(ctx, "Failed to restore schema on host",
			"host", w.hosts[i],
			"error", err,
		)
	}

	return parallel.Run(len(w.hosts), target.Parallel, f, notify)
}

func (w *schemaWorker) initHosts(ctx context.Context) error {
	status, err := w.Client.Status(ctx)
	if err != nil {
		return errors.Wrap(err, "get client status")
	}

	remotePath := w.location.RemotePath("")
	checkedNodes, err := w.Client.GetLiveNodesWithLocationAccess(ctx, status, remotePath)
	if err != nil {
		return errors.Wrap(err, "no live nodes with location access")
	}

	w.hosts = make([]string, 0)
	for _, host := range checkedNodes {
		w.hosts = append(w.hosts, host.Addr)
	}

	w.Logger.Info(ctx, "Initialized restore hosts", "hosts", w.hosts)
	return nil
}

func (w *schemaWorker) initRenamedID(sstables []string) {
	bundles := make(map[string][]string)
	for _, sst := range sstables {
		id := sstableID(sst)
		bundles[id] = append(bundles[id], sst)
	}

	w.renamedSSTables = make(map[string]string)
	for _, b := range bundles {
		newID := int(w.generationCnt.Add(1))
		for _, sst := range b {
			w.renamedSSTables[sst] = renameSSTableID(sst, strconv.Itoa(newID))
		}
	}
}

func (w *schemaWorker) continuePrevRun() {}
