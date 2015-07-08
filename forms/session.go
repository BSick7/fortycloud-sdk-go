package forms

import (
	"errors"
	"net/http"
	"strconv"
)

type Session struct {
	initiator         *FormInitiator
	authenticator     *FormAuthenticator
	username          string
	password          string
	authenticityToken string
	csrfToken         string
	userId            int
	accountId         int
}

func NewSession(url string, client *http.Client) *Session {
	session := &Session{
		userId:    -1,
		accountId: -1,
	}
	session.initiator = NewFormInitiator(url, client)
	session.authenticator = NewFormAuthenticator(url, client)
	return session
}

func (session *Session) Set(username string, password string) {
	session.username = username
	session.password = password
}

func (session *Session) SecureRequest(method string, endpoint string, req *http.Request) error {
	if err := session.ensure(); err != nil {
		return err
	}
	req.Header.Set("X-CSRF-Token", session.csrfToken)
	if session.accountId > -1 {
		values := req.URL.Query()
		values.Add("account", strconv.Itoa(session.accountId))
		req.URL.RawQuery = values.Encode()
	}
	return nil
}

func (session *Session) ensure() error {
	if session.isAuthenticated() {
		return nil
	}

	session.userId = -1
	session.accountId = -1
	result, err := session.initiator.Initiate()
	if err != nil {
		return err
	}
	session.authenticityToken = result.AuthenticityToken
	session.csrfToken = result.CsrfToken

	result2, err := session.authenticator.Authenticate(session.username, session.password, session.authenticityToken)
	if err != nil {
		return err
	}
	session.userId = result2.UserId
	session.accountId = result2.AccountId
	
	if result2.AccountId < 0 {
		return errors.New("Could not authenticate.")
	}
	return nil
}

func (session *Session) isAuthenticated() bool {
	return session.userId > -1
}
