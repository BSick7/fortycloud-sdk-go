package forms

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type PrivateSubnetsEndpoint struct {
	service *internal.FormService
	container *internal.CookieContainer
	url string
}

func NewPrivateSubnetsEndpoint(service *internal.FormService) *PrivateSubnetsEndpoint {
	return &PrivateSubnetsEndpoint {
		service: service,
		url: "/EntityPrivateSubnet",
	}
}

type privateSubnetsAllResult struct {
	EntityAllResult
	Objects []*PrivateSubnet `json:"objects"`
}
type PrivateSubnet struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Source string `json:"source"`
	Subnet string `json:"subnet"`
	Version string `json:"version"`
}
func (endpoint *PrivateSubnetsEndpoint) All() ([]*PrivateSubnet, error) {
	var result privateSubnetsAllResult
	err := endpoint.service.Post(endpoint.url, nil, &result)
	if err != nil {
		return nil, err
	}
	return result.Objects, nil
}