package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

// confidentialClient represents an application that has a client secret.
type confidentialClient struct {
	*baseClient
}

// newConfidentialClient creates a new confidentialClient.
func newConfidentialClient(clientID, clientSecret string, authority Authority, options ...clientOption) (*confidentialClient, error) {
	if clientID == "" {
		return nil, EmptyClientID
	}
	if clientSecret == "" {
		return nil, EmptyClientSecret
	}
	if authority == "" {
		return nil, EmptyBaseURL
	}

	opts := defaultOptions()
	for _, opt := range options {
		opt(&opts)
	}

	oauthClient := &oauth2.Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoints: &oauth2.Endpoints{
			TokenURL:         authority.TokenURL(),
			AuthURL:          authority.AuthURL(),
			DeviceURL:        "",
			RevocationURL:    authority.RevocationURL(),
			IntrospectionURL: "",
		},
		AuthMethod: oauth2.AuthMethodClientSecretPost, // Default to Post
		HTTPClient: opts.httpClient,
	}

	return &confidentialClient{
		baseClient: &baseClient{
			oauthClient: oauthClient,
		},
	}, nil
}

// acquireTokenByCode acquires a token using the authorization code flow.
func (c *confidentialClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeCode(ctx, code, redirectURI, "", state)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByClientCredentials acquires a token using the client credentials flow.
func (c *confidentialClient) acquireTokenByClientCredentials(ctx context.Context, scopes []string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeClientCredentials(ctx, scopes)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// getAuthorizationURL returns the authorization URL for the authorization code flow.
func (c *confidentialClient) getAuthorizationURL(redirectURI, state string, scopes []string) (string, error) {
	return c.oauthClient.AuthCodeURL(redirectURI, state, "", "", scopes)
}
