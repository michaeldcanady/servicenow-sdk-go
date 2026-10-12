// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/microsoft/kiota-abstractions-go/authentication"
)

type jwtClient interface {
	Initialize(baseURL string)
	acquireTokenByJWT(ctx context.Context, assertion string) (*AccessToken, error)
	revokeToken(ctx context.Context, token, tokenTypeHint string) error
}

// JWTCredential implements the OAuth2 JWT Bearer Token flow.
type JWTCredential struct {
	*BaseAccessTokenProvider
	tokenProvider authentication.AccessTokenProvider
	client        jwtClient
}

func validateJWT(rawToken string) error {
	// Parse without verifying the signature
	token, _, err := jwt.NewParser(
		jwt.WithoutClaimsValidation(),
	).ParseUnverified(rawToken, jwt.MapClaims{})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid JWT claims format")
	}

	// Validate required claims exist
	required := []string{"iss", "sub", "aud", "exp", "iat", "jti"}
	for _, key := range required {
		if _, exists := claims[key]; !exists {
			return errors.New("missing required claim: " + key)
		}
	}

	// Validate issued-at
	if iat, ok := claims["iat"].(float64); ok {
		if int64(iat) > time.Now().Unix() {
			return errors.New("JWT issued in the future")
		}
	} else {
		return errors.New("iat claim is not numeric")
	}

	// Validate algorithm
	if token.Method.Alg() != jwt.SigningMethodRS256.Alg() {
		return errors.New("unexpected signing algorithm: " + token.Method.Alg())
	}

	return nil
}

// NewJWTCredential creates a new JWTCredential.
func NewJWTCredential(client jwtClient, tokenProvider authentication.AccessTokenProvider) (*JWTCredential, error) {
	c := &JWTCredential{
		client:        client,
		tokenProvider: tokenProvider,
	}

	base := newBaseAccessTokenProvider()
	base.retrieveInitialToken = c.GetToken
	base.revokeToken = client.revokeToken

	c.BaseAccessTokenProvider = base

	return c, nil
}

// Initialize initializes the provider and its internal client with the base URL.
func (c *JWTCredential) Initialize(baseURL string) {
	c.BaseAccessTokenProvider.Initialize(baseURL)
	c.client.Initialize(baseURL)
}

// GetToken acquires a token using the JWT assertion.
func (c *JWTCredential) GetToken(ctx context.Context, uri *url.URL, additionalAuthenticationContext map[string]interface{}) (*AccessToken, error) {
	assertion, err := c.tokenProvider.GetAuthorizationToken(ctx, uri, additionalAuthenticationContext)
	if err != nil {
		return nil, err
	}

	if err := validateJWT(assertion); err != nil {
		return nil, err
	}

	return c.client.acquireTokenByJWT(ctx, assertion)
}

// NewJWTProvider creates a new AuthenticationProvider for the JWT Bearer Token flow.
func NewJWTProvider(clientID, clientSecret string, tokenProvider authentication.AccessTokenProvider, opts JWTOptions) (authentication.AuthenticationProvider, error) {
	if clientID == "" {
		return nil, snerrors.ErrEmptyClientID
	}
	if clientSecret == "" {
		return nil, snerrors.ErrEmptyClientSecret
	}
	if tokenProvider == nil {
		return nil, snerrors.ErrEmptyTokenProvider
	}

	client, err := newConfidentialClient(clientID, clientSecret, "", func(co *clientOptions) {
		co.httpClient = opts.HTTPClient
	})
	if err != nil {
		return nil, err
	}

	snTokenProvider, err := NewJWTCredential(client, tokenProvider)
	if err != nil {
		return nil, err
	}

	snTokenProvider.clientID = clientID
	snTokenProvider.SetAllowedHosts(opts.AllowedHosts)

	if opts.TokenStore != nil {
		snTokenProvider.SetTokenStore(opts.TokenStore)
	}

	return NewBearerTokenAuthenticationProvider(snTokenProvider), nil
}
