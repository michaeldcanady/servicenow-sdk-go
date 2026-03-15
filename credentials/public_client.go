package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

// publicClient represents an application that does not have a client secret.
type publicClient struct {
	oauthClient *oauth2.Client
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

	oauthClient := &oauth2.Client{
		ClientID: clientID,
		Endpoints: &oauth2.Endpoints{
			TokenURL: authority.TokenURL(),
			AuthURL:  authority.AuthURL(),
		},
		AuthMethod: oauth2.AuthMethodClientSecretPost, // Public clients don't use secret, but we need a method that doesn't REQUIRE it in internal/oauth2
		HTTPClient: opts.httpClient,
	}

	return &publicClient{
		oauthClient: oauthClient,
	}, nil
}

// acquireTokenByUsernamePassword acquires a token using the ROPC flow.
func (c *publicClient) acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangePassword(ctx, username, password, nil)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}
