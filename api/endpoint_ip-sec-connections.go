package api

import (
	"fmt"
	"github.com/BSick7/fortycloud-sdk-go/internal"
)

type IPSecConnection struct {
	Name                 string `json:"name"`
	Description          string `json:"description"`
	StateAB              string `json:"stateAB"`
	ForceIPSecOverUDP    bool   `json:"forceIPSecOverUDP"`
	Enable               bool   `json:"enable"`
	GatewayA             string `json:"gatewayA"`
	NetworkArchitectureA string `json:"networkArchitectureA"`
	GatewayB             string `json:"gatewayB"`
	NetworkArchitectureB string `json:"networkArchitectureB"`
	IKEPhaseAlgorithm    string `json:"iKEPhaseAlgorithm"`
	Pfs                  bool   `json:"pfs"`
	DiffieHellmanGroup   string `json:"diffieHellmanGroup"`
	IKESessionLifetime   string `json:"iKESessionLifetime"`
	IPSecPhaseAlgorithm  string `json:"iPSecPhaseAlgorithm"`
	IPSecSessionLifetime string `json:"iPSecSessionLifetime"`
	DPDTimeout           int    `json:"dPDTimeout"`
	DPDAction            string `json:"dPDAction"`
	Id                   string `json:"id,omitempty"`
}

type IPSecConnectionsEndpoint struct {
	service *internal.JsonService
	url     string
}

func NewIPSecConnectionsEndpoint(service *internal.JsonService) *IPSecConnectionsEndpoint {
	return &IPSecConnectionsEndpoint{
		service: service,
		url:     "/ip-sec-connections",
	}
}

func (endpoint *IPSecConnectionsEndpoint) All() ([]IPSecConnection, error) {
	type body struct {
		IPSecConnection IPSecConnection `json:"iPSecConnection"`
	}

	type result struct {
		IPSecConnections []IPSecConnection `json:"iPSecConnections"`
	}
	var res result
	_, err := endpoint.service.Get(endpoint.url, &res)
	if err != nil {
		return nil, err
	}
	return res.IPSecConnections, nil
}

func (endpoint *IPSecConnectionsEndpoint) Get(id string) (*IPSecConnection, error) {
	type result struct {
		IPSecConnection IPSecConnection `json:"iPSecConnection"`
	}
	var res result
	_, err := endpoint.service.Get(fmt.Sprintf("%s/%s", endpoint.url, id), &res)
	if err != nil {
		return nil, err
	}
	return &res.IPSecConnection, nil
}

func (endpoint *IPSecConnectionsEndpoint) Create(connection *IPSecConnection) (*IPSecConnection, error) {
	type body struct {
		IPSecConnection IPSecConnection `json:"iPSecConnection"`
	}
	type result struct {
		IPSecConnection IPSecConnection `json:"iPSecConnection"`
	}
	var res result
	_, err := endpoint.service.Post(endpoint.url, &body{IPSecConnection: *connection}, &res)
	if err != nil {
		return nil, err
	}
	return &res.IPSecConnection, nil
}

func (endpoint *IPSecConnectionsEndpoint) Update(id string, connection *IPSecConnection) (*IPSecConnection, error) {
	type body struct {
		IPSecConnection IPSecConnection `json:"iPSecConnection"`
	}
	type result struct {
		IPSecConnection IPSecConnection `json:"iPSecConnection"`
	}
	var res result
	_, err := endpoint.service.Put(endpoint.url, id, &body{IPSecConnection: *connection}, &res)
	if err != nil {
		return nil, err
	}
	return &res.IPSecConnection, nil
}

func (endpoint *IPSecConnectionsEndpoint) Delete(id string) error {
	type result struct{}
	var res result
	_, err := endpoint.service.Delete(fmt.Sprintf("%s/%s", endpoint.url, id), nil, &res)
	return err
}
