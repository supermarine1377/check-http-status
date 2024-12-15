package client

import (
	"context"
	"net/http"

	"github.com/supermarine1377/check-http-status/internal/models"
	"github.com/supermarine1377/check-http-status/timectx"
)

type Client struct {
	base *http.Client
}

func New(transport http.RoundTripper) *Client {
	return &Client{base: &http.Client{
		Transport: transport,
	}}
}

func (c *Client) Get(ctx context.Context, req *models.Request) (*models.Response, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		req.RawURL,
		nil,
	)
	if err != nil {
		return nil, err
	}
	now1 := timectx.Now(ctx)
	res, err := c.base.Do(httpReq)
	if err != nil {
		return nil, err
	}
	now2 := timectx.Now(ctx)

	return &models.Response{
		Status:       res.Status,
		ReceivedAt:   timectx.Now(ctx),
		ResponseTime: now2.Sub(now1),
	}, nil
}
