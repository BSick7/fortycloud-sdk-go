package fortycloud

import (
    "net/http"
    "io/ioutil"
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

func (s *Service) Get(url string) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("X-Auth-Token", s.Auth.Token)
    
    res, err := s.Client.Do(req)
    if err != nil {
        return nil, err
    }
    
    defer res.Body.Close()
    
    return ioutil.ReadAll(res.Body)
}