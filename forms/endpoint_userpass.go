package forms

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type UserPassEndpoint struct {
	service *internal.JsonService
}

func NewUserPassEndpoint(service *internal.JsonService) *UserPassEndpoint {
	return &UserPassEndpoint {
		service: service,
	}
}