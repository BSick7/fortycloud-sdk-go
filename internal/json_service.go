package internal

import (
	"log"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type JsonService struct {
	url string
	client *http.Client
	requestSites []RequestSite
}

func NewJsonService(url string, client *http.Client) *JsonService {
    if client == nil {
        client = http.DefaultClient
    }
	svc := &JsonService {
		url: url,
		client: client,
	}
	svc.InjectRequest(func (method string, endpoint string, req *http.Request) error {
    	req.Header.Set("Content-Type", "application/json")
		return nil
	})
	return svc
}

func (svc *JsonService) InjectRequest(action RequestSite) {
	svc.requestSites = append(svc.requestSites, action)
}

func (svc *JsonService) Get(endpoint string, result interface{}) error {
	return svc.do("GET", endpoint, nil, result)
}
func (svc *JsonService) Post(endpoint string, body interface{}, result interface{}) error {
	json, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return svc.do("POST", endpoint, json, result)
}
func (svc *JsonService) Put(endpoint string, id string, body interface{}, result interface{}) error {
	json, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return svc.do("PUT", endpoint + "/" + id, json, result)
}
func (svc *JsonService) Delete(endpoint string, result interface{}) error {
	return svc.do("DELETE", endpoint, nil, result)
}

func (svc *JsonService) do(method string, endpoint string, body []byte, result interface{}) error {
	req, err := http.NewRequest(method, svc.url + endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	
	for _, site := range svc.requestSites {
		err := site(method, endpoint, req)
		if err != nil {
		return err
		}
	}
	
	log.Println(method, svc.url + endpoint)
	res, err := svc.client.Do(req)
	if err != nil {
		return err
	}
	
    defer res.Body.Close()
    resbody, _ := ioutil.ReadAll(res.Body)
	
	return json.Unmarshal(resbody, result)
}