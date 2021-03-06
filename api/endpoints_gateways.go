package api

import (
	"fmt"
	"github.com/BSick7/fortycloud-sdk-go/internal"
)

type Gateway struct {
	Fqdn                 string   `json:"fqdn"`
	PublicIP             string   `json:"publicIP"`
	OverlayAddress       string   `json:"overlayAddress"`
	VpnUsersSubnet       string   `json:"vpnUsersSubnet"`
	Region               string   `json:"region"`
	Enable               bool     `json:"enable"`
	AllowSSHToEveryone   bool     `json:"allowSSHToEveryone"`
	RouteAllTrafficViaGW bool     `json:"routeAllTrafficViaGW"`
	PrivateIP            string   `json:"privateIP"`
	State                string   `json:"state"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	SecurityGroup        string   `json:"securityGroup"`
	OpenVPNProtocol      string   `json:"openVPNProtocol"`
	ClientRoutes         []string `json:"clientRoutes"`
	GatewayAsDNS         bool     `json:"gatewayAsDNS"`
	DirectRoutesOnly     bool     `json:"directRoutesOnly"`
	HaState              string   `json:"haState"`
	Id                   string   `json:"id,omitempty"`
	Release              string   `json:"release,omitempty"`
	IdentityServerName   string   `json:"identityServerName,omitempty"`
}

type GatewaysEndpoint struct {
	service *internal.JsonService
	url     string
}

func NewGatewaysEndpoint(service *internal.JsonService) *GatewaysEndpoint {
	return &GatewaysEndpoint{
		service: service,
		url:     "/gateways",
	}
}

func (endpoint *GatewaysEndpoint) All() ([]Gateway, error) {
	type result struct {
		Gateways []Gateway `json:"gateways"`
	}
	var res result
	_, err := endpoint.service.Get(endpoint.url, &res)
	if err != nil {
		return nil, err
	}
	return res.Gateways, nil
}

func (endpoint *GatewaysEndpoint) Get(id string) (*Gateway, error) {
	type result struct {
		Gateway Gateway `json:"gateway"`
	}
	var res result
	_, err := endpoint.service.Get(fmt.Sprintf("%s/%s", endpoint.url, id), &res)
	if err != nil {
		if IsErrorObjectNotExists(err) {
			return nil, nil
		}
		return nil, err
	}
	return &res.Gateway, nil
}

func (endpoint *GatewaysEndpoint) GetInstallationScript(platform string, license string) (string, error) {
	type body struct {
		InstallationPlatform string `json:"installationPlatform"`
		LicenseName          string `json:"licenseName"`
	}
	type result struct {
		InstallationScript string `json:"installationScript"`
	}
	var res result
	_, err := endpoint.service.Post("/installation-scripts", &body{InstallationPlatform: platform, LicenseName: license}, &res)
	if err != nil {
		return "", err
	}
	return res.InstallationScript, nil
}

func (endpoint *GatewaysEndpoint) GetRegistrationToken(platform string, license string) (string, error) {
	type body struct {
		InstallationPlatform string `json:"installationPlatform"`
		LicenseName          string `json:"licenseName"`
	}
	type result struct {
		RegistrationToken string `json:"registrationToken"`
	}
	var res result
	_, err := endpoint.service.Post("/registration-tokens", &body{InstallationPlatform: platform, LicenseName: license}, &res)
	if err != nil {
		return "", err
	}
	return res.RegistrationToken, nil
}

func (endpoint *GatewaysEndpoint) Update(id string, gateway *Gateway) (*Gateway, error) {
	copy := *gateway
	// These fields can't be included in the request body
	copy.IdentityServerName = ""
	copy.Release = ""

	type body struct {
		Gateway Gateway `json:"gateway"`
	}
	type result struct {
		Gateway Gateway `json:"gateway"`
	}
	var res result
	_, err := endpoint.service.Put(endpoint.url, id, &body{Gateway: copy}, &res)
	if err != nil {
		return nil, err
	}
	return &res.Gateway, nil
}

func (endpoint *GatewaysEndpoint) Delete(id string) error {
	type result struct{}
	var res result
	_, err := endpoint.service.Delete(fmt.Sprintf("%s/%s", endpoint.url, id), nil, &res)
	return err
}
