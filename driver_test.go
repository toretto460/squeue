// Code generated by MockGen. DO NOT EDIT.
// Source: driver/driver.go

// Package squeue_test is a generated GoMock package.
package squeue_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	driver "github.com/toretto460/squeue/driver"
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
func (m *MockDriver) Consume(ctx context.Context, topic string) (chan driver.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", ctx, topic)
	ret0, _ := ret[0].(chan driver.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockDriverMockRecorder) Consume(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockDriver)(nil).Consume), ctx, topic)
}

// Enqueue mocks base method.
func (m *MockDriver) Enqueue(queue string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enqueue", queue, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enqueue indicates an expected call of Enqueue.
func (mr *MockDriverMockRecorder) Enqueue(queue, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockDriver)(nil).Enqueue), queue, data)
}
