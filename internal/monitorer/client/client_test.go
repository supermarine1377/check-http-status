package client_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/client"
	"github.com/supermarine1377/check-http-status/timectx/timectxtest"
)

var now = time.Date(2024, time.December, 14, 0, 0, 0, 0, time.Local)

type MockTransport struct {
	status string
	err    error
}

func (mt MockTransport) RoundTrip(_ *http.Request) (*http.Response, error) {
	if mt.err != nil {
		return nil, mt.err
	}
	return &http.Response{Status: mt.status}, nil
}

func TestClient_Get(t *testing.T) {
	type args struct {
		req models.Request
	}
	tests := []struct {
		name      string
		args      args
		transport MockTransport
		want      *models.Response
		wantErr   bool
	}{
		{
			name: "do method returns error",
			args: args{
				req: models.Request{RawURL: "http://localhost"},
			},
			transport: MockTransport{
				status: "500 internal server error",
				err:    errors.New("error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "do method doesn't return error",
			args: args{
				req: models.Request{RawURL: "http://localhost"},
			},
			transport: MockTransport{
				status: "200 OK",
				err:    nil,
			},
			want: &models.Response{
				ReceivedAt: now,
				Status:     "200 OK",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := client.New(tt.transport)
			ctx := context.Background()
			ctx = timectxtest.WithFixedNow(t, ctx, now)
			got, err := c.Get(ctx, &tt.args.req)
			if err == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
