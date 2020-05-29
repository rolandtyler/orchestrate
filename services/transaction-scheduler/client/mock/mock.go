// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/services/transaction-scheduler/service/types"
	reflect "reflect"
)

// MockTransactionClient is a mock of TransactionClient interface.
type MockTransactionClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionClientMockRecorder
}

// MockTransactionClientMockRecorder is the mock recorder for MockTransactionClient.
type MockTransactionClientMockRecorder struct {
	mock *MockTransactionClient
}

// NewMockTransactionClient creates a new mock instance.
func NewMockTransactionClient(ctrl *gomock.Controller) *MockTransactionClient {
	mock := &MockTransactionClient{ctrl: ctrl}
	mock.recorder = &MockTransactionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionClient) EXPECT() *MockTransactionClientMockRecorder {
	return m.recorder
}

// SendTransaction mocks base method.
func (m *MockTransactionClient) SendTransaction(ctx context.Context, chainUUID string, request *types.SendTransactionRequest) (*types.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, chainUUID, request)
	ret0, _ := ret[0].(*types.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockTransactionClientMockRecorder) SendTransaction(ctx, chainUUID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockTransactionClient)(nil).SendTransaction), ctx, chainUUID, request)
}

// MockScheduleClient is a mock of ScheduleClient interface.
type MockScheduleClient struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleClientMockRecorder
}

// MockScheduleClientMockRecorder is the mock recorder for MockScheduleClient.
type MockScheduleClientMockRecorder struct {
	mock *MockScheduleClient
}

// NewMockScheduleClient creates a new mock instance.
func NewMockScheduleClient(ctrl *gomock.Controller) *MockScheduleClient {
	mock := &MockScheduleClient{ctrl: ctrl}
	mock.recorder = &MockScheduleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduleClient) EXPECT() *MockScheduleClientMockRecorder {
	return m.recorder
}

// GetSchedule mocks base method.
func (m *MockScheduleClient) GetSchedule(ctx context.Context, scheduleUUID string) (*types.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", ctx, scheduleUUID)
	ret0, _ := ret[0].(*types.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule.
func (mr *MockScheduleClientMockRecorder) GetSchedule(ctx, scheduleUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockScheduleClient)(nil).GetSchedule), ctx, scheduleUUID)
}

// GetSchedules mocks base method.
func (m *MockScheduleClient) GetSchedules(ctx context.Context) ([]*types.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedules", ctx)
	ret0, _ := ret[0].([]*types.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedules indicates an expected call of GetSchedules.
func (mr *MockScheduleClientMockRecorder) GetSchedules(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedules", reflect.TypeOf((*MockScheduleClient)(nil).GetSchedules), ctx)
}

// CreateSchedule mocks base method.
func (m *MockScheduleClient) CreateSchedule(ctx context.Context, request *types.CreateScheduleRequest) (*types.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", ctx, request)
	ret0, _ := ret[0].(*types.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchedule indicates an expected call of CreateSchedule.
func (mr *MockScheduleClientMockRecorder) CreateSchedule(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockScheduleClient)(nil).CreateSchedule), ctx, request)
}

// MockJobClient is a mock of JobClient interface.
type MockJobClient struct {
	ctrl     *gomock.Controller
	recorder *MockJobClientMockRecorder
}

// MockJobClientMockRecorder is the mock recorder for MockJobClient.
type MockJobClientMockRecorder struct {
	mock *MockJobClient
}

// NewMockJobClient creates a new mock instance.
func NewMockJobClient(ctrl *gomock.Controller) *MockJobClient {
	mock := &MockJobClient{ctrl: ctrl}
	mock.recorder = &MockJobClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobClient) EXPECT() *MockJobClientMockRecorder {
	return m.recorder
}

// GetJob mocks base method.
func (m *MockJobClient) GetJob(ctx context.Context, jobUUID string) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", ctx, jobUUID)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob.
func (mr *MockJobClientMockRecorder) GetJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockJobClient)(nil).GetJob), ctx, jobUUID)
}

// GetJobs mocks base method.
func (m *MockJobClient) GetJobs(ctx context.Context) ([]*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", ctx)
	ret0, _ := ret[0].([]*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockJobClientMockRecorder) GetJobs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockJobClient)(nil).GetJobs), ctx)
}

// CreateJob mocks base method.
func (m *MockJobClient) CreateJob(ctx context.Context, request *types.CreateJobRequest) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, request)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockJobClientMockRecorder) CreateJob(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockJobClient)(nil).CreateJob), ctx, request)
}

