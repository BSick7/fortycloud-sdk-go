package fortycloud

import (
	"github.com/mdl/fortycloud-sdk-go/api"
)

func NewApi(url string, formUrl string) *api.Api {
	return api.NewApi(url, formUrl)
}