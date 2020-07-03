// Code generated by MockGen. DO NOT EDIT.
// Source: get_method_signatures.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	abi "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/abi"
	reflect "reflect"
)

// MockGetMethodSignaturesUseCase is a mock of GetMethodSignaturesUseCase interface
type MockGetMethodSignaturesUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetMethodSignaturesUseCaseMockRecorder
}

// MockGetMethodSignaturesUseCaseMockRecorder is the mock recorder for MockGetMethodSignaturesUseCase
type MockGetMethodSignaturesUseCaseMockRecorder struct {
	mock *MockGetMethodSignaturesUseCase
}

// NewMockGetMethodSignaturesUseCase creates a new mock instance
func NewMockGetMethodSignaturesUseCase(ctrl *gomock.Controller) *MockGetMethodSignaturesUseCase {
	mock := &MockGetMethodSignaturesUseCase{ctrl: ctrl}
	mock.recorder = &MockGetMethodSignaturesUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetMethodSignaturesUseCase) EXPECT() *MockGetMethodSignaturesUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetMethodSignaturesUseCase) Execute(ctx context.Context, contract *abi.ContractId, methodName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, contract, methodName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetMethodSignaturesUseCaseMockRecorder) Execute(ctx, contract, methodName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetMethodSignaturesUseCase)(nil).Execute), ctx, contract, methodName)
}