// Code generated by MockGen. DO NOT EDIT.
// Source: tviso-scrapper/pkg/tviso (interfaces: ReadRepository,WriteRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	tviso "tviso-scrapper/pkg/tviso"

	gomock "github.com/golang/mock/gomock"
)

// MockReadRepository is a mock of ReadRepository interface
type MockReadRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReadRepositoryMockRecorder
}

// MockReadRepositoryMockRecorder is the mock recorder for MockReadRepository
type MockReadRepositoryMockRecorder struct {
	mock *MockReadRepository
}

// NewMockReadRepository creates a new mock instance
func NewMockReadRepository(ctrl *gomock.Controller) *MockReadRepository {
	mock := &MockReadRepository{ctrl: ctrl}
	mock.recorder = &MockReadRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadRepository) EXPECT() *MockReadRepositoryMockRecorder {
	return m.recorder
}

// GetMediaInfo mocks base method
func (m *MockReadRepository) GetMediaInfo(arg0 *tviso.Media) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMediaInfo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetMediaInfo indicates an expected call of GetMediaInfo
func (mr *MockReadRepositoryMockRecorder) GetMediaInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediaInfo", reflect.TypeOf((*MockReadRepository)(nil).GetMediaInfo), arg0)
}

// GetUserCollection mocks base method
func (m *MockReadRepository) GetUserCollection() ([]tviso.Media, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserCollection")
	ret0, _ := ret[0].([]tviso.Media)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserCollection indicates an expected call of GetUserCollection
func (mr *MockReadRepositoryMockRecorder) GetUserCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserCollection", reflect.TypeOf((*MockReadRepository)(nil).GetUserCollection))
}

// MockWriteRepository is a mock of WriteRepository interface
type MockWriteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWriteRepositoryMockRecorder
}

// MockWriteRepositoryMockRecorder is the mock recorder for MockWriteRepository
type MockWriteRepositoryMockRecorder struct {
	mock *MockWriteRepository
}

// NewMockWriteRepository creates a new mock instance
func NewMockWriteRepository(ctrl *gomock.Controller) *MockWriteRepository {
	mock := &MockWriteRepository{ctrl: ctrl}
	mock.recorder = &MockWriteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriteRepository) EXPECT() *MockWriteRepositoryMockRecorder {
	return m.recorder
}

// StoreCollection mocks base method
func (m *MockWriteRepository) StoreCollection(arg0 []tviso.Media) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreCollection", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreCollection indicates an expected call of StoreCollection
func (mr *MockWriteRepositoryMockRecorder) StoreCollection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreCollection", reflect.TypeOf((*MockWriteRepository)(nil).StoreCollection), arg0)
}
