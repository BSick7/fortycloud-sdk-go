package internal

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestSigner struct {
}

func (signer *RequestSigner) GetSignature(req *http.Request, secret string) (string, error) {
	body, err := cloneBody(req)
	if err != nil {
		return "", err
	}

	//Build fingerprint
	msg := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		req.Method,
		req.Header.Get("Content-Type"),
		req.Header.Get("Date"),
		req.URL.Path,
		body)

	return computeHmac256(msg, secret), nil
}

// http://www.jokecamp.com/blog/examples-of-creating-base64-hashes-using-hmac-sha256-in-different-languages/#go
func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func cloneBody(req *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	newBody := io.Reader(bytes.NewBuffer(body))
	rc, ok := newBody.(io.ReadCloser)
	if !ok && newBody != nil {
		rc = ioutil.NopCloser(newBody)
	}
	req.Body = rc
	return
}
