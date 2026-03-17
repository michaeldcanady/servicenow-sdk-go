package credentials

import "net/http"

type authConfig struct {
	instance     string
	baseURL      string
	scopes       []string
	allowedHosts []string
	tokenStore   TokenStore
	httpClient   *http.Client
}

func defaultAuthConfig() *authConfig {
	return &authConfig{
		httpClient: http.DefaultClient,
	}
}

// AuthOption is a functional option for configuring authentication providers.
type AuthOption func(*authConfig)

// WithInstance sets the ServiceNow instance name (e.g., "dev12345").
func WithInstance(instance string) AuthOption {
	return func(c *authConfig) {
		c.instance = instance
	}
}

// WithURL sets the full base URL for the ServiceNow instance.
func WithURL(url string) AuthOption {
	return func(c *authConfig) {
		c.baseURL = url
	}
}

// WithScopes sets the OAuth2 scopes.
func WithScopes(scopes ...string) AuthOption {
	return func(c *authConfig) {
		c.scopes = scopes
	}
}

// WithAllowedHosts sets the allowed hosts for the provider.
func WithAllowedHosts(hosts ...string) AuthOption {
	return func(c *authConfig) {
		c.allowedHosts = hosts
	}
}

// WithTokenStore sets the token store for the provider.
func WithTokenStore(store TokenStore) AuthOption {
	return func(c *authConfig) {
		c.tokenStore = store
	}
}

// WithHTTPClient sets the HTTP client for the provider.
func WithHTTPClient(client *http.Client) AuthOption {
	return func(c *authConfig) {
		c.httpClient = client
	}
}
