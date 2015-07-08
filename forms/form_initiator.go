package forms

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type FormInitiator struct {
	url    string
	client *http.Client
}

func NewFormInitiator(url string, client *http.Client) *FormInitiator {
	if client == nil {
		client = http.DefaultClient
	}
	return &FormInitiator{
		url:    url,
		client: client,
	}
}

type InitiatorResult struct {
	AuthenticityToken string
}

func (initiator *FormInitiator) Initiate() (*InitiatorResult, error) {
	req, err := http.NewRequest("GET", initiator.url+"/login", nil)
	if err != nil {
		return nil, err
	}

	log.Println("GET ", req.URL)
	res, err := initiator.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resbody, _ := ioutil.ReadAll(res.Body)

	token, err := findAuthenticityToken(string(resbody))
	if err != nil {
		return nil, err
	}

	result := &InitiatorResult{
		AuthenticityToken: token,
	}

	return result, nil
}

func findAuthenticityToken(body string) (string, error) {
	re, err := regexp.Compile(`input\stype="hidden"\sname="authenticityToken"\svalue="(?P<value>[^\"]*)"`)
	if err != nil {
		return "", err
	}
	match := re.FindStringSubmatch(body)
	if match == nil {
		return "", nil
	}
	for i, name := range re.SubexpNames() {
		if name == "value" {
			return match[i], nil
		}
	}
	return "", nil
}
