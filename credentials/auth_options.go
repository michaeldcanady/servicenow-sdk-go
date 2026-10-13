// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"fmt"
	"net/http"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// CommonOptions contains options shared by all authentication providers.
type CommonOptions struct {
	// AllowedHosts restricts the hosts to which the provider will send credentials.
	AllowedHosts []string
	// HTTPClient is the HTTP client used for token requests.
	HTTPClient *http.Client
	// TokenStore persists tokens across process restarts.
	TokenStore KeyedTokenStore
	// Scopes are the OAuth2 scopes requested for the token.
	Scopes []string
}

// apply copies the shared options into a legacy AuthConfig.
func (o CommonOptions) apply(c *AuthConfig) {
	c.allowedHosts = o.AllowedHosts
	c.httpClient = o.HTTPClient
	c.tokenStore = o.TokenStore
	c.scopes = o.Scopes
}

// ClientCredentialsOptions configures the client-credentials flow.
type ClientCredentialsOptions struct {
	CommonOptions
}

// JWTOptions configures the JWT bearer flow.
type JWTOptions struct {
	CommonOptions
	// TokenProvider supplies the JWT assertion.
	TokenProvider authentication.AccessTokenProvider
}

// AuthorizationCodeOptions configures the authorization-code flow.
type AuthorizationCodeOptions struct {
	CommonOptions
	// Port is the port for the local callback server.
	Port int
	// StateGenerator generates the OAuth2 state parameter.
	StateGenerator func() string
	// URLOpener opens the authorization URL in the user's browser.
	URLOpener func(string) error
	// ServerFactory creates the local callback server.
	ServerFactory ServerFactory
}

// apply copies the options into a legacy AuthConfig.
func (o AuthorizationCodeOptions) apply(c *AuthConfig) {
	o.CommonOptions.apply(c)
	c.port = o.Port
	c.stateGenerator = o.StateGenerator
	c.urlOpener = o.URLOpener
	c.serverFactory = o.ServerFactory
}

// ROPCOptions configures the ROPC flow.
type ROPCOptions struct {
	CommonOptions
}

// ExternalOIDCTokenSource specifies which token to use from the OIDC response.
type ExternalOIDCTokenSource int

const (
	// ExternalOIDCTokenSourceIDToken uses the ID token as the bearer token (default).
	ExternalOIDCTokenSourceIDToken ExternalOIDCTokenSource = iota
	// ExternalOIDCTokenSourceAccessToken uses the access token as the bearer token.
	ExternalOIDCTokenSourceAccessToken
)

// ExternalOIDCOptions configures the external OIDC provider.
type ExternalOIDCOptions struct {
	CommonOptions
	// Issuer is the OIDC issuer URL or the full metadata URL.
	Issuer string
	// TokenSource specifies which token to present as the bearer token.
	TokenSource ExternalOIDCTokenSource
	// Audience is the expected audience for the access token (when using access token).
	Audience string
}

// AuthorizationCodeServer is the interface for the local callback server used in Authorization Code flow.
type AuthorizationCodeServer interface {
	GetAddr() string
	Result(ctx context.Context) (code string, state string, err error)
	Shutdown(ctx context.Context) error
}

// ServerFactory is a function that creates a new AuthorizationCodeServer.
type ServerFactory func(state string, port int) (AuthorizationCodeServer, error)

// AuthConfig contains options shared by all authentication providers.
//
// Deprecated: pass a per-flow option struct such as ClientCredentialsOptions
// to the provider constructor instead.
type AuthConfig struct {
	baseURL        string
	allowedHosts   []string
	tokenStore     KeyedTokenStore
	httpClient     *http.Client
	scopes         []string
	port           int
	stateGenerator func() string
	urlOpener      func(string) error
	serverFactory  ServerFactory
}

// AuthOption configures an authentication provider.
//
// Deprecated: pass a per-flow option struct such as ClientCredentialsOptions
// to the provider constructor instead. AuthOption is satisfied both by those
// structs and by the WithInstance, WithURL, and related helpers below.
type AuthOption interface {
	apply(*AuthConfig)
}

// authOptionFunc adapts a function to AuthOption so the deprecated WithX
// helpers keep their functional-option shape.
type authOptionFunc func(*AuthConfig)

// apply calls the wrapped function.
func (f authOptionFunc) apply(c *AuthConfig) { f(c) }

// defaultBaseURL is the suffix appended to a bare instance name.
const defaultBaseURL = "service-now.com"

// newAuthConfig returns an AuthConfig with the default HTTP client set.
func newAuthConfig() *AuthConfig {
	return &AuthConfig{httpClient: http.DefaultClient}
}

// buildAuthConfig folds options into an AuthConfig. Nil options are skipped so
// callers can pass a conditionally built option.
func buildAuthConfig(opts []AuthOption) *AuthConfig {
	c := newAuthConfig()
	for _, opt := range opts {
		if opt != nil {
			opt.apply(c)
		}
	}
	return c
}

// WithInstance sets the ServiceNow instance name (for example, "dev12345").
//
// Deprecated: register the instance on the client once instead.
func WithInstance(instance string) AuthOption {
	return WithURL(fmt.Sprintf("https://%s.%s", instance, defaultBaseURL))
}

// WithURL sets the full base URL for the ServiceNow instance.
//
// Deprecated: register the instance on the client once instead.
func WithURL(rawURL string) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.baseURL = rawURL
	})
}

// WithAllowedHosts sets the allowed hosts for the provider.
//
// Deprecated: set CommonOptions.AllowedHosts instead.
func WithAllowedHosts(hosts ...string) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.allowedHosts = hosts
	})
}

// WithTokenStore sets the token store for the provider.
//
// Deprecated: set CommonOptions.TokenStore to a KeyedTokenStore instead.
func WithTokenStore(store TokenStore) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		if store != nil {
			c.tokenStore = clientIDTokenStore{store: store}
		}
	})
}

// WithHTTPClient sets the HTTP client for the provider.
//
// Deprecated: set CommonOptions.HTTPClient instead.
func WithHTTPClient(client *http.Client) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.httpClient = client
	})
}

// WithScopes sets the OAuth2 scopes.
//
// Deprecated: set CommonOptions.Scopes instead.
func WithScopes(scopes ...string) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.scopes = scopes
	})
}

// WithPort sets the port for the local callback server.
//
// Deprecated: set AuthorizationCodeOptions.Port instead.
func WithPort(port int) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.port = port
	})
}

// WithStateGenerator sets the state generator function.
//
// Deprecated: set AuthorizationCodeOptions.StateGenerator instead.
func WithStateGenerator(generator func() string) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.stateGenerator = generator
	})
}

// WithURLOpener sets the function to open the authorization URL.
//
// Deprecated: set AuthorizationCodeOptions.URLOpener instead.
func WithURLOpener(opener func(string) error) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.urlOpener = opener
	})
}

// WithServerFactory sets the factory function for creating the local callback server.
//
// Deprecated: set AuthorizationCodeOptions.ServerFactory instead.
func WithServerFactory(factory ServerFactory) AuthOption {
	return authOptionFunc(func(c *AuthConfig) {
		c.serverFactory = factory
	})
}
