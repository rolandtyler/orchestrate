// Code generated by MockGen. DO NOT EDIT.
// Source: use-cases.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	core "github.com/ethereum/go-ethereum/signer/core"
	gomock "github.com/golang/mock/gomock"
	usecases "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/services/key-manager/key-manager/use-cases"
	reflect "reflect"
)

// MockUseCases is a mock of UseCases interface
type MockUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockUseCasesMockRecorder
}

// MockUseCasesMockRecorder is the mock recorder for MockUseCases
type MockUseCasesMockRecorder struct {
	mock *MockUseCases
}

// NewMockUseCases creates a new mock instance
func NewMockUseCases(ctrl *gomock.Controller) *MockUseCases {
	mock := &MockUseCases{ctrl: ctrl}
	mock.recorder = &MockUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCases) EXPECT() *MockUseCasesMockRecorder {
	return m.recorder
}

// SignTypedData mocks base method
func (m *MockUseCases) SignTypedData() usecases.SignTypedDataUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTypedData")
	ret0, _ := ret[0].(usecases.SignTypedDataUseCase)
	return ret0
}

// SignTypedData indicates an expected call of SignTypedData
func (mr *MockUseCasesMockRecorder) SignTypedData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTypedData", reflect.TypeOf((*MockUseCases)(nil).SignTypedData))
}

// MockSignTypedDataUseCase is a mock of SignTypedDataUseCase interface
type MockSignTypedDataUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSignTypedDataUseCaseMockRecorder
}

// MockSignTypedDataUseCaseMockRecorder is the mock recorder for MockSignTypedDataUseCase
type MockSignTypedDataUseCaseMockRecorder struct {
	mock *MockSignTypedDataUseCase
}

// NewMockSignTypedDataUseCase creates a new mock instance
func NewMockSignTypedDataUseCase(ctrl *gomock.Controller) *MockSignTypedDataUseCase {
	mock := &MockSignTypedDataUseCase{ctrl: ctrl}
	mock.recorder = &MockSignTypedDataUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSignTypedDataUseCase) EXPECT() *MockSignTypedDataUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSignTypedDataUseCase) Execute(ctx context.Context, address, namespace string, typedData *core.TypedData) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, address, namespace, typedData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockSignTypedDataUseCaseMockRecorder) Execute(ctx, address, namespace, typedData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSignTypedDataUseCase)(nil).Execute), ctx, address, namespace, typedData)
}
