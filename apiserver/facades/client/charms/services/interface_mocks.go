// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/charms/services (interfaces: StateBackend,ModelBackend,Storage,UploadedCharm)

// Package services is a generated GoMock package.
package services

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v8 "github.com/juju/charm/v8"
	controller "github.com/juju/juju/controller"
	config "github.com/juju/juju/environs/config"
	state "github.com/juju/juju/state"
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
func (m *MockStateBackend) PrepareCharmUpload(arg0 *v8.URL) (UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareCharmUpload", arg0)
	ret0, _ := ret[0].(UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareCharmUpload indicates an expected call of PrepareCharmUpload.
func (mr *MockStateBackendMockRecorder) PrepareCharmUpload(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareCharmUpload", reflect.TypeOf((*MockStateBackend)(nil).PrepareCharmUpload), arg0)
}

// UpdateUploadedCharm mocks base method.
func (m *MockStateBackend) UpdateUploadedCharm(arg0 state.CharmInfo) (UploadedCharm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUploadedCharm", arg0)
	ret0, _ := ret[0].(UploadedCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUploadedCharm indicates an expected call of UpdateUploadedCharm.
func (mr *MockStateBackendMockRecorder) UpdateUploadedCharm(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUploadedCharm", reflect.TypeOf((*MockStateBackend)(nil).UpdateUploadedCharm), arg0)
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

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Put mocks base method.
func (m *MockStorage) Put(arg0 string, arg1 io.Reader, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockStorageMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockStorage)(nil).Put), arg0, arg1, arg2)
}

// Remove mocks base method.
func (m *MockStorage) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStorageMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStorage)(nil).Remove), arg0)
}

// MockUploadedCharm is a mock of UploadedCharm interface.
type MockUploadedCharm struct {
	ctrl     *gomock.Controller
	recorder *MockUploadedCharmMockRecorder
}

// MockUploadedCharmMockRecorder is the mock recorder for MockUploadedCharm.
type MockUploadedCharmMockRecorder struct {
	mock *MockUploadedCharm
}

// NewMockUploadedCharm creates a new mock instance.
func NewMockUploadedCharm(ctrl *gomock.Controller) *MockUploadedCharm {
	mock := &MockUploadedCharm{ctrl: ctrl}
	mock.recorder = &MockUploadedCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadedCharm) EXPECT() *MockUploadedCharmMockRecorder {
	return m.recorder
}

// IsUploaded mocks base method.
func (m *MockUploadedCharm) IsUploaded() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUploaded")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUploaded indicates an expected call of IsUploaded.
func (mr *MockUploadedCharmMockRecorder) IsUploaded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUploaded", reflect.TypeOf((*MockUploadedCharm)(nil).IsUploaded))
}
