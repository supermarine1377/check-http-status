package monitorer

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/mock"
	"go.uber.org/mock/gomock"
)

const targetURL = "https://localhost"

func mockOption(ctrl *gomock.Controller) *mock.MockOption {
	mo := mock.NewMockOption(ctrl)
	mo.EXPECT().TimeoutSeconds().Return(10).AnyTimes()
	return mo
}

func TestMonitorer_result(t *testing.T) {
	tests := []struct {
		name           string
		mockHTTPClient func(ctrl *gomock.Controller) *mock.MockHTTPClient
		want           *models.Response
		wantErr        bool
	}{
		{
			name: "200 OK",
			mockHTTPClient: func(ctrl *gomock.Controller) *mock.MockHTTPClient {
				mc := mock.NewMockHTTPClient(ctrl)
				req := &models.Request{
					RawURL: targetURL,
				}
				res := &models.Response{
					Status: "200 OK",
				}
				mc.EXPECT().Get(gomock.Any(), req).Return(res, nil)
				return mc
			},
			want: &models.Response{
				Status: "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mc := tt.mockHTTPClient(ctrl)
			opt := mockOption(ctrl)

			m := New(mc, nil, nil, targetURL, opt)

			got, err := m.result(context.Background())
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, got, tt.want)
		})
	}
}

func createMockSleeper(ctrl *gomock.Controller, done chan struct{}, sleepCount int) *mock.MockSleeper {
	ms := mock.NewMockSleeper(ctrl)
	callCount := 0
	ms.EXPECT().Sleep().Do(func() {
		callCount++
		if callCount == sleepCount {
			done <- struct{}{}
		}
	}).AnyTimes()
	return ms
}

func createMockLogger(ctrl *gomock.Controller, doSummarize bool) *mock.MockLogger {
	ml := mock.NewMockLogger(ctrl)
	ml.EXPECT().LogResponse(gomock.Any(), gomock.Any()).AnyTimes()
	ml.EXPECT().LogError(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	ml.EXPECT().LogErrorResponse(gomock.Any(), gomock.Any()).AnyTimes()
	if doSummarize {
		ml.EXPECT().SummarizeResults(gomock.Any()).Return()
	}
	return ml
}

func TestMonitorer_Do(t *testing.T) {
	tests := []struct {
		name           string
		mockHTTPClient func(ctrl *gomock.Controller) *mock.MockHTTPClient
		mockSleeper    func(ctrl *gomock.Controller, done chan struct{}) *mock.MockSleeper
		mockLogger     func(ctrl *gomock.Controller) *mock.MockLogger
		want           *models.Response
		wantErr        bool
	}{
		{
			name: "200 OK",
			mockHTTPClient: func(ctrl *gomock.Controller) *mock.MockHTTPClient {
				mc := mock.NewMockHTTPClient(ctrl)
				req := &models.Request{
					RawURL: targetURL,
				}
				res := &models.Response{
					Status: "200 OK",
				}
				mc.EXPECT().Get(gomock.Any(), req).Return(res, nil).AnyTimes()
				return mc
			},
			mockSleeper: func(ctrl *gomock.Controller, done chan struct{}) *mock.MockSleeper {
				return createMockSleeper(ctrl, done, 3)
			},
			mockLogger: func(ctrl *gomock.Controller) *mock.MockLogger {
				return createMockLogger(ctrl, true)
			},
		},
		{
			name: "2 successes followed by 1 failure",
			mockHTTPClient: func(ctrl *gomock.Controller) *mock.MockHTTPClient {
				mc := mock.NewMockHTTPClient(ctrl)
				req := &models.Request{
					RawURL: targetURL,
				}
				successRes := &models.Response{Status: "200 OK"}
				var callCount int
				mc.EXPECT().Get(gomock.Any(), req).DoAndReturn(func(ctx context.Context, req *models.Request) (*models.Response, error) {
					callCount++
					if callCount <= 2 {
						return successRes, nil
					}
					return nil, errors.New("network error")
				}).AnyTimes()
				return mc
			},
			mockSleeper: func(ctrl *gomock.Controller, done chan struct{}) *mock.MockSleeper {
				return createMockSleeper(ctrl, done, 3)
			},
			mockLogger: func(ctrl *gomock.Controller) *mock.MockLogger {
				return createMockLogger(ctrl, false)
			},
		},
		{
			name: "2 successes followed by 500 response",
			mockHTTPClient: func(ctrl *gomock.Controller) *mock.MockHTTPClient {
				mc := mock.NewMockHTTPClient(ctrl)
				req := &models.Request{
					RawURL: targetURL,
				}
				successRes := &models.Response{Status: "200 OK"}
				var callCount int
				mc.EXPECT().Get(gomock.Any(), req).DoAndReturn(func(ctx context.Context, req *models.Request) (*models.Response, error) {
					callCount++
					if callCount <= 2 {
						return successRes, nil
					}
					return &models.Response{Status: "500 Internal server error"}, nil
				}).AnyTimes()
				return mc
			},
			mockSleeper: func(ctrl *gomock.Controller, done chan struct{}) *mock.MockSleeper {
				return createMockSleeper(ctrl, done, 3)
			},
			mockLogger: func(ctrl *gomock.Controller) *mock.MockLogger {
				return createMockLogger(ctrl, true)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)

			mc := tt.mockHTTPClient(ctrl)
			ml := tt.mockLogger(ctrl)
			done := make(chan struct{})
			ms := tt.mockSleeper(ctrl, done)
			opt := mockOption(ctrl)

			m := New(mc, ml, ms, targetURL, opt)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			go func() {
				<-done
				cancel()
			}()

			m.Do(ctx)

		})
	}
}
