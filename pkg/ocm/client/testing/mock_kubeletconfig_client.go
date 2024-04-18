// Code generated by MockGen. DO NOT EDIT.
// Source: kubeletconfig_client.go
//
// Generated by this command:
//
//	mockgen -source=kubeletconfig_client.go -package=testing -destination=testing/mock_kubeletconfig_client.go
//

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockKubeletConfigClient is a mock of KubeletConfigClient interface.
type MockKubeletConfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeletConfigClientMockRecorder
}

// MockKubeletConfigClientMockRecorder is the mock recorder for MockKubeletConfigClient.
type MockKubeletConfigClientMockRecorder struct {
	mock *MockKubeletConfigClient
}

// NewMockKubeletConfigClient creates a new mock instance.
func NewMockKubeletConfigClient(ctrl *gomock.Controller) *MockKubeletConfigClient {
	mock := &MockKubeletConfigClient{ctrl: ctrl}
	mock.recorder = &MockKubeletConfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeletConfigClient) EXPECT() *MockKubeletConfigClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockKubeletConfigClient) Create(ctx context.Context, clusterId string, instance *v1.KubeletConfig) (*v1.KubeletConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, clusterId, instance)
	ret0, _ := ret[0].(*v1.KubeletConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockKubeletConfigClientMockRecorder) Create(ctx, clusterId, instance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKubeletConfigClient)(nil).Create), ctx, clusterId, instance)
}

// Delete mocks base method.
func (m *MockKubeletConfigClient) Delete(ctx context.Context, clusterId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, clusterId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockKubeletConfigClientMockRecorder) Delete(ctx, clusterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKubeletConfigClient)(nil).Delete), ctx, clusterId)
}

// Exists mocks base method.
func (m *MockKubeletConfigClient) Exists(ctx context.Context, clusterId string) (bool, *v1.KubeletConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, clusterId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*v1.KubeletConfig)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Exists indicates an expected call of Exists.
func (mr *MockKubeletConfigClientMockRecorder) Exists(ctx, clusterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockKubeletConfigClient)(nil).Exists), ctx, clusterId)
}

// Get mocks base method.
func (m *MockKubeletConfigClient) Get(ctx context.Context, clusterId string) (*v1.KubeletConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, clusterId)
	ret0, _ := ret[0].(*v1.KubeletConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKubeletConfigClientMockRecorder) Get(ctx, clusterId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKubeletConfigClient)(nil).Get), ctx, clusterId)
}

// Update mocks base method.
func (m *MockKubeletConfigClient) Update(ctx context.Context, clusterId string, instance *v1.KubeletConfig) (*v1.KubeletConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, clusterId, instance)
	ret0, _ := ret[0].(*v1.KubeletConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockKubeletConfigClientMockRecorder) Update(ctx, clusterId, instance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockKubeletConfigClient)(nil).Update), ctx, clusterId, instance)
}
