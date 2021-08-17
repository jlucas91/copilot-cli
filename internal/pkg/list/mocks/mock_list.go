// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/list/list.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	config "github.com/aws/copilot-cli/internal/pkg/config"
	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetApplication mocks base method.
func (m *MockStore) GetApplication(appName string) (*config.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", appName)
	ret0, _ := ret[0].(*config.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockStoreMockRecorder) GetApplication(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockStore)(nil).GetApplication), appName)
}

// ListJobs mocks base method.
func (m *MockStore) ListJobs(appName string) ([]*config.Workload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJobs", appName)
	ret0, _ := ret[0].([]*config.Workload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJobs indicates an expected call of ListJobs.
func (mr *MockStoreMockRecorder) ListJobs(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockStore)(nil).ListJobs), appName)
}

// ListServices mocks base method.
func (m *MockStore) ListServices(appName string) ([]*config.Workload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", appName)
	ret0, _ := ret[0].([]*config.Workload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockStoreMockRecorder) ListServices(appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockStore)(nil).ListServices), appName)
}

// MockWorkspace is a mock of Workspace interface.
type MockWorkspace struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceMockRecorder
}

// MockWorkspaceMockRecorder is the mock recorder for MockWorkspace.
type MockWorkspaceMockRecorder struct {
	mock *MockWorkspace
}

// NewMockWorkspace creates a new mock instance.
func NewMockWorkspace(ctrl *gomock.Controller) *MockWorkspace {
	mock := &MockWorkspace{ctrl: ctrl}
	mock.recorder = &MockWorkspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspace) EXPECT() *MockWorkspaceMockRecorder {
	return m.recorder
}

// JobNames mocks base method.
func (m *MockWorkspace) JobNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JobNames indicates an expected call of JobNames.
func (mr *MockWorkspaceMockRecorder) JobNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobNames", reflect.TypeOf((*MockWorkspace)(nil).JobNames))
}

// ServiceNames mocks base method.
func (m *MockWorkspace) ServiceNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceNames indicates an expected call of ServiceNames.
func (mr *MockWorkspaceMockRecorder) ServiceNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceNames", reflect.TypeOf((*MockWorkspace)(nil).ServiceNames))
}