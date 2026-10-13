// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oidc"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/pkg/browser"
)

// externalOIDCCredential implements the external OIDC Authorization Code + PKCE flow.
type externalOIDCCredential struct {
	*baseAccessTokenProvider
	client             authorizationCodeClient
	issuer             string
	tokenSource        ExternalOIDCTokenSource
	audience           string
	scopes             []string
	idToken            string
	storedRefreshToken string
	tokenMutex         sync.Mutex
}

// NewExternalOIDCProvider creates a new AuthenticationProvider for external OIDC
// using the Authorization Code + PKCE flow. The issuer is resolved through
// discovery at construction time. The default token source is the ID token;
// use ExternalOIDCTokenSourceAccessToken to present the access token instead.
// Default scopes are "openid".
func NewExternalOIDCProvider(clientID string, opts ExternalOIDCOptions) (authentication.AuthenticationProvider, error) {
	if clientID == "" {
		return nil, snerrors.ErrEmptyClientID
	}
	if strings.TrimSpace(opts.Issuer) == "" {
		return nil, snerrors.ErrEmptyIssuer
	}

	// Resolve issuer through discovery. Discovery uses the same HTTP client as
	// the token requests so callers can inject one client for both.
	discoveryClient := opts.HTTPClient
	if discoveryClient == nil {
		discoveryClient = http.DefaultClient
	}
	endpoints, err := oidc.Discover(context.Background(), opts.Issuer, discoveryClient)
	if err != nil {
		return nil, fmt.Errorf("external OIDC discovery failed: %w", err)
	}

	// Default scopes for OIDC
	scopes := opts.Scopes
	if len(scopes) == 0 {
		scopes = []string{"openid"}
	}

	// Build the authorization code client (public client with PKCE)
	client, err := newPublicClient(clientID, "", withPKCEChallenge(pkce.MethodS256), func(co *clientOptions) {
		co.httpClient = opts.HTTPClient
	})
	if err != nil {
		return nil, err
	}

	// Initialize the client with discovered endpoints
	client.Initialize(endpoints.AuthURL)

	tokenSource := opts.TokenSource
	audience := opts.Audience

	cred := &externalOIDCCredential{
		client:      client,
		issuer:      endpoints.AuthURL,
		tokenSource: tokenSource,
		audience:    audience,
		scopes:      scopes,
	}

	base := newBaseAccessTokenProvider()
	base.retrieveInitialToken = cred.retrieveInitialToken
	base.refreshToken = cred.refreshTokenFunc
	base.revokeToken = cred.revokeTokenFunc
	cred.baseAccessTokenProvider = base

	cred.SetAllowedHosts(opts.AllowedHosts)
	if opts.TokenStore != nil {
		cred.setKeyedTokenStore(opts.TokenStore)
	}

	provider, err := NewBearerTokenAuthenticationProvider(cred, BearerOptions{})
	if err != nil {
		return nil, err
	}
	return provider, nil
}

// retrieveInitialToken performs the authorization code flow and returns the chosen token.
func (c *externalOIDCCredential) retrieveInitialToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (*AccessToken, error) {
	// Generate a state
	stateGen := uuid.NewString
	stateVal := stateGen()

	// Create a redirect server
	server, err := defaultServerFactory(stateVal, 5001)
	if err != nil {
		return nil, fmt.Errorf("failed to create redirect server: %w", err)
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = server.Shutdown(shutdownCtx)
	}()

	// Build the authorization URL with the server's redirect URI
	// This internally generates the PKCE challenge and returns the verifier
	authURL, verifier, err := c.client.getAuthorizationURL(server.GetAddr(), stateVal, c.scopes)
	if err != nil {
		return nil, fmt.Errorf("authorization URL generation failed: %w", err)
	}

	// Open the authorization URL
	if err := browser.OpenURL(authURL); err != nil {
		return nil, fmt.Errorf("failed to open authorization URL: %w", err)
	}

	// Wait for the callback
	code, returnedState, err := server.Result(context.Background())
	if err != nil {
		return nil, fmt.Errorf("authorization code callback failed: %w", err)
	}
	if returnedState != stateVal {
		return nil, fmt.Errorf("state mismatch: expected %q, got %q", stateVal, returnedState)
	}

	// Exchange the code for tokens using the verifier returned by getAuthorizationURL
	token, err := c.client.acquireTokenByCode(ctx, code, server.GetAddr(), stateVal, verifier)
	if err != nil {
		return nil, fmt.Errorf("token exchange failed: %w", err)
	}

	// Store the ID token and refresh token for later use
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	if c.tokenSource == ExternalOIDCTokenSourceIDToken {
		// Extract ID token from the token response
		c.idToken = extractIDToken(token)
		c.storedRefreshToken = token.RefreshToken
		return &AccessToken{
			AccessToken:  c.idToken,
			ExpiresAt:    token.ExpiresAt,
			RefreshToken: token.RefreshToken,
			TokenType:    "Bearer",
		}, nil
	} else {
		// Use access token
		c.storedRefreshToken = token.RefreshToken
		return token, nil
	}
}

// refreshTokenFunc refreshes the access token using the refresh token.
func (c *externalOIDCCredential) refreshTokenFunc(ctx context.Context, refreshToken string) (*AccessToken, error) {
	token, err := c.client.acquireTokenByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	if c.tokenSource == ExternalOIDCTokenSourceIDToken {
		c.idToken = extractIDToken(token)
		c.storedRefreshToken = token.RefreshToken
		return &AccessToken{
			AccessToken:  c.idToken,
			ExpiresAt:    token.ExpiresAt,
			RefreshToken: token.RefreshToken,
			TokenType:    "Bearer",
		}, nil
	} else {
		c.storedRefreshToken = token.RefreshToken
		return token, nil
	}
}

// revokeTokenFunc revokes the current token.
func (c *externalOIDCCredential) revokeTokenFunc(ctx context.Context, token, tokenTypeHint string) error {
	err := c.client.revokeToken(ctx, token, tokenTypeHint)
	if err != nil {
		// A missing revocation endpoint is a best-effort success: clear local state.
		if errors.Is(err, snerrors.ErrRevocationEndpointNotSet) {
			// Best-effort: clear local state and succeed
			c.tokenMutex.Lock()
			c.idToken = ""
			c.storedRefreshToken = ""
			c.tokenMutex.Unlock()
			return nil
		}
		return err
	}

	c.tokenMutex.Lock()
	c.idToken = ""
	c.storedRefreshToken = ""
	c.tokenMutex.Unlock()
	return nil
}

// extractIDToken returns the ID token from the token response.
func extractIDToken(token *AccessToken) string {
	if token == nil {
		return ""
	}
	return token.IDToken
}
