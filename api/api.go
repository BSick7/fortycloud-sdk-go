package api

import (
	"github.com/bsick7/fortycloud-sdk-go/internal"
	"net/http"
)

type Api struct {
	session          *Session
	Gateways         *GatewaysEndpoint
	IPAddressSets    *IpAddressSetsEndpoint
	IPSecConnections *IPSecConnectionsEndpoint
	Subnets          *SubnetsEndpoint
}

func NewApi(config *ApiConfig) *Api {
	if config == nil {
		config = DefaultApiConfig()
	}
	ap := &Api{
		session: DefaultSession(),
	}
	ap.SetAccessCredentials(config.AccessKey, config.SecretKey)
	ap.SetURL(config.URL)
	return ap
}

func (ap *Api) SetAccessCredentials(key string, secret string) {
	ap.session.Set(key, secret)
}

func (ap *Api) SetURL(url string) {
	svc := internal.NewJsonService(url, nil)
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		return ap.session.SignRequest(req)
	})
	ap.Gateways = NewGatewaysEndpoint(svc)
	ap.IPAddressSets = NewIpAddressSetsEndpoint(svc)
	ap.IPSecConnections = NewIPSecConnectionsEndpoint(svc)
	ap.Subnets = NewSubnetsEndpoint(svc)
}
