// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/instancemutater (interfaces: ModelCache,ModelCacheMachine)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	instancemutater "github.com/juju/juju/apiserver/facades/agent/instancemutater"
	cache "github.com/juju/juju/core/cache"
	reflect "reflect"
)

// MockModelCache is a mock of ModelCache interface
type MockModelCache struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheMockRecorder
}

// MockModelCacheMockRecorder is the mock recorder for MockModelCache
type MockModelCacheMockRecorder struct {
	mock *MockModelCache
}

// NewMockModelCache creates a new mock instance
func NewMockModelCache(ctrl *gomock.Controller) *MockModelCache {
	mock := &MockModelCache{ctrl: ctrl}
	mock.recorder = &MockModelCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCache) EXPECT() *MockModelCacheMockRecorder {
	return m.recorder
}

// Machine mocks base method
func (m *MockModelCache) Machine(arg0 string) (instancemutater.ModelCacheMachine, error) {
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(instancemutater.ModelCacheMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine
func (mr *MockModelCacheMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockModelCache)(nil).Machine), arg0)
}

// WatchMachines mocks base method
func (m *MockModelCache) WatchMachines() cache.StringsWatcher {
	ret := m.ctrl.Call(m, "WatchMachines")
	ret0, _ := ret[0].(cache.StringsWatcher)
	return ret0
}

// WatchMachines indicates an expected call of WatchMachines
func (mr *MockModelCacheMockRecorder) WatchMachines() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchMachines", reflect.TypeOf((*MockModelCache)(nil).WatchMachines))
}

// MockModelCacheMachine is a mock of ModelCacheMachine interface
type MockModelCacheMachine struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheMachineMockRecorder
}

// MockModelCacheMachineMockRecorder is the mock recorder for MockModelCacheMachine
type MockModelCacheMachineMockRecorder struct {
	mock *MockModelCacheMachine
}

// NewMockModelCacheMachine creates a new mock instance
func NewMockModelCacheMachine(ctrl *gomock.Controller) *MockModelCacheMachine {
	mock := &MockModelCacheMachine{ctrl: ctrl}
	mock.recorder = &MockModelCacheMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCacheMachine) EXPECT() *MockModelCacheMachineMockRecorder {
	return m.recorder
}

// WatchApplicationLXDProfiles mocks base method
func (m *MockModelCacheMachine) WatchApplicationLXDProfiles() cache.NotifyWatcher {
	ret := m.ctrl.Call(m, "WatchApplicationLXDProfiles")
	ret0, _ := ret[0].(cache.NotifyWatcher)
	return ret0
}

// WatchApplicationLXDProfiles indicates an expected call of WatchApplicationLXDProfiles
func (mr *MockModelCacheMachineMockRecorder) WatchApplicationLXDProfiles() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplicationLXDProfiles", reflect.TypeOf((*MockModelCacheMachine)(nil).WatchApplicationLXDProfiles))
}
