package forms

import (
	"log"
	"net/http"
	"io/ioutil"
	"regexp"
	"strconv"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type FormAdminCollector struct {
	url string
	client *http.Client
}

func NewFormAdminCollector(url string, client *http.Client) *FormAdminCollector {
    if client == nil {
        client = http.DefaultClient
    }
	return &FormAdminCollector{
		url: url,
		client: client,
	}
}

type FormAdminCollectorResult struct {
	UserId int
	AccountId int
}

func (collector *FormAdminCollector) Collect(cookies *internal.CookieContainer) (*FormAdminCollectorResult, error) {
	req, err := http.NewRequest("GET", collector.url, nil)
	if err != nil {
		return nil, err
	}
	cookies.AddToRequest(req)
	
	log.Println("GET ", collector.url)
	res, err := collector.client.Do(req)
	if err != nil {
		return nil, err
	}
	
    defer res.Body.Close()
    resbody, _ := ioutil.ReadAll(res.Body)
	cookies.Merge(res.Cookies())
	
	return findAdminInfo(string(resbody))
}

func findAdminInfo(body string) (*FormAdminCollectorResult, error) {
	re, err := regexp.Compile(`ng-init="init\((?P<user>[\d]*), (?P<account>[\d]*), true\)"`)
	if err != nil {
		return nil, err
	}
	result := &FormAdminCollectorResult{}
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