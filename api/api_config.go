package api

import "os"

const (
	DEFAULT_API_URL = "https://api.fortycloud.net/restapi/v0.4"
)

type ApiConfig struct {
	URL       string
	AccessKey string
	SecretKey string
}

func DefaultApiConfig() *ApiConfig {
	return &ApiConfig{
		URL:       DEFAULT_API_URL,
		AccessKey: os.Getenv("FORTYCLOUD_ACCESS_KEY"),
		SecretKey: os.Getenv("FORTYCLOUD_SECRET_KEY"),
	}
}
