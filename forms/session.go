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
	session := new(Session)
	session.userpass = NewUserPassEndpoint(internal.NewJsonService(url, nil))
	return session
}

func (session *Session) Set(username string, password string, tenantName string) {
    session.username = username
    session.password = password
    session.tenantName = tenantName
}

func (session *Session) SecureRequest(method string, endpoint string, req *http.Request) error {
	return nil
}