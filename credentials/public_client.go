package credentials

import (
	"context"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
)

// publicClient represents an application that does not have a client secret.
type publicClient struct {
	*baseClient
	method   pkce.Method
	verifier string
}

// newPublicClient creates a new publicClient.
func newPublicClient(clientID string, authority Authority, options ...clientOption) (*publicClient, error) {
	if clientID == "" {
		return nil, EmptyClientID
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
func (c *publicClient) acquireTokenByCode(ctx context.Context, code, redirectURI, state string) (*AccessToken, error) {
	client, err := c.getOAuthClient()
	if err != nil {
		return nil, err
	}
	token, err := client.ExchangeCode(ctx, code, redirectURI, c.verifier, state)
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
	client, err := c.getOAuthClient()
	if err != nil {
		return "", err
	}
	challenge, err := c.generateChallenge()
	if err != nil {
		return "", err
	}

	return client.AuthCodeURL(redirectURI, state, challenge, c.method.String(), scopes)
}

func (c *publicClient) Initialize(baseURL string) {
	c.baseClient.Initialize(baseURL)
}
