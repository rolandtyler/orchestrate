// Code generated by MockGen. DO NOT EDIT.
// Source: services/multi-vault/keystore/keystore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
	types0 "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/ethereum/types"
	big "math/big"
	reflect "reflect"
)

// MockKeyStore is a mock of KeyStore interface
type MockKeyStore struct {
	ctrl     *gomock.Controller
	recorder *MockKeyStoreMockRecorder
}

// MockKeyStoreMockRecorder is the mock recorder for MockKeyStore
type MockKeyStoreMockRecorder struct {
	mock *MockKeyStore
}

// NewMockKeyStore creates a new mock instance
func NewMockKeyStore(ctrl *gomock.Controller) *MockKeyStore {
	mock := &MockKeyStore{ctrl: ctrl}
	mock.recorder = &MockKeyStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyStore) EXPECT() *MockKeyStoreMockRecorder {
	return m.recorder
}

// SignTx mocks base method
func (m *MockKeyStore) SignTx(ctx context.Context, chain *big.Int, a common.Address, tx *types.Transaction) ([]byte, *common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTx", ctx, chain, a, tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*common.Hash)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignTx indicates an expected call of SignTx
func (mr *MockKeyStoreMockRecorder) SignTx(ctx, chain, a, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTx", reflect.TypeOf((*MockKeyStore)(nil).SignTx), ctx, chain, a, tx)
}

// SignPrivateEEATx mocks base method
func (m *MockKeyStore) SignPrivateEEATx(ctx context.Context, chain *big.Int, a common.Address, tx *types.Transaction, privateArgs *types0.PrivateArgs) ([]byte, *common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignPrivateEEATx", ctx, chain, a, tx, privateArgs)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*common.Hash)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignPrivateEEATx indicates an expected call of SignPrivateEEATx
func (mr *MockKeyStoreMockRecorder) SignPrivateEEATx(ctx, chain, a, tx, privateArgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignPrivateEEATx", reflect.TypeOf((*MockKeyStore)(nil).SignPrivateEEATx), ctx, chain, a, tx, privateArgs)
}

// SignPrivateTesseraTx mocks base method
func (m *MockKeyStore) SignPrivateTesseraTx(ctx context.Context, chain *big.Int, a common.Address, tx *types.Transaction) ([]byte, *common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignPrivateTesseraTx", ctx, chain, a, tx)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*common.Hash)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignPrivateTesseraTx indicates an expected call of SignPrivateTesseraTx
func (mr *MockKeyStoreMockRecorder) SignPrivateTesseraTx(ctx, chain, a, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignPrivateTesseraTx", reflect.TypeOf((*MockKeyStore)(nil).SignPrivateTesseraTx), ctx, chain, a, tx)
}

// SignMsg mocks base method
func (m *MockKeyStore) SignMsg(ctx context.Context, a common.Address, msg string) ([]byte, *common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignMsg", ctx, a, msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*common.Hash)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignMsg indicates an expected call of SignMsg
func (mr *MockKeyStoreMockRecorder) SignMsg(ctx, a, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignMsg", reflect.TypeOf((*MockKeyStore)(nil).SignMsg), ctx, a, msg)
}

// SignRawHash mocks base method
func (m *MockKeyStore) SignRawHash(a common.Address, hash []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignRawHash", a, hash)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignRawHash indicates an expected call of SignRawHash
func (mr *MockKeyStoreMockRecorder) SignRawHash(a, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignRawHash", reflect.TypeOf((*MockKeyStore)(nil).SignRawHash), a, hash)
}

// GenerateAccount mocks base method
func (m *MockKeyStore) GenerateAccount(ctx context.Context) (*common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccount", ctx)
	ret0, _ := ret[0].(*common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccount indicates an expected call of GenerateAccount
func (mr *MockKeyStoreMockRecorder) GenerateAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccount", reflect.TypeOf((*MockKeyStore)(nil).GenerateAccount), ctx)
}

// ImportPrivateKey mocks base method
func (m *MockKeyStore) ImportPrivateKey(ctx context.Context, priv string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportPrivateKey", ctx, priv)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportPrivateKey indicates an expected call of ImportPrivateKey
func (mr *MockKeyStoreMockRecorder) ImportPrivateKey(ctx, priv interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportPrivateKey", reflect.TypeOf((*MockKeyStore)(nil).ImportPrivateKey), ctx, priv)
}
