package fortycloud

import (
	"net/http"
	"github.com/mdl/fortycloud-sdk-go/internal"
	"github.com/mdl/fortycloud-sdk-go/forms"
	"github.com/mdl/fortycloud-sdk-go/rest"
)

type Api struct {
	session1 *rest.Session
	session2 *forms.Session
	Tokens *rest.TokensEndpoint
	Scripts *rest.ScriptsEndpoint
	Servers *rest.ServersEndpoint
}

func NewApi(url string, formUrl string) *Api {
	ap := new(Api)
	configureRestApi(ap, url)
	configureFormsApi(ap, formUrl)
	return ap
}

func (ap *Api) SetCredentials(username string, password string, tenantName string) {
	ap.session1.Set(username, password, tenantName)
	ap.session2.Set(username, password, tenantName)
}

func configureRestApi(ap *Api, url string) {
	ap.session1 = rest.NewSession(url)
	svc := internal.NewJsonService(url, nil)
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		return ap.session1.SecureRequest(method, endpoint, req)
	})
	ap.Scripts = rest.NewScriptsEndpoint(svc)
	ap.Servers = rest.NewServersEndpoint(svc)
}

func configureFormsApi(ap *Api, url string) {
	ap.session2 = forms.NewSession(url + "/authenticate")
	svc := internal.NewFormService(url + "/api", nil)
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		return ap.session2.SecureRequest(method, endpoint, req)
	})
}