package api

import (
	"github.com/mdl/fortycloud-sdk-go/internal"
)

type ScriptsEndpoint struct {
	service *internal.JsonService
    url string
}

func NewScriptsEndpoint(service *internal.JsonService) *ScriptsEndpoint {
    return &ScriptsEndpoint {
        service: service,
		url: "/enrollment-scripts",
    }
}

type scriptRequest struct {
	GlobalSettingName string `json:"globalSettingName"`
	IsGateway bool `json:"isGateway"`
}
type scriptResponse struct {
	EnrollmentScript string `json:"enrollmentScript"`
}
func (endpoint *ScriptsEndpoint) Get(setting string, isGateway bool) (string, error) {
	var result scriptResponse
	err := endpoint.service.Post(endpoint.url, &scriptRequest{
		GlobalSettingName: setting,
		IsGateway: isGateway,
	}, &result)
	if err != nil {
		return "", err
	}
	return result.EnrollmentScript, nil
}