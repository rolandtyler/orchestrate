// Code generated by MockGen. DO NOT EDIT.
// Source: contracts.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entities "github.com/consensys/orchestrate/pkg/types/entities"
	usecases "github.com/consensys/orchestrate/services/api/business/use-cases"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockContractUseCases is a mock of ContractUseCases interface
type MockContractUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockContractUseCasesMockRecorder
}

// MockContractUseCasesMockRecorder is the mock recorder for MockContractUseCases
type MockContractUseCasesMockRecorder struct {
	mock *MockContractUseCases
}

// NewMockContractUseCases creates a new mock instance
func NewMockContractUseCases(ctrl *gomock.Controller) *MockContractUseCases {
	mock := &MockContractUseCases{ctrl: ctrl}
	mock.recorder = &MockContractUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContractUseCases) EXPECT() *MockContractUseCasesMockRecorder {
	return m.recorder
}

// GetContractsCatalog mocks base method
func (m *MockContractUseCases) GetContractsCatalog() usecases.GetContractsCatalogUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractsCatalog")
	ret0, _ := ret[0].(usecases.GetContractsCatalogUseCase)
	return ret0
}

// GetContractsCatalog indicates an expected call of GetContractsCatalog
func (mr *MockContractUseCasesMockRecorder) GetContractsCatalog() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractsCatalog", reflect.TypeOf((*MockContractUseCases)(nil).GetContractsCatalog))
}

// GetContract mocks base method
func (m *MockContractUseCases) GetContract() usecases.GetContractUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract")
	ret0, _ := ret[0].(usecases.GetContractUseCase)
	return ret0
}

// GetContract indicates an expected call of GetContract
func (mr *MockContractUseCasesMockRecorder) GetContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockContractUseCases)(nil).GetContract))
}

// GetContractEvents mocks base method
func (m *MockContractUseCases) GetContractEvents() usecases.GetContractEventsUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractEvents")
	ret0, _ := ret[0].(usecases.GetContractEventsUseCase)
	return ret0
}

// GetContractEvents indicates an expected call of GetContractEvents
func (mr *MockContractUseCasesMockRecorder) GetContractEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractEvents", reflect.TypeOf((*MockContractUseCases)(nil).GetContractEvents))
}

// GetContractMethodSignatures mocks base method
func (m *MockContractUseCases) GetContractMethodSignatures() usecases.GetContractMethodSignaturesUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractMethodSignatures")
	ret0, _ := ret[0].(usecases.GetContractMethodSignaturesUseCase)
	return ret0
}

// GetContractMethodSignatures indicates an expected call of GetContractMethodSignatures
func (mr *MockContractUseCasesMockRecorder) GetContractMethodSignatures() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractMethodSignatures", reflect.TypeOf((*MockContractUseCases)(nil).GetContractMethodSignatures))
}

// GetContractMethods mocks base method
func (m *MockContractUseCases) GetContractMethods() usecases.GetContractMethodsUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractMethods")
	ret0, _ := ret[0].(usecases.GetContractMethodsUseCase)
	return ret0
}

// GetContractMethods indicates an expected call of GetContractMethods
func (mr *MockContractUseCasesMockRecorder) GetContractMethods() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractMethods", reflect.TypeOf((*MockContractUseCases)(nil).GetContractMethods))
}

// GetContractTags mocks base method
func (m *MockContractUseCases) GetContractTags() usecases.GetContractTagsUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractTags")
	ret0, _ := ret[0].(usecases.GetContractTagsUseCase)
	return ret0
}

// GetContractTags indicates an expected call of GetContractTags
func (mr *MockContractUseCasesMockRecorder) GetContractTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractTags", reflect.TypeOf((*MockContractUseCases)(nil).GetContractTags))
}

// SetContractCodeHash mocks base method
func (m *MockContractUseCases) SetContractCodeHash() usecases.SetContractCodeHashUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetContractCodeHash")
	ret0, _ := ret[0].(usecases.SetContractCodeHashUseCase)
	return ret0
}

// SetContractCodeHash indicates an expected call of SetContractCodeHash
func (mr *MockContractUseCasesMockRecorder) SetContractCodeHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContractCodeHash", reflect.TypeOf((*MockContractUseCases)(nil).SetContractCodeHash))
}

// RegisterContract mocks base method
func (m *MockContractUseCases) RegisterContract() usecases.RegisterContractUseCase {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterContract")
	ret0, _ := ret[0].(usecases.RegisterContractUseCase)
	return ret0
}

