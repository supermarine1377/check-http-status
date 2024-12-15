package logger

import (
	"context"
	"fmt"
	"io"
	"os"

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
	s := r.ReceivedAt.Format(timeFormat) + " " + r.Status
	b := []byte(s + "\n")
	_, _ = l.file.Write(b)
}

func (l *Logger) Error(ctx context.Context, err error) {
	t := timectx.Now(ctx).Format(timeFormat)
	s := t + " " + err.Error()
	fmt.Fprintln(os.Stderr, s)
	b := []byte(err.Error())
	_, _ = l.file.Write(b)
}

func fileName(ctx context.Context) string {
	t := timectx.Now(ctx).Format(timeFormat)
	return "check-http-status_" + t + ".log"
}
