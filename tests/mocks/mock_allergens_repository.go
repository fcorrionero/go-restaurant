// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fcorrionero/go-restaurant/src/domain (interfaces: AllergensRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/fcorrionero/go-restaurant/src/domain"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockAllergensRepository is a mock of AllergensRepository interface.
type MockAllergensRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAllergensRepositoryMockRecorder
}

// MockAllergensRepositoryMockRecorder is the mock recorder for MockAllergensRepository.
type MockAllergensRepositoryMockRecorder struct {
	mock *MockAllergensRepository
}

// NewMockAllergensRepository creates a new mock instance.
func NewMockAllergensRepository(ctrl *gomock.Controller) *MockAllergensRepository {
	mock := &MockAllergensRepository{ctrl: ctrl}
	mock.recorder = &MockAllergensRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllergensRepository) EXPECT() *MockAllergensRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockAllergensRepository) FindAll() []*domain.Allergen {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]*domain.Allergen)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAllergensRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAllergensRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockAllergensRepository) FindById(arg0 uuid.UUID) *domain.Allergen {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*domain.Allergen)
	return ret0
}

// FindById indicates an expected call of FindById.
func (mr *MockAllergensRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAllergensRepository)(nil).FindById), arg0)
}

// FindByName mocks base method.
func (m *MockAllergensRepository) FindByName(arg0 string) *domain.Allergen {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0)
	ret0, _ := ret[0].(*domain.Allergen)
	return ret0
}

// FindByName indicates an expected call of FindByName.
func (mr *MockAllergensRepositoryMockRecorder) FindByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockAllergensRepository)(nil).FindByName), arg0)
}

// Save mocks base method.
func (m *MockAllergensRepository) Save(arg0 *domain.Allergen) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockAllergensRepositoryMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAllergensRepository)(nil).Save), arg0)
}
