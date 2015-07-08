package rest

import (
	"errors"
	"fmt"
	"github.com/mdl/fortycloud-sdk-go/internal"
	"log"
	"net/http"
	"time"
)

type Session struct {
	tokens     *TokensEndpoint
	username   string
	password   string
	tenantName string
	token      string
	expires    time.Time
}

func NewSession(url string) *Session {
	return &Session{
		tokens: NewTokensEndpoint(internal.NewJsonService(url, nil)),
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
	req.Header.Add("X-Auth-Token", session.token)
	return nil
}

func (session *Session) ensure() error {
	if session.isAuthenticated() {
		return nil
	}

	result, err := session.tokens.Post(session.username, session.password, session.tenantName)
	if err != nil {
		return err
	}

	fault := result.IdentityFault
	if len(fault.Code) > 0 {
		return errors.New(fmt.Sprintf("[%s] '%s': %s", fault.Code, fault.Message, fault.Details))
	}

	session.token = result.Access.Token.Id
	session.expires, err = time.Parse("2006-01-02T15:04:05-0700", result.Access.Token.Expires)
	if err != nil {
		log.Println("Could not parse auth expiration. ", err)
	}
	return nil
}

func (session *Session) isAuthenticated() bool {
	if len(session.token) <= 0 {
		log.Println("No auth token")
		return false
	}
	if session.expires.Sub(time.Now()) < 0 {
		log.Println("Expired token")
		return false
	}
	return true
}
