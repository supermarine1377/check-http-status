package timectx

import (
	"context"

	"github.com/supermarine1377/check-http-status/timectx/internal"
)

const timeFormat = "2006-01-02_15-04-05"

func NowStr(ctx context.Context) string {
	return internal.Now(ctx).Format(timeFormat)
}
