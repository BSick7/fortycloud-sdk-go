package forms

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type PrivateSubnetsEndpoint struct {
	service *internal.JsonService
	url     string
}

func NewPrivateSubnetsEndpoint(service *internal.JsonService) *PrivateSubnetsEndpoint {
	return &PrivateSubnetsEndpoint{
		service: service,
		url:     "/EntityPrivateSubnet",
	}
}

type privateSubnetsAllRequest struct {
	Match struct {} `json:"match"`
	Offset int `json:"offset"`
	Order string `json:"order"`
	OrderBy string `json:"orderBy"`
	Rows int `json:"rows"`
	Where []string `json:"where"`
}
type privateSubnetsAllResult struct {
	EntityAllResult
	Objects []*PrivateSubnet `json:"objects"`
}
type PrivateSubnet struct {
	Id      int `json:"id"`
	Name    string `json:"name"`
	Source  string `json:"source"`
	Subnet  string `json:"subnet"`
	Version int `json:"version"`
}

func (endpoint *PrivateSubnetsEndpoint) All() ([]*PrivateSubnet, error) {
	body := &privateSubnetsAllRequest{
		Offset: 0,
		Order: "DESC",
		OrderBy: "id",
		Rows: 100,
		Where: []string{},
	}
	
	var result privateSubnetsAllResult
	_, err := endpoint.service.Post(endpoint.url, body, &result)
	if err != nil {
		return nil, err
	}
	return result.Objects, nil
}
