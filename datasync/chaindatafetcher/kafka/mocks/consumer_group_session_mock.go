// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kaiachain/kaia/datasync/chaindatafetcher/kafka (interfaces: ConsumerGroupSession)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
)

// MockConsumerGroupSession is a mock of ConsumerGroupSession interface
type MockConsumerGroupSession struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerGroupSessionMockRecorder
}

// MockConsumerGroupSessionMockRecorder is the mock recorder for MockConsumerGroupSession
type MockConsumerGroupSessionMockRecorder struct {
	mock *MockConsumerGroupSession
}

// NewMockConsumerGroupSession creates a new mock instance
func NewMockConsumerGroupSession(ctrl *gomock.Controller) *MockConsumerGroupSession {
	mock := &MockConsumerGroupSession{ctrl: ctrl}
	mock.recorder = &MockConsumerGroupSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsumerGroupSession) EXPECT() *MockConsumerGroupSessionMockRecorder {
	return m.recorder
}

// MarkMessage mocks base method
func (m *MockConsumerGroupSession) MarkMessage(arg0 *sarama.ConsumerMessage, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkMessage", arg0, arg1)
}

// MarkMessage indicates an expected call of MarkMessage
func (mr *MockConsumerGroupSessionMockRecorder) MarkMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkMessage", reflect.TypeOf((*MockConsumerGroupSession)(nil).MarkMessage), arg0, arg1)
}

// MarkOffset mocks base method
func (m *MockConsumerGroupSession) MarkOffset(arg0 string, arg1 int32, arg2 int64, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkOffset", arg0, arg1, arg2, arg3)
}

// MarkOffset indicates an expected call of MarkOffset
func (mr *MockConsumerGroupSessionMockRecorder) MarkOffset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkOffset", reflect.TypeOf((*MockConsumerGroupSession)(nil).MarkOffset), arg0, arg1, arg2, arg3)
}