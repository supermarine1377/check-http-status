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
	LogResponse(ctx context.Context, r *models.Response)
	LogError(ctx context.Context, format string, args ...interface{})
	LogErrorResponse(ctx context.Context, r *models.Response)
	SummarizeResults(ctx context.Context)
}

func (m *Monitorer) Do(ctx context.Context) {
Loop:
	for {
		select {
		case <-ctx.Done():
			m.SummarizeResults(ctx)
			break Loop
		default:
			r, err := m.result(ctx)
			if err != nil {
				m.LogError(ctx, "%w", err)
				continue
			}
			if r.IsOK() {
				m.LogResponse(ctx, r)
			} else {
				m.LogErrorResponse(ctx, r)
			}
			m.Sleep()
		}
	}
}

func (m *Monitorer) result(ctx context.Context) (*models.Response, error) {
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
