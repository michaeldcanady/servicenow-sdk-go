package credentials

import (
	"net/http"
	"time"
)

// clientOptions contains configuration for both public and confidential clients.
type clientOptions struct {
	httpClient *http.Client
}

// clientOption is a functional option for configuring clients.
type clientOption func(*clientOptions)

// withHTTPClient sets a custom HTTP client.
func withHTTPClient(client *http.Client) clientOption {
	return func(o *clientOptions) {
		o.httpClient = client
	}
}

func defaultOptions() clientOptions {
	return clientOptions{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
