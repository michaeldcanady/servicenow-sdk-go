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

func defaultOptions() clientOptions {
	return clientOptions{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}
