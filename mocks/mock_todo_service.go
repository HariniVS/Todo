// Code generated by MockGen. DO NOT EDIT.
// Source: service/todo_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/todo/model"
	reflect "reflect"
)

// MockTodoService is a mock of TodoService interface
type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

// MockTodoServiceMockRecorder is the mock recorder for MockTodoService
type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

// NewMockTodoService creates a new mock instance
func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

// GetTodo mocks base method
func (m *MockTodoService) GetTodo(id int) (model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodo", id)
	ret0, _ := ret[0].(model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo
func (mr *MockTodoServiceMockRecorder) GetTodo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockTodoService)(nil).GetTodo), id)
}

// GetAllTodos mocks base method
func (m *MockTodoService) GetAllTodos() ([]model.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodos")
	ret0, _ := ret[0].([]model.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTodos indicates an expected call of GetAllTodos
func (mr *MockTodoServiceMockRecorder) GetAllTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodos", reflect.TypeOf((*MockTodoService)(nil).GetAllTodos))
}
