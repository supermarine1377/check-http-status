package monitorer

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/supermarine1377/check-http-status/internal/log_files"
	"github.com/supermarine1377/check-http-status/internal/monitorer/sleeper"
	"github.com/supermarine1377/check-http-status/timeutil"
)

type Monitorer struct {
	httpClient HTTPClient
	targetURL  string
	*Options
	Sleeper
}

func New(client HTTPClient, targetURL string, options *Options) *Monitorer {
	d := time.Second * time.Duration(options.IntervalSeconds())
	return &Monitorer{
		httpClient: client,
		targetURL:  targetURL,
		Options:    options,
		Sleeper:    sleeper.New(d),
	}
}

type Options struct {
	files []io.Writer
	Flags
}

func NewOptions(flags Flags) (*Options, error) {
	files, err := log_files.New(flags.CreateLogFile())
	if err != nil {
		return nil, err
	}
	return &Options{
		files: files,
		Flags: flags,
	}, nil
}

type Flags interface {
	IntervalSeconds() int
	CreateLogFile() bool
	TimeoutSeconds() int
}

type Sleeper interface {
	Sleep()
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (m *Monitorer) Do(ctx context.Context) {
Loop:
	for {
		select {
		case <-ctx.Done():
			break Loop
		default:
			r, err := m.result(ctx)
			if err != nil {
				m.logln(err.Error())
				continue
			}
			m.logln(r)
			m.Sleep()
		}
	}
}

func (m *Monitorer) result(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.TimeoutSeconds())*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		m.targetURL,
		nil,
	)
	if err != nil {
		return "", err
	}
	t := timeutil.NowStr()
	res, err := m.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	s := t + res.Status
	return s, nil
}

func (m *Monitorer) logln(s string) {
	b := []byte(s + "\n")
	for _, f := range m.files {
		f.Write(b)
	}
}
