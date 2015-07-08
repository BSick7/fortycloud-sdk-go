package forms

import (
	"fmt"
	"errors"
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
	Where []FilterClause `json:"where"`
}
type privateSubnetsAllResult struct {
	EntityAllResult
	Objects []*PrivateSubnet `json:"objects"`
}
type PrivateSubnet struct {
	Id      int `json:"id"`
	Name    string `json:"name"`
	Description string `json:"description"`
	ActualSubnet string `json:"actualSubnet"`
	Source  string `json:"source"`
	Subnet  string `json:"subnet"`
	NatDisabled bool `json:"sNatDisabled"`
	Version int `json:"version"`
}

func (endpoint *PrivateSubnetsEndpoint) All(filters []FilterClause) ([]*PrivateSubnet, error) {
	if filters == nil {
		filters = []FilterClause{}
	}
	body := &privateSubnetsAllRequest{
		Offset: 0,
		Order: "DESC",
		OrderBy: "id",
		Rows: 100,
		Where: filters,
	}
	
	var result privateSubnetsAllResult
	_, err := endpoint.service.Post(endpoint.url, body, &result)
	if err != nil {
		return nil, err
	}
	return result.Objects, nil
}

func (endpoint *PrivateSubnetsEndpoint) Create(subnet *PrivateSubnet) (*PrivateSubnet, error) {
	err := endpoint.put(&subnetPutObject {
		Name: subnet.Name,
		Description: subnet.Description,
		ActualSubnet: subnet.ActualSubnet,
		Source: subnet.Source,
		Subnet: subnet.Subnet,
		NatDisabled: subnet.NatDisabled,
	})
	if err != nil {
		return nil, err
	}
	
	filters := []FilterClause{
		NewFilterLike("name", subnet.Name+"%"),
		NewFilterLike("subnet", subnet.Subnet+"%"),
	}
	matches, err := endpoint.All(filters)
	if err != nil {
		return nil, err
	}
	
	if len(matches) <= 0 {
		return nil, errors.New("Could not get created subnet.")
	}
	return matches[0], nil
}

type privateSubnetDeleteResult struct {
	Result string `json:"result"`
	Total int `json:"total"`
}
func (endpoint *PrivateSubnetsEndpoint) Delete(id int) error {
	var result privateSubnetDeleteResult
	_, err := endpoint.service.Delete(endpoint.url, []int{id}, &result)
	if err != nil {
		return err
	}
	if result.Result != "OK" {
		return errors.New(fmt.Sprintf("Failed subnet delete: %s", result.Result))
	}
	return nil
}

type subnetPutObject struct {
	Name string `json:"name"`
	Description string `json:"description"`
	ActualSubnet string `json:"actualSubnet"`
	GwId string `json:"gw.id"`
	Source string `json:"source"`
	Subnet string `json:"subnet"`
	NatDisabled bool `json:"sNatDisabled"`
	SubnetRoleId string `json:"subnetRole.id"`
}
type privateSubnetPutResult struct {
	Result string `json:"result"`
	Total int `json:"total"`
}
func (endpoint *PrivateSubnetsEndpoint) put(subnet *subnetPutObject) error {
	var result privateSubnetPutResult
	_, err := endpoint.service.Put(endpoint.url, "", subnet, &result)
	if err != nil {
		return err
	}
	if result.Result != "OK" {
		return errors.New(fmt.Sprintf("Failed subnet put: %s", result.Result))
	}
	return nil
}