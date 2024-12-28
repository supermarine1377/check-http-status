package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/timectx"
)

type PlainTextFormatter struct {
	TimeFormat string
}

const defaultTimeFormat = "2006-01-02_15-04-05"

func NewPlainTextFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{
		TimeFormat: defaultTimeFormat,
	}
}

func (p *PlainTextFormatter) Format(ctx context.Context, res *models.Response, message string) string {
	timestamp := timectx.Now(ctx).Format(p.TimeFormat)
	fields := []string{
		fmt.Sprintf("Timestamp=%s", timestamp),
		fmt.Sprintf("Response time=%s", res.ResponseTime),
		fmt.Sprintf("Status=%s", res.Status),
	}
	if message != "" {
		fields = append(fields, fmt.Sprintf("Message=%s", message))
	}
	return strings.Join(fields, ", ")
}
