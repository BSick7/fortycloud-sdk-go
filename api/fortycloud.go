package fortycloud

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "errors"
    "fmt"
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

func (s *Service) Post(url string, body interface{}) ([]byte, error) {
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("X-Auth-Token", s.Auth.Token)
    
    res, err := s.Client.Do(req)
    if err != nil {
        return nil, err
    }
    
    defer res.Body.Close()
    
    return ioutil.ReadAll(res.Body)
}

func (s *Service) Put(url string, id string, body interface{}) ([]byte, error) {
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    
    req, err := http.NewRequest("POST", url + "/" + id, bytes.NewBuffer(jsonBody))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("X-Auth-Token", s.Auth.Token)
    
    res, err := s.Client.Do(req)
    if err != nil {
        return nil, err
    }
    
    defer res.Body.Close()
    
    return ioutil.ReadAll(res.Body)
}

func (s *Service) Delete(url string, id string) error {
    req, err := http.NewRequest("DELETE", url + "/" + id, nil)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Add("X-Auth-Token", s.Auth.Token)
    
    res, err := s.Client.Do(req)
    if err != nil {
        return err
    }
    
    defer res.Body.Close()
    
    content, errc := ioutil.ReadAll(res.Body)
    if res.StatusCode != 200 {
        return errors.New(fmt.Sprintf("Could not delete [%s] '%s': %s", url + "/" + id, res.Status, content))
    }
    return errc
}