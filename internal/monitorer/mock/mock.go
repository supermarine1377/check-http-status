// Code generated by MockGen. DO NOT EDIT.
// Source: monitorer.go
//
// Generated by this command:
//
//	mockgen -source=monitorer.go -package=mock -destination=mock/mock.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/supermarine1377/check-http-status/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockOption is a mock of Option interface.
type MockOption struct {
	ctrl     *gomock.Controller
	recorder *MockOptionMockRecorder
	isgomock struct{}
}

// MockOptionMockRecorder is the mock recorder for MockOption.
type MockOptionMockRecorder struct {
	mock *MockOption
}

// NewMockOption creates a new mock instance.
func NewMockOption(ctrl *gomock.Controller) *MockOption {
	mock := &MockOption{ctrl: ctrl}
	mock.recorder = &MockOptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOption) EXPECT() *MockOptionMockRecorder {
	return m.recorder
}

// TimeoutSeconds mocks base method.
func (m *MockOption) TimeoutSeconds() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeoutSeconds")
	ret0, _ := ret[0].(int)
	return ret0
}

// TimeoutSeconds indicates an expected call of TimeoutSeconds.
func (mr *MockOptionMockRecorder) TimeoutSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutSeconds", reflect.TypeOf((*MockOption)(nil).TimeoutSeconds))
}

// MockSleeper is a mock of Sleeper interface.
type MockSleeper struct {
	ctrl     *gomock.Controller
	recorder *MockSleeperMockRecorder
	isgomock struct{}
}

// MockSleeperMockRecorder is the mock recorder for MockSleeper.
type MockSleeperMockRecorder struct {
	mock *MockSleeper
}

// NewMockSleeper creates a new mock instance.
func NewMockSleeper(ctrl *gomock.Controller) *MockSleeper {
	mock := &MockSleeper{ctrl: ctrl}
	mock.recorder = &MockSleeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSleeper) EXPECT() *MockSleeperMockRecorder {
	return m.recorder
}

// Sleep mocks base method.
func (m *MockSleeper) Sleep() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sleep")
}

// Sleep indicates an expected call of Sleep.
func (mr *MockSleeperMockRecorder) Sleep() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockSleeper)(nil).Sleep))
}

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
	isgomock struct{}
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHTTPClient) Get(ctx context.Context, req *models.Request) (*models.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*models.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockHTTPClientMockRecorder) Get(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClient)(nil).Get), ctx, req)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
	isgomock struct{}
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// LogError mocks base method.
func (m *MockLogger) LogError(ctx context.Context, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogError", ctx, err)
}

// LogError indicates an expected call of LogError.
func (mr *MockLoggerMockRecorder) LogError(ctx, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockLogger)(nil).LogError), ctx, err)
}

// LogErrorResponse mocks base method.
func (m *MockLogger) LogErrorResponse(ctx context.Context, r *models.Response) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogErrorResponse", ctx, r)
}

// LogErrorResponse indicates an expected call of LogErrorResponse.
func (mr *MockLoggerMockRecorder) LogErrorResponse(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogErrorResponse", reflect.TypeOf((*MockLogger)(nil).LogErrorResponse), ctx, r)
}

// LogResponse mocks base method.
func (m *MockLogger) LogResponse(ctx context.Context, r *models.Response) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogResponse", ctx, r)
}

// LogResponse indicates an expected call of LogResponse.
func (mr *MockLoggerMockRecorder) LogResponse(ctx, r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogResponse", reflect.TypeOf((*MockLogger)(nil).LogResponse), ctx, r)
}

// SummarizeResults mocks base method.
func (m *MockLogger) SummarizeResults(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SummarizeResults", ctx)
}

// SummarizeResults indicates an expected call of SummarizeResults.
func (mr *MockLoggerMockRecorder) SummarizeResults(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SummarizeResults", reflect.TypeOf((*MockLogger)(nil).SummarizeResults), ctx)
}