// RegisterContract indicates an expected call of RegisterContract
func (mr *MockContractUseCasesMockRecorder) RegisterContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterContract", reflect.TypeOf((*MockContractUseCases)(nil).RegisterContract))
}

// MockGetContractsCatalogUseCase is a mock of GetContractsCatalogUseCase interface
type MockGetContractsCatalogUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetContractsCatalogUseCaseMockRecorder
}

// MockGetContractsCatalogUseCaseMockRecorder is the mock recorder for MockGetContractsCatalogUseCase
type MockGetContractsCatalogUseCaseMockRecorder struct {
	mock *MockGetContractsCatalogUseCase
}

// NewMockGetContractsCatalogUseCase creates a new mock instance
func NewMockGetContractsCatalogUseCase(ctrl *gomock.Controller) *MockGetContractsCatalogUseCase {
	mock := &MockGetContractsCatalogUseCase{ctrl: ctrl}
	mock.recorder = &MockGetContractsCatalogUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetContractsCatalogUseCase) EXPECT() *MockGetContractsCatalogUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetContractsCatalogUseCase) Execute(ctx context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetContractsCatalogUseCaseMockRecorder) Execute(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetContractsCatalogUseCase)(nil).Execute), ctx)
}

// MockGetContractUseCase is a mock of GetContractUseCase interface
type MockGetContractUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetContractUseCaseMockRecorder
}

// MockGetContractUseCaseMockRecorder is the mock recorder for MockGetContractUseCase
type MockGetContractUseCaseMockRecorder struct {
	mock *MockGetContractUseCase
}

// NewMockGetContractUseCase creates a new mock instance
func NewMockGetContractUseCase(ctrl *gomock.Controller) *MockGetContractUseCase {
	mock := &MockGetContractUseCase{ctrl: ctrl}
	mock.recorder = &MockGetContractUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetContractUseCase) EXPECT() *MockGetContractUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetContractUseCase) Execute(ctx context.Context, name, tag string) (*entities.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, name, tag)
	ret0, _ := ret[0].(*entities.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetContractUseCaseMockRecorder) Execute(ctx, name, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetContractUseCase)(nil).Execute), ctx, name, tag)
}

// MockGetContractEventsUseCase is a mock of GetContractEventsUseCase interface
type MockGetContractEventsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetContractEventsUseCaseMockRecorder
}

// MockGetContractEventsUseCaseMockRecorder is the mock recorder for MockGetContractEventsUseCase
type MockGetContractEventsUseCaseMockRecorder struct {
	mock *MockGetContractEventsUseCase
}

// NewMockGetContractEventsUseCase creates a new mock instance
func NewMockGetContractEventsUseCase(ctrl *gomock.Controller) *MockGetContractEventsUseCase {
	mock := &MockGetContractEventsUseCase{ctrl: ctrl}
	mock.recorder = &MockGetContractEventsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetContractEventsUseCase) EXPECT() *MockGetContractEventsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetContractEventsUseCase) Execute(ctx context.Context, chainID string, address common.Address, sighash string, indexedInputCount uint32) (string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, chainID, address, sighash, indexedInputCount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Execute indicates an expected call of Execute
func (mr *MockGetContractEventsUseCaseMockRecorder) Execute(ctx, chainID, address, sighash, indexedInputCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetContractEventsUseCase)(nil).Execute), ctx, chainID, address, sighash, indexedInputCount)
}

// MockGetContractMethodSignaturesUseCase is a mock of GetContractMethodSignaturesUseCase interface
type MockGetContractMethodSignaturesUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetContractMethodSignaturesUseCaseMockRecorder
}

// MockGetContractMethodSignaturesUseCaseMockRecorder is the mock recorder for MockGetContractMethodSignaturesUseCase
type MockGetContractMethodSignaturesUseCaseMockRecorder struct {
	mock *MockGetContractMethodSignaturesUseCase
}

// NewMockGetContractMethodSignaturesUseCase creates a new mock instance
func NewMockGetContractMethodSignaturesUseCase(ctrl *gomock.Controller) *MockGetContractMethodSignaturesUseCase {
	mock := &MockGetContractMethodSignaturesUseCase{ctrl: ctrl}
	mock.recorder = &MockGetContractMethodSignaturesUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetContractMethodSignaturesUseCase) EXPECT() *MockGetContractMethodSignaturesUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetContractMethodSignaturesUseCase) Execute(ctx context.Context, name, tag, methodName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, name, tag, methodName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetContractMethodSignaturesUseCaseMockRecorder) Execute(ctx, name, tag, methodName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetContractMethodSignaturesUseCase)(nil).Execute), ctx, name, tag, methodName)
}

