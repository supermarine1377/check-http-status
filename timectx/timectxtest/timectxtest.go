package timectxtest

import (
	"context"
	"testing"
	"time"

	"github.com/supermarine1377/check-http-status/timectx/internal"
)

type ctxkey struct{}

func init() {
	if testing.Testing() {
		internal.Now = nowForTest
	}
}

func WithFixedNow(t *testing.T, ctx context.Context, fix time.Time) context.Context {
	t.Helper()
	return context.WithValue(ctx, ctxkey{}, fix)
}

func nowForTest(ctx context.Context) time.Time {
	now, ok := ctx.Value(ctxkey{}).(time.Time)
	if ok {
		return now
	}
	return internal.DefaultNow(ctx)
}
