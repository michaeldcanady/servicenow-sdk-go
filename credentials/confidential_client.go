// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
)

// confidentialClient represents an application that has a client secret.
type confidentialClient struct {
	*baseClient
}

// newConfidentialClient creates a new confidentialClient.
func newConfidentialClient(clientID, clientSecret string, authority Authority, options ...clientOption) (*confidentialClient, error) {
	if clientID == "" {
		return nil, snerrors.ErrEmptyClientID
	}
	if clientSecret == "" {
		return nil, snerrors.ErrEmptyClientSecret
	}

	opts := defaultOptions()
	for _, opt := range options {
		opt(&opts)
	}

	c := &confidentialClient{
		baseClient: &baseClient{
			clientID:     clientID,
			clientSecret: clientSecret,
			httpClient:   opts.httpClient,
		},
	}

	if authority != "" {
		c.Initialize(string(authority))
	}

	return c, nil
}

// acquireTokenByCode acquires a token using the authorization code flow.
func (c *confidentialClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state, verifier string) (*AccessToken, error) {
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

// acquireTokenByClientCredentials acquires a token using the client credentials flow.
func (c *confidentialClient) acquireTokenByClientCredentials(ctx context.Context, scopes []string) (*AccessToken, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return nil, err
	}
	token, err := client.ExchangeClientCredentials(ctx, scopes)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// getAuthorizationURL returns the authorization URL for the authorization code flow.
func (c *confidentialClient) getAuthorizationURL(redirectURI, state string, scopes []string) (string, string, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return "", "", err
	}
	authURL, err := client.AuthCodeURL(redirectURI, state, "", "", scopes)
	if err != nil {
		return "", "", err
	}
	return authURL, "", nil
}

func (c *confidentialClient) Initialize(baseURL string) {
	c.baseClient.Initialize(baseURL)
}
