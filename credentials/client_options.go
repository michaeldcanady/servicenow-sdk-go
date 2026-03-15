package credentials

import (
	"net/http"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
)

// clientOptions contains configuration for both public and confidential clients.
type clientOptions struct {
	httpClient *http.Client
	method     pkce.Method
}

// clientOption is a functional option for configuring clients.
type clientOption func(*clientOptions)

func defaultOptions() clientOptions {
	return clientOptions{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		method: pkce.MethodUnset,
	}
}

func withPKCEChallenge(method pkce.Method) clientOption {
	return func(co *clientOptions) {
		co.method = method
	}
}
