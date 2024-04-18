// Code generated by MockGen. DO NOT EDIT.
// Source: machinepool_client.go
//
// Generated by this command:
//
//	mockgen -source=machinepool_client.go -package=testing -destination=testing/mock_machinepool_client.go
//

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	client "github.com/openshift-online/ocm-common/pkg/ocm/client"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockMachinePoolClient is a mock of MachinePoolClient interface.
type MockMachinePoolClient struct {
	ctrl     *gomock.Controller
	recorder *MockMachinePoolClientMockRecorder
}

// MockMachinePoolClientMockRecorder is the mock recorder for MockMachinePoolClient.
type MockMachinePoolClientMockRecorder struct {
	mock *MockMachinePoolClient
}

// NewMockMachinePoolClient creates a new mock instance.
func NewMockMachinePoolClient(ctrl *gomock.Controller) *MockMachinePoolClient {
	mock := &MockMachinePoolClient{ctrl: ctrl}
	mock.recorder = &MockMachinePoolClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachinePoolClient) EXPECT() *MockMachinePoolClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMachinePoolClient) Create(ctx context.Context, clusterId string, instance *v1.MachinePool) (*v1.MachinePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, clusterId, instance)
	ret0, _ := ret[0].(*v1.MachinePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMachinePoolClientMockRecorder) Create(ctx, clusterId, instance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMachinePoolClient)(nil).Create), ctx, clusterId, instance)
}

// Delete mocks base method.
func (m *MockMachinePoolClient) Delete(ctx context.Context, clusterId, instanceId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, clusterId, instanceId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMachinePoolClientMockRecorder) Delete(ctx, clusterId, instanceId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMachinePoolClient)(nil).Delete), ctx, clusterId, instanceId)
}

// Exists mocks base method.
func (m *MockMachinePoolClient) Exists(ctx context.Context, clusterId, instanceId string) (bool, *v1.MachinePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, clusterId, instanceId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*v1.MachinePool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Exists indicates an expected call of Exists.
func (mr *MockMachinePoolClientMockRecorder) Exists(ctx, clusterId, instanceId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockMachinePoolClient)(nil).Exists), ctx, clusterId, instanceId)
}

// Get mocks base method.
func (m *MockMachinePoolClient) Get(ctx context.Context, clusterId, instanceId string) (*v1.MachinePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, clusterId, instanceId)
	ret0, _ := ret[0].(*v1.MachinePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMachinePoolClientMockRecorder) Get(ctx, clusterId, instanceId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMachinePoolClient)(nil).Get), ctx, clusterId, instanceId)
}

// List mocks base method.
func (m *MockMachinePoolClient) List(ctx context.Context, clusterId string, paging client.Paging) ([]*v1.MachinePool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, clusterId, paging)
	ret0, _ := ret[0].([]*v1.MachinePool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockMachinePoolClientMockRecorder) List(ctx, clusterId, paging any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMachinePoolClient)(nil).List), ctx, clusterId, paging)
}

// Update mocks base method.
func (m *MockMachinePoolClient) Update(ctx context.Context, clusterId string, instance *v1.MachinePool) (*v1.MachinePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, clusterId, instance)
	ret0, _ := ret[0].(*v1.MachinePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMachinePoolClientMockRecorder) Update(ctx, clusterId, instance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMachinePoolClient)(nil).Update), ctx, clusterId, instance)
}
