package api

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type ServersEndpoint struct {
	service *internal.JsonService
	url string
}

type Server struct {
	ResourceGroupName string `json:"resourceGroupName"`
	PublicIP string `json:"publicIP"`
	OverlayAddress string `json:"overlayAddress"`
	Region string `json:"region"`
	Enable bool `json:"enable"`
	AllowSSHToEveryone bool `json:"allowSSHToEveryone"`
	PermitRules string `json:"permitRules"`
	RouteAllTrafficViaGW bool `json:"routeAllTrafficViaGW"`
	CloudAccount string `json:"cloudAccount"`
	IsGW bool `json:"isGW"`
	PrivateIP string `json:"privateIP"`
	ForcePublicIP bool `json:"forcePublicIP"`
	IdentityServerName string `json:"identityServerName"`
	State string `json:"state"`
	Name string `json:"name"`
	Description string `json:"description"`
	Id string `json:"id"`
}

type serversAllResult struct {
	Servers []Server `json:"servers"`	
}

func NewServersEndpoint(service *internal.JsonService) *ServersEndpoint {
    return &ServersEndpoint {
        service: service,
		url: "/servers",
    }
}

func (endpoint *ServersEndpoint) All() ([]Server, error) {
	var result serversAllResult
	err := endpoint.service.Get(endpoint.url, &result)
	if err != nil {
		return nil, err
	}
	return result.Servers, nil
}