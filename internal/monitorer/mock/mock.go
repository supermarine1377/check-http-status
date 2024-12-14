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

// MockFlags is a mock of Flags interface.
type MockFlags struct {
	ctrl     *gomock.Controller
	recorder *MockFlagsMockRecorder
	isgomock struct{}
}

// MockFlagsMockRecorder is the mock recorder for MockFlags.
type MockFlagsMockRecorder struct {
	mock *MockFlags
}

// NewMockFlags creates a new mock instance.
func NewMockFlags(ctrl *gomock.Controller) *MockFlags {
	mock := &MockFlags{ctrl: ctrl}
	mock.recorder = &MockFlagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlags) EXPECT() *MockFlagsMockRecorder {
	return m.recorder
}

// CreateLogFile mocks base method.
func (m *MockFlags) CreateLogFile() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLogFile")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CreateLogFile indicates an expected call of CreateLogFile.
func (mr *MockFlagsMockRecorder) CreateLogFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLogFile", reflect.TypeOf((*MockFlags)(nil).CreateLogFile))
}

// IntervalSeconds mocks base method.
func (m *MockFlags) IntervalSeconds() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntervalSeconds")
	ret0, _ := ret[0].(int)
	return ret0
}

// IntervalSeconds indicates an expected call of IntervalSeconds.
func (mr *MockFlagsMockRecorder) IntervalSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntervalSeconds", reflect.TypeOf((*MockFlags)(nil).IntervalSeconds))
}

// TimeoutSeconds mocks base method.
func (m *MockFlags) TimeoutSeconds() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeoutSeconds")
	ret0, _ := ret[0].(int)
	return ret0
}

// TimeoutSeconds indicates an expected call of TimeoutSeconds.
func (mr *MockFlagsMockRecorder) TimeoutSeconds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutSeconds", reflect.TypeOf((*MockFlags)(nil).TimeoutSeconds))
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
