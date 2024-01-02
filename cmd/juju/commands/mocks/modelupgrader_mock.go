// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/commands (interfaces: ModelUpgraderAPI)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/modelupgrader_mock.go github.com/juju/juju/cmd/juju/commands ModelUpgraderAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	tools "github.com/juju/juju/tools"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockModelUpgraderAPI is a mock of ModelUpgraderAPI interface.
type MockModelUpgraderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockModelUpgraderAPIMockRecorder
}

// MockModelUpgraderAPIMockRecorder is the mock recorder for MockModelUpgraderAPI.
type MockModelUpgraderAPIMockRecorder struct {
	mock *MockModelUpgraderAPI
}

// NewMockModelUpgraderAPI creates a new mock instance.
func NewMockModelUpgraderAPI(ctrl *gomock.Controller) *MockModelUpgraderAPI {
	mock := &MockModelUpgraderAPI{ctrl: ctrl}
	mock.recorder = &MockModelUpgraderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelUpgraderAPI) EXPECT() *MockModelUpgraderAPIMockRecorder {
	return m.recorder
}

// AbortModelUpgrade mocks base method.
func (m *MockModelUpgraderAPI) AbortModelUpgrade(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortModelUpgrade", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AbortModelUpgrade indicates an expected call of AbortModelUpgrade.
func (mr *MockModelUpgraderAPIMockRecorder) AbortModelUpgrade(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortModelUpgrade", reflect.TypeOf((*MockModelUpgraderAPI)(nil).AbortModelUpgrade), arg0)
}

// Close mocks base method.
func (m *MockModelUpgraderAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockModelUpgraderAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockModelUpgraderAPI)(nil).Close))
}

// UpgradeModel mocks base method.
func (m *MockModelUpgraderAPI) UpgradeModel(arg0 string, arg1 version.Number, arg2 string, arg3, arg4 bool) (version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeModel", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeModel indicates an expected call of UpgradeModel.
func (mr *MockModelUpgraderAPIMockRecorder) UpgradeModel(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeModel", reflect.TypeOf((*MockModelUpgraderAPI)(nil).UpgradeModel), arg0, arg1, arg2, arg3, arg4)
}

// UploadTools mocks base method.
func (m *MockModelUpgraderAPI) UploadTools(arg0 io.ReadSeeker, arg1 version.Binary) (tools.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadTools", arg0, arg1)
	ret0, _ := ret[0].(tools.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadTools indicates an expected call of UploadTools.
func (mr *MockModelUpgraderAPIMockRecorder) UploadTools(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTools", reflect.TypeOf((*MockModelUpgraderAPI)(nil).UploadTools), arg0, arg1)
}
