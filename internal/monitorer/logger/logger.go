package logger

import (
	"context"
	"io"
	"os"

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

func (l *Logger) Logln(s string) {
	b := []byte(s + "\n")
	for _, f := range l.files {
		f.Write(b)
	}
}

func fileName(ctx context.Context) string {
	return "check-http-status_" + timectx.NowStr(ctx) + ".log"
}
