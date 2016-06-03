package api

import (
	"fmt"
	"github.com/BSick7/fortycloud-sdk-go/internal"
	"net/http"
	"time"
)

type Session struct {
	key    string
	secret string
	signer *internal.RequestSigner
}

func DefaultSession() *Session {
	return &Session{
		signer: &internal.RequestSigner{},
	}
}

func (session *Session) Set(key string, secret string) {
	session.key = key
	session.secret = secret
}

func (session *Session) SignRequest(req *http.Request) error {
	dt := getDateHeader()
	req.Header.Add("Date", dt)

	signature, err := session.signer.GetSignature(req, session.secret)
	if err != nil {
		return err
	}
	auth := fmt.Sprintf("FCRestAPI AccessKey=%s SignatureType=HmacSHA256 Signature=%s", session.key, signature)
	req.Header.Add("Authorization", auth)

	return nil
}

func getDateHeader() string {
	return fmt.Sprintf("%s GMT", time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05"))
}
