package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

// confidentialClient represents an application that has a client secret.
type confidentialClient struct {
	oauthClient *oauth2.Client
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
			TokenURL: authority.TokenURL(),
			AuthURL:  authority.AuthURL(),
		},
		AuthMethod: oauth2.AuthMethodClientSecretPost, // Default to Post
		HTTPClient: opts.httpClient,
	}

	return &confidentialClient{
		oauthClient: oauthClient,
	}, nil
}

// acquireTokenByUsernamePassword acquires a token using the ROPC flow.
func (c *confidentialClient) acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangePassword(ctx, username, password, nil)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByRefreshToken acquires a new token using a refresh token.
func (c *confidentialClient) acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// acquireTokenByCode acquires a token using the authorization code flow.
func (c *confidentialClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeCode(ctx, code, redirectURI, "", state)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

// getAuthorizationURL returns the authorization URL for the authorization code flow.
func (c *confidentialClient) getAuthorizationURL(redirectURI, state string, scopes []string) (string, error) {
	return c.oauthClient.AuthCodeURL(redirectURI, state, "", "", scopes)
}
