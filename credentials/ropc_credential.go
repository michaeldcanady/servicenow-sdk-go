package credentials

import (
	"context"
	"net/url"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type ropcStrategy interface {
	acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// ROPCCredential implements the ROPC (Resource Owner Password Credentials) flow.
type ROPCCredential struct {
	*baseAccessTokenProvider
	client ropcStrategy
}

// NewROPCCredential creates a new ROPCCredential.
func NewROPCCredential(client ropcStrategy, username, password string, allowedHosts []string) (*ROPCCredential, error) {
	initialFunc := func(ctx context.Context, _ *url.URL, _ map[string]interface{}) (*AccessToken, error) {
		return client.acquireTokenByUsernamePassword(ctx, username, password)
	}

	refreshFunc := func(ctx context.Context, refreshToken string) (*AccessToken, error) {
		return client.acquireTokenByRefreshToken(ctx, refreshToken)
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.refreshToken = refreshFunc
	base.revokeToken = client.revokeToken

	return &ROPCCredential{
		baseAccessTokenProvider: base,
		client:                  client,
	}, nil
}

// NewROPCAuthenticationProvider creates a new AuthenticationProvider for the ROPC flow.
func NewROPCAuthenticationProvider(clientID, clientSecret, username, password string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority)
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewROPCCredential(client, username, password, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
