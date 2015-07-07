package rest

import (
    "log"
    "time"
    "net/http"
    "errors"
    "fmt"
)

type Authentication struct {
    tokens *TokensEndpoint
    username string
    password string
    tenantName string
    token string
    expires time.Time
}

func NewAuthentication(tokens *TokensEndpoint) *Authentication {
    return &Authentication{
        tokens: tokens,
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
    
    result, err := auth.tokens.Post(auth.username, auth.password, auth.tenantName)
    if err != nil {
        return err
    }
    
    fault := result.IdentityFault
    if len(fault.Code) > 0 {
        return errors.New(fmt.Sprintf("[%s] '%s': %s", fault.Code, fault.Message, fault.Details))   
    }
    
    auth.token = result.Access.Token.Id
    auth.expires, err = time.Parse("2006-01-02T15:04:05-0700", result.Access.Token.Expires)
    if err != nil {
        log.Println("Could not parse auth expiration. ", err)
    }
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
    if len(auth.token) <= 0 {
        log.Println("No auth token")
        return false
    }
    if auth.expires.Sub(time.Now()) < 0 {
        log.Println("Expired token")
        return false
    }
    return true
}