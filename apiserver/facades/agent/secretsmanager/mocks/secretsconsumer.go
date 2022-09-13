// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/secretsmanager (interfaces: SecretsConsumer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	secrets "github.com/juju/juju/core/secrets"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
)

// MockSecretsConsumer is a mock of SecretsConsumer interface.
type MockSecretsConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsConsumerMockRecorder
}

// MockSecretsConsumerMockRecorder is the mock recorder for MockSecretsConsumer.
type MockSecretsConsumerMockRecorder struct {
	mock *MockSecretsConsumer
}

// NewMockSecretsConsumer creates a new mock instance.
func NewMockSecretsConsumer(ctrl *gomock.Controller) *MockSecretsConsumer {
	mock := &MockSecretsConsumer{ctrl: ctrl}
	mock.recorder = &MockSecretsConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsConsumer) EXPECT() *MockSecretsConsumerMockRecorder {
	return m.recorder
}

// GetSecretConsumer mocks base method.
func (m *MockSecretsConsumer) GetSecretConsumer(arg0 *secrets.URI, arg1 string) (*secrets.SecretConsumerMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretConsumer", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretConsumerMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretConsumer indicates an expected call of GetSecretConsumer.
func (mr *MockSecretsConsumerMockRecorder) GetSecretConsumer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretConsumer", reflect.TypeOf((*MockSecretsConsumer)(nil).GetSecretConsumer), arg0, arg1)
}

// GrantSecretAccess mocks base method.
func (m *MockSecretsConsumer) GrantSecretAccess(arg0 *secrets.URI, arg1 state.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantSecretAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantSecretAccess indicates an expected call of GrantSecretAccess.
func (mr *MockSecretsConsumerMockRecorder) GrantSecretAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantSecretAccess", reflect.TypeOf((*MockSecretsConsumer)(nil).GrantSecretAccess), arg0, arg1)
}

// RevokeSecretAccess mocks base method.
func (m *MockSecretsConsumer) RevokeSecretAccess(arg0 *secrets.URI, arg1 state.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecretAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeSecretAccess indicates an expected call of RevokeSecretAccess.
func (mr *MockSecretsConsumerMockRecorder) RevokeSecretAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecretAccess", reflect.TypeOf((*MockSecretsConsumer)(nil).RevokeSecretAccess), arg0, arg1)
}

// SaveSecretConsumer mocks base method.
func (m *MockSecretsConsumer) SaveSecretConsumer(arg0 *secrets.URI, arg1 string, arg2 *secrets.SecretConsumerMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSecretConsumer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSecretConsumer indicates an expected call of SaveSecretConsumer.
func (mr *MockSecretsConsumerMockRecorder) SaveSecretConsumer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSecretConsumer", reflect.TypeOf((*MockSecretsConsumer)(nil).SaveSecretConsumer), arg0, arg1, arg2)
}

// SecretAccess mocks base method.
func (m *MockSecretsConsumer) SecretAccess(arg0 *secrets.URI, arg1 names.Tag) (secrets.SecretRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretAccess", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretAccess indicates an expected call of SecretAccess.
func (mr *MockSecretsConsumerMockRecorder) SecretAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretAccess", reflect.TypeOf((*MockSecretsConsumer)(nil).SecretAccess), arg0, arg1)
}

// WatchConsumedSecretsChanges mocks base method.
func (m *MockSecretsConsumer) WatchConsumedSecretsChanges(arg0 string) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchConsumedSecretsChanges", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchConsumedSecretsChanges indicates an expected call of WatchConsumedSecretsChanges.
func (mr *MockSecretsConsumerMockRecorder) WatchConsumedSecretsChanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchConsumedSecretsChanges", reflect.TypeOf((*MockSecretsConsumer)(nil).WatchConsumedSecretsChanges), arg0)
}

// WatchObsolete mocks base method.
func (m *MockSecretsConsumer) WatchObsolete(arg0 string) (state.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchObsolete", arg0)
	ret0, _ := ret[0].(state.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchObsolete indicates an expected call of WatchObsolete.
func (mr *MockSecretsConsumerMockRecorder) WatchObsolete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchObsolete", reflect.TypeOf((*MockSecretsConsumer)(nil).WatchObsolete), arg0)
}
