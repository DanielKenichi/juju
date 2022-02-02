// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/tools/cache (interfaces: SharedIndexInformer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	cache "k8s.io/client-go/tools/cache"
)

// MockSharedIndexInformer is a mock of SharedIndexInformer interface.
type MockSharedIndexInformer struct {
	ctrl     *gomock.Controller
	recorder *MockSharedIndexInformerMockRecorder
}

// MockSharedIndexInformerMockRecorder is the mock recorder for MockSharedIndexInformer.
type MockSharedIndexInformerMockRecorder struct {
	mock *MockSharedIndexInformer
}

// NewMockSharedIndexInformer creates a new mock instance.
func NewMockSharedIndexInformer(ctrl *gomock.Controller) *MockSharedIndexInformer {
	mock := &MockSharedIndexInformer{ctrl: ctrl}
	mock.recorder = &MockSharedIndexInformerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharedIndexInformer) EXPECT() *MockSharedIndexInformerMockRecorder {
	return m.recorder
}

// AddEventHandler mocks base method.
func (m *MockSharedIndexInformer) AddEventHandler(arg0 cache.ResourceEventHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEventHandler", arg0)
}

// AddEventHandler indicates an expected call of AddEventHandler.
func (mr *MockSharedIndexInformerMockRecorder) AddEventHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandler", reflect.TypeOf((*MockSharedIndexInformer)(nil).AddEventHandler), arg0)
}

// AddEventHandlerWithResyncPeriod mocks base method.
func (m *MockSharedIndexInformer) AddEventHandlerWithResyncPeriod(arg0 cache.ResourceEventHandler, arg1 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEventHandlerWithResyncPeriod", arg0, arg1)
}

// AddEventHandlerWithResyncPeriod indicates an expected call of AddEventHandlerWithResyncPeriod.
func (mr *MockSharedIndexInformerMockRecorder) AddEventHandlerWithResyncPeriod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventHandlerWithResyncPeriod", reflect.TypeOf((*MockSharedIndexInformer)(nil).AddEventHandlerWithResyncPeriod), arg0, arg1)
}

// AddIndexers mocks base method.
func (m *MockSharedIndexInformer) AddIndexers(arg0 cache.Indexers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIndexers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddIndexers indicates an expected call of AddIndexers.
func (mr *MockSharedIndexInformerMockRecorder) AddIndexers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIndexers", reflect.TypeOf((*MockSharedIndexInformer)(nil).AddIndexers), arg0)
}

// GetController mocks base method.
func (m *MockSharedIndexInformer) GetController() cache.Controller {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetController")
	ret0, _ := ret[0].(cache.Controller)
	return ret0
}

// GetController indicates an expected call of GetController.
func (mr *MockSharedIndexInformerMockRecorder) GetController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetController", reflect.TypeOf((*MockSharedIndexInformer)(nil).GetController))
}

// GetIndexer mocks base method.
func (m *MockSharedIndexInformer) GetIndexer() cache.Indexer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexer")
	ret0, _ := ret[0].(cache.Indexer)
	return ret0
}

// GetIndexer indicates an expected call of GetIndexer.
func (mr *MockSharedIndexInformerMockRecorder) GetIndexer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexer", reflect.TypeOf((*MockSharedIndexInformer)(nil).GetIndexer))
}

// GetStore mocks base method.
func (m *MockSharedIndexInformer) GetStore() cache.Store {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore")
	ret0, _ := ret[0].(cache.Store)
	return ret0
}

// GetStore indicates an expected call of GetStore.
func (mr *MockSharedIndexInformerMockRecorder) GetStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockSharedIndexInformer)(nil).GetStore))
}

// HasSynced mocks base method.
func (m *MockSharedIndexInformer) HasSynced() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSynced")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSynced indicates an expected call of HasSynced.
func (mr *MockSharedIndexInformerMockRecorder) HasSynced() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSynced", reflect.TypeOf((*MockSharedIndexInformer)(nil).HasSynced))
}

// LastSyncResourceVersion mocks base method.
func (m *MockSharedIndexInformer) LastSyncResourceVersion() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastSyncResourceVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// LastSyncResourceVersion indicates an expected call of LastSyncResourceVersion.
func (mr *MockSharedIndexInformerMockRecorder) LastSyncResourceVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastSyncResourceVersion", reflect.TypeOf((*MockSharedIndexInformer)(nil).LastSyncResourceVersion))
}

// Run mocks base method.
func (m *MockSharedIndexInformer) Run(arg0 <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", arg0)
}

// Run indicates an expected call of Run.
func (mr *MockSharedIndexInformerMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockSharedIndexInformer)(nil).Run), arg0)
}

// SetWatchErrorHandler mocks base method.
func (m *MockSharedIndexInformer) SetWatchErrorHandler(arg0 cache.WatchErrorHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWatchErrorHandler", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWatchErrorHandler indicates an expected call of SetWatchErrorHandler.
func (mr *MockSharedIndexInformerMockRecorder) SetWatchErrorHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWatchErrorHandler", reflect.TypeOf((*MockSharedIndexInformer)(nil).SetWatchErrorHandler), arg0)
}
