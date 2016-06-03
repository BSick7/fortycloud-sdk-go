package api

import (
	"fmt"
	"github.com/BSick7/fortycloud-sdk-go/internal"
)

type IpAddressSetsEndpoint struct {
	service *internal.JsonService
	url     string
}
type IpAddressSet struct {
	Id            string `json:"id,omitempty"`
	Name          string `json:"name"`
	Ips           string `json:"ips"`
	Active        bool   `json:"active"`
	ResourceGroup string `json:"resourceGroup"`
}

func NewIpAddressSetsEndpoint(service *internal.JsonService) *IpAddressSetsEndpoint {
	return &IpAddressSetsEndpoint{
		service: service,
		url:     "/ip-address-sets",
	}
}

func (endpoint *IpAddressSetsEndpoint) All() ([]IpAddressSet, error) {
	type body struct {
		IpAddressSet IpAddressSet `json:"iPAddressSet"`
	}
	type result struct {
		IpAddressSets []IpAddressSet `json:"iPAddressSets"`
	}
	var res result
	_, err := endpoint.service.Get(endpoint.url, &res)
	if err != nil {
		return nil, err
	}
	return res.IpAddressSets, nil
}

func (endpoint *IpAddressSetsEndpoint) Get(id string) (*IpAddressSet, error) {
	type result struct {
		IpAddressSet IpAddressSet `json:"iPAddressSet"`
	}
	var res result
	_, err := endpoint.service.Get(fmt.Sprintf("%s/%s", endpoint.url, id), &res)
	if err != nil {
		return nil, err
	}
	return &res.IpAddressSet, nil
}

func (endpoint *IpAddressSetsEndpoint) Create(set *IpAddressSet) (*IpAddressSet, error) {
	type body struct {
		IpAddressSet IpAddressSet `json:"iPAddressSet"`
	}
	type result struct {
		IpAddressSet IpAddressSet `json:"iPAddressSet"`
	}
	var res result
	_, err := endpoint.service.Post(endpoint.url, &body{IpAddressSet: *set}, &res)
	if err != nil {
		return nil, err
	}
	return &res.IpAddressSet, nil
}

func (endpoint *IpAddressSetsEndpoint) Update(set *IpAddressSet) (*IpAddressSet, error) {
	type body struct {
		IpAddressSet IpAddressSet `json:"iPAddressSet"`
	}
	type result struct {
		IpAddressSet IpAddressSet `json:"iPAddressSet"`
	}
	var res result
	_, err := endpoint.service.Put(endpoint.url, set.Id, &body{IpAddressSet: *set}, &res)
	if err != nil {
		return nil, err
	}
	return &res.IpAddressSet, nil
}

func (endpoint *IpAddressSetsEndpoint) Delete(id string) error {
	type result struct{}
	var res result
	_, err := endpoint.service.Delete(fmt.Sprintf("%s/%s", endpoint.url, id), nil, &res)
	return err
}
