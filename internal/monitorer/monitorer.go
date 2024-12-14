package monitorer

import (
	"context"
	"time"

	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/timeutil"
)

type Monitorer struct {
	httpClient HTTPClient
	targetURL  string
	Sleeper
	Logger
	Flags
}

func New(client HTTPClient, logger Logger, sleeper Sleeper, targetURL string, flags Flags) *Monitorer {
	return &Monitorer{
		httpClient: client,
		targetURL:  targetURL,
		Logger:     logger,
		Sleeper:    sleeper,
		Flags:      flags,
	}
}

//go:generate mockgen -source=$GOFILE -package=mock -destination=mock/mock.go
type Flags interface {
	TimeoutSeconds() int
}

type Sleeper interface {
	Sleep()
}

type HTTPClient interface {
	Get(ctx context.Context, req *models.Request) (*models.Response, error)
}

type Logger interface {
	Logln(s string)
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
				m.Logln(err.Error())
				continue
			}
			m.Logln(r)
			m.Sleep()
		}
	}
}

func (m *Monitorer) result(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(m.TimeoutSeconds())*time.Second,
	)
	defer cancel()
	req, err := models.NewRequest(m.targetURL)
	if err != nil {
		return "", err
	}
	res, err := m.httpClient.Get(ctx, req)
	if err != nil {
		return "", err
	}

	t := timeutil.NowStr()
	s := t + " " + res.Status
	return s, nil
}
