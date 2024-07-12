// Code generated by MockGen. DO NOT EDIT.
// Source: ../../controlplane/controlplane.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination=./mock_controlplane.go -source=../../controlplane/controlplane.go -build_flags=-mod=mod
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	controlplane "github.com/platform9/cluster-api-sdk-go/controlplane"
	gomock "go.uber.org/mock/gomock"
)

// MockControlPlaneProvider is a mock of ControlPlaneProvider interface.
type MockControlPlaneProvider struct {
	ctrl     *gomock.Controller
	recorder *MockControlPlaneProviderMockRecorder
}

// MockControlPlaneProviderMockRecorder is the mock recorder for MockControlPlaneProvider.
type MockControlPlaneProviderMockRecorder struct {
	mock *MockControlPlaneProvider
}

// NewMockControlPlaneProvider creates a new mock instance.
func NewMockControlPlaneProvider(ctrl *gomock.Controller) *MockControlPlaneProvider {
	mock := &MockControlPlaneProvider{ctrl: ctrl}
	mock.recorder = &MockControlPlaneProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControlPlaneProvider) EXPECT() *MockControlPlaneProviderMockRecorder {
	return m.recorder
}

// CreateControlPlane mocks base method.
func (m *MockControlPlaneProvider) CreateControlPlane(ctx context.Context, input controlplane.CreateControlPlaneInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateControlPlane", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateControlPlane indicates an expected call of CreateControlPlane.
func (mr *MockControlPlaneProviderMockRecorder) CreateControlPlane(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateControlPlane", reflect.TypeOf((*MockControlPlaneProvider)(nil).CreateControlPlane), ctx, input)
}

// DeleteControlPlane mocks base method.
func (m *MockControlPlaneProvider) DeleteControlPlane(ctx context.Context, input controlplane.DeleteControlPlaneInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteControlPlane", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteControlPlane indicates an expected call of DeleteControlPlane.
func (mr *MockControlPlaneProviderMockRecorder) DeleteControlPlane(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteControlPlane", reflect.TypeOf((*MockControlPlaneProvider)(nil).DeleteControlPlane), ctx, input)
}

// GetControlPlane mocks base method.
func (m *MockControlPlaneProvider) GetControlPlane(ctx context.Context, input controlplane.GetControlPlaneInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControlPlane", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetControlPlane indicates an expected call of GetControlPlane.
func (mr *MockControlPlaneProviderMockRecorder) GetControlPlane(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControlPlane", reflect.TypeOf((*MockControlPlaneProvider)(nil).GetControlPlane), ctx, input)
}

// MockGetControlPlaneInput is a mock of GetControlPlaneInput interface.
type MockGetControlPlaneInput struct {
	ctrl     *gomock.Controller
	recorder *MockGetControlPlaneInputMockRecorder
}

// MockGetControlPlaneInputMockRecorder is the mock recorder for MockGetControlPlaneInput.
type MockGetControlPlaneInputMockRecorder struct {
	mock *MockGetControlPlaneInput
}

// NewMockGetControlPlaneInput creates a new mock instance.
func NewMockGetControlPlaneInput(ctrl *gomock.Controller) *MockGetControlPlaneInput {
	mock := &MockGetControlPlaneInput{ctrl: ctrl}
	mock.recorder = &MockGetControlPlaneInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetControlPlaneInput) EXPECT() *MockGetControlPlaneInputMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockGetControlPlaneInput) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockGetControlPlaneInputMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockGetControlPlaneInput)(nil).GetName))
}

// MockCreateControlPlaneInput is a mock of CreateControlPlaneInput interface.
type MockCreateControlPlaneInput struct {
	ctrl     *gomock.Controller
	recorder *MockCreateControlPlaneInputMockRecorder
}

// MockCreateControlPlaneInputMockRecorder is the mock recorder for MockCreateControlPlaneInput.
type MockCreateControlPlaneInputMockRecorder struct {
	mock *MockCreateControlPlaneInput
}

// NewMockCreateControlPlaneInput creates a new mock instance.
func NewMockCreateControlPlaneInput(ctrl *gomock.Controller) *MockCreateControlPlaneInput {
	mock := &MockCreateControlPlaneInput{ctrl: ctrl}
	mock.recorder = &MockCreateControlPlaneInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateControlPlaneInput) EXPECT() *MockCreateControlPlaneInputMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockCreateControlPlaneInput) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockCreateControlPlaneInputMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockCreateControlPlaneInput)(nil).GetName))
}

// MockDeleteControlPlaneInput is a mock of DeleteControlPlaneInput interface.
type MockDeleteControlPlaneInput struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteControlPlaneInputMockRecorder
}

// MockDeleteControlPlaneInputMockRecorder is the mock recorder for MockDeleteControlPlaneInput.
type MockDeleteControlPlaneInputMockRecorder struct {
	mock *MockDeleteControlPlaneInput
}

// NewMockDeleteControlPlaneInput creates a new mock instance.
func NewMockDeleteControlPlaneInput(ctrl *gomock.Controller) *MockDeleteControlPlaneInput {
	mock := &MockDeleteControlPlaneInput{ctrl: ctrl}
	mock.recorder = &MockDeleteControlPlaneInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteControlPlaneInput) EXPECT() *MockDeleteControlPlaneInputMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockDeleteControlPlaneInput) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockDeleteControlPlaneInputMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockDeleteControlPlaneInput)(nil).GetName))
}
