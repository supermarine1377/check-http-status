package logger

import (
	"context"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/timectx"
)

type Logger struct {
	file io.Writer
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

	return &Logger{file: w}, nil
}

const timeFormat = "2006-01-02_15-04-05"

func (l *Logger) Logln(ctx context.Context, r *models.Response) {
	t := r.ReceivedAt.Format(timeFormat)
	resTime := r.ResponseTime.String()
	fontColor := color.New(color.FgGreen)
	_, _ = fontColor.Fprintln(l.file, t, resTime, r.Status)
}

func (l *Logger) ErrorRes(ctx context.Context, r *models.Response) {
	t := timectx.Now(ctx).Format(timeFormat)
	fontColor := color.New(color.FgRed)
	_, _ = fontColor.Fprintln(l.file, t, r.Status)
}

func (l *Logger) Error(ctx context.Context, err error) {
	t := timectx.Now(ctx).Format(timeFormat)
	fontColor := color.New(color.FgRed)
	_, _ = fontColor.Fprintln(l.file, t, err)
}

func fileName(ctx context.Context) string {
	t := timectx.Now(ctx).Format(timeFormat)
	return "check-http-status_" + t + ".log"
}
