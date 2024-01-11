// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/controller/charmdownloader (interfaces: StateBackend,ModelBackend,Application,Charm,Downloader,AuthChecker,ResourcesBackend)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mocks.go github.com/juju/juju/apiserver/facades/controller/charmdownloader StateBackend,ModelBackend,Application,Charm,Downloader,AuthChecker,ResourcesBackend
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	charm "github.com/juju/charm/v12"
	charmdownloader "github.com/juju/juju/apiserver/facades/controller/charmdownloader"
	controller "github.com/juju/juju/controller"
	charm0 "github.com/juju/juju/core/charm"
	status "github.com/juju/juju/core/status"
	config "github.com/juju/juju/environs/config"
	services "github.com/juju/juju/internal/charm/services"
	state "github.com/juju/juju/state"
	worker "github.com/juju/worker/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockStateBackend is a mock of StateBackend interface.
type MockStateBackend struct {
	ctrl     *gomock.Controller
	recorder *MockStateBackendMockRecorder
}

// MockStateBackendMockRecorder is the mock recorder for MockStateBackend.
type MockStateBackendMockRecorder struct {
	mock *MockStateBackend
}

// NewMockStateBackend creates a new mock instance.
func NewMockStateBackend(ctrl *gomock.Controller) *MockStateBackend {
	mock := &MockStateBackend{ctrl: ctrl}
	mock.recorder = &MockStateBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateBackend) EXPECT() *MockStateBackendMockRecorder {
	return m.recorder
}

// Application mocks base method.
func (m *MockStateBackend) Application(arg0 string) (charmdownloader.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(charmdownloader.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application.
func (mr *MockStateBackendMockRecorder) Application(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockStateBackend)(nil).Application), arg0)
}

// ControllerConfig mocks base method.
func (m *MockStateBackend) ControllerConfig() (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig")
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockStateBackendMockRecorder) ControllerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockStateBackend)(nil).ControllerConfig))
}

// ModelUUID mocks base method.
func (m *MockStateBackend) ModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockStateBackendMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockStateBackend)(nil).ModelUUID))
}

// PrepareCharmUpload mocks base method.
func (m *MockStateBackend) PrepareCharmUpload(arg0 string) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockStateBackendMockRecorder) PrepareCharmUpload(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockStateBackend)(nil).PrepareCharmUpload), arg0)
}

// UpdateUploadedCharm mocks base method.
func (m *MockStateBackend) UpdateUploadedCharm(arg0 state.CharmInfo) (services.UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(services.UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockStateBackendMockRecorder) UpdateUploadedCharm(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockStateBackend)(nil).UpdateUploadedCharm), arg0)
}

// WatchApplicationsWithPendingCharms mocks base method.
func (m *MockStateBackend) WatchApplicationsWithPendingCharms() state.StringsWatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplicationsWithPendingCharms")
	ret0, _ := ret[0].(state.StringsWatcher)
	return ret0
}

// WatchApplicationsWithPendingCharms indicates an expected call of WatchApplicationsWithPendingCharms.
func (mr *MockStateBackendMockRecorder) WatchApplicationsWithPendingCharms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplicationsWithPendingCharms", reflect.TypeOf((*MockStateBackend)(nil).WatchApplicationsWithPendingCharms))
}

// MockModelBackend is a mock of ModelBackend interface.
type MockModelBackend struct {
	ctrl     *gomock.Controller
	recorder *MockModelBackendMockRecorder
}

// MockModelBackendMockRecorder is the mock recorder for MockModelBackend.
type MockModelBackendMockRecorder struct {
	mock *MockModelBackend
}

// NewMockModelBackend creates a new mock instance.
func NewMockModelBackend(ctrl *gomock.Controller) *MockModelBackend {
	mock := &MockModelBackend{ctrl: ctrl}
	mock.recorder = &MockModelBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelBackend) EXPECT() *MockModelBackendMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockModelBackend) Config() (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockModelBackendMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockModelBackend)(nil).Config))
}

// MockApplication is a mock of Application interface.
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication.
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance.
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Charm mocks base method.
func (m *MockApplication) Charm() (charmdownloader.Charm, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm")
	ret0, _ := ret[0].(charmdownloader.Charm)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Charm indicates an expected call of Charm.
func (mr *MockApplicationMockRecorder) Charm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockApplication)(nil).Charm))
}

// CharmOrigin mocks base method.
func (m *MockApplication) CharmOrigin() *charm0.Origin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmOrigin")
	ret0, _ := ret[0].(*charm0.Origin)
	return ret0
}

// CharmOrigin indicates an expected call of CharmOrigin.
func (mr *MockApplicationMockRecorder) CharmOrigin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmOrigin", reflect.TypeOf((*MockApplication)(nil).CharmOrigin))
}

// CharmPendingToBeDownloaded mocks base method.
func (m *MockApplication) CharmPendingToBeDownloaded() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmPendingToBeDownloaded")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CharmPendingToBeDownloaded indicates an expected call of CharmPendingToBeDownloaded.
func (mr *MockApplicationMockRecorder) CharmPendingToBeDownloaded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmPendingToBeDownloaded", reflect.TypeOf((*MockApplication)(nil).CharmPendingToBeDownloaded))
}

// SetDownloadedIDAndHash mocks base method.
func (m *MockApplication) SetDownloadedIDAndHash(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDownloadedIDAndHash", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDownloadedIDAndHash indicates an expected call of SetDownloadedIDAndHash.
func (mr *MockApplicationMockRecorder) SetDownloadedIDAndHash(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDownloadedIDAndHash", reflect.TypeOf((*MockApplication)(nil).SetDownloadedIDAndHash), arg0, arg1)
}

// SetStatus mocks base method.
func (m *MockApplication) SetStatus(arg0 status.StatusInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockApplicationMockRecorder) SetStatus(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockApplication)(nil).SetStatus), arg0)
}

// MockCharm is a mock of Charm interface.
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm.
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance.
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// URL mocks base method.
func (m *MockCharm) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockCharmMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockCharm)(nil).URL))
}

// MockDownloader is a mock of Downloader interface.
type MockDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockDownloaderMockRecorder
}

// MockDownloaderMockRecorder is the mock recorder for MockDownloader.
type MockDownloaderMockRecorder struct {
	mock *MockDownloader
}

// NewMockDownloader creates a new mock instance.
func NewMockDownloader(ctrl *gomock.Controller) *MockDownloader {
	mock := &MockDownloader{ctrl: ctrl}
	mock.recorder = &MockDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloader) EXPECT() *MockDownloaderMockRecorder {
	return m.recorder
}

// DownloadAndStore mocks base method.
func (m *MockDownloader) DownloadAndStore(arg0 context.Context, arg1 *charm.URL, arg2 charm0.Origin, arg3 bool) (charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadAndStore", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm0.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadAndStore indicates an expected call of DownloadAndStore.
func (mr *MockDownloaderMockRecorder) DownloadAndStore(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadAndStore", reflect.TypeOf((*MockDownloader)(nil).DownloadAndStore), arg0, arg1, arg2, arg3)
}

// MockAuthChecker is a mock of AuthChecker interface.
type MockAuthChecker struct {
	ctrl     *gomock.Controller
	recorder *MockAuthCheckerMockRecorder
}

// MockAuthCheckerMockRecorder is the mock recorder for MockAuthChecker.
type MockAuthCheckerMockRecorder struct {
	mock *MockAuthChecker
}

// NewMockAuthChecker creates a new mock instance.
func NewMockAuthChecker(ctrl *gomock.Controller) *MockAuthChecker {
	mock := &MockAuthChecker{ctrl: ctrl}
	mock.recorder = &MockAuthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthChecker) EXPECT() *MockAuthCheckerMockRecorder {
	return m.recorder
}

// AuthController mocks base method.
func (m *MockAuthChecker) AuthController() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthController")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthController indicates an expected call of AuthController.
func (mr *MockAuthCheckerMockRecorder) AuthController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthController", reflect.TypeOf((*MockAuthChecker)(nil).AuthController))
}

// MockResourcesBackend is a mock of ResourcesBackend interface.
type MockResourcesBackend struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesBackendMockRecorder
}

// MockResourcesBackendMockRecorder is the mock recorder for MockResourcesBackend.
type MockResourcesBackendMockRecorder struct {
	mock *MockResourcesBackend
}

// NewMockResourcesBackend creates a new mock instance.
func NewMockResourcesBackend(ctrl *gomock.Controller) *MockResourcesBackend {
	mock := &MockResourcesBackend{ctrl: ctrl}
	mock.recorder = &MockResourcesBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourcesBackend) EXPECT() *MockResourcesBackendMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockResourcesBackend) Register(arg0 worker.Worker) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockResourcesBackendMockRecorder) Register(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockResourcesBackend)(nil).Register), arg0)
}
