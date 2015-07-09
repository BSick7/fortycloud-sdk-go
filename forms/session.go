package forms

import (
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"
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
	gate              *sync.Mutex
}

const (
	retries        = 3
	retrySleepSecs = 2
)

func NewSession(url string, client *http.Client) *Session {
	session := &Session{
		userId:    -1,
		accountId: -1,
		gate:      new(sync.Mutex),
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
	session.gate.Lock()
	err := session.ensure()
	session.gate.Unlock()
	if err != nil {
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
	for i := 0; i < retries; i++ {
		result, err := session.tryAuth()
		if err != nil {
			return err
		}
		if result {
			return nil
		}
		session.initiator.Reset()
		time.Sleep(retrySleepSecs * time.Second)
	}

	return errors.New("Could not authenticate.")
}

func (session *Session) isAuthenticated() bool {
	return session.userId > -1
}

func (session *Session) tryAuth() (bool, error) {
	iresult, err := session.initiator.Initiate()
	if err != nil {
		return false, err
	}
	session.authenticityToken = iresult.AuthenticityToken
	session.csrfToken = iresult.CsrfToken

	aresult, err2 := session.authenticator.Authenticate(session.username, session.password, session.authenticityToken)
	if err2 != nil {
		return false, err2
	}
	session.userId = aresult.UserId
	session.accountId = aresult.AccountId
	return aresult.AccountId > -1, nil
}
