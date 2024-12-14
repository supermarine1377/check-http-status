package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/models"
)

func TestNewRequest(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Request
		wantErr bool
	}{
		{
			name: "Valid URL",
			args: args{
				rawURL: "https://example.com",
			},
			want: &models.Request{
				RawURL: "https://example.com",
			},
			wantErr: false,
		},
		{
			name: "Without scheme #1",
			args: args{
				rawURL: "example",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Without scheme",
			args: args{
				rawURL: "example.com #2",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Without host",
			args: args{
				rawURL: "https://",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := models.NewRequest(tt.args.rawURL)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
