package monitorer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/mock"
	"github.com/supermarine1377/check-http-status/timectx/timectxtest"
	"go.uber.org/mock/gomock"
)

var now = time.Date(2024, time.December, 14, 0, 0, 0, 0, time.Local)

const targetURL = "https://localhost"

func prepareMockFlags(m *mock.MockFlags) {
	m.EXPECT().TimeoutSeconds().Return(10)
}
func TestMonitorer_result(t *testing.T) {
	tests := []struct {
		name                  string
		prepareMockHTTPClient func(mc *mock.MockHTTPClient)
		want                  string
		wantErr               bool
	}{
		{
			name: "200 OK",
			prepareMockHTTPClient: func(mc *mock.MockHTTPClient) {
				req := &models.Request{
					RawURL: targetURL,
				}
				res := &models.Response{
					Status: "200 OK",
				}
				mc.EXPECT().Get(gomock.Any(), req).Return(res, nil)
			},
			want:    "2024-12-14_00-00-00 200 OK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mc := mock.NewMockHTTPClient(ctrl)
			tt.prepareMockHTTPClient(mc)

			flags := mock.NewMockFlags(ctrl)
			prepareMockFlags(flags)

			m := New(mc, nil, nil, targetURL, flags)
			ctx := timectxtest.WithFixedNow(t, context.Background(), now)

			got, err := m.result(ctx)
			if !tt.wantErr {
				require.NoError(t, err)
			}

			assert.Equal(t, got, tt.want)
		})
	}
}
