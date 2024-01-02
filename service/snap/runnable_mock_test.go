// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/service/snap (interfaces: Runnable)
//
// Generated by this command:
//
//	mockgen -package snap -destination runnable_mock_test.go github.com/juju/juju/service/snap Runnable
//

// Package snap is a generated GoMock package.
package snap

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRunnable is a mock of Runnable interface.
type MockRunnable struct {
	ctrl     *gomock.Controller
	recorder *MockRunnableMockRecorder
}

// MockRunnableMockRecorder is the mock recorder for MockRunnable.
type MockRunnableMockRecorder struct {
	mock *MockRunnable
}

// NewMockRunnable creates a new mock instance.
func NewMockRunnable(ctrl *gomock.Controller) *MockRunnable {
	mock := &MockRunnable{ctrl: ctrl}
	mock.recorder = &MockRunnableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunnable) EXPECT() *MockRunnableMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockRunnable) Execute(arg0 string, arg1 ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockRunnableMockRecorder) Execute(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRunnable)(nil).Execute), varargs...)
}
