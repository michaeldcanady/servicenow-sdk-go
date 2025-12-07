package oauth2

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Helper to build HTTP response
func newResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       ioNopCloser{strings.NewReader(body)},
	}
}

// Minimal io.ReadCloser
type ioNopCloser struct{ *strings.Reader }

func (ioNopCloser) Close() error { return nil }

// --------------------
// Tests by method
// --------------------

func TestClient_ExchangeClientCredentials(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"access_token":"abc","token_type":"bearer","expires_in":3600}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{
					ClientID:     "id",
					ClientSecret: "secret",
					Endpoints: &Endpoints{
						TokenURL: "http://token",
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeClientCredentials(context.Background())
				require.NoError(t, err)
				require.Equal(t, "abc", tok.AccessToken)
				require.Equal(t, int64(3600), tok.ExpiresIn)

				mockClient.AssertExpectations(t)
			},
		},
		{
			name: "non-2xx returns TokenError",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(400, `{"error":"invalid_request","error_description":"bad"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{
					ClientID:     "id",
					ClientSecret: "secret",
					Endpoints: &Endpoints{
						TokenURL: "http://token",
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				_, err := client.ExchangeClientCredentials(context.Background())
				require.Error(t, err)
				_, ok := err.(*TokenError)
				require.True(t, ok)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_ExchangeRefreshToken(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"access_token":"def","token_type":"bearer"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{
					ClientID:     "id",
					ClientSecret: "secret",
					Endpoints: &Endpoints{
						TokenURL: "http://token",
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeRefreshToken(context.Background(), "refresh123")
				require.NoError(t, err)
				require.Equal(t, "def", tok.AccessToken)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_ExchangePassword(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"access_token":"ghi","token_type":"bearer"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{
					ClientID:     "id",
					ClientSecret: "secret",
					Endpoints: &Endpoints{
						TokenURL: "http://token",
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangePassword(context.Background(), "user", "pass")
				require.NoError(t, err)
				require.Equal(t, "ghi", tok.AccessToken)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_ExchangeCode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"access_token":"jkl","token_type":"bearer"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{
					ClientID:     "id",
					ClientSecret: "secret",
					Endpoints: &Endpoints{
						TokenURL: "http://token",
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeCode(context.Background(), "code123", "http://redirect", "verifier")
				require.NoError(t, err)
				require.Equal(t, "jkl", tok.AccessToken)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_ExchangeJWT(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"access_token":"mno","token_type":"bearer"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{
					ClientID:     "id",
					ClientSecret: "secret",
					Endpoints: &Endpoints{
						TokenURL: "http://token",
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeJWT(context.Background(), "assertion")
				require.NoError(t, err)
				require.Equal(t, "mno", tok.AccessToken)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_AuthCodeURL(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "builds correct URL",
			test: func(t *testing.T) {
				client := &Client{
					ClientID: "id",
					Endpoints: &Endpoints{
						AuthURL: "http://auth",
					},
				}
				u, err := client.AuthCodeURL("http://redirect", "state123", "challenge", PKCEMethodS256.String(), []string{"openid"})
				require.NoError(t, err)

				parsed, _ := url.Parse(u)
				q := parsed.Query()
				require.Equal(t, "id", q.Get(ClientIDKey))
				require.Equal(t, "http://redirect", q.Get(RedirectURIKey))
				require.Equal(t, "state123", q.Get(StateKey))
				require.Equal(t, "challenge", q.Get(CodeChallengeKey))
				require.Equal(t, PKCEMethodS256.String(), q.Get(CodeChallengeMethodKey))
				require.Equal(t, "openid", q.Get(ScopeKey))
			},
		},
		{
			name: "builds correct URL",
			test: func(t *testing.T) {
				client := &Client{
					ClientID: "id",
					Endpoints: &Endpoints{
						AuthURL: "",
					},
				}
				u, err := client.AuthCodeURL("http://redirect", "state123", "challenge", PKCEMethodS256.String(), []string{"openid"})
				require.Error(t, errors.New("authorization endpoint is not set"), err)
				require.Empty(t, u)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
