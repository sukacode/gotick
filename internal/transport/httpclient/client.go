package httpclient

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	APIKey  string
	HTTP    *http.Client
}

func New(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		HTTP: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}