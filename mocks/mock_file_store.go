// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yurifrl/logapi (interfaces: FileStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFileStore is a mock of FileStore interface
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// Bump mocks base method
func (m *MockFileStore) Bump(arg0 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bump", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bump indicates an expected call of Bump
func (mr *MockFileStoreMockRecorder) Bump(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bump", reflect.TypeOf((*MockFileStore)(nil).Bump), arg0)
}

// GetAll mocks base method
func (m *MockFileStore) GetAll() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockFileStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockFileStore)(nil).GetAll))
}
