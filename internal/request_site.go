package internal

import (
	"net/http"
)

type RequestSite func(method string, endpoint string, req *http.Request) error
