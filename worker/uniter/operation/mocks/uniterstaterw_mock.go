// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/uniter/operation (interfaces: UnitStateReadWriter)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/uniterstaterw_mock.go github.com/juju/juju/worker/uniter/operation UnitStateReadWriter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	params "github.com/juju/juju/rpc/params"
	gomock "go.uber.org/mock/gomock"
)

// MockUnitStateReadWriter is a mock of UnitStateReadWriter interface.
type MockUnitStateReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockUnitStateReadWriterMockRecorder
}

// MockUnitStateReadWriterMockRecorder is the mock recorder for MockUnitStateReadWriter.
type MockUnitStateReadWriterMockRecorder struct {
	mock *MockUnitStateReadWriter
}

// NewMockUnitStateReadWriter creates a new mock instance.
func NewMockUnitStateReadWriter(ctrl *gomock.Controller) *MockUnitStateReadWriter {
	mock := &MockUnitStateReadWriter{ctrl: ctrl}
	mock.recorder = &MockUnitStateReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitStateReadWriter) EXPECT() *MockUnitStateReadWriterMockRecorder {
	return m.recorder
}

// SetState mocks base method.
func (m *MockUnitStateReadWriter) SetState(arg0 params.SetUnitStateArg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetState indicates an expected call of SetState.
func (mr *MockUnitStateReadWriterMockRecorder) SetState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockUnitStateReadWriter)(nil).SetState), arg0)
}

// State mocks base method.
func (m *MockUnitStateReadWriter) State() (params.UnitStateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(params.UnitStateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// State indicates an expected call of State.
func (mr *MockUnitStateReadWriterMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockUnitStateReadWriter)(nil).State))
}
