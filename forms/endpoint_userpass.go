package forms

import (
	"net/http"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type UserPassEndpoint struct {
	service *internal.FormService
    url string
}

type UserPassResult struct {
	Cookies []*http.Cookie
}

func NewUserPassEndpoint(service *internal.FormService) *UserPassEndpoint {
	return &UserPassEndpoint {
		service: service,
		url: "/userpass",
	}
}

func (endpoint *UserPassEndpoint) Post(username string, password string, tenantName string) (*UserPassResult, error) {
	body := map[string]string {
		"username": username,
		"password": password,
	}
	var result UserPassResult
	cookies, err := endpoint.service.Post(endpoint.url, body, &result)
	result.Cookies = cookies
	return &result, err
}