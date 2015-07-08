package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type FormService struct {
	url           string
	client        *http.Client
	requestSites  []RequestSite
	responseSites []ResponseSite
}

func NewFormService(url string, client *http.Client) *FormService {
	if client == nil {
		client = http.DefaultClient
	}
	svc := &FormService{
		url:    url,
		client: client,
	}
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return nil
	})
	return svc
}

func (svc *FormService) InjectRequest(action RequestSite) {
	svc.requestSites = append(svc.requestSites, action)
}
func (svc *FormService) InjectResponse(action ResponseSite) {
	svc.responseSites = append(svc.responseSites, action)
}

func (svc *FormService) Get(endpoint string, result interface{}) error {
	return svc.do("GET", endpoint, nil, result)
}
func (svc *FormService) Post(endpoint string, body map[string]string, result interface{}) error {
	data := url.Values{}
	if body != nil {
		for k, v := range body {
			data.Set(k, v)
		}
	}
	return svc.do("POST", endpoint, data, result)
}
func (svc *FormService) Put(endpoint string, id string, body map[string]string, result interface{}) error {
	data := url.Values{}
	if body != nil {
		for k, v := range body {
			data.Set(k, v)
		}
	}
	return svc.do("PUT", endpoint+"/"+id, data, result)
}
func (svc *FormService) Delete(endpoint string, result interface{}) error {
	return svc.do("DELETE", endpoint, nil, result)
}

func (svc *FormService) do(method string, endpoint string, data url.Values, result interface{}) error {
	req, err := http.NewRequest(method, svc.url+endpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))

	for _, site := range svc.requestSites {
		err := site(method, endpoint, req)
		if err != nil {
			return err
		}
	}

	log.Println(method, svc.url+endpoint)
	res, err := svc.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	resbody, _ := ioutil.ReadAll(res.Body)

	for _, site := range svc.responseSites {
		err := site(method, endpoint, res)
		if err != nil {
			return err
		}
	}

	json.Unmarshal(resbody, result)
	return nil
}
