package forms

import (
	"net/http"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type Session struct {
	userpass *UserPassEndpoint
	username string
	password string
	tenantName string
}

func NewSession(url string) *Session {
	return &Session{
		userpass: NewUserPassEndpoint(internal.NewFormService(url, nil)),
	}
}

func (session *Session) Set(username string, password string, tenantName string) {
    session.username = username
    session.password = password
    session.tenantName = tenantName
}

func (session *Session) SecureRequest(method string, endpoint string, req *http.Request) error {
	if err := session.ensure(); err != nil {
        return err
    }
	//TODO: Add auth stuff to request
	return nil
}

func (session *Session) ensure() error {
	if session.isAuthenticated() {
        return nil
    }
    
    _, err := session.userpass.Post(session.username, session.password, session.tenantName)
    if err != nil {
        return err
    }
	
	//TODO: Handle response
	
    return nil
}

func (session *Session) isAuthenticated() bool {
	return false
}