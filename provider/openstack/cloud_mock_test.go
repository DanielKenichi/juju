// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cloudconfig/cloudinit (interfaces: NetworkingConfig)
//
// Generated by this command:
//
//	mockgen -package openstack -destination cloud_mock_test.go github.com/juju/juju/cloudconfig/cloudinit NetworkingConfig
//

// Package openstack is a generated GoMock package.
package openstack

import (
	reflect "reflect"

	network "github.com/juju/juju/core/network"
	gomock "go.uber.org/mock/gomock"
)

// MockNetworkingConfig is a mock of NetworkingConfig interface.
type MockNetworkingConfig struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkingConfigMockRecorder
}

// MockNetworkingConfigMockRecorder is the mock recorder for MockNetworkingConfig.
type MockNetworkingConfigMockRecorder struct {
	mock *MockNetworkingConfig
}

// NewMockNetworkingConfig creates a new mock instance.
func NewMockNetworkingConfig(ctrl *gomock.Controller) *MockNetworkingConfig {
	mock := &MockNetworkingConfig{ctrl: ctrl}
	mock.recorder = &MockNetworkingConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkingConfig) EXPECT() *MockNetworkingConfigMockRecorder {
	return m.recorder
}

// AddNetworkConfig mocks base method.
func (m *MockNetworkingConfig) AddNetworkConfig(arg0 network.InterfaceInfos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNetworkConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNetworkConfig indicates an expected call of AddNetworkConfig.
func (mr *MockNetworkingConfigMockRecorder) AddNetworkConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNetworkConfig", reflect.TypeOf((*MockNetworkingConfig)(nil).AddNetworkConfig), arg0)
}
