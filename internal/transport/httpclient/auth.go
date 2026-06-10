package httpclient

//authentication middleware

import (
	"context"
	"net/http"
)

func (c *Client) NewRequest(
	ctx context.Context,
	method string,
	url string,
) (*http.Request, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.BaseURL+url,
		nil,
	)

	if err != nil {
		return nil, err
	}

	// Sesuaikan nanti dengan auth resmi API tiket.com
	req.Header.Set("Content-Type", "application/json")

	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}

	return req, nil
}