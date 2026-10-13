// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package credentials

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPublicClient(t *testing.T) {
	tests := []struct {
		name     string
		clientID string
		wantErr  error
	}{
		{
			name:     "empty client id returns error",
			clientID: "",
			wantErr:  snerrors.ErrEmptyClientID,
		},
		{
			name:     "valid client id succeeds",
			clientID: "client-id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := newPublicClient(test.clientID, Authority("https://dev12345.service-now.com"))

			if test.wantErr != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.wantErr)
				assert.Nil(t, client)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, client)
			assert.NotNil(t, client.oauthClient)
		})
	}

	t.Run("empty authority does not initialize the client", func(t *testing.T) {
		client, err := newPublicClient("client-id", "")
		require.NoError(t, err)
		require.NotNil(t, client)
		assert.Nil(t, client.oauthClient)
	})
}

func TestPublicClient_AcquireTokenByCode(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		wantErr    string
	}{
		{
			name:       "uninitialized client returns error",
			initClient: false,
			wantErr:    "OAuth2 client not initialized",
		},
		{
			name:       "exchange error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{err: errors.New("code exchange failed")},
			wantErr:    "code exchange failed",
		},
		{
			name:       "successful exchange returns converted token",
			initClient: true,
			fake:       &baseClientTestOAuthClient{token: &oauth2.Token{AccessToken: "code-access"}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &publicClient{baseClient: &baseClient{clientID: "id"}}
			if test.initClient {
				c.oauthClient = test.fake
			}

			token, err := c.acquireTokenByCode(context.Background(), "code", "https://redirect", "state", "verifier")

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "code-access", token.AccessToken)
			assert.Equal(t, "code", test.fake.lastCode)
		})
	}
}

func TestPublicClient_GenerateChallenge(t *testing.T) {
	tests := []struct {
		name      string
		method    pkce.Method
		wantEmpty bool
		wantErr   bool
	}{
		{
			name:      "unset method returns empty challenge",
			method:    pkce.MethodUnset,
			wantEmpty: true,
		},
		{
			name:   "S256 method generates a challenge and verifier",
			method: pkce.MethodS256,
		},
		{
			name:   "plain method generates a challenge and verifier",
			method: pkce.MethodPlain,
		},
		{
			name:    "unknown method returns error",
			method:  pkce.MethodUnknown,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &publicClient{baseClient: &baseClient{clientID: "id"}, method: test.method}

			challenge, verifier, err := c.generateChallenge()

			if test.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			if test.wantEmpty {
				assert.Empty(t, challenge)
				assert.Empty(t, verifier)
				return
			}
			assert.NotEmpty(t, challenge)
			assert.NotEmpty(t, verifier)
		})
	}
}

func TestPublicClient_GetAuthorizationURL(t *testing.T) {
	tests := []struct {
		name       string
		initClient bool
		fake       *baseClientTestOAuthClient
		method     pkce.Method
		wantErr    string
	}{
		{
			name:    "uninitialized client returns error",
			wantErr: "OAuth2 client not initialized",
		},
		{
			name:       "challenge generation error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{},
			method:     pkce.MethodUnknown,
			wantErr:    "unsupported PKCE method",
		},
		{
			name:       "auth code url error propagates",
			initClient: true,
			fake:       &baseClientTestOAuthClient{authCodeErr: errors.New("auth url failed")},
			method:     pkce.MethodS256,
			wantErr:    "auth url failed",
		},
		{
			name:       "successful url generation",
			initClient: true,
			fake:       &baseClientTestOAuthClient{authCodeURL: "https://auth.example.com"},
			method:     pkce.MethodS256,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &publicClient{baseClient: &baseClient{clientID: "id"}, method: test.method}
			if test.initClient {
				c.oauthClient = test.fake
			}

			url, _, err := c.getAuthorizationURL("https://redirect", "state", nil)

			if test.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "https://auth.example.com", url)
		})
	}
}

func TestPublicClient_Initialize(t *testing.T) {
	c := &publicClient{baseClient: &baseClient{clientID: "id"}}
	c.Initialize("https://dev12345.service-now.com")

	_, err := c.getOAuthClient()
	require.NoError(t, err)
}

// challengeEchoOAuthClient echoes the PKCE challenge into the authorization
// URL and records the verifier presented at code exchange.
type challengeEchoOAuthClient struct {
	*baseClientTestOAuthClient
	authURL            string
	exchangedCodes     []string
	exchangedVerifiers []string
}

