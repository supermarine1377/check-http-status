package format

import (
	"context"

	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/logger/format/internal"
)

type Format int

const (
	PLAIN_TEXT Format = iota
)

type Formatter interface {
	Format(ctx context.Context, r *models.Response, message string) string
}

func New(f Format) Formatter {
	switch f {
	case PLAIN_TEXT:
		return internal.NewPlainTextFormatter()
	default:
		return internal.NewPlainTextFormatter()
	}
}
