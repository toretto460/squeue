// Code generated by MockGen. DO NOT EDIT.
// Source: driver/driver.go

// Package squeue_test is a generated GoMock package.
package squeue_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/simodima/squeue/driver"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *MockDriver) Ack(queue, messageID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", queue, messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *MockDriverMockRecorder) Ack(queue, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockDriver)(nil).Ack), queue, messageID)
}

// Consume mocks base method.
func (m *MockDriver) Consume(queue string, opts ...func(any)) (*driver.ConsumerController, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{queue}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Consume", varargs...)
	ret0, _ := ret[0].(*driver.ConsumerController)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockDriverMockRecorder) Consume(queue interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{queue}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockDriver)(nil).Consume), varargs...)
}

// Enqueue mocks base method.
func (m *MockDriver) Enqueue(queue string, data []byte, opts ...func(any)) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{queue, data}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Enqueue", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockDriverMockRecorder) Enqueue(queue, data interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{queue, data}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockDriver)(nil).Enqueue), varargs...)
}
