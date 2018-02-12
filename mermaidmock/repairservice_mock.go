// Code generated by MockGen. DO NOT EDIT.
// Source: repair.go

// Package mermaidmock is a generated GoMock package.
package mermaidmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	repair "github.com/scylladb/mermaid/repair"
	uuid "github.com/scylladb/mermaid/uuid"
	reflect "reflect"
)

// MockRepairService is a mock of RepairService interface
type MockRepairService struct {
	ctrl     *gomock.Controller
	recorder *MockRepairServiceMockRecorder
}

// MockRepairServiceMockRecorder is the mock recorder for MockRepairService
type MockRepairServiceMockRecorder struct {
	mock *MockRepairService
}

// NewMockRepairService creates a new mock instance
func NewMockRepairService(ctrl *gomock.Controller) *MockRepairService {
	mock := &MockRepairService{ctrl: ctrl}
	mock.recorder = &MockRepairServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepairService) EXPECT() *MockRepairServiceMockRecorder {
	return m.recorder
}

// GetUnit mocks base method
func (m *MockRepairService) GetUnit(ctx context.Context, clusterID uuid.UUID, idOrName string) (*repair.Unit, error) {
	ret := m.ctrl.Call(m, "GetUnit", ctx, clusterID, idOrName)
	ret0, _ := ret[0].(*repair.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnit indicates an expected call of GetUnit
func (mr *MockRepairServiceMockRecorder) GetUnit(ctx, clusterID, idOrName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnit", reflect.TypeOf((*MockRepairService)(nil).GetUnit), ctx, clusterID, idOrName)
}

// PutUnit mocks base method
func (m *MockRepairService) PutUnit(ctx context.Context, u *repair.Unit) error {
	ret := m.ctrl.Call(m, "PutUnit", ctx, u)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutUnit indicates an expected call of PutUnit
func (mr *MockRepairServiceMockRecorder) PutUnit(ctx, u interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUnit", reflect.TypeOf((*MockRepairService)(nil).PutUnit), ctx, u)
}

// DeleteUnit mocks base method
func (m *MockRepairService) DeleteUnit(ctx context.Context, clusterID, ID uuid.UUID) error {
	ret := m.ctrl.Call(m, "DeleteUnit", ctx, clusterID, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnit indicates an expected call of DeleteUnit
func (mr *MockRepairServiceMockRecorder) DeleteUnit(ctx, clusterID, ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnit", reflect.TypeOf((*MockRepairService)(nil).DeleteUnit), ctx, clusterID, ID)
}

// ListUnits mocks base method
func (m *MockRepairService) ListUnits(ctx context.Context, clusterID uuid.UUID, f *repair.UnitFilter) ([]*repair.Unit, error) {
	ret := m.ctrl.Call(m, "ListUnits", ctx, clusterID, f)
	ret0, _ := ret[0].([]*repair.Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnits indicates an expected call of ListUnits
func (mr *MockRepairServiceMockRecorder) ListUnits(ctx, clusterID, f interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnits", reflect.TypeOf((*MockRepairService)(nil).ListUnits), ctx, clusterID, f)
}

// GetConfig mocks base method
func (m *MockRepairService) GetConfig(ctx context.Context, src repair.ConfigSource) (*repair.LegacyConfig, error) {
	ret := m.ctrl.Call(m, "GetConfig", ctx, src)
	ret0, _ := ret[0].(*repair.LegacyConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockRepairServiceMockRecorder) GetConfig(ctx, src interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockRepairService)(nil).GetConfig), ctx, src)
}

// PutConfig mocks base method
func (m *MockRepairService) PutConfig(ctx context.Context, src repair.ConfigSource, c *repair.LegacyConfig) error {
	ret := m.ctrl.Call(m, "PutConfig", ctx, src, c)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutConfig indicates an expected call of PutConfig
func (mr *MockRepairServiceMockRecorder) PutConfig(ctx, src, c interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutConfig", reflect.TypeOf((*MockRepairService)(nil).PutConfig), ctx, src, c)
}

// DeleteConfig mocks base method
func (m *MockRepairService) DeleteConfig(ctx context.Context, src repair.ConfigSource) error {
	ret := m.ctrl.Call(m, "DeleteConfig", ctx, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfig indicates an expected call of DeleteConfig
func (mr *MockRepairServiceMockRecorder) DeleteConfig(ctx, src interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfig", reflect.TypeOf((*MockRepairService)(nil).DeleteConfig), ctx, src)
}

// GetRun mocks base method
func (m *MockRepairService) GetRun(ctx context.Context, u *repair.Unit, runID uuid.UUID) (*repair.Run, error) {
	ret := m.ctrl.Call(m, "GetRun", ctx, u, runID)
	ret0, _ := ret[0].(*repair.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun
func (mr *MockRepairServiceMockRecorder) GetRun(ctx, u, runID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockRepairService)(nil).GetRun), ctx, u, runID)
}

// GetLastRun mocks base method
func (m *MockRepairService) GetLastRun(ctx context.Context, u *repair.Unit) (*repair.Run, error) {
	ret := m.ctrl.Call(m, "GetLastRun", ctx, u)
	ret0, _ := ret[0].(*repair.Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastRun indicates an expected call of GetLastRun
func (mr *MockRepairServiceMockRecorder) GetLastRun(ctx, u interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastRun", reflect.TypeOf((*MockRepairService)(nil).GetLastRun), ctx, u)
}

// GetProgress mocks base method
func (m *MockRepairService) GetProgress(ctx context.Context, u *repair.Unit, runID uuid.UUID, hosts ...string) ([]*repair.RunProgress, error) {
	varargs := []interface{}{ctx, u, runID}
	for _, a := range hosts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProgress", varargs...)
	ret0, _ := ret[0].([]*repair.RunProgress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProgress indicates an expected call of GetProgress
func (mr *MockRepairServiceMockRecorder) GetProgress(ctx, u, runID interface{}, hosts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, u, runID}, hosts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProgress", reflect.TypeOf((*MockRepairService)(nil).GetProgress), varargs...)
}
