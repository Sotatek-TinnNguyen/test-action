// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kaiachain/kaia/datasync/chaindatafetcher/kas (interfaces: BlockchainAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/kaiachain/kaia/api"
	common "github.com/kaiachain/kaia/common"
	hexutil "github.com/kaiachain/kaia/common/hexutil"
	rpc "github.com/kaiachain/kaia/networks/rpc"
)

// MockBlockchainAPI is a mock of BlockchainAPI interface
type MockBlockchainAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainAPIMockRecorder
}

// MockBlockchainAPIMockRecorder is the mock recorder for MockBlockchainAPI
type MockBlockchainAPIMockRecorder struct {
	mock *MockBlockchainAPI
}

// NewMockBlockchainAPI creates a new mock instance
func NewMockBlockchainAPI(ctrl *gomock.Controller) *MockBlockchainAPI {
	mock := &MockBlockchainAPI{ctrl: ctrl}
	mock.recorder = &MockBlockchainAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchainAPI) EXPECT() *MockBlockchainAPIMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockBlockchainAPI) Call(arg0 context.Context, arg1 api.CallArgs, arg2 rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1, arg2)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockBlockchainAPIMockRecorder) Call(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockBlockchainAPI)(nil).Call), arg0, arg1, arg2)
}

// GetCode mocks base method
func (m *MockBlockchainAPI) GetCode(arg0 context.Context, arg1 common.Address, arg2 rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", arg0, arg1, arg2)
	ret0, _ := ret[0].(hexutil.Bytes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCode indicates an expected call of GetCode
func (mr *MockBlockchainAPIMockRecorder) GetCode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockBlockchainAPI)(nil).GetCode), arg0, arg1, arg2)
}