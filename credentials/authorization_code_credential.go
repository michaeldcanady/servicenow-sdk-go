package credentials

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/pkg/browser"
)

var _ authentication.AccessTokenProvider = (*AuthorizationCodeCredential)(nil)

type authorizationCodeClient interface {
	getAuthorizationURL(redirectURI, state string, scopes []string) (string, error)
	acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error)
	acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// AuthorizationCodeCredential implements the OAuth2 Authorization Code flow.
type AuthorizationCodeCredential struct {
	*BaseAccessTokenProvider
	client         authorizationCodeClient
	port           int
	stateGenerator func() string
	urlOpener      func(string) error
	serverFactory  ServerFactory
}

// NewAuthorizationCodeCredential creates a new AuthorizationCodeCredential.
// This credential uses the provided authorization code to acquire an access token.
// The code is one-time use. Subsequent token acquisitions will use the refresh token.
func NewAuthorizationCodeCredential(client authorizationCodeClient, allowedHosts []string, port int, stateGenerator func() string, urlOpener func(string) error, serverFactory ServerFactory) (*AuthorizationCodeCredential, error) {
	if port == 0 {
		port = 5001
	}
	if stateGenerator == nil {
		stateGenerator = uuid.NewString
	}
	if urlOpener == nil {
		urlOpener = browser.OpenURL
	}
	if serverFactory == nil {
		serverFactory = defaultServerFactory
	}

	c := &AuthorizationCodeCredential{
		client:         client,
		port:           port,
		stateGenerator: stateGenerator,
		urlOpener:      urlOpener,
		serverFactory:  serverFactory,
	}

	base := newBaseAccessTokenProvider(allowedHosts)
	base.retrieveInitialToken = c.GetToken
	base.refreshToken = client.acquireTokenByRefreshToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// GetToken acquires a token using the authorization code flow.
func (c *AuthorizationCodeCredential) GetToken(ctx context.Context, _ *url.URL, _ map[string]interface{}) (token *AccessToken, err error) {
	state := c.stateGenerator()

	server, err := c.serverFactory(state, c.port)
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

	authURL, err := c.client.getAuthorizationURL(server.GetAddr(), state, nil)
	if err != nil {
		return nil, err
	}

	if err := c.urlOpener(authURL); err != nil {
		return nil, err
	}

	code, _, err := server.Result(ctx)
	if err != nil {
		return nil, err
	}

	token, err = c.client.acquireTokenByCode(ctx, code, server.GetAddr(), state)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// NewAuthorizationCodeProvider creates a new AuthenticationProvider for the Authorization Code flow using functional options.
func NewAuthorizationCodeProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	config := &AuthConfig{
		httpClient: http.DefaultClient,
		port:       5001,
	}
	for _, opt := range opts {
		opt(config)
	}

	authority := Authority(config.baseURL)

	var (
		client authorizationCodeClient
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

	tokenProvider, err := NewAuthorizationCodeCredential(client, config.allowedHosts, config.port, config.stateGenerator, config.urlOpener, config.serverFactory)
	if err != nil {
		return nil, err
	}

	tokenProvider.Initialize(string(authority))

	if config.tokenStore != nil {
		tokenProvider.SetTokenStore(config.tokenStore)
	}

	return NewBearerTokenAuthenticationProvider(tokenProvider), nil
}

func NewPrivateAuthorizationCodeProvider(clientID, clientSecret string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	return NewAuthorizationCodeProvider(clientID, clientSecret, opts...)
}

func NewPublicAuthorizationCodeProvider(clientID string, opts ...AuthOption) (authentication.AuthenticationProvider, error) {
	return NewAuthorizationCodeProvider(clientID, "", opts...)
}
