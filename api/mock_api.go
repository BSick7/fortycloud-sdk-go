package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type MockApi struct {
	Api
	server  *httptest.Server
	handler http.HandlerFunc
}

func NewMockApi() *MockApi {
	ma := &MockApi{
		Api:     *NewApi(DefaultApiConfig()),
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}

	ma.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ma.handler(w, r)
	}))

	ma.Api.SetURL(fmt.Sprintf("%s/restapi/v0.4", ma.server.URL))

	return ma
}

func (b *MockApi) Handle(handler http.HandlerFunc) {
	b.handler = handler
}

func (b *MockApi) Close() {
	if b.server != nil {
		b.server.Close()
		b.server = nil
	}
}
