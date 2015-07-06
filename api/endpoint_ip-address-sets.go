package api

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type IpAddressSetsEndpoint struct {
	service *internal.JsonService
	url string
}
type IpAddressSet struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Ips string `json:"ips"`
	Active bool `json:"active"`
	ResourceGroup string `json:"resourceGroup"` 
}

func NewIpAddressSetsEndpoint(service *internal.JsonService) *IpAddressSetsEndpoint {
    return &IpAddressSetsEndpoint {
        service: service,
		url: "/ip-address-sets",
    }
}

type ipAddressSetsAllResult struct {
	IpAddressSets []IpAddressSet `json:"iPAddressSets"`
}
func (endpoint *IpAddressSetsEndpoint) All() ([]IpAddressSet, error) {
	var result ipAddressSetsAllResult
	err := endpoint.service.Get(endpoint.url, &result)
	if err != nil {
		return nil, err
	}
	return result.IpAddressSets, nil
}

type ipAddressSetsGetResult struct {
	IpAddressSet IpAddressSet `json:"iPAddressSet"`
}
func (endpoint *IpAddressSetsEndpoint) Get(id string) (*IpAddressSet, error) {
	var result ipAddressSetsGetResult
	err := endpoint.service.Get(endpoint.url + "/" + id, &result)
	if err != nil {
		return nil, err
	}
	return result.IpAddressSet, nil
}

type ipAddressSetsPostResult struct {
	IpAddressSet IpAddressSet `json:"iPAddressSet"`
}
func (endpoint *IpAddressSetsEndpoint) Post(set *IpAddressSet) (*IpAddressSet, error) {
	var result ipAddressSetsPostResult
	err := endpoint.service.Post(endpoint.url, set, &result)
	if err != nil {
		return nil, err
	}
	return result.IpAddressSet, nil
}

type ipAddressSetsPutResult struct {
	IpAddressSet IpAddressSet `json:"iPAddressSet"`
}
func (endpoint *IpAddressSetsEndpoint) Put(set *IpAddressSet) (*IpAddressSet, error) {
	var result ipAddressSetsPutResult
	err := endpoint.service.Put(endpoint.url, set.Id, set, &result)
	if err != nil {
		return nil, err
	}
	return result.IpAddressSet, nil
}

func (endpoint *IpAddressSetsEndpoint) Delete(id string) error {
	return endpoint.service.Delete(endpoint.url, id)
}