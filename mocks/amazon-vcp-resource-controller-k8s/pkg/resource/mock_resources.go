// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-resource-controller-k8s/pkg/resource (interfaces: ResourceManager)

// Package mock_resource is a generated GoMock package.
package mock_resource

import (
	reflect "reflect"

	handler "github.com/aws/amazon-vpc-resource-controller-k8s/pkg/handler"
	provider "github.com/aws/amazon-vpc-resource-controller-k8s/pkg/provider"
	gomock "github.com/golang/mock/gomock"
)

// MockResourceManager is a mock of ResourceManager interface.
type MockResourceManager struct {
	ctrl     *gomock.Controller
	recorder *MockResourceManagerMockRecorder
}

// MockResourceManagerMockRecorder is the mock recorder for MockResourceManager.
type MockResourceManagerMockRecorder struct {
	mock *MockResourceManager
}

// NewMockResourceManager creates a new mock instance.
func NewMockResourceManager(ctrl *gomock.Controller) *MockResourceManager {
	mock := &MockResourceManager{ctrl: ctrl}
	mock.recorder = &MockResourceManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceManager) EXPECT() *MockResourceManagerMockRecorder {
	return m.recorder
}

// GetResourceHandler mocks base method.
func (m *MockResourceManager) GetResourceHandler(arg0 string) (handler.Handler, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceHandler", arg0)
	ret0, _ := ret[0].(handler.Handler)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetResourceHandler indicates an expected call of GetResourceHandler.
func (mr *MockResourceManagerMockRecorder) GetResourceHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceHandler", reflect.TypeOf((*MockResourceManager)(nil).GetResourceHandler), arg0)
}

// GetResourceProvider mocks base method.
func (m *MockResourceManager) GetResourceProvider(arg0 string) (provider.ResourceProvider, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceProvider", arg0)
	ret0, _ := ret[0].(provider.ResourceProvider)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetResourceProvider indicates an expected call of GetResourceProvider.
func (mr *MockResourceManagerMockRecorder) GetResourceProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceProvider", reflect.TypeOf((*MockResourceManager)(nil).GetResourceProvider), arg0)
}

// GetResourceProviders mocks base method.
func (m *MockResourceManager) GetResourceProviders() map[string]provider.ResourceProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceProviders")
	ret0, _ := ret[0].(map[string]provider.ResourceProvider)
	return ret0
}

// GetResourceProviders indicates an expected call of GetResourceProviders.
func (mr *MockResourceManagerMockRecorder) GetResourceProviders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceProviders", reflect.TypeOf((*MockResourceManager)(nil).GetResourceProviders))
}
