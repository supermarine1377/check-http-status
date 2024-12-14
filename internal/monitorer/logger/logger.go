package logger

import (
	"io"
	"os"

	"github.com/supermarine1377/check-http-status/timeutil"
)

type Logger struct {
	files []io.Writer
}

func New(createLogFile bool) (*Logger, error) {
	files := make([]io.Writer, 0, 2)
	files = append(files, os.Stdout)

	if createLogFile {
		logFile, err := os.Create(fileName())
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

func fileName() string {
	return "check-http-status_" + timeutil.NowStr() + ".log"
}
