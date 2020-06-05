// Code generated by MockGen. DO NOT EDIT.
// Source: search_jobs.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entities "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/transaction-scheduler/entities"
	reflect "reflect"
)

// MockSearchJobsUseCase is a mock of SearchJobsUseCase interface.
type MockSearchJobsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSearchJobsUseCaseMockRecorder
}

// MockSearchJobsUseCaseMockRecorder is the mock recorder for MockSearchJobsUseCase.
type MockSearchJobsUseCaseMockRecorder struct {
	mock *MockSearchJobsUseCase
}

// NewMockSearchJobsUseCase creates a new mock instance.
func NewMockSearchJobsUseCase(ctrl *gomock.Controller) *MockSearchJobsUseCase {
	mock := &MockSearchJobsUseCase{ctrl: ctrl}
	mock.recorder = &MockSearchJobsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchJobsUseCase) EXPECT() *MockSearchJobsUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockSearchJobsUseCase) Execute(ctx context.Context, filters *entities.JobFilters, tenantID string) ([]*entities.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, filters, tenantID)
	ret0, _ := ret[0].([]*entities.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockSearchJobsUseCaseMockRecorder) Execute(ctx, filters, tenantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockSearchJobsUseCase)(nil).Execute), ctx, filters, tenantID)
}
