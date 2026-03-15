package credentials

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/authentication"
)

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
type AuthorizationCodeCredential struct {
	*baseAccessTokenProvider
	client *confidentialClient
}

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
// This credential uses the provided authorization code to acquire an access token.
// The code is one-time use. Subsequent token acquisitions will use the refresh token.
func NewAuthorizationCodeCredential(clientID, clientSecret, code, redirectURI string, authority Authority, allowedHosts []string) (*AuthorizationCodeCredential, error) {
	client, err := newConfidentialClient(clientID, clientSecret, authority)
	if err != nil {
		return nil, err
	}

	initialFunc := func(ctx context.Context) (*AccessToken, error) {
		token, err := client.acquireTokenByCode(ctx, code, redirectURI)
		if err != nil {
			return nil, err
		}
		return token, nil
	}

	refreshFunc := func(ctx context.Context, refreshToken string) (*AccessToken, error) {
		return client.acquireTokenByRefreshToken(ctx, refreshToken)
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = initialFunc
	base.refreshToken = refreshFunc

	return &AuthorizationCodeCredential{
		baseAccessTokenProvider: base,
		client:                  client,
	}, nil
}

// GetAuthorizationCodeURL generates the authorization URL for the 3-legged OAuth flow.
// This URL is used to redirect the user to the identity provider to sign in and consent to permissions.
func GetAuthorizationCodeURL(clientID, redirectURI string, authority Authority, state string, scopes []string) (string, error) {
	client, err := newPublicClient(clientID, authority)
	if err != nil {
		return "", err
	}
	return client.oauthClient.AuthCodeURL(redirectURI, state, "", "", scopes)
}

// NewAuthorizationCodeAuthenticationProvider creates a new AuthenticationProvider for the Authorization Code flow.
func NewAuthorizationCodeAuthenticationProvider(clientID, clientSecret, code, redirectURI string, authority Authority, allowedHosts []string) (authentication.AuthenticationProvider, error) {
	tokenProvider, err := NewAuthorizationCodeCredential(clientID, clientSecret, code, redirectURI, authority, allowedHosts)
	if err != nil {
		return nil, err
	}
	return authentication.NewBaseBearerTokenAuthenticationProvider(tokenProvider), nil
}
