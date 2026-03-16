package credentials

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// ROPCCredential implements the ROPC (Resource Owner Password Credentials) flow.
type ROPCCredential struct {
	*baseAccessTokenProvider
}

// NewROPCCredential creates a new ROPCCredential.
// If clientSecret is empty, it uses a public client, otherwise it uses a confidential client.
func NewROPCCredential(clientID, clientSecret, username, password string, authority Authority, allowedHosts []string) (*ROPCCredential, error) {
	var initialFunc func(ctx context.Context) (*AccessToken, error)
	var refreshFunc func(ctx context.Context, refreshToken string) (*AccessToken, error)
	var revokeToken func(ctx context.Context, token, tokenTypeHint string) error

	if clientSecret == "" {
		client, err := newPublicClient(clientID, authority)
		if err != nil {
			return nil, err
		}
		initialFunc = func(ctx context.Context) (*AccessToken, error) {
			return client.acquireTokenByUsernamePassword(ctx, username, password)
		}
		revokeToken = client.revokeToken
		// Public clients usually don't have a simple refresh_token grant in some OAuth2 implementations
		// but we can add it if needed.
	} else {
		client, err := newConfidentialClient(clientID, clientSecret, authority)
		if err != nil {
			return nil, err
		}
		initialFunc = func(ctx context.Context) (*AccessToken, error) {
			return client.acquireTokenByUsernamePassword(ctx, username, password)
		}
		refreshFunc = func(ctx context.Context, refreshToken string) (*AccessToken, error) {
			return client.acquireTokenByRefreshToken(ctx, refreshToken)
		}
		revokeToken = client.revokeToken
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.refreshToken = refreshFunc
	base.revokeToken = revokeToken

	return &ROPCCredential{
		baseAccessTokenProvider: base,
	}, nil
}

// NewUsernamePasswordAuthenticationProvider creates a new AuthenticationProvider for the ROPC flow.
func NewROPCAuthenticationProvider(clientID, clientSecret, username, password string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	tokenProvider, err := NewROPCCredential(clientID, clientSecret, username, password, authority, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
