package timectx

import (
	"context"
	"time"

	"github.com/supermarine1377/check-http-status/timectx/internal"
)

func Now(ctx context.Context) time.Time {
	return internal.Now(ctx)
}