// UpdateJob mocks base method.
func (m *MockJobClient) UpdateJob(ctx context.Context, jobUUID string, request *types.UpdateJobRequest) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJob", ctx, jobUUID, request)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJob indicates an expected call of UpdateJob.
func (mr *MockJobClientMockRecorder) UpdateJob(ctx, jobUUID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockJobClient)(nil).UpdateJob), ctx, jobUUID, request)
}

// StartJob mocks base method.
func (m *MockJobClient) StartJob(ctx context.Context, jobUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartJob", ctx, jobUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartJob indicates an expected call of StartJob.
func (mr *MockJobClientMockRecorder) StartJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartJob", reflect.TypeOf((*MockJobClient)(nil).StartJob), ctx, jobUUID)
}

// MockTransactionSchedulerClient is a mock of TransactionSchedulerClient interface.
type MockTransactionSchedulerClient struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionSchedulerClientMockRecorder
}

// MockTransactionSchedulerClientMockRecorder is the mock recorder for MockTransactionSchedulerClient.
type MockTransactionSchedulerClientMockRecorder struct {
	mock *MockTransactionSchedulerClient
}

// NewMockTransactionSchedulerClient creates a new mock instance.
func NewMockTransactionSchedulerClient(ctrl *gomock.Controller) *MockTransactionSchedulerClient {
	mock := &MockTransactionSchedulerClient{ctrl: ctrl}
	mock.recorder = &MockTransactionSchedulerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionSchedulerClient) EXPECT() *MockTransactionSchedulerClientMockRecorder {
	return m.recorder
}

// SendTransaction mocks base method.
func (m *MockTransactionSchedulerClient) SendTransaction(ctx context.Context, chainUUID string, request *types.SendTransactionRequest) (*types.TransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", ctx, chainUUID, request)
	ret0, _ := ret[0].(*types.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockTransactionSchedulerClientMockRecorder) SendTransaction(ctx, chainUUID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).SendTransaction), ctx, chainUUID, request)
}

// GetSchedule mocks base method.
func (m *MockTransactionSchedulerClient) GetSchedule(ctx context.Context, scheduleUUID string) (*types.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", ctx, scheduleUUID)
	ret0, _ := ret[0].(*types.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule.
func (mr *MockTransactionSchedulerClientMockRecorder) GetSchedule(ctx, scheduleUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetSchedule), ctx, scheduleUUID)
}

// GetSchedules mocks base method.
func (m *MockTransactionSchedulerClient) GetSchedules(ctx context.Context) ([]*types.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedules", ctx)
	ret0, _ := ret[0].([]*types.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedules indicates an expected call of GetSchedules.
func (mr *MockTransactionSchedulerClientMockRecorder) GetSchedules(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedules", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetSchedules), ctx)
}

// CreateSchedule mocks base method.
func (m *MockTransactionSchedulerClient) CreateSchedule(ctx context.Context, request *types.CreateScheduleRequest) (*types.ScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", ctx, request)
	ret0, _ := ret[0].(*types.ScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchedule indicates an expected call of CreateSchedule.
func (mr *MockTransactionSchedulerClientMockRecorder) CreateSchedule(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).CreateSchedule), ctx, request)
}

// GetJob mocks base method.
func (m *MockTransactionSchedulerClient) GetJob(ctx context.Context, jobUUID string) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", ctx, jobUUID)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob.
func (mr *MockTransactionSchedulerClientMockRecorder) GetJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetJob), ctx, jobUUID)
}

// GetJobs mocks base method.
func (m *MockTransactionSchedulerClient) GetJobs(ctx context.Context) ([]*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", ctx)
	ret0, _ := ret[0].([]*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockTransactionSchedulerClientMockRecorder) GetJobs(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).GetJobs), ctx)
}

// CreateJob mocks base method.
func (m *MockTransactionSchedulerClient) CreateJob(ctx context.Context, request *types.CreateJobRequest) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, request)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockTransactionSchedulerClientMockRecorder) CreateJob(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).CreateJob), ctx, request)
}

// UpdateJob mocks base method.
func (m *MockTransactionSchedulerClient) UpdateJob(ctx context.Context, jobUUID string, request *types.UpdateJobRequest) (*types.JobResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJob", ctx, jobUUID, request)
	ret0, _ := ret[0].(*types.JobResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateJob indicates an expected call of UpdateJob.
func (mr *MockTransactionSchedulerClientMockRecorder) UpdateJob(ctx, jobUUID, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).UpdateJob), ctx, jobUUID, request)
}

// StartJob mocks base method.
func (m *MockTransactionSchedulerClient) StartJob(ctx context.Context, jobUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartJob", ctx, jobUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartJob indicates an expected call of StartJob.
func (mr *MockTransactionSchedulerClientMockRecorder) StartJob(ctx, jobUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartJob", reflect.TypeOf((*MockTransactionSchedulerClient)(nil).StartJob), ctx, jobUUID)
}