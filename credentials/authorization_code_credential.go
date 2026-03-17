package credentials

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/pkg/browser"
)

var _ authentication.AccessTokenProvider = (*AuthorizationCodeCredential)(nil)

type authorizationCodeStrategy interface {
	getAuthorizationURL(redirectURI, state string, scopes []string) (string, error)
	acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
type AuthorizationCodeCredential struct {
	*BaseAccessTokenProvider
	client authorizationCodeStrategy
}

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
// This credential uses the provided authorization code to acquire an access token.
// The code is one-time use. Subsequent token acquisitions will use the refresh token.
func NewAuthorizationCodeCredential(client authorizationCodeStrategy, allowedHosts []string) (*AuthorizationCodeCredential, error) {
	c := &AuthorizationCodeCredential{
		client: client,
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

func (c *AuthorizationCodeCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (token *AccessToken, err error) {
	port := 5001

	state := uuid.NewString()

	server, err := oauth2.NewServer(state, port)
	if err != nil {
		return nil, err
	}

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if shutdownErr := server.Shutdown(shutdownCtx); shutdownErr != nil && err == nil {
			err = shutdownErr
		}
	}()

	authURL, err := c.client.getAuthorizationURL(server.Addr, state, nil)
	if err != nil {
		return nil, err
	}

	if err := browser.OpenURL(authURL); err != nil {
		return nil, err
	}

	result := server.Result(ctx)

	if err := result.Err; err != nil {
		return nil, err
	}

	token, err = c.client.acquireTokenByCode(ctx, result.Code, server.Addr, state)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// NewAuthorizationCodeProvider creates a new AuthenticationProvider for the Authorization Code flow using functional options.
func NewAuthorizationCodeProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := defaultAuthConfig()
	for _, opt := range opts {
		opt(config)
	}

	authority := Authority(config.baseURL)
	if authority == "" && config.instance != "" {
		authority = NewInstanceAuthority(config.instance)
	}

	var (
		client authorizationCodeStrategy
		err    error
	)

	if strings.TrimSpace(clientSecret) != "" {
		client, err = newConfidentialClient(clientID, clientSecret, authority, func(co *clientOptions) {
			co.httpClient = config.httpClient
		})
	} else {
		client, err = newPublicClient(clientID, authority, withPKCEChallenge(pkce.MethodS256), func(co *clientOptions) {
			co.httpClient = config.httpClient
		})
	}
	if err != nil {
		return nil, err
	}

	tokenProvider, err := NewAuthorizationCodeCredential(client, config.allowedHosts)
	if err != nil {
		return nil, err
	}

	if config.tokenStore != nil {
		tokenProvider.SetTokenStore(config.tokenStore)
	}

	return NewBearerTokenAuthenticationProvider(tokenProvider), nil
}
