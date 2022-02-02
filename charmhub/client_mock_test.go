// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/charmhub (interfaces: Transport,RESTClient,FileSystem,Logger)

// Package charmhub is a generated GoMock package.
package charmhub

import (
	context "context"
	http "net/http"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	path "github.com/juju/juju/charmhub/path"
	loggo "github.com/juju/loggo"
)

// MockTransport is a mock of Transport interface.
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *MockTransportMockRecorder
}

// MockTransportMockRecorder is the mock recorder for MockTransport.
type MockTransportMockRecorder struct {
	mock *MockTransport
}

// NewMockTransport creates a new mock instance.
func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &MockTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransport) EXPECT() *MockTransportMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockTransport) Do(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockTransportMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockTransport)(nil).Do), arg0)
}

// MockRESTClient is a mock of RESTClient interface.
type MockRESTClient struct {
	ctrl     *gomock.Controller
	recorder *MockRESTClientMockRecorder
}

// MockRESTClientMockRecorder is the mock recorder for MockRESTClient.
type MockRESTClientMockRecorder struct {
	mock *MockRESTClient
}

// NewMockRESTClient creates a new mock instance.
func NewMockRESTClient(ctrl *gomock.Controller) *MockRESTClient {
	mock := &MockRESTClient{ctrl: ctrl}
	mock.recorder = &MockRESTClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRESTClient) EXPECT() *MockRESTClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRESTClient) Get(arg0 context.Context, arg1 path.Path, arg2 interface{}) (RESTResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(RESTResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRESTClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRESTClient)(nil).Get), arg0, arg1, arg2)
}

// Post mocks base method.
func (m *MockRESTClient) Post(arg0 context.Context, arg1 path.Path, arg2 http.Header, arg3, arg4 interface{}) (RESTResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(RESTResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockRESTClientMockRecorder) Post(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockRESTClient)(nil).Post), arg0, arg1, arg2, arg3, arg4)
}

// MockFileSystem is a mock of FileSystem interface.
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemMockRecorder
}

// MockFileSystemMockRecorder is the mock recorder for MockFileSystem.
type MockFileSystemMockRecorder struct {
	mock *MockFileSystem
}

// NewMockFileSystem creates a new mock instance.
func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &MockFileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSystem) EXPECT() *MockFileSystemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFileSystem) Create(arg0 string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFileSystemMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFileSystem)(nil).Create), arg0)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// ChildWithLabels mocks base method.
func (m *MockLogger) ChildWithLabels(arg0 string, arg1 ...string) loggo.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChildWithLabels", varargs...)
	ret0, _ := ret[0].(loggo.Logger)
	return ret0
}

// ChildWithLabels indicates an expected call of ChildWithLabels.
func (mr *MockLoggerMockRecorder) ChildWithLabels(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChildWithLabels", reflect.TypeOf((*MockLogger)(nil).ChildWithLabels), varargs...)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// IsTraceEnabled mocks base method.
func (m *MockLogger) IsTraceEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTraceEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTraceEnabled indicates an expected call of IsTraceEnabled.
func (mr *MockLoggerMockRecorder) IsTraceEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTraceEnabled", reflect.TypeOf((*MockLogger)(nil).IsTraceEnabled))
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}
