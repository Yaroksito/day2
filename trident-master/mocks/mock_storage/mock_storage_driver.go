// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/storage (interfaces: Driver)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	roaring "github.com/RoaringBitmap/roaring"
	gomock "github.com/golang/mock/gomock"
	config "github.com/netapp/trident/config"
	storage "github.com/netapp/trident/storage"
	storageattribute "github.com/netapp/trident/storage_attribute"
	storagedrivers "github.com/netapp/trident/storage_drivers"
	utils "github.com/netapp/trident/utils"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// BackendName mocks base method.
func (m *MockDriver) BackendName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackendName")
	ret0, _ := ret[0].(string)
	return ret0
}

// BackendName indicates an expected call of BackendName.
func (mr *MockDriverMockRecorder) BackendName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackendName", reflect.TypeOf((*MockDriver)(nil).BackendName))
}

// CanSnapshot mocks base method.
func (m *MockDriver) CanSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanSnapshot indicates an expected call of CanSnapshot.
func (mr *MockDriverMockRecorder) CanSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSnapshot", reflect.TypeOf((*MockDriver)(nil).CanSnapshot), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockDriver) Create(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 storage.Pool, arg3 map[string]storageattribute.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDriverMockRecorder) Create(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDriver)(nil).Create), arg0, arg1, arg2, arg3)
}

// CreateClone mocks base method.
func (m *MockDriver) CreateClone(arg0 context.Context, arg1, arg2 *storage.VolumeConfig, arg3 storage.Pool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClone", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClone indicates an expected call of CreateClone.
func (mr *MockDriverMockRecorder) CreateClone(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClone", reflect.TypeOf((*MockDriver)(nil).CreateClone), arg0, arg1, arg2, arg3)
}

// CreateFollowup mocks base method.
func (m *MockDriver) CreateFollowup(arg0 context.Context, arg1 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFollowup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFollowup indicates an expected call of CreateFollowup.
func (mr *MockDriverMockRecorder) CreateFollowup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFollowup", reflect.TypeOf((*MockDriver)(nil).CreateFollowup), arg0, arg1)
}

// CreatePrepare mocks base method.
func (m *MockDriver) CreatePrepare(arg0 context.Context, arg1 *storage.VolumeConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreatePrepare", arg0, arg1)
}

// CreatePrepare indicates an expected call of CreatePrepare.
func (mr *MockDriverMockRecorder) CreatePrepare(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrepare", reflect.TypeOf((*MockDriver)(nil).CreatePrepare), arg0, arg1)
}

// CreateSnapshot mocks base method.
func (m *MockDriver) CreateSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) (*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockDriverMockRecorder) CreateSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockDriver)(nil).CreateSnapshot), arg0, arg1, arg2)
}

// DeleteSnapshot mocks base method.
func (m *MockDriver) DeleteSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockDriverMockRecorder) DeleteSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockDriver)(nil).DeleteSnapshot), arg0, arg1, arg2)
}

// Destroy mocks base method.
func (m *MockDriver) Destroy(arg0 context.Context, arg1 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockDriverMockRecorder) Destroy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockDriver)(nil).Destroy), arg0, arg1)
}

// Get mocks base method.
func (m *MockDriver) Get(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockDriverMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDriver)(nil).Get), arg0, arg1)
}

// GetCommonConfig mocks base method.
func (m *MockDriver) GetCommonConfig(arg0 context.Context) *storagedrivers.CommonStorageDriverConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommonConfig", arg0)
	ret0, _ := ret[0].(*storagedrivers.CommonStorageDriverConfig)
	return ret0
}

// GetCommonConfig indicates an expected call of GetCommonConfig.
func (mr *MockDriverMockRecorder) GetCommonConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommonConfig", reflect.TypeOf((*MockDriver)(nil).GetCommonConfig), arg0)
}

// GetExternalConfig mocks base method.
func (m *MockDriver) GetExternalConfig(arg0 context.Context) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalConfig", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetExternalConfig indicates an expected call of GetExternalConfig.
func (mr *MockDriverMockRecorder) GetExternalConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalConfig", reflect.TypeOf((*MockDriver)(nil).GetExternalConfig), arg0)
}

// GetInternalVolumeName mocks base method.
func (m *MockDriver) GetInternalVolumeName(arg0 context.Context, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternalVolumeName", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInternalVolumeName indicates an expected call of GetInternalVolumeName.
func (mr *MockDriverMockRecorder) GetInternalVolumeName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternalVolumeName", reflect.TypeOf((*MockDriver)(nil).GetInternalVolumeName), arg0, arg1)
}

// GetProtocol mocks base method.
func (m *MockDriver) GetProtocol(arg0 context.Context) config.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtocol", arg0)
	ret0, _ := ret[0].(config.Protocol)
	return ret0
}

// GetProtocol indicates an expected call of GetProtocol.
func (mr *MockDriverMockRecorder) GetProtocol(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtocol", reflect.TypeOf((*MockDriver)(nil).GetProtocol), arg0)
}

// GetSnapshot mocks base method.
func (m *MockDriver) GetSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) (*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshot indicates an expected call of GetSnapshot.
func (mr *MockDriverMockRecorder) GetSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockDriver)(nil).GetSnapshot), arg0, arg1, arg2)
}

// GetSnapshots mocks base method.
func (m *MockDriver) GetSnapshots(arg0 context.Context, arg1 *storage.VolumeConfig) ([]*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshots", arg0, arg1)
	ret0, _ := ret[0].([]*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshots indicates an expected call of GetSnapshots.
func (mr *MockDriverMockRecorder) GetSnapshots(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshots", reflect.TypeOf((*MockDriver)(nil).GetSnapshots), arg0, arg1)
}

// GetStorageBackendPhysicalPoolNames mocks base method.
func (m *MockDriver) GetStorageBackendPhysicalPoolNames(arg0 context.Context) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageBackendPhysicalPoolNames", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStorageBackendPhysicalPoolNames indicates an expected call of GetStorageBackendPhysicalPoolNames.
func (mr *MockDriverMockRecorder) GetStorageBackendPhysicalPoolNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageBackendPhysicalPoolNames", reflect.TypeOf((*MockDriver)(nil).GetStorageBackendPhysicalPoolNames), arg0)
}

// GetStorageBackendSpecs mocks base method.
func (m *MockDriver) GetStorageBackendSpecs(arg0 context.Context, arg1 storage.Backend) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStorageBackendSpecs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStorageBackendSpecs indicates an expected call of GetStorageBackendSpecs.
func (mr *MockDriverMockRecorder) GetStorageBackendSpecs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageBackendSpecs", reflect.TypeOf((*MockDriver)(nil).GetStorageBackendSpecs), arg0, arg1)
}

// GetUpdateType mocks base method.
func (m *MockDriver) GetUpdateType(arg0 context.Context, arg1 storage.Driver) *roaring.Bitmap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateType", arg0, arg1)
	ret0, _ := ret[0].(*roaring.Bitmap)
	return ret0
}

// GetUpdateType indicates an expected call of GetUpdateType.
func (mr *MockDriverMockRecorder) GetUpdateType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateType", reflect.TypeOf((*MockDriver)(nil).GetUpdateType), arg0, arg1)
}

// GetVolumeExternal mocks base method.
func (m *MockDriver) GetVolumeExternal(arg0 context.Context, arg1 string) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeExternal", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeExternal indicates an expected call of GetVolumeExternal.
func (mr *MockDriverMockRecorder) GetVolumeExternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeExternal", reflect.TypeOf((*MockDriver)(nil).GetVolumeExternal), arg0, arg1)
}

// GetVolumeExternalWrappers mocks base method.
func (m *MockDriver) GetVolumeExternalWrappers(arg0 context.Context, arg1 chan *storage.VolumeExternalWrapper) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetVolumeExternalWrappers", arg0, arg1)
}

// GetVolumeExternalWrappers indicates an expected call of GetVolumeExternalWrappers.
func (mr *MockDriverMockRecorder) GetVolumeExternalWrappers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeExternalWrappers", reflect.TypeOf((*MockDriver)(nil).GetVolumeExternalWrappers), arg0, arg1)
}

// Import mocks base method.
func (m *MockDriver) Import(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Import", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Import indicates an expected call of Import.
func (mr *MockDriverMockRecorder) Import(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Import", reflect.TypeOf((*MockDriver)(nil).Import), arg0, arg1, arg2)
}

// Initialize mocks base method.
func (m *MockDriver) Initialize(arg0 context.Context, arg1 config.DriverContext, arg2 string, arg3 *storagedrivers.CommonStorageDriverConfig, arg4 map[string]string, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockDriverMockRecorder) Initialize(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockDriver)(nil).Initialize), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialized mocks base method.
func (m *MockDriver) Initialized() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialized")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Initialized indicates an expected call of Initialized.
func (mr *MockDriverMockRecorder) Initialized() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialized", reflect.TypeOf((*MockDriver)(nil).Initialized))
}

// Name mocks base method.
func (m *MockDriver) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDriverMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDriver)(nil).Name))
}

// Publish mocks base method.
func (m *MockDriver) Publish(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 *utils.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockDriverMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockDriver)(nil).Publish), arg0, arg1, arg2)
}

// ReconcileNodeAccess mocks base method.
func (m *MockDriver) ReconcileNodeAccess(arg0 context.Context, arg1 []*utils.Node, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNodeAccess", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileNodeAccess indicates an expected call of ReconcileNodeAccess.
func (mr *MockDriverMockRecorder) ReconcileNodeAccess(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNodeAccess", reflect.TypeOf((*MockDriver)(nil).ReconcileNodeAccess), arg0, arg1, arg2, arg3)
}

// Rename mocks base method.
func (m *MockDriver) Rename(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename.
func (mr *MockDriverMockRecorder) Rename(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockDriver)(nil).Rename), arg0, arg1, arg2)
}

// Resize mocks base method.
func (m *MockDriver) Resize(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resize", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Resize indicates an expected call of Resize.
func (mr *MockDriverMockRecorder) Resize(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resize", reflect.TypeOf((*MockDriver)(nil).Resize), arg0, arg1, arg2)
}

// RestoreSnapshot mocks base method.
func (m *MockDriver) RestoreSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreSnapshot indicates an expected call of RestoreSnapshot.
func (mr *MockDriverMockRecorder) RestoreSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSnapshot", reflect.TypeOf((*MockDriver)(nil).RestoreSnapshot), arg0, arg1, arg2)
}

// StoreConfig mocks base method.
func (m *MockDriver) StoreConfig(arg0 context.Context, arg1 *storage.PersistentStorageBackendConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StoreConfig", arg0, arg1)
}

// StoreConfig indicates an expected call of StoreConfig.
func (mr *MockDriverMockRecorder) StoreConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreConfig", reflect.TypeOf((*MockDriver)(nil).StoreConfig), arg0, arg1)
}

// Terminate mocks base method.
func (m *MockDriver) Terminate(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Terminate", arg0, arg1)
}

// Terminate indicates an expected call of Terminate.
func (mr *MockDriverMockRecorder) Terminate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MockDriver)(nil).Terminate), arg0, arg1)
}