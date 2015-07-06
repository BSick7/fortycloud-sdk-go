package api

import (
	"net/http"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type Api struct {
	Auth *Authentication
	Tokens *TokensEndpoint
	Scripts *ScriptsEndpoint
	Servers *ServersEndpoint
}

func NewApi(url string, formUrl string) *Api {
	ap := new(Api)
	ap.Auth = NewAuthentication(ap)
	ap.Tokens = NewTokensEndpoint(internal.NewJsonService(url, nil))
	
	svc := internal.NewJsonService(url, nil)
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		return ap.Auth.SecureRequest(method, endpoint, req)
	})
	ap.Scripts = NewScriptsEndpoint(svc)
	ap.Servers = NewServersEndpoint(svc)
	
	//formSvc := internal.NewJsonService(formUrl, nil)
	
	return ap
}