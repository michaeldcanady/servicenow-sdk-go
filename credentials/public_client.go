// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
)

// publicClient represents an application that does not have a client secret.
type publicClient struct {
	*baseClient
	method pkce.Method
}

// newPublicClient creates a new publicClient.
func newPublicClient(clientID string, authority authority, options ...clientOption) (*publicClient, error) {
	if clientID == "" {
		return nil, snerrors.ErrEmptyClientID
	}

	opts := defaultOptions()
	for _, opt := range options {
		opt(&opts)
	}

	c := &publicClient{
		baseClient: &baseClient{
			clientID:   clientID,
			httpClient: opts.httpClient,
		},
		method: opts.method,
	}

	if authority != "" {
		c.Initialize(string(authority))
	}

	return c, nil
}

// acquireTokenByCode acquires a token using the authorization code flow.
func (c *publicClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state, verifier string) (*AccessToken, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return nil, err
	}
	token, err := client.ExchangeCode(ctx, code, redirectURI, verifier, state)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// generateChallenge creates a PKCE challenge and the verifier it was derived
// from. The verifier is returned to the caller rather than stored on the
// client so each authorization presents its own proof.
func (c *publicClient) generateChallenge() (string, string, error) {
	if c.method == pkce.MethodUnset {
		return "", "", nil
	}

	gen := pkce.NewVerifierGenerator(pkce.DefaultVerifierEntropy)

	verifier, err := gen.Generate()
	if err != nil {
		return "", "", err
	}

	challenge, err := pkce.NewPKCEChallenge(c.method, verifier)
	if err != nil {
		return "", "", err
	}

	return challenge, verifier, nil
}

// getAuthorizationURL returns the authorization URL for the authorization code flow
// along with the PKCE verifier that must be presented when exchanging its code.
func (c *publicClient) getAuthorizationURL(redirectURI, state string, scopes []string) (string, string, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return "", "", err
	}
	challenge, verifier, err := c.generateChallenge()
	if err != nil {
		return "", "", err
	}

	authURL, err := client.AuthCodeURL(redirectURI, state, challenge, c.method.String(), scopes)
	if err != nil {
		return "", "", err
	}

	return authURL, verifier, nil
}

func (c *publicClient) Initialize(baseURL string) {
	c.baseClient.Initialize(baseURL)
}
