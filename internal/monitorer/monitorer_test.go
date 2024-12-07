package monitorer

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/monitorer/mock"
	"go.uber.org/mock/gomock"
)

const targetURL = "/"

func prepareMockFlags(m *mock.MockFlags) {
	m.EXPECT().IntervalSeconds().Return(1)
	m.EXPECT().CreateLogFile().Return(false)
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
			name: "",
			prepareMockHTTPClient: func(mc *mock.MockHTTPClient) {
				res := &http.Response{
					Status: "200 OK",
				}
				mc.EXPECT().Do(gomock.Any()).Return(res, nil)
			},
			want:    "",
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
			opt, err := NewOptions(flags)
			if !tt.wantErr {
				require.NoError(t, err)
			}
			m := New(mc, targetURL, opt)
			got, err := m.result(context.Background())
			if !tt.wantErr {
				require.NoError(t, err)
			}

			assert.Contains(t, got, "200 OK")
		})
	}
}
