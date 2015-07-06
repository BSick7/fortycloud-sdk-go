package fortycloud

import (
    "net/http"
    "encoding/json"
)

const (
    API_URL = "https://api.fortycloud.net/restapi"
)

type Service struct {
	client *http.Client
    auth *Authentication
}

func NewService(c *http.Client, auth *Authentication) *Service {
    if c == nil {
        c = http.DefaultClient
    }
    return &Service{
        client: c,
        auth: auth
    }
}