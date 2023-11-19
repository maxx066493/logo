// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/signer (interfaces: Signer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/signer.go -package=mocks github.com/Layr-Labs/eigensdk-go/signer Signer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockSigner is a mock of Signer interface.
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner.
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance.
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// GetTxOpts mocks base method.
func (m *MockSigner) GetTxOpts() *bind.TransactOpts {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxOpts")
	ret0, _ := ret[0].(*bind.TransactOpts)
	return ret0
}

// GetTxOpts indicates an expected call of GetTxOpts.
func (mr *MockSignerMockRecorder) GetTxOpts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxOpts", reflect.TypeOf((*MockSigner)(nil).GetTxOpts))
}

// SendToExternal mocks base method.
func (m *MockSigner) SendToExternal(arg0 context.Context, arg1 *types.Transaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendToExternal", arg0, arg1)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendToExternal indicates an expected call of SendToExternal.
func (mr *MockSignerMockRecorder) SendToExternal(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendToExternal", reflect.TypeOf((*MockSigner)(nil).SendToExternal), arg0, arg1)
}
