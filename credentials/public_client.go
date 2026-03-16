package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
)

// publicClient represents an application that does not have a client secret.
type publicClient struct {
	oauthClient *oauth2.Client
	method      pkce.Method
	verifier    string
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
		AuthMethod: oauth2.AuthMethodClientSecretPost,
		HTTPClient: opts.httpClient,
	}

	return &publicClient{
		oauthClient: oauthClient,
		method:      opts.method,
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

// acquireTokenByCode acquires a token using the authorization code flow.
func (c *publicClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeCode(ctx, code, redirectURI, c.verifier, state)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}

func (c *publicClient) generateChallenge() (string, error) {
	if c.method == pkce.MethodUnset {
		return "", nil
	}

	gen := pkce.NewVerifierGenerator(pkce.DefaultVerifierEntropy)

	verifier, err := gen.Generate()
	if err != nil {
		return "", err
	}

	c.verifier = verifier

	challenge, err := pkce.NewPKCEChallenge(c.method, verifier)
	if err != nil {
		return "", err
	}

	return challenge, nil
}

// getAuthorizationURL returns the authorization URL for the authorization code flow.
func (c *publicClient) getAuthorizationURL(redirectURI, state string, scopes []string) (string, error) {
	challenge, err := c.generateChallenge()
	if err != nil {
		return "", err
	}

	return c.oauthClient.AuthCodeURL(redirectURI, state, challenge, c.method.String(), scopes)
}

// acquireTokenByRefreshToken acquires a new token using a refresh token.
func (c *publicClient) acquireTokenByRefreshToken(ctx context.Context, refreshToken string) (*AccessToken, error) {
	token, err := c.oauthClient.ExchangeRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return convertToken(token), nil
}
