package forms

import (
	"fmt"
	"errors"
	"strconv"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type PrivateSubnetsEndpoint struct {
	service *internal.JsonService
	url     string
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

func (endpoint *PrivateSubnetsEndpoint) Get(id int) (*PrivateSubnet, error) {
	subnets, err := endpoint.All(nil)
	if err != nil {
		return nil, err
	}
	
	if len(subnets) <= 0 {
		return nil, nil
	}
	
	for _,subnet := range subnets {
		if subnet.Id == id {
			return subnet, nil
		}
	}
	
	return nil, nil
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

func (endpoint *PrivateSubnetsEndpoint) Update(subnet *PrivateSubnet) (*PrivateSubnet, error) {
	err := endpoint.post(&subnetPostObject {
		Type: "EntityPrivateSubnet",
		Id: subnet.Id,
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
		return nil, errors.New("Could not get updated subnet.")
	}
	for _,match := range matches {
		if match.Id == subnet.Id {
			return match, nil
		}
	}
	return nil, errors.New("Could not find updated subnet.")
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

type subnetPostObject struct {
	Type string `json:"_type"`
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ActualSubnet string `json:"actualSubnet"`
	GwId string `json:"gw.id"`
	Source string `json:"source"`
	Subnet string `json:"subnet"`
	NatDisabled bool `json:"sNatDisabled"`
	SubnetRoleId string `json:"subnetRole.id"`
}
type privateSubnetPostResult struct {
	Result string `json:"result"`
	Total int `json:"total"`
}
func (endpoint *PrivateSubnetsEndpoint) post(subnet *subnetPostObject) error {
	var result privateSubnetPostResult
	_, err := endpoint.service.Post(endpoint.url+"/"+strconv.Itoa(subnet.Id), subnet, &result)
	if err != nil {
		return err
	}
	if result.Result != "OK" {
		return errors.New(fmt.Sprintf("Failed subnet post: %s", result.Result))
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