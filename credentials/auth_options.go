package credentials

import (
	"context"
	"fmt"
	"net/http"
)

// AuthorizationCodeServer is the interface for the local callback server used in Authorization Code flow.
type AuthorizationCodeServer interface {
	GetAddr() string
	Result(ctx context.Context) (code string, state string, err error)
	Shutdown(ctx context.Context) error
}

// ServerFactory is a function that creates a new AuthorizationCodeServer.
type ServerFactory func(state string, port int) (AuthorizationCodeServer, error)

// AuthConfig contains options shared by all authentication providers.
type AuthConfig struct {
	baseURL        string
	allowedHosts   []string
	tokenStore     TokenStore
	httpClient     *http.Client
	scopes         []string
	port           int
	stateGenerator func() string
	urlOpener      func(string) error
	serverFactory  ServerFactory
}

// AuthOption is a function type that modifies the AuthConfig.
type AuthOption func(*AuthConfig)

// --- Shared Options ---

const defaultBaseURL = "service-now.com"

// WithInstance sets the ServiceNow instance name (e.g., "dev12345").
func WithInstance(instance string) AuthOption {
	return WithURL(fmt.Sprintf("https://%s.%s", instance, defaultBaseURL))
}

// WithURL sets the full base URL for the ServiceNow instance.
func WithURL(url string) AuthOption {
	return func(c *AuthConfig) {
		c.baseURL = url
	}
}

// WithAllowedHosts sets the allowed hosts for the provider.
func WithAllowedHosts(hosts ...string) AuthOption {
	return func(c *AuthConfig) {
		c.allowedHosts = hosts
	}
}

// WithTokenStore sets the token store for the provider.
func WithTokenStore(store TokenStore) AuthOption {
	return func(c *AuthConfig) {
		c.tokenStore = store
	}
}

// WithHTTPClient sets the HTTP client for the provider.
func WithHTTPClient(client *http.Client) AuthOption {
	return func(c *AuthConfig) {
		c.httpClient = client
	}
}

// WithScopes sets the OAuth2 scopes.
func WithScopes(scopes ...string) AuthOption {
	return func(c *AuthConfig) {
		c.scopes = scopes
	}
}

// --- Auth Code Specific Options ---

// WithPort sets the port for the local callback server.
func WithPort(port int) AuthOption {
	return func(c *AuthConfig) {
		c.port = port
	}
}

// WithStateGenerator sets the state generator function.
func WithStateGenerator(generator func() string) AuthOption {
	return func(c *AuthConfig) {
		c.stateGenerator = generator
	}
}

// WithURLOpener sets the function to open the authorization URL.
func WithURLOpener(opener func(string) error) AuthOption {
	return func(c *AuthConfig) {
		c.urlOpener = opener
	}
}

// WithServerFactory sets the factory function for creating the local callback server.
func WithServerFactory(factory ServerFactory) AuthOption {
	return func(c *AuthConfig) {
		c.serverFactory = factory
	}
}
