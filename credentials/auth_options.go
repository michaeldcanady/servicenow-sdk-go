package credentials

import (
	"context"
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

// baseAuthConfig contains options shared by all authentication providers.
type baseAuthConfig struct {
	instance     string
	baseURL      string
	allowedHosts []string
	tokenStore   TokenStore
	httpClient   *http.Client
}

func (c *baseAuthConfig) getBase() *baseAuthConfig {
	return c
}

// oauth2Config contains options shared by all OAuth2 providers.
type oauth2Config struct {
	baseAuthConfig
	scopes []string
}

// ropcConfig contains options specific to the ROPC flow.
type ropcConfig struct {
	oauth2Config
}

// clientCredentialsConfig contains options specific to the Client Credentials flow.
type clientCredentialsConfig struct {
	oauth2Config
}

// authCodeConfig contains options specific to the Authorization Code flow.
type authCodeConfig struct {
	oauth2Config
	port           int
	stateGenerator func() string
	urlOpener      func(string) error
	serverFactory  ServerFactory
}

// jwtConfig contains options specific to the JWT flow.
type jwtConfig struct {
	oauth2Config
}

type baseConfigurable interface {
	getBase() *baseAuthConfig
}

type oauth2Configurable interface {
	baseConfigurable
	getOAuth2() *oauth2Config
}

func (c *oauth2Config) getOAuth2() *oauth2Config {
	return c
}

// --- Shared Options (Generics) ---

// WithInstance sets the ServiceNow instance name (e.g., "dev12345").
func WithInstance[T baseConfigurable](instance string) func(T) {
	return func(c T) {
		c.getBase().instance = instance
	}
}

// WithURL sets the full base URL for the ServiceNow instance.
func WithURL[T baseConfigurable](url string) func(T) {
	return func(c T) {
		c.getBase().baseURL = url
	}
}

// WithAllowedHosts sets the allowed hosts for the provider.
func WithAllowedHosts[T baseConfigurable](hosts ...string) func(T) {
	return func(c T) {
		c.getBase().allowedHosts = hosts
	}
}

// WithTokenStore sets the token store for the provider.
func WithTokenStore[T baseConfigurable](store TokenStore) func(T) {
	return func(c T) {
		c.getBase().tokenStore = store
	}
}

// WithHTTPClient sets the HTTP client for the provider.
func WithHTTPClient[T baseConfigurable](client *http.Client) func(T) {
	return func(c T) {
		c.getBase().httpClient = client
	}
}

// WithScopes sets the OAuth2 scopes.
func WithScopes[T oauth2Configurable](scopes ...string) func(T) {
	return func(c T) {
		c.getOAuth2().scopes = scopes
	}
}

// --- Auth Code Specific Options ---

// WithPort sets the port for the local callback server.
func WithPort(port int) func(*authCodeConfig) {
	return func(c *authCodeConfig) {
		c.port = port
	}
}

// WithStateGenerator sets the state generator function.
func WithStateGenerator(generator func() string) func(*authCodeConfig) {
	return func(c *authCodeConfig) {
		c.stateGenerator = generator
	}
}

// WithURLOpener sets the function to open the authorization URL.
func WithURLOpener(opener func(string) error) func(*authCodeConfig) {
	return func(c *authCodeConfig) {
		c.urlOpener = opener
	}
}

// WithServerFactory sets the factory function for creating the local callback server.
func WithServerFactory(factory ServerFactory) func(*authCodeConfig) {
	return func(c *authCodeConfig) {
		c.serverFactory = factory
	}
}
