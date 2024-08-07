// Code generated by MockGen. DO NOT EDIT.
// Source: ../../capi/capi.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination=./mock_core.go -source=../../capi/capi.go -build_flags=-mod=mod
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	capi "github.com/platform9/cluster-api-sdk-go/capi"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockCAPICoreComponents is a mock of CAPICoreComponents interface.
type MockCAPICoreComponents struct {
	ctrl     *gomock.Controller
	recorder *MockCAPICoreComponentsMockRecorder
}

// MockCAPICoreComponentsMockRecorder is the mock recorder for MockCAPICoreComponents.
type MockCAPICoreComponentsMockRecorder struct {
	mock *MockCAPICoreComponents
}

// NewMockCAPICoreComponents creates a new mock instance.
func NewMockCAPICoreComponents(ctrl *gomock.Controller) *MockCAPICoreComponents {
	mock := &MockCAPICoreComponents{ctrl: ctrl}
	mock.recorder = &MockCAPICoreComponentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCAPICoreComponents) EXPECT() *MockCAPICoreComponentsMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method.
func (m *MockCAPICoreComponents) CreateCluster(ctx context.Context, input *capi.CreateClusterInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster.
func (mr *MockCAPICoreComponentsMockRecorder) CreateCluster(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockCAPICoreComponents)(nil).CreateCluster), ctx, input)
}

// CreateClusterResourceSet mocks base method.
func (m *MockCAPICoreComponents) CreateClusterResourceSet(ctx context.Context, input *capi.CreateClusterResourceSetInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClusterResourceSet", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClusterResourceSet indicates an expected call of CreateClusterResourceSet.
func (mr *MockCAPICoreComponentsMockRecorder) CreateClusterResourceSet(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClusterResourceSet", reflect.TypeOf((*MockCAPICoreComponents)(nil).CreateClusterResourceSet), ctx, input)
}

// CreateMachine mocks base method.
func (m *MockCAPICoreComponents) CreateMachine(ctx context.Context, input *capi.CreateMachineInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachine", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMachine indicates an expected call of CreateMachine.
func (mr *MockCAPICoreComponentsMockRecorder) CreateMachine(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachine", reflect.TypeOf((*MockCAPICoreComponents)(nil).CreateMachine), ctx, input)
}

// CreateMachineDeployment mocks base method.
func (m *MockCAPICoreComponents) CreateMachineDeployment(ctx context.Context, input *capi.CreateMachineDeploymentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachineDeployment", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMachineDeployment indicates an expected call of CreateMachineDeployment.
func (mr *MockCAPICoreComponentsMockRecorder) CreateMachineDeployment(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachineDeployment", reflect.TypeOf((*MockCAPICoreComponents)(nil).CreateMachineDeployment), ctx, input)
}

// CreateMachinePool mocks base method.
func (m *MockCAPICoreComponents) CreateMachinePool(ctx context.Context, input *capi.CreateMachinePoolInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMachinePool", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMachinePool indicates an expected call of CreateMachinePool.
func (mr *MockCAPICoreComponentsMockRecorder) CreateMachinePool(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMachinePool", reflect.TypeOf((*MockCAPICoreComponents)(nil).CreateMachinePool), ctx, input)
}

// DeleteCluster mocks base method.
func (m *MockCAPICoreComponents) DeleteCluster(ctx context.Context, input *capi.DeleteClusterInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster.
func (mr *MockCAPICoreComponentsMockRecorder) DeleteCluster(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockCAPICoreComponents)(nil).DeleteCluster), ctx, input)
}

// GetCluster mocks base method.
func (m *MockCAPICoreComponents) GetCluster(ctx context.Context, input *capi.GetClusterInput) (*v1beta1.Cluster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", ctx, input)
	ret0, _ := ret[0].(*v1beta1.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster.
func (mr *MockCAPICoreComponentsMockRecorder) GetCluster(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockCAPICoreComponents)(nil).GetCluster), ctx, input)
}

// GetMachineDeployment mocks base method.
func (m *MockCAPICoreComponents) GetMachineDeployment(ctx context.Context, input *capi.GetMachineDeploymentInput) (*v1beta1.MachineDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineDeployment", ctx, input)
	ret0, _ := ret[0].(*v1beta1.MachineDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineDeployment indicates an expected call of GetMachineDeployment.
func (mr *MockCAPICoreComponentsMockRecorder) GetMachineDeployment(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineDeployment", reflect.TypeOf((*MockCAPICoreComponents)(nil).GetMachineDeployment), ctx, input)
}

// ListMachineDeployment mocks base method.
func (m *MockCAPICoreComponents) ListMachineDeployment(ctx context.Context, input *capi.ListMachineDeploymentInput) (*v1beta1.MachineDeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMachineDeployment", ctx, input)
	ret0, _ := ret[0].(*v1beta1.MachineDeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMachineDeployment indicates an expected call of ListMachineDeployment.
func (mr *MockCAPICoreComponentsMockRecorder) ListMachineDeployment(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMachineDeployment", reflect.TypeOf((*MockCAPICoreComponents)(nil).ListMachineDeployment), ctx, input)
}

// UpdateMachineDeployment mocks base method.
func (m *MockCAPICoreComponents) UpdateMachineDeployment(ctx context.Context, input *capi.UpdateMachineDeploymentInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMachineDeployment", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMachineDeployment indicates an expected call of UpdateMachineDeployment.
func (mr *MockCAPICoreComponentsMockRecorder) UpdateMachineDeployment(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMachineDeployment", reflect.TypeOf((*MockCAPICoreComponents)(nil).UpdateMachineDeployment), ctx, input)
}
