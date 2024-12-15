package models

import (
	"errors"
	"net/url"
	"time"
)

type Response struct {
	ReceivedAt   time.Time
	Status       string
	ResponseTime time.Duration
}

type Request struct {
	RawURL string
}

func NewRequest(rawURL string) (*Request, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	if url.Scheme != "http" && url.Scheme != "https" {
		return nil, errors.New("the given string has unsupported scheme")
	}
	if url.Host == "" {
		return nil, errors.New("the given string doesn't have a host")
	}
	return &Request{RawURL: rawURL}, nil
}