func (c *challengeEchoOAuthClient) AuthCodeURL(_, _, codeChallenge, codeChallengeMethod string, _ []string) (string, error) {
	u, err := url.Parse(c.authURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("code_challenge", codeChallenge)
	q.Set("code_challenge_method", codeChallengeMethod)
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (c *challengeEchoOAuthClient) ExchangeCode(_ context.Context, code, _, verifier, _ string) (*oauth2.Token, error) {
	c.exchangedCodes = append(c.exchangedCodes, code)
	c.exchangedVerifiers = append(c.exchangedVerifiers, verifier)
	return &oauth2.Token{AccessToken: "token-" + code}, nil
}

func assertChallengeForVerifier(t *testing.T, authURL string, method pkce.Method, verifier string) {
	t.Helper()
	u, err := url.Parse(authURL)
	require.NoError(t, err)
	challenge := u.Query().Get("code_challenge")
	require.NotEmpty(t, challenge)
	want, err := pkce.NewPKCEChallenge(method, verifier)
	require.NoError(t, err)
	assert.Equal(t, want, challenge)
}

// capturingTokenEndpoint records the form body and Authorization header of the
// last request and replies with a canned token response.
type capturingTokenEndpoint struct {
	body           []byte
	lastForm       url.Values
	lastAuthHeader string
}

func (rt *capturingTokenEndpoint) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := req.ParseForm(); err == nil {
		rt.lastForm = req.Form
	}
	rt.lastAuthHeader = req.Header.Get("Authorization")
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       io.NopCloser(bytes.NewReader(rt.body)),
		Header:     make(http.Header),
		Request:    req,
	}, nil
}

func TestPublicClient_ExchangeOmitsClientSecret(t *testing.T) {
	rt := &capturingTokenEndpoint{body: []byte(`{"access_token":"t","expires_in":3600}`)}
	c := &publicClient{baseClient: &baseClient{clientID: "public-id", httpClient: &http.Client{Transport: rt}}}
	c.Initialize("https://dev12345.service-now.com")

	_, verifier, err := c.getAuthorizationURL("https://redirect", "state", nil)
	require.NoError(t, err)
	_, err = c.acquireTokenByCode(context.Background(), "code", "https://redirect", "state", verifier)
	require.NoError(t, err)

	require.NotNil(t, rt.lastForm)
	assert.Equal(t, "public-id", rt.lastForm.Get("client_id"))
	assert.Empty(t, rt.lastForm.Get("client_secret"), "a public client must not send a client_secret")
	assert.Empty(t, rt.lastAuthHeader, "a public client must not send an Authorization header")
}

func TestPublicClient_PKCEVerifierIsPerAuthorization(t *testing.T) {
	fake := &challengeEchoOAuthClient{
		baseClientTestOAuthClient: &baseClientTestOAuthClient{},
		authURL:                   "https://auth.example.com/authorize",
	}
	c := &publicClient{baseClient: &baseClient{clientID: "id"}, method: pkce.MethodS256}
	c.oauthClient = fake

	url1, verifier1, err := c.getAuthorizationURL("https://redirect", "state-1", nil)
	require.NoError(t, err)
	url2, verifier2, err := c.getAuthorizationURL("https://redirect", "state-2", nil)
	require.NoError(t, err)

	assert.NotEmpty(t, verifier1)
	assert.NotEqual(t, verifier1, verifier2, "each authorization must have its own verifier")

	assertChallengeForVerifier(t, url1, pkce.MethodS256, verifier1)
	assertChallengeForVerifier(t, url2, pkce.MethodS256, verifier2)

	_, err = c.acquireTokenByCode(context.Background(), "code-1", "https://redirect", "state-1", verifier1)
	require.NoError(t, err)
	_, err = c.acquireTokenByCode(context.Background(), "code-2", "https://redirect", "state-2", verifier2)
	require.NoError(t, err)

	assert.Equal(t, []string{"code-1", "code-2"}, fake.exchangedCodes)
	assert.Equal(t, []string{verifier1, verifier2}, fake.exchangedVerifiers)
}

func TestNewPublicClient_UsesNoSecretAuthMethod(t *testing.T) {
	client, err := newPublicClient("client-id", Authority("https://dev12345.service-now.com"))
	require.NoError(t, err)
	require.NotNil(t, client)

	oauthClient, ok := client.oauthClient.(*oauth2.Client)
	require.True(t, ok)
	assert.Equal(t, oauth2.AuthMethodNone, oauthClient.AuthMethod)
}

func TestNewConfidentialClient_UsesClientSecretPost(t *testing.T) {
	client, err := newConfidentialClient("client-id", "secret", Authority("https://dev12345.service-now.com"))
	require.NoError(t, err)
	require.NotNil(t, client)

	oauthClient, ok := client.oauthClient.(*oauth2.Client)
	require.True(t, ok)
	assert.Equal(t, oauth2.AuthMethodClientSecretPost, oauthClient.AuthMethod)
}
