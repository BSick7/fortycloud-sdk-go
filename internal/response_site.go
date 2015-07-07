package internal

import (
	"net/http"
)

type ResponseSite func(method string, endpoint string, res *http.Response) error