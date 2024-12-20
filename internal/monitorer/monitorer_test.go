package monitorer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/mock"
	"go.uber.org/mock/gomock"
)

const targetURL = "https://localhost"

func prepareMockOption(mo *mock.MockOption) {
	mo.EXPECT().TimeoutSeconds().Return(10)
}
func TestMonitorer_result(t *testing.T) {
	tests := []struct {
		name                  string
		prepareMockHTTPClient func(mc *mock.MockHTTPClient)
		want                  *models.Response
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
			want: &models.Response{
				Status: "200 OK",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mc := mock.NewMockHTTPClient(ctrl)
			tt.prepareMockHTTPClient(mc)

			opt := mock.NewMockOption(ctrl)
			prepareMockOption(opt)

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
