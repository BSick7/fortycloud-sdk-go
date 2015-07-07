package forms

import (
	"net/http"
	"strconv"
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type Session struct {
	Cookies *internal.CookieContainer
	initiator *FormInitiator
	userpass *UserPassEndpoint
	admin *FormAdminCollector
	username string
	password string
	tenantName string
	authenticityToken string
	authenticated bool
	userId int
	accountId int
}

func NewSession(url string) *Session {
	session := new(Session)
	session.Cookies = new(internal.CookieContainer)
	session.accountId = -1
	svc := internal.NewFormService(url, nil)
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		session.Cookies.AddToRequest(req)
		return nil
	})
	svc.InjectRequest(func(method string, endpoint string, req *http.Request) error {
		if endpoint == "/authenticate/userpass" {
			req.Header.Set("Referer", url + "/login")
		}
		return nil
	})
	svc.InjectResponse(func(method string, endpoint string, res *http.Response) error {
		session.Cookies.Merge(res.Cookies())
		return nil
	})
	session.initiator = NewFormInitiator(url, nil)
	session.userpass = NewUserPassEndpoint(svc)
	session.admin = NewFormAdminCollector(url, nil)
	return session
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
	session.Cookies.AddToRequest(req)
	if (session.accountId > -1) {
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
	result, err := session.initiator.Initiate(session.Cookies)
	if err != nil {
		return err
	}
	session.authenticityToken = result.AuthenticityToken
    
    err = session.userpass.Post(session.username, session.password, session.tenantName, session.authenticityToken)
    if err != nil {
        return err
    }
	cookie := session.Cookies.Get("FORTYCLOUD_FLASH")
	session.authenticated = cookie != nil && cookie.Value == "%00requestAfterLogin%3Atrue%00"
	log.Println("Form Authentication status: ", session.authenticated)
	
	result3, err := session.admin.Collect(session.Cookies)
	if err != nil {
		return err
	}
	session.userId = result3.UserId
	session.accountId = result3.AccountId
    return nil
}

func (session *Session) isAuthenticated() bool {
	return session.authenticated
}