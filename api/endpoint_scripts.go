package fortycloud

import (
	"encoding/json"
)

type ScriptsEndpoint struct {
	service *Service
	url string
}

func (s *Service) Scripts() *ScriptsEndpoint {
	return &ScriptsEndpoint{
		service: s,
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
	body, err := endpoint.service.Post(endpoint.url, &scriptRequest{
		GlobalSettingName: setting,
		IsGateway: isGateway,
	})
	if err != nil {
		return "", err
	}
	
	var result scriptResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	return result.EnrollmentScript, nil
}