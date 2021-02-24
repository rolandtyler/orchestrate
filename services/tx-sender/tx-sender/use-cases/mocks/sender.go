// Code generated by MockGen. DO NOT EDIT.
// Source: sender.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "github.com/ConsenSys/orchestrate/pkg/types/entities"
	reflect "reflect"
)

// MockSendETHRawTxUseCase is a mock of SendETHRawTxUseCase interface
type MockSendETHRawTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendETHRawTxUseCaseMockRecorder
}

// MockSendETHRawTxUseCaseMockRecorder is the mock recorder for MockSendETHRawTxUseCase
type MockSendETHRawTxUseCaseMockRecorder struct {
	mock *MockSendETHRawTxUseCase
}

// NewMockSendETHRawTxUseCase creates a new mock instance
func NewMockSendETHRawTxUseCase(ctrl *gomock.Controller) *MockSendETHRawTxUseCase {
	mock := &MockSendETHRawTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendETHRawTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendETHRawTxUseCase) EXPECT() *MockSendETHRawTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendETHRawTxUseCase) Execute(ctx context.Context, job *entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSendETHRawTxUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendETHRawTxUseCase)(nil).Execute), ctx, job)
}

// MockSendETHTxUseCase is a mock of SendETHTxUseCase interface
type MockSendETHTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendETHTxUseCaseMockRecorder
}

// MockSendETHTxUseCaseMockRecorder is the mock recorder for MockSendETHTxUseCase
type MockSendETHTxUseCaseMockRecorder struct {
	mock *MockSendETHTxUseCase
}

// NewMockSendETHTxUseCase creates a new mock instance
func NewMockSendETHTxUseCase(ctrl *gomock.Controller) *MockSendETHTxUseCase {
	mock := &MockSendETHTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendETHTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendETHTxUseCase) EXPECT() *MockSendETHTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendETHTxUseCase) Execute(ctx context.Context, job *entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSendETHTxUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendETHTxUseCase)(nil).Execute), ctx, job)
}

// MockSendEEAPrivateTxUseCase is a mock of SendEEAPrivateTxUseCase interface
type MockSendEEAPrivateTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendEEAPrivateTxUseCaseMockRecorder
}

// MockSendEEAPrivateTxUseCaseMockRecorder is the mock recorder for MockSendEEAPrivateTxUseCase
type MockSendEEAPrivateTxUseCaseMockRecorder struct {
	mock *MockSendEEAPrivateTxUseCase
}

// NewMockSendEEAPrivateTxUseCase creates a new mock instance
func NewMockSendEEAPrivateTxUseCase(ctrl *gomock.Controller) *MockSendEEAPrivateTxUseCase {
	mock := &MockSendEEAPrivateTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendEEAPrivateTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendEEAPrivateTxUseCase) EXPECT() *MockSendEEAPrivateTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendEEAPrivateTxUseCase) Execute(ctx context.Context, job *entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSendEEAPrivateTxUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendEEAPrivateTxUseCase)(nil).Execute), ctx, job)
}

// MockSendTesseraPrivateTxUseCase is a mock of SendTesseraPrivateTxUseCase interface
type MockSendTesseraPrivateTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendTesseraPrivateTxUseCaseMockRecorder
}

// MockSendTesseraPrivateTxUseCaseMockRecorder is the mock recorder for MockSendTesseraPrivateTxUseCase
type MockSendTesseraPrivateTxUseCaseMockRecorder struct {
	mock *MockSendTesseraPrivateTxUseCase
}

// NewMockSendTesseraPrivateTxUseCase creates a new mock instance
func NewMockSendTesseraPrivateTxUseCase(ctrl *gomock.Controller) *MockSendTesseraPrivateTxUseCase {
	mock := &MockSendTesseraPrivateTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendTesseraPrivateTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendTesseraPrivateTxUseCase) EXPECT() *MockSendTesseraPrivateTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendTesseraPrivateTxUseCase) Execute(ctx context.Context, job *entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSendTesseraPrivateTxUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendTesseraPrivateTxUseCase)(nil).Execute), ctx, job)
}

// MockSendTesseraMarkingTxUseCase is a mock of SendTesseraMarkingTxUseCase interface
type MockSendTesseraMarkingTxUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSendTesseraMarkingTxUseCaseMockRecorder
}

// MockSendTesseraMarkingTxUseCaseMockRecorder is the mock recorder for MockSendTesseraMarkingTxUseCase
type MockSendTesseraMarkingTxUseCaseMockRecorder struct {
	mock *MockSendTesseraMarkingTxUseCase
}

// NewMockSendTesseraMarkingTxUseCase creates a new mock instance
func NewMockSendTesseraMarkingTxUseCase(ctrl *gomock.Controller) *MockSendTesseraMarkingTxUseCase {
	mock := &MockSendTesseraMarkingTxUseCase{ctrl: ctrl}
	mock.recorder = &MockSendTesseraMarkingTxUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSendTesseraMarkingTxUseCase) EXPECT() *MockSendTesseraMarkingTxUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSendTesseraMarkingTxUseCase) Execute(ctx context.Context, job *entities.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, job)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSendTesseraMarkingTxUseCaseMockRecorder) Execute(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSendTesseraMarkingTxUseCase)(nil).Execute), ctx, job)
}
