package logger

import (
	"context"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/logger/format"
	"github.com/supermarine1377/check-http-status/internal/monitorer/logger/metrics"
	"github.com/supermarine1377/check-http-status/timectx"
)

type Logger struct {
	w io.Writer
	f format.Formatter
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
		f: format.New(format.PLAIN_TEXT),
	}, nil
}

var (
	colorDefault = color.New(color.FgWhite)
	colorInfo    = color.New(color.FgGreen)
	colorError   = color.New(color.FgRed)
)

func (l *Logger) LogResponse(ctx context.Context, r *models.Response) {
	l.m.Update(r)
	l.LogInfo(ctx, l.f.Format(ctx, r, ""))
}

func (l *Logger) LogInfo(ctx context.Context, message string) {
	l.log(colorInfo, message)
}

func (l *Logger) LogError(ctx context.Context, err error) {
	l.log(colorError, err.Error())
}

func (l *Logger) LogDefault(_ context.Context, format string, args ...interface{}) {
	_, _ = colorDefault.Fprintf(l.w, format, args...)
}

func (l *Logger) LogErrorResponse(ctx context.Context, r *models.Response) {
	l.m.Update(r)
	l.log(colorError, l.f.Format(ctx, r, "client received non 200 response"))
}

const timeFormat = "2006-01-02_15-04-05"

func (l *Logger) log(color *color.Color, message string) {
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
