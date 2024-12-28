package monitorer

import (
	"context"
	"time"

	"github.com/supermarine1377/check-http-status/internal/models"
)

type Monitorer struct {
	httpClient HTTPClient
	targetURL  string
	Sleeper
	Logger
	Option
}

func New(client HTTPClient, logger Logger, sleeper Sleeper, targetURL string, opt Option) *Monitorer {
	return &Monitorer{
		httpClient: client,
		targetURL:  targetURL,
		Logger:     logger,
		Sleeper:    sleeper,
		Option:     opt,
	}
}

//go:generate mockgen -source=$GOFILE -package=mock -destination=mock/mock.go
type Option interface {
	TimeoutSeconds() int
}

type Sleeper interface {
	Sleep()
}

type HTTPClient interface {
	Get(ctx context.Context, req *models.Request) (*models.Response, error)
}

type Logger interface {
	LogResponse(ctx context.Context, r *models.Response)
	LogError(ctx context.Context, format string, args ...interface{})
	LogErrorResponse(ctx context.Context, r *models.Response)
	SummarizeResults(ctx context.Context)
}

func (m *Monitorer) Do(ctx context.Context) {
	for {
		if err := m.once(ctx); err != nil {
			return
		}
		m.Sleep()
	}
}

func (m *Monitorer) once(ctx context.Context) error {
	select {
	case <-ctx.Done():
		m.SummarizeResults(ctx)
		return ctx.Err()
	default:
		return m.handleResult(ctx)
	}
}

func (m *Monitorer) handleResult(ctx context.Context) error {
	res, err := m.fetchResult(ctx)
	if err != nil {
		m.LogError(ctx, "%w", err)
		return err
	}
	if res.IsOK() {
		m.LogResponse(ctx, res)
	} else {
		m.LogErrorResponse(ctx, res)
	}
	return nil
}

func (m *Monitorer) fetchResult(ctx context.Context) (*models.Response, error) {
	ctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(m.TimeoutSeconds())*time.Second,
	)
	defer cancel()
	req, err := models.NewRequest(m.targetURL)
	if err != nil {
		return nil, err
	}
	res, err := m.httpClient.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
