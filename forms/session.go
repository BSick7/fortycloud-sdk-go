package forms

import (
	"net/http"
)

type Session struct {
}

func NewSession() *Session {
	return &Session{
	}
}

func (session *Session) SetCredentials(username string, password string) {
	
}

func (session *Session) SecureRequest(method string, endpoint string, req *http.Request) error {
	return nil
}