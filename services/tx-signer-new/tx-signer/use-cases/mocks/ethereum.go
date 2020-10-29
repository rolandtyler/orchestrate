// Code generated by MockGen. DO NOT EDIT.
// Source: ethereum.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	entities "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/entities"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/tx-signer-new/tx-signer/use-cases"
	reflect "reflect"
)

// MockEthereumUseCases is a mock of EthereumUseCases interface
type MockEthereumUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockEthereumUseCasesMockRecorder
}

// MockEthereumUseCasesMockRecorder is the mock recorder for MockEthereumUseCases
type MockEthereumUseCasesMockRecorder struct {
	mock *MockEthereumUseCases
}

// NewMockEthereumUseCases creates a new mock instance
func NewMockEthereumUseCases(ctrl *gomock.Controller) *MockEthereumUseCases {
	mock := &MockEthereumUseCases{ctrl: ctrl}
	mock.recorder = &MockEthereumUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEthereumUseCases) EXPECT() *MockEthereumUseCasesMockRecorder {
	return m.recorder
}

// SignTransaction mocks base method
func (m *MockEthereumUseCases) SignTransaction() usecases.SignTransactionUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTransaction")
	ret0, _ := ret[0].(usecases.SignTransactionUseCase)
	return ret0
}

// SignTransaction indicates an expected call of SignTransaction
func (mr *MockEthereumUseCasesMockRecorder) SignTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTransaction", reflect.TypeOf((*MockEthereumUseCases)(nil).SignTransaction))
}

// SendEnvelope mocks base method
func (m *MockEthereumUseCases) SendEnvelope() usecases.SendEnvelopeUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEnvelope")
	ret0, _ := ret[0].(usecases.SendEnvelopeUseCase)
	return ret0
}

// SendEnvelope indicates an expected call of SendEnvelope
func (mr *MockEthereumUseCasesMockRecorder) SendEnvelope() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEnvelope", reflect.TypeOf((*MockEthereumUseCases)(nil).SendEnvelope))
}

// MockSignTransactionUseCase is a mock of SignTransactionUseCase interface
type MockSignTransactionUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSignTransactionUseCaseMockRecorder
}

// MockSignTransactionUseCaseMockRecorder is the mock recorder for MockSignTransactionUseCase
type MockSignTransactionUseCaseMockRecorder struct {
	mock *MockSignTransactionUseCase
}

// NewMockSignTransactionUseCase creates a new mock instance
func NewMockSignTransactionUseCase(ctrl *gomock.Controller) *MockSignTransactionUseCase {
	mock := &MockSignTransactionUseCase{ctrl: ctrl}
	mock.recorder = &MockSignTransactionUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSignTransactionUseCase) EXPECT() *MockSignTransactionUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSignTransactionUseCase) Execute(ctx context.Context, job *entities.Job) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Execute indicates an expected call of Execute
func (mr *MockSignTransactionUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSignTransactionUseCase)(nil).Execute), ctx, job)
}

// MockSendEnvelopeUseCase is a mock of SendEnvelopeUseCase interface
type MockSendEnvelopeUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendEnvelopeUseCaseMockRecorder
}

// MockSendEnvelopeUseCaseMockRecorder is the mock recorder for MockSendEnvelopeUseCase
type MockSendEnvelopeUseCaseMockRecorder struct {
	mock *MockSendEnvelopeUseCase
}

// NewMockSendEnvelopeUseCase creates a new mock instance
func NewMockSendEnvelopeUseCase(ctrl *gomock.Controller) *MockSendEnvelopeUseCase {
	mock := &MockSendEnvelopeUseCase{ctrl: ctrl}
	mock.recorder = &MockSendEnvelopeUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendEnvelopeUseCase) EXPECT() *MockSendEnvelopeUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendEnvelopeUseCase) Execute(ctx context.Context, protoMessage proto.Message, topic, partitionKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, protoMessage, topic, partitionKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSendEnvelopeUseCaseMockRecorder) Execute(ctx, protoMessage, topic, partitionKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendEnvelopeUseCase)(nil).Execute), ctx, protoMessage, topic, partitionKey)
}
