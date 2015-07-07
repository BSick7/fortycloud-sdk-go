package fortycloud

import (
	"net/http"
	"github.com/mdl/fortycloud-sdk-go/internal"
	"github.com/mdl/fortycloud-sdk-go/forms"
	"github.com/mdl/fortycloud-sdk-go/rest"
)

type Api struct {
	auth *rest.Authentication
	session *forms.Session
	Tokens *rest.TokensEndpoint
	Scripts *rest.ScriptsEndpoint
	Servers *rest.ServersEndpoint
}

func NewApi(url string, formUrl string) *Api {
	ap := new(Api)
	ap.Tokens = rest.NewTokensEndpoint(internal.NewJsonService(url, nil))
	ap.auth = rest.NewAuthentication(ap.Tokens)
	
	svc := internal.NewJsonService(url, nil)
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		return ap.auth.SecureRequest(method, endpoint, req)
	})
	ap.Scripts = rest.NewScriptsEndpoint(svc)
	ap.Servers = rest.NewServersEndpoint(svc)
	
	
	ap.session = forms.NewSession()
	formSvc := internal.NewJsonService(formUrl, nil)
	formSvc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		return ap.session.SecureRequest(method, endpoint, req)
	})
	
	return ap
}

func (ap *Api) SetCredentials(username string, password string, tenantName string) {
	ap.auth.Set(username, password, tenantName)
	ap.session.SetCredentials(username, password)
}