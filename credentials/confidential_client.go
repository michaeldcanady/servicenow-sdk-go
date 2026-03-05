package credentials

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// confidentialClient represents an application that has a client secret.
type confidentialClient struct {
	clientID     string
	clientSecret string
	tokenRequester
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

	return &confidentialClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenRequester: &defaultTokenRequester{
			authority: authority,
			options:   opts,
		},
	}, nil
}

// acquireTokenByUsernamePassword acquires a token using the ROPC flow.
func (c *confidentialClient) acquireTokenByUsernamePassword(ctx context.Context, username, password string) (*AccessToken, error) {
	params := url.Values{}
	params.Set("grant_type", "password")
	params.Set("client_id", c.clientID)
	params.Set("client_secret", c.clientSecret)
	params.Set("username", username)
	params.Set("password", password)

	return c.requestToken(ctx, params)
}

// acquireTokenByRefreshToken acquires a new token using a refresh token.
func (c *confidentialClient) acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error) {
	params := url.Values{}
	params.Set("grant_type", "refresh_token")
	params.Set("client_id", c.clientID)
	params.Set("client_secret", c.clientSecret)
	params.Set("refresh_token", refreshToken)

	return c.requestToken(ctx, params)
}

type defaultTokenRequester struct {
	authority Authority
	options   clientOptions
}

func (c *defaultTokenRequester) requestToken(ctx context.Context, params url.Values) (*AccessToken, error) {
	tokenURL := c.authority.TokenURL()

	req, err := http.NewRequestWithContext(ctx, "POST", tokenURL, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// nolint: gosec //G704 external requests can't be processed here
	resp, err := c.options.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body) // ignore error here
		resp.Body.Close()
		return nil, fmt.Errorf("token request failed with status code %d, body: %s", resp.StatusCode, body)
	}

	return decodeAccessToken(resp)
}

func (c *confidentialClient) requestToken(ctx context.Context, params url.Values) (*AccessToken, error) {
	// Since we know c.tokenRequester is a *defaultTokenRequester (internal knowledge)
	tr := c.tokenRequester.(*defaultTokenRequester)
	return tr.requestToken(ctx, params)
}
