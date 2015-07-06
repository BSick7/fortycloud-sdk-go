package fortycloud

import (
    "net/http"
)

const (
    API_URL = "https://api.fortycloud.net/restapi"
)

type Service struct {
	Client *http.Client
    Auth *Authentication
}

func NewService(client *http.Client, auth *Authentication) *Service {
    if client == nil {
        client = http.DefaultClient
    }
    return &Service{
        Client: client,
        Auth: auth,
    }
}

func (s *Service) Authenticate() error {
    return s.Auth.Do(s)
}