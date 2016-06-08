package api

import (
	"encoding/json"
	"github.com/BSick7/fortycloud-sdk-go/internal"
)

type badRequestResponse struct {
	BadRequest struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"badRequest"`
}

func IsErrorObjectNotExists(err error) bool {
	re, ok := err.(*internal.ResponseError)
	if !ok {
		return false
	}
	if re.Code != 400 {
		return false
	}
	var body badRequestResponse
	if err := json.Unmarshal([]byte(re.Body), &body); err != nil {
		return false
	}
	if body.BadRequest.Details == "Object does not exist" {
		return true
	}
	return false
}
