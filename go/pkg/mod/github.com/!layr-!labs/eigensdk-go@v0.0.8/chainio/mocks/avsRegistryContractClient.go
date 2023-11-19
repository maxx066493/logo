// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/clients (interfaces: AVSRegistryContractsClient)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsRegistryContractClient.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients AVSRegistryContractsClient
//
// Package mocks is a generated GoMock package.
package mocks

import (
	big "math/big"
	reflect "reflect"

	contractBLSOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	contractBLSPubkeyRegistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	contractBLSRegistryCoordinatorWithIndices "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockAVSRegistryContractsClient is a mock of AVSRegistryContractsClient interface.
type MockAVSRegistryContractsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAVSRegistryContractsClientMockRecorder
}

// MockAVSRegistryContractsClientMockRecorder is the mock recorder for MockAVSRegistryContractsClient.
type MockAVSRegistryContractsClientMockRecorder struct {
	mock *MockAVSRegistryContractsClient
}

// NewMockAVSRegistryContractsClient creates a new mock instance.
func NewMockAVSRegistryContractsClient(ctrl *gomock.Controller) *MockAVSRegistryContractsClient {
	mock := &MockAVSRegistryContractsClient{ctrl: ctrl}
	mock.recorder = &MockAVSRegistryContractsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAVSRegistryContractsClient) EXPECT() *MockAVSRegistryContractsClientMockRecorder {
	return m.recorder
}

// DeregisterOperator mocks base method.
func (m *MockAVSRegistryContractsClient) DeregisterOperator(arg0 *bind.TransactOpts, arg1 common.Address, arg2 []byte, arg3 contractBLSPubkeyRegistry.BN254G1Point) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterOperator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeregisterOperator indicates an expected call of DeregisterOperator.
func (mr *MockAVSRegistryContractsClientMockRecorder) DeregisterOperator(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterOperator", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).DeregisterOperator), arg0, arg1, arg2, arg3)
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockAVSRegistryContractsClient) GetCheckSignaturesIndices(arg0 *bind.CallOpts, arg1 uint32, arg2 []byte, arg3 [][32]byte) (contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockAVSRegistryContractsClientMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetCurrentOperatorStakeForQuorum mocks base method.
func (m *MockAVSRegistryContractsClient) GetCurrentOperatorStakeForQuorum(arg0 *bind.CallOpts, arg1 [32]byte, arg2 byte) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentOperatorStakeForQuorum", arg0, arg1, arg2)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentOperatorStakeForQuorum indicates an expected call of GetCurrentOperatorStakeForQuorum.
func (mr *MockAVSRegistryContractsClientMockRecorder) GetCurrentOperatorStakeForQuorum(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentOperatorStakeForQuorum", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).GetCurrentOperatorStakeForQuorum), arg0, arg1, arg2)
}

// GetOperatorId mocks base method.
func (m *MockAVSRegistryContractsClient) GetOperatorId(arg0 *bind.CallOpts, arg1 common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorId", arg0, arg1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorId indicates an expected call of GetOperatorId.
func (mr *MockAVSRegistryContractsClientMockRecorder) GetOperatorId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorId", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).GetOperatorId), arg0, arg1)
}

// GetOperatorQuorumsAtCurrentBlock mocks base method.
func (m *MockAVSRegistryContractsClient) GetOperatorQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 [32]byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorQuorumsAtCurrentBlock indicates an expected call of GetOperatorQuorumsAtCurrentBlock.
func (mr *MockAVSRegistryContractsClientMockRecorder) GetOperatorQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorQuorumsAtCurrentBlock", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).GetOperatorQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsAtBlock mocks base method.
func (m *MockAVSRegistryContractsClient) GetOperatorsStakeInQuorumsAtBlock(arg0 *bind.CallOpts, arg1 []byte, arg2 uint32) ([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtBlock indicates an expected call of GetOperatorsStakeInQuorumsAtBlock.
func (mr *MockAVSRegistryContractsClientMockRecorder) GetOperatorsStakeInQuorumsAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtBlock", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).GetOperatorsStakeInQuorumsAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock mocks base method.
func (m *MockAVSRegistryContractsClient) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0 *bind.CallOpts, arg1 [32]byte, arg2 uint32) ([]byte, [][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtBlock.
func (mr *MockAVSRegistryContractsClientMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).GetOperatorsStakeInQuorumsOfOperatorAtBlock), arg0, arg1, arg2)
}

// RegisterOperatorWithCoordinator mocks base method.
func (m *MockAVSRegistryContractsClient) RegisterOperatorWithCoordinator(arg0 *bind.TransactOpts, arg1 []byte, arg2 contractBLSRegistryCoordinatorWithIndices.BN254G1Point, arg3 string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterOperatorWithCoordinator", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterOperatorWithCoordinator indicates an expected call of RegisterOperatorWithCoordinator.
func (mr *MockAVSRegistryContractsClientMockRecorder) RegisterOperatorWithCoordinator(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOperatorWithCoordinator", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).RegisterOperatorWithCoordinator), arg0, arg1, arg2, arg3)
}

// UpdateStakes mocks base method.
func (m *MockAVSRegistryContractsClient) UpdateStakes(arg0 *bind.TransactOpts, arg1 []common.Address) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStakes", arg0, arg1)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStakes indicates an expected call of UpdateStakes.
func (mr *MockAVSRegistryContractsClientMockRecorder) UpdateStakes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStakes", reflect.TypeOf((*MockAVSRegistryContractsClient)(nil).UpdateStakes), arg0, arg1)
}