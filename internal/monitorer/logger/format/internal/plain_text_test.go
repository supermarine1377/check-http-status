package internal

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/timectx"
	"github.com/supermarine1377/check-http-status/timectx/timectxtest"
)

var now = time.Date(2024, time.December, 14, 0, 0, 0, 0, time.Local)

func TestPlainTextFormatter_Format(t *testing.T) {
	ctx := timectxtest.WithFixedNow(t, context.Background(), now)

	type args struct {
		res     *models.Response
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "No message",
			args: args{
				res: &models.Response{Status: "200 OK", ReceivedAt: timectx.Now(ctx), ResponseTime: time.Second},
			},
			want: "Timestamp=2024-12-14_00-00-00, Response time=1s, Status=200 OK",
		},
		{
			name: "No message",
			args: args{
				res:     &models.Response{Status: "200 OK", ReceivedAt: timectx.Now(ctx), ResponseTime: time.Second},
				message: "message",
			},
			want: "Timestamp=2024-12-14_00-00-00, Response time=1s, Status=200 OK, Message=message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ptf := NewPlainTextFormatter()

			got := ptf.Format(ctx, tt.args.res, tt.args.message)
			assert.Equal(t, tt.want, got)
		})
	}
}
