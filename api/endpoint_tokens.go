package api

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type TokensEndpoint struct {
	service *internal.JsonService
    url string
}
type tokensRequest struct {
	Auth struct {
        Credentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        } `json:"passwordCredentials"`
        TenantName string `json:"tenantName"`   
    } `json:"auth"`
}
type TokensResult struct {
    Access struct {
        Token struct {
            Id string `json:"id"`
            Expires string `json:"expires"`
        } `json:"token"`
    } `json:"access"`
    IdentityFault struct {
        Code string `json:"code"`
        Message string `json:"message"`
        Details string `json:"details"`
    } `json:"identityFault"`
}

func NewTokensEndpoint(service *internal.JsonService) *TokensEndpoint {
    return &TokensEndpoint {
        service: service,
        url: "/tokens"
    }
}

func (endpoint *TokensEndpoint) Post(username string, password string, tenantName string) (*TokensResult, error) {
    body := new(tokensRequest)
    body.Auth.Credentials.Username = username
    body.Auth.Credentials.Password = password
    body.Auth.TenantName = tenantName
	var result TokensResult
	err := endpoint.service.Post(endpoint.url, &result, body)
	return result, err
}