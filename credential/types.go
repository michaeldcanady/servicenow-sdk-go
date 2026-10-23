// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

// Package credential provides authentication providers for ServiceNow REST
// APIs. Providers receive the instance URL from the ServiceNow client; a flow
// config never contains an instance URL.
package credential

import (
	"context"
	"net/http"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// Initializer is implemented by providers that need the ServiceNow instance
// URL before they can acquire a token.
type Initializer interface {
	// Initialize stores the base URL used to reach ServiceNow.
	Initialize(baseURL string) error
}

// Revoker is implemented by token providers that can revoke their current
// access token.
type Revoker interface {
	// Revoke revokes the current token and clears the local cache.
	Revoke(ctx context.Context) error
}

// ProviderConfig carries transport settings shared by every provider.
type ProviderConfig struct {
	// HTTPClient sends token and discovery requests. A nil client uses the
	// package default.
	HTTPClient *http.Client
	// AllowedHosts limits authentication to these hosts. An empty list allows
	// any HTTPS host.
	AllowedHosts []string
}

// Validate checks the shared transport settings.
func (c ProviderConfig) Validate() error {
	_, err := authentication.NewAllowedHostsValidatorErrorCheck(c.AllowedHosts)
	return err
}

// Token is an access token and its associated metadata.
type Token struct {
	// AccessToken is the token presented to ServiceNow.
	AccessToken string
	// RefreshToken obtains a new access token when the authority issues one.
	RefreshToken string
	// IDToken is the OpenID Connect identity token when the authority issues one.
	IDToken string
	// TokenType is the token type reported by the authority.
	TokenType string
	// Scope is the scope reported by the authority.
	Scope string
	// ExpiresAt is the absolute expiry time. A zero value means the token has
	// no advertised expiry.
	ExpiresAt time.Time
}

// TokenKey identifies a cached token. Scopes is the sorted, space-joined scope
// string. SessionKey is empty for machine flows and required for user-specific
// flows when a token store is configured.
type TokenKey struct {
	// BaseURL is the normalized ServiceNow origin.
	BaseURL string
	// ClientID is the OAuth client identifier.
	ClientID string
	// Flow is the flow that acquired the token.
	Flow Flow
	// Scopes is the sorted, space-joined scope string.
	Scopes string
	// SessionKey identifies the user session for user-specific flows.
	SessionKey string
}

// TokenStore persists tokens across providers or processes.
type TokenStore interface {
	// Load returns the token for key, or (nil, nil) when no token exists.
	Load(ctx context.Context, key TokenKey) (*Token, error)
	// Save stores token for key.
	Save(ctx context.Context, key TokenKey, token *Token) error
	// Delete removes the token for key.
	Delete(ctx context.Context, key TokenKey) error
}

// ClientAuthMethod is the client authentication method used at the token
// endpoint.
type ClientAuthMethod int

const (
	// ClientAuthMethodDefault uses client_secret_post for native ServiceNow
	// flows.
	ClientAuthMethodDefault ClientAuthMethod = iota
	// ClientAuthMethodClientSecretPost sends the client secret in the request
	// body.
	ClientAuthMethodClientSecretPost
	// ClientAuthMethodClientSecretBasic sends the client secret with HTTP Basic
	// authentication.
	ClientAuthMethodClientSecretBasic
	// ClientAuthMethodNone sends no client secret. Use it only for public
	// clients.
	ClientAuthMethodNone
	// ClientAuthMethodUnknown is an invalid value.
	ClientAuthMethodUnknown ClientAuthMethod = -1
)

// String returns the OAuth name of the client authentication method.
func (m ClientAuthMethod) String() string {
	switch m {
	case ClientAuthMethodDefault, ClientAuthMethodClientSecretPost:
		return "client_secret_post"
	case ClientAuthMethodClientSecretBasic:
		return "client_secret_basic"
	case ClientAuthMethodNone:
		return "none"
	default:
		return "unknown"
	}
}

// AuthLocation is where a direct provider places its authentication value.
type AuthLocation int

const (
	// AuthLocationHeader places the value in an HTTP header.
	AuthLocationHeader AuthLocation = iota
	// AuthLocationQuery places the value in a query parameter.
	AuthLocationQuery
	// AuthLocationUnknown is an invalid value.
	AuthLocationUnknown AuthLocation = -1
)

// String returns the stored name of the authentication location.
func (l AuthLocation) String() string {
	switch l {
	case AuthLocationHeader:
		return "header"
	case AuthLocationQuery:
		return "query"
	default:
		return "unknown"
	}
}

func isNilTokenStore(store TokenStore) bool {
	return conversion.IsNil(store)
}
