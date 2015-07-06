package api

import (
    "time"
    "net/http"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type Authentication struct {
	api *Api
    username string
    password string
    tenantName string
    token string
    expires Time
}

func NewAuthentication(api *Api) *Authentication {
    return &Authentication{
        api: Api,
    }
}

func (auth *Authentication) Set(username string, password string, tenantName string) {
    auth.username = username
    auth.password = password
    auth.tenantName = tenantName
}

func (auth *Authentication) Ensure() error {
    if auth.isAuthenticated() {
        return nil
    }
    
    result, err := api.Tokens.Post(auth.username, auth.password, auth.tenantName)
    if err != nil {
        return err
    }
    
    fault := result.IdentityFault
    if len(fault.Code) > 0 {
        return errors.New(fmt.Sprintf("[%s] '%s': %s", fault.Code, fault.Message, fault.Details))   
    }
    
    auth.token = result.Access.Token.Id
    auth.expires = time.Parse(time.RFC3339, result.Access.Token.Expires)
    return nil
}

func (auth *Authentication) SecureRequest(method string, endpoint string, req *http.Request) error {
	if err := auth.Ensure(); err != nil {
        return err
    }
    req.Header.Add("X-Auth-Token", auth.token)
	return nil
}

func (auth *Authentication) isAuthenticated() bool {
    if len(auth.token) > 0 {
        //TODO: check expiration
        return true
    }
    return false
}