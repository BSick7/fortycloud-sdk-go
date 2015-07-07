package forms

type EntityAllResult struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Result string `json:"result"`
	Total int `json:"total"`
	Type string `json:"type"`
}