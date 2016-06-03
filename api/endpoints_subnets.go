package api

import (
	"fmt"
	"github.com/mdl/fortycloud-sdk-go/internal"
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
		return nil, err
	}
	return &res.Subnet, nil
}

func (endpoint *SubnetsEndpoint) Create(subnet *Subnet) (*Subnet, error) {
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

func (endpoint *SubnetsEndpoint) AssignResourceGroup(subnetId string, resourceGroupId string) (*Subnet, error) {
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
