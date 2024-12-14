package client

import (
	"context"
	"net/http"

	"github.com/supermarine1377/check-http-status/internal/models"
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
	res, err := c.base.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return &models.Response{Status: res.Status}, nil
}
