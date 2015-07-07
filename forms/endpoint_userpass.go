package forms

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type UserPassEndpoint struct {
	service *internal.FormService
    url string
}

type UserPassResult struct {
}

func NewUserPassEndpoint(service *internal.FormService) *UserPassEndpoint {
	return &UserPassEndpoint {
		service: service,
		url: "/authenticate/userpass",
	}
}

func (endpoint *UserPassEndpoint) Post(username string, password string, tenantName string, authenticityToken string) error {
	body := map[string]string {
		"username": username,
		"password": password,
		"authenticityToken": authenticityToken,
	}
	var result UserPassResult
	return endpoint.service.Post(endpoint.url, body, &result)
}