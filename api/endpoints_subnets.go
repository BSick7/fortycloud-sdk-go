package api

import (
	"fmt"
	"github.com/BSick7/fortycloud-sdk-go/internal"
	"strings"
)

type Subnet struct {
	Id               string `json:"id,omitempty"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Cidr             string `json:"cidr"`
	DisableAutoNAT   bool   `json:"disableAutoNAT"`
	GatewayRef       string `json:"gatewayRef,omitempty"`
	ResourceGroupRef string `json:"resourceGroupRef,omitempty"`
}

func (s *Subnet) GatewayID() string {
	// expected GatewayRef formats:
	//   - https://api.fortycloud.net/restapi/v0.4/gateways/4013
	//   - https://api.fortycloud.net/servers/46859
	tokens := strings.Split(s.GatewayRef, "/")
	if len(tokens) < 2 {
		return ""
	}
	prevToken := tokens[len(tokens)-2]
	if prevToken == "gateways" || prevToken == "servers" {
		return tokens[len(tokens)-1]
	}
	return ""
}

func (s *Subnet) SetGatewayID(id string) {
	if id == "" {
		s.GatewayRef = ""
		return
	}
	s.GatewayRef = fmt.Sprintf("https://api.fortycloud.net/restapi/v0.4/gateways/%s", id)
}

func (s *Subnet) ResourceGroupID() string {
	// expected ResourceGroupRef format: https://api.fortycloud.net/restapi/v0.4/resource-groups/2069
	tokens := strings.Split(s.ResourceGroupRef, "/")
	if len(tokens) < 2 || tokens[len(tokens)-2] != "resource-groups" {
		return ""
	}
	return tokens[len(tokens)-1]
}

func (s *Subnet) SetResourceGroupID(id string) {
	if id == "" {
		s.ResourceGroupRef = ""
		return
	}
	s.ResourceGroupRef = fmt.Sprintf("https://api.fortycloud.net/restapi/v0.4/resource-groups/%s", id)
}

type SubnetsEndpoint struct {
	service *internal.JsonService
	url     string
}

func NewSubnetsEndpoint(service *internal.JsonService) *SubnetsEndpoint {
	return &SubnetsEndpoint{
		service: service,
		url:     "/subnets",
	}
}

func (endpoint *SubnetsEndpoint) All() ([]Subnet, error) {
	type result struct {
		Subnets []Subnet `json:"subnets"`
	}
	var res result
	_, err := endpoint.service.Get(endpoint.url, &res)
	if err != nil {
		return nil, err
	}
	return res.Subnets, nil
}

func (endpoint *SubnetsEndpoint) Get(id string) (*Subnet, error) {
	type result struct {
		Subnet Subnet `json:"subnet"`
	}
	var res result
	_, err := endpoint.service.Get(fmt.Sprintf("%s/%s", endpoint.url, id), &res)
	if err != nil {
		if IsErrorObjectNotExists(err) {
			return nil, nil
		}
		return nil, err
	}
	return &res.Subnet, nil
}

func (endpoint *SubnetsEndpoint) Create(subnet *Subnet) (*Subnet, error) {
	copy := *subnet
	// These fields can't be included in the request body
	copy.GatewayRef = ""
	copy.ResourceGroupRef = ""

	type body struct {
		Subnet Subnet `json:"subnet"`
	}
	type result struct {
		Subnet Subnet `json:"subnet"`
	}
	var res result
	_, err := endpoint.service.Post(endpoint.url, &body{Subnet: *subnet}, &res)
	if err != nil {
		return nil, err
	}
	return &res.Subnet, nil
}

func (endpoint *SubnetsEndpoint) Update(id string, subnet *Subnet) (*Subnet, error) {
	copy := *subnet
	// These fields can't be included in the request body
	copy.GatewayRef = ""
	copy.ResourceGroupRef = ""

	type body struct {
		Subnet Subnet `json:"subnet"`
	}
	type result struct {
		Subnet Subnet `json:"subnet"`
	}
	var res result
	_, err := endpoint.service.Put(endpoint.url, id, &body{Subnet: *subnet}, &res)
	if err != nil {
		return nil, err
	}
	return &res.Subnet, nil
}

func (endpoint *SubnetsEndpoint) Delete(id string) error {
	type result struct{}
	var res result
	_, err := endpoint.service.Delete(fmt.Sprintf("%s/%s", endpoint.url, id), nil, &res)
	return err
}

func (endpoint *SubnetsEndpoint) AssignGateway(subnetId string, gatewayId string) (*Subnet, error) {
	if gatewayId == "" {
		return endpoint.ClearGateway(subnetId)
	}
	type result struct {
		Subnet Subnet `json:"subnet"`
	}
	var res result
	_, err := endpoint.service.Put(fmt.Sprintf("%s/%s/gateways/%s", endpoint.url, subnetId, gatewayId), "", "", &res)
	if err != nil {
		return nil, err
	}
	return &res.Subnet, nil
}

func (endpoint *SubnetsEndpoint) ClearGateway(subnetId string) (*Subnet, error) {
	type result struct {
		Subnet Subnet `json:"subnet"`
	}
	var res result
	_, err := endpoint.service.Put(fmt.Sprintf("%s/%s/gateways", endpoint.url, subnetId), "", "", &res)
	if err != nil {
		return nil, err
	}
	return &res.Subnet, nil
}

func (endpoint *SubnetsEndpoint) AssignResourceGroup(subnetId string, resourceGroupId string) (*Subnet, error) {
	if resourceGroupId == "" {
		return nil, nil
	}
	type result struct {
		Subnet Subnet `json:"subnet"`
	}
	var res result
	_, err := endpoint.service.Put(fmt.Sprintf("%s/%s/resource-groups/%s", endpoint.url, subnetId, resourceGroupId), "", "", &res)
	if err != nil {
		return nil, err
	}
	return &res.Subnet, nil
}
