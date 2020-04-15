// Code generated by MockGen. DO NOT EDIT.
// Source: get_methods.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	common "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/pkg/types/common"
	reflect "reflect"
)

// MockGetMethodsUseCase is a mock of GetMethodsUseCase interface.
type MockGetMethodsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetMethodsUseCaseMockRecorder
}

// MockGetMethodsUseCaseMockRecorder is the mock recorder for MockGetMethodsUseCase.
type MockGetMethodsUseCaseMockRecorder struct {
	mock *MockGetMethodsUseCase
}

// NewMockGetMethodsUseCase creates a new mock instance.
func NewMockGetMethodsUseCase(ctrl *gomock.Controller) *MockGetMethodsUseCase {
	mock := &MockGetMethodsUseCase{ctrl: ctrl}
	mock.recorder = &MockGetMethodsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetMethodsUseCase) EXPECT() *MockGetMethodsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetMethodsUseCase) Execute(ctx context.Context, account *common.AccountInstance, selector []byte) (string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, account, selector)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Execute indicates an expected call of Execute.
func (mr *MockGetMethodsUseCaseMockRecorder) Execute(ctx, account, selector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetMethodsUseCase)(nil).Execute), ctx, account, selector)
}
