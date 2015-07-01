package fortycloud

import (
    "net/http"
)

type Service struct {
	client *http.Client
}

func NewService(c *http.Client) *Service {
    if c == nil {
        c = http.DefaultClient
    }
    return &Service{
        client: c,
    }
}