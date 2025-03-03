// Code generated by MockGen. DO NOT EDIT.
// Source: ./cli/actions/actions.go

// Package mock_actions is a generated GoMock package.
package mock_actions

import (
	reflect "reflect"

	actions "github.com/NethermindEth/sedge/cli/actions"
	gomock "github.com/golang/mock/gomock"
)

// MockSedgeActions is a mock of SedgeActions interface.
type MockSedgeActions struct {
	ctrl     *gomock.Controller
	recorder *MockSedgeActionsMockRecorder
}

// MockSedgeActionsMockRecorder is the mock recorder for MockSedgeActions.
type MockSedgeActionsMockRecorder struct {
	mock *MockSedgeActions
}

// NewMockSedgeActions creates a new mock instance.
func NewMockSedgeActions(ctrl *gomock.Controller) *MockSedgeActions {
	mock := &MockSedgeActions{ctrl: ctrl}
	mock.recorder = &MockSedgeActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSedgeActions) EXPECT() *MockSedgeActionsMockRecorder {
	return m.recorder
}

// CreateJWTSecrets mocks base method.
func (m *MockSedgeActions) CreateJWTSecrets(arg0 actions.CreateJWTSecretOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJWTSecrets", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJWTSecrets indicates an expected call of CreateJWTSecrets.
func (mr *MockSedgeActionsMockRecorder) CreateJWTSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJWTSecrets", reflect.TypeOf((*MockSedgeActions)(nil).CreateJWTSecrets), arg0)
}

// ExportSlashingInterchangeData mocks base method.
func (m *MockSedgeActions) ExportSlashingInterchangeData(arg0 actions.SlashingExportOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportSlashingInterchangeData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExportSlashingInterchangeData indicates an expected call of ExportSlashingInterchangeData.
func (mr *MockSedgeActionsMockRecorder) ExportSlashingInterchangeData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportSlashingInterchangeData", reflect.TypeOf((*MockSedgeActions)(nil).ExportSlashingInterchangeData), arg0)
}

// Generate mocks base method.
func (m *MockSedgeActions) Generate(arg0 actions.GenerateOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockSedgeActionsMockRecorder) Generate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockSedgeActions)(nil).Generate), arg0)
}

// ImportSlashingInterchangeData mocks base method.
func (m *MockSedgeActions) ImportSlashingInterchangeData(arg0 actions.SlashingImportOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportSlashingInterchangeData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportSlashingInterchangeData indicates an expected call of ImportSlashingInterchangeData.
func (mr *MockSedgeActionsMockRecorder) ImportSlashingInterchangeData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportSlashingInterchangeData", reflect.TypeOf((*MockSedgeActions)(nil).ImportSlashingInterchangeData), arg0)
}

// InstallDependencies mocks base method.
func (m *MockSedgeActions) InstallDependencies(arg0 actions.InstallDependenciesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallDependencies", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallDependencies indicates an expected call of InstallDependencies.
func (mr *MockSedgeActionsMockRecorder) InstallDependencies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallDependencies", reflect.TypeOf((*MockSedgeActions)(nil).InstallDependencies), arg0)
}

// RunContainers mocks base method.
func (m *MockSedgeActions) RunContainers(arg0 actions.RunContainersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunContainers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunContainers indicates an expected call of RunContainers.
func (mr *MockSedgeActionsMockRecorder) RunContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunContainers", reflect.TypeOf((*MockSedgeActions)(nil).RunContainers), arg0)
}

// SetupContainers mocks base method.
func (m *MockSedgeActions) SetupContainers(arg0 actions.SetupContainersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupContainers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupContainers indicates an expected call of SetupContainers.
func (mr *MockSedgeActionsMockRecorder) SetupContainers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupContainers", reflect.TypeOf((*MockSedgeActions)(nil).SetupContainers), arg0)
}
