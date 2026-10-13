// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
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
	TokenStore TokenStore
	// Scopes are the OAuth2 scopes requested for the token.
	Scopes []string
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

// ROPCOptions configures the ROPC flow.
type ROPCOptions struct {
	CommonOptions
}

// AuthorizationCodeServer is the interface for the local callback server used in Authorization Code flow.
type AuthorizationCodeServer interface {
	GetAddr() string
	Result(ctx context.Context) (code string, state string, err error)
	Shutdown(ctx context.Context) error
}

// ServerFactory is a function that creates a new AuthorizationCodeServer.
type ServerFactory func(state string, port int) (AuthorizationCodeServer, error)