// MockGetContractMethodsUseCase is a mock of GetContractMethodsUseCase interface
type MockGetContractMethodsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetContractMethodsUseCaseMockRecorder
}

// MockGetContractMethodsUseCaseMockRecorder is the mock recorder for MockGetContractMethodsUseCase
type MockGetContractMethodsUseCaseMockRecorder struct {
	mock *MockGetContractMethodsUseCase
}

// NewMockGetContractMethodsUseCase creates a new mock instance
func NewMockGetContractMethodsUseCase(ctrl *gomock.Controller) *MockGetContractMethodsUseCase {
	mock := &MockGetContractMethodsUseCase{ctrl: ctrl}
	mock.recorder = &MockGetContractMethodsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetContractMethodsUseCase) EXPECT() *MockGetContractMethodsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetContractMethodsUseCase) Execute(ctx context.Context, chainID string, address common.Address, selector []byte) (string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, chainID, address, selector)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Execute indicates an expected call of Execute
func (mr *MockGetContractMethodsUseCaseMockRecorder) Execute(ctx, chainID, address, selector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetContractMethodsUseCase)(nil).Execute), ctx, chainID, address, selector)
}

// MockGetContractTagsUseCase is a mock of GetContractTagsUseCase interface
type MockGetContractTagsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetContractTagsUseCaseMockRecorder
}

// MockGetContractTagsUseCaseMockRecorder is the mock recorder for MockGetContractTagsUseCase
type MockGetContractTagsUseCaseMockRecorder struct {
	mock *MockGetContractTagsUseCase
}

// NewMockGetContractTagsUseCase creates a new mock instance
func NewMockGetContractTagsUseCase(ctrl *gomock.Controller) *MockGetContractTagsUseCase {
	mock := &MockGetContractTagsUseCase{ctrl: ctrl}
	mock.recorder = &MockGetContractTagsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGetContractTagsUseCase) EXPECT() *MockGetContractTagsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockGetContractTagsUseCase) Execute(ctx context.Context, name string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, name)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockGetContractTagsUseCaseMockRecorder) Execute(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetContractTagsUseCase)(nil).Execute), ctx, name)
}

// MockRegisterContractUseCase is a mock of RegisterContractUseCase interface
type MockRegisterContractUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterContractUseCaseMockRecorder
}

// MockRegisterContractUseCaseMockRecorder is the mock recorder for MockRegisterContractUseCase
type MockRegisterContractUseCaseMockRecorder struct {
	mock *MockRegisterContractUseCase
}

// NewMockRegisterContractUseCase creates a new mock instance
func NewMockRegisterContractUseCase(ctrl *gomock.Controller) *MockRegisterContractUseCase {
	mock := &MockRegisterContractUseCase{ctrl: ctrl}
	mock.recorder = &MockRegisterContractUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegisterContractUseCase) EXPECT() *MockRegisterContractUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockRegisterContractUseCase) Execute(ctx context.Context, contract *entities.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, contract)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockRegisterContractUseCaseMockRecorder) Execute(ctx, contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRegisterContractUseCase)(nil).Execute), ctx, contract)
}

// MockSetContractCodeHashUseCase is a mock of SetContractCodeHashUseCase interface
type MockSetContractCodeHashUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSetContractCodeHashUseCaseMockRecorder
}

// MockSetContractCodeHashUseCaseMockRecorder is the mock recorder for MockSetContractCodeHashUseCase
type MockSetContractCodeHashUseCaseMockRecorder struct {
	mock *MockSetContractCodeHashUseCase
}

// NewMockSetContractCodeHashUseCase creates a new mock instance
func NewMockSetContractCodeHashUseCase(ctrl *gomock.Controller) *MockSetContractCodeHashUseCase {
	mock := &MockSetContractCodeHashUseCase{ctrl: ctrl}
	mock.recorder = &MockSetContractCodeHashUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSetContractCodeHashUseCase) EXPECT() *MockSetContractCodeHashUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockSetContractCodeHashUseCase) Execute(ctx context.Context, chainID string, address common.Address, hash string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, chainID, address, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockSetContractCodeHashUseCaseMockRecorder) Execute(ctx, chainID, address, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSetContractCodeHashUseCase)(nil).Execute), ctx, chainID, address, hash)
}
