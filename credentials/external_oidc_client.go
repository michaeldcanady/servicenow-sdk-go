// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
)

// externalOIDCClient is the internal OIDC client for an external identity
// provider. It builds the PKCE-protected authorization URL against the
// provider's endpoints and exchanges the returned code (and refresh token) for
// the token the SDK presents to ServiceNow.
//
// The client is a public client: it authenticates with AuthMethodNone and never
// sends a client secret.
type externalOIDCClient struct {
	clientID string
	// scopes are the scopes requested from the provider. Defaults to ["openid"].
	scopes []string
	// useAccessToken selects the access token instead of the ID token as the
	// token presented to ServiceNow. It is set when the caller configures scopes.
	useAccessToken bool
	// method is the PKCE challenge method. v1 always uses S256.
	method pkce.Method
	// verifier is the PKCE code verifier generated for the current authorization
	// request, sent on the matching code exchange.
	verifier    string
	oauthClient *oauth2.Client
}

// newExternalOIDCClient builds a client for the given provider endpoints.
func newExternalOIDCClient(clientID string, endpoints *oauth2.Endpoints, scopes []string, useAccessToken bool, httpClient *http.Client) (*externalOIDCClient, error) {
	if strings.TrimSpace(clientID) == "" {
		return nil, ErrEmptyClientID
	}

	options := []oauth2.Option{oauth2.WithAuthMethod(oauth2.AuthMethodNone)}
	if httpClient != nil {
		options = append(options, oauth2.WithHTTPClient(httpClient))
	}

	return &externalOIDCClient{
		clientID:       clientID,
		scopes:         scopes,
		useAccessToken: useAccessToken,
		method:         pkce.MethodS256,
		oauthClient:    oauth2.NewClient(clientID, endpoints, options...),
	}, nil
}

// Initialize is a no-op. The external provider's endpoints come from OIDC
// discovery, not from the ServiceNow instance URL, so there is nothing to bind.
func (c *externalOIDCClient) Initialize(_ string) {}

// getAuthorizationURL builds the PKCE-protected authorization URL for the
// provider, generating a fresh code verifier for the exchange.
func (c *externalOIDCClient) getAuthorizationURL(redirectURI, state string, _ []string) (string, error) {
	challenge, verifier, err := generatePKCEChallenge(c.method)
	if err != nil {
		return "", err
	}

	c.verifier = verifier

	return c.oauthClient.AuthCodeURL(redirectURI, state, challenge, c.method.String(), c.scopes)
}

// acquireTokenByCode exchanges the authorization code for the presented token.
func (c *externalOIDCClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeCode(ctx, code, redirectURI, c.verifier, state)
	if err != nil {
		return nil, err
	}

	return c.toAccessToken(token)
}

// acquireTokenByRefreshToken refreshes the presented token using the provider's
// refresh token.
func (c *externalOIDCClient) acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return c.toAccessToken(token)
}

// revokeToken revokes a token at the identity provider. Revocation is
// best-effort: when the provider's discovery document does not advertise a
// revocation endpoint there is nothing to revoke at the provider, so this is a
// successful no-op and the shared base provider clears the cached token.
func (c *externalOIDCClient) revokeToken(ctx context.Context, token, tokenTypeHint string) error {
	if !c.revocationSupported() {
		return nil
	}

	return c.oauthClient.Revoke(ctx, token, tokenTypeHint)
}

// revocationSupported reports whether the identity provider advertised a
// revocation endpoint in its discovery document.
func (c *externalOIDCClient) revocationSupported() bool {
	return c.oauthClient.Endpoints != nil && strings.TrimSpace(c.oauthClient.Endpoints.RevocationURL) != ""
}

// toAccessToken maps the provider's token response onto the SDK token, choosing
// the token to present to ServiceNow and deriving its expiry from the token's
// exp claim.
func (c *externalOIDCClient) toAccessToken(token *oauth2.Token) (*AccessToken, error) {
	presented, err := c.presentedToken(token)
	if err != nil {
		return nil, err
	}

	expiresAt, err := tokenExpiry(presented)
	if err != nil {
		return nil, err
	}

	return &AccessToken{
		AccessToken:  presented,
		RefreshToken: token.RefreshToken,
		Scope:        token.Scope,
		TokenType:    token.TokenType,
		ExpiresAt:    expiresAt,
		ExpiresIn:    int(time.Until(expiresAt).Seconds()),
	}, nil
}

// presentedToken selects the token to present to ServiceNow: the ID token by
// default, or the access token when the caller configured scopes.
func (c *externalOIDCClient) presentedToken(token *oauth2.Token) (string, error) {
	if token == nil {
		return "", ErrMissingIDToken
	}

	if c.useAccessToken {
		if strings.TrimSpace(token.AccessToken) == "" {
			return "", ErrMissingAccessToken
		}

		return token.AccessToken, nil
	}

	idToken, _ := token.Raw["id_token"].(string)
	if strings.TrimSpace(idToken) == "" {
		return "", ErrMissingIDToken
	}

	return idToken, nil
}

// tokenExpiry reads the exp claim from a JWT, the SDK's only sanity check on the
// provider's token (ADR-014).
func tokenExpiry(token string) (time.Time, error) {
	claims, err := decodeJWTClaims(token)
	if err != nil {
		return time.Time{}, err
	}

	rawExp, ok := claims["exp"]
	if !ok {
		return time.Time{}, ErrMissingTokenExpiry
	}

	seconds, ok := rawExp.(float64)
	if !ok {
		return time.Time{}, ErrMissingTokenExpiry
	}

	return time.Unix(int64(seconds), 0), nil
}

// decodeJWTClaims decodes the claims of a JWT without verifying its signature.
func decodeJWTClaims(token string) (map[string]any, error) {
	segments := strings.Split(token, ".")
	if len(segments) != 3 {
		return nil, ErrUnparseableToken
	}

	payload, err := decodeJWTPayload(segments[1])
	if err != nil {
		return nil, ErrUnparseableToken
	}

	claims := map[string]any{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, ErrUnparseableToken
	}

	return claims, nil
}

// decodeJWTPayload decodes a base64url JWT segment, tolerating padded and
// unpadded encodings.
func decodeJWTPayload(segment string) ([]byte, error) {
	if decoded, err := base64.RawURLEncoding.DecodeString(segment); err == nil {
		return decoded, nil
	}

	return base64.URLEncoding.DecodeString(segment)
}
