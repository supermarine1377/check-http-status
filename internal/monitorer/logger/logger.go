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
	files []io.Writer
}

func New(createLogFile bool) (*Logger, error) {
	files := make([]io.Writer, 0, 2)
	files = append(files, os.Stdout)

	if createLogFile {
		logFile, err := os.Create(fileName(context.Background()))
		if err != nil {
			return nil, err
		}
		files = append(files, logFile)
	}

	return &Logger{files: files}, nil
}

func (l *Logger) Logln(ctx context.Context, r *models.Response) {
	t := timectx.NowStr(ctx)
	s := t + " " + r.Status
	b := []byte(s + "\n")
	for _, f := range l.files {
		_, _ = f.Write(b)
	}
}

func (l *Logger) Error(ctx context.Context, err error) {
	t := timectx.NowStr(ctx)
	s := t + " " + err.Error()
	fmt.Fprintln(os.Stderr, s)
	b := []byte(err.Error())
	for _, f := range l.files {
		_, _ = f.Write([]byte(b))
	}
}

func fileName(ctx context.Context) string {
	return "check-http-status_" + timectx.NowStr(ctx) + ".log"
}
