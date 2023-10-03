// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/changestream/eventqueue (interfaces: Stream)

// Package eventqueue is a generated GoMock package.
package eventqueue

import (
	reflect "reflect"

	changestream "github.com/juju/juju/core/changestream"
	gomock "go.uber.org/mock/gomock"
)

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Changes mocks base method.
func (m *MockStream) Changes() <-chan changestream.ChangeEvent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Changes")
	ret0, _ := ret[0].(<-chan changestream.ChangeEvent)
	return ret0
}

// Changes indicates an expected call of Changes.
func (mr *MockStreamMockRecorder) Changes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Changes", reflect.TypeOf((*MockStream)(nil).Changes))
}