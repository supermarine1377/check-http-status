package logger_test

import (
	"bytes"
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/internal/monitorer/logger"
	"github.com/supermarine1377/check-http-status/timectx"
	"github.com/supermarine1377/check-http-status/timectx/timectxtest"
)

var now = time.Date(2024, time.December, 14, 0, 0, 0, 0, time.Local)

func TestLogger_Logln(t *testing.T) {
	out := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var captured bytes.Buffer
	var wg sync.WaitGroup
	// io.Pipeのio.Reader.Read()とio.Writer.Write()は同期的に実行できない
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, _ = captured.ReadFrom(r)
	}()

	l, err := logger.New(false)
	require.NoError(t, err)

	ctx := timectxtest.WithFixedNow(t, context.Background(), now)
	res := &models.Response{Status: "200 OK", ReceivedAt: timectx.Now(ctx), ResponseTime: time.Second}
	l.Logln(ctx, res)

	w.Close()
	wg.Wait()

	os.Stdout = out

	assert.Equal(t, "2024-12-14_00-00-00 1s 200 OK\n", captured.String())
}
