package logger

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/logger/metrics"
	"github.com/supermarine1377/check-http-status/timectx"
)

type Logger struct {
	w io.Writer
	m metrics.Metrics
}

func New(createLogFile bool) (*Logger, error) {
	var w io.Writer = os.Stdout

	if createLogFile {
		logFile, err := os.Create(fileName(context.Background()))
		if err != nil {
			return nil, err
		}
		w = io.MultiWriter(w, logFile)
	}

	return &Logger{
		w: w,
	}, nil
}

var (
	colorDefault = color.New(color.FgWhite)
	colorInfo    = color.New(color.FgGreen)
	colorError   = color.New(color.FgRed)
)

func (l *Logger) LogResponse(ctx context.Context, r *models.Response) {
	l.m.Update(r)
	l.LogInfo(ctx, "Response time=%s, Status=%s", r.ResponseTime, r.Status)
}

func (l *Logger) LogInfo(ctx context.Context, format string, args ...interface{}) {
	l.log(ctx, colorInfo, format, args...)
}

func (l *Logger) LogError(ctx context.Context, format string, args ...interface{}) {
	l.log(ctx, colorError, format, args...)
}

func (l *Logger) LogDefault(_ context.Context, format string, args ...interface{}) {
	_, _ = colorDefault.Fprintf(l.w, format, args...)
}

func (l *Logger) LogErrorResponse(ctx context.Context, r *models.Response) {
	l.m.Update(r)
	l.LogError(ctx, "Response time:%s, Status:%s", r.ResponseTime, r.Status)
}

const timeFormat = "2006-01-02_15-04-05"

func (l *Logger) log(ctx context.Context, color *color.Color, format string, args ...interface{}) {
	timestamp := timectx.Now(ctx).Format(timeFormat)
	message := fmt.Sprintf("Timestamp=%s, %s", timestamp, fmt.Sprintf(format, args...))
	_, _ = color.Fprintln(l.w, message)
}

func fileName(ctx context.Context) string {
	t := timectx.Now(ctx).Format(timeFormat)
	return "check-http-status_" + t + ".log"
}

func (l *Logger) SummarizeResults(ctx context.Context) {
	summery, err := l.m.Summarize()
	if err == nil {
		l.LogDefault(ctx, "\n----------Summary----------\n%s", summery)
	}
}
