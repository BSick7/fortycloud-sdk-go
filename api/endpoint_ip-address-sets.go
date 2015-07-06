package fortycloud

import (
	"encoding/json"
)

type IpAddressSetsEndpoint struct {
	service *Service
	url string
}

func (s *Service) IpAddressSets() *IpAddressSetsEndpoint {
	return &IpAddressSetsEndpoint{
		service: s,
		url: "/ip-address-sets",
	}
}

type IpAddressSet struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Ips string `json:"ips"`
	Active bool `json:"active"`
	ResourceGroup string `json:"resourceGroup"` 
}

type ipAddressSetsAllResult struct {
	IpAddressSets []IpAddressSet `json:"iPAddressSets"`
}
func (endpoint *IpAddressSetsEndpoint) All() ([]IpAddressSet, error) {
	body, err := endpoint.service.Get(endpoint.url)
	if err != nil {
		return nil, err
	}
	
	var result ipAddressSetsAllResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	
	return result.IpAddressSets, nil
}

type ipAddressSetsGetResult struct {
	IpAddressSet IpAddressSet `json:"iPAddressSet"`
}
func (endpoint *IpAddressSetsEndpoint) Get(id string) (*IpAddressSet, error) {
	body, err := endpoint.service.Get(endpoint.url + "/" + id)
	if err != nil {
		return nil, err
	}
	
	var result ipAddressSetsGetResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	
	return &result.IpAddressSet, nil
}

type ipAddressSetsPostResult struct {
	IpAddressSet IpAddressSet `json:"iPAddressSet"`
}
func (endpoint *IpAddressSetsEndpoint) Post(set *IpAddressSet) (*IpAddressSet, error) {
	body, err := endpoint.service.Post(endpoint.url, set)
	if err != nil {
		return nil, err
	}
	
	var result ipAddressSetsPostResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	
	return &result.IpAddressSet, nil
}

type ipAddressSetsPutResult struct {
	IpAddressSet IpAddressSet `json:"iPAddressSet"`
}
func (endpoint *IpAddressSetsEndpoint) Put(set *IpAddressSet) (*IpAddressSet, error) {
	body, err := endpoint.service.Put(endpoint.url, set.Id, set)
	if err != nil {
		return nil, err
	}
	
	var result ipAddressSetsPutResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	
	return &result.IpAddressSet, nil
}

func (endpoint *IpAddressSetsEndpoint) Delete(id string) error {
	return endpoint.service.Delete(endpoint.url, id)
}