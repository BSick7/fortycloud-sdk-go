package fortycloud

import (
	"encoding/json"
)

type ServersEndpoint struct {
	service *Service
	url string
}

func (s *Service) Servers() *ServersEndpoint {
	return &ServersEndpoint{
		service: s,
		url: "/servers",
	}
}

type serversAllResult struct {
	Servers []Server `json:"servers"`
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
func (endpoint *ServersEndpoint) All() ([]Server, error) {
	body, err := endpoint.service.Get(endpoint.url)
	if err != nil {
		return nil, err
	}
	
	var result serversAllResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	
	return result.Servers, nil
}