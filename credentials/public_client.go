package credentials

import (
	"context"
	"net/url"
)

type tokenRequester interface {
	requestToken(ctx context.Context, params url.Values) (*AccessToken, error)
}

// publicClient represents an application that does not have a client secret.
type publicClient struct {
	clientID string
	tokenRequester
}

// newPublicClient creates a new publicClient.
func newPublicClient(clientID string, authority Authority, options ...clientOption) (*publicClient, error) {
	if clientID == "" {
		return nil, EmptyClientID
	}
	if authority == "" {
		return nil, EmptyBaseURL
	}

	opts := defaultOptions()
	for _, opt := range options {
		opt(&opts)
	}

	return &publicClient{
		clientID: clientID,
		tokenRequester: &defaultTokenRequester{
			authority: authority,
			options:   opts,
		},
	}, nil
}

// acquireTokenByUsernamePassword acquires a token using the ROPC flow.
func (c *publicClient) acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error) {
	params := url.Values{}
	params.Set("grant_type", "password")
	params.Set("client_id", c.clientID)
	params.Set("username", username)
	params.Set("password", password)

	return c.requestToken(ctx, params)
}
