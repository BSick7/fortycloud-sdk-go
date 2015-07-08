package forms

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type FormAuthenticator struct {
	url    string
	client *http.Client
}

func NewFormAuthenticator(url string, client *http.Client) *FormAuthenticator {
	if client == nil {
		client = http.DefaultClient
	}
	return &FormAuthenticator{
		url:    url,
		client: client,
	}
}

type FormAuthenticatorResult struct {
	UserId    int
	AccountId int
}

func (authenticator *FormAuthenticator) Authenticate(username string, password string, authenticityToken string) (*FormAuthenticatorResult, error) {
	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("authenticityToken", authenticityToken)
	encoded := data.Encode()

	url := authenticator.url + "/authenticate/userpass"
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(encoded))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cache-Control", "no-cache")

	log.Println("POST ", url)
	res, err := authenticator.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resbody, _ := ioutil.ReadAll(res.Body)

	return findAdminInfo(string(resbody))
}

func findAdminInfo(body string) (*FormAuthenticatorResult, error) {
	re, err := regexp.Compile(`ng-init="init\((?P<user>[\d]*), (?P<account>[\d]*), true\)"`)
	if err != nil {
		return nil, err
	}
	result := &FormAuthenticatorResult{
		UserId:    -1,
		AccountId: -1,
	}
	match := re.FindStringSubmatch(body)
	if match == nil {
		return result, nil
	}
	for i, name := range re.SubexpNames() {
		if name == "user" {
			result.UserId, _ = strconv.Atoi(match[i])
		}
		if name == "account" {
			result.AccountId, _ = strconv.Atoi(match[i])
		}
	}
	return result, nil
}