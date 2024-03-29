// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yurifrl/logapi (interfaces: File)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFile is a mock of File interface
type MockFile struct {
	ctrl     *gomock.Controller
	recorder *MockFileMockRecorder
}

// MockFileMockRecorder is the mock recorder for MockFile
type MockFileMockRecorder struct {
	mock *MockFile
}

// NewMockFile creates a new mock instance
func NewMockFile(ctrl *gomock.Controller) *MockFile {
	mock := &MockFile{ctrl: ctrl}
	mock.recorder = &MockFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFile) EXPECT() *MockFileMockRecorder {
	return m.recorder
}

// Sync mocks base method
func (m *MockFile) Sync(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockFileMockRecorder) Sync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockFile)(nil).Sync), arg0)
}
