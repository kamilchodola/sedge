// Code generated by MockGen. DO NOT EDIT.
// Source: ./cli/prompts/prompt.go

// Package mock_prompts is a generated GoMock package.
package mock_prompts

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPrompt is a mock of Prompt interface.
type MockPrompt struct {
	ctrl     *gomock.Controller
	recorder *MockPromptMockRecorder
}

// MockPromptMockRecorder is the mock recorder for MockPrompt.
type MockPromptMockRecorder struct {
	mock *MockPrompt
}

// NewMockPrompt creates a new mock instance.
func NewMockPrompt(ctrl *gomock.Controller) *MockPrompt {
	mock := &MockPrompt{ctrl: ctrl}
	mock.recorder = &MockPromptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrompt) EXPECT() *MockPromptMockRecorder {
	return m.recorder
}

// Eth1Withdrawal mocks base method.
func (m *MockPrompt) Eth1Withdrawal() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eth1Withdrawal")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Eth1Withdrawal indicates an expected call of Eth1Withdrawal.
func (mr *MockPromptMockRecorder) Eth1Withdrawal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eth1Withdrawal", reflect.TypeOf((*MockPrompt)(nil).Eth1Withdrawal))
}

// ExistingVal mocks base method.
func (m *MockPrompt) ExistingVal() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistingVal")
	ret0, _ := ret[0].(int64)
	return ret0
}

// ExistingVal indicates an expected call of ExistingVal.
func (mr *MockPromptMockRecorder) ExistingVal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistingVal", reflect.TypeOf((*MockPrompt)(nil).ExistingVal))
}

// FeeRecipient mocks base method.
func (m *MockPrompt) FeeRecipient() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FeeRecipient")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FeeRecipient indicates an expected call of FeeRecipient.
func (mr *MockPromptMockRecorder) FeeRecipient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FeeRecipient", reflect.TypeOf((*MockPrompt)(nil).FeeRecipient))
}

// NumberVal mocks base method.
func (m *MockPrompt) NumberVal() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumberVal")
	ret0, _ := ret[0].(int64)
	return ret0
}

// NumberVal indicates an expected call of NumberVal.
func (mr *MockPromptMockRecorder) NumberVal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumberVal", reflect.TypeOf((*MockPrompt)(nil).NumberVal))
}

// Passphrase mocks base method.
func (m *MockPrompt) Passphrase() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Passphrase")
	ret0, _ := ret[0].(string)
	return ret0
}

// Passphrase indicates an expected call of Passphrase.
func (mr *MockPromptMockRecorder) Passphrase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Passphrase", reflect.TypeOf((*MockPrompt)(nil).Passphrase))
}
