package oauth2

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2/pkce"
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

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeClientCredentials(context.Background(), []string{"scope1"})
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

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				_, err := client.ExchangeClientCredentials(context.Background(), nil)
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

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
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

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangePassword(context.Background(), "user", "pass", []string{"scope"})
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

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeCode(context.Background(), "code123", "http://redirect", "verifier", "")
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

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
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

func TestClient_RequestDeviceAuthorization(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"device_code":"dc","user_code":"uc","verification_uri":"http://verify","expires_in":300,"interval":5}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID: "id",
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
					},
					HTTPClient: mockClient,
				}
				dar, err := client.RequestDeviceAuthorization(context.Background(), []string{"scope"})
				require.NoError(t, err)
				require.Equal(t, "dc", dar.DeviceCode)
				require.Equal(t, "uc", dar.UserCode)
				require.Equal(t, "http://verify", dar.VerificationURI)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_ExchangeDeviceCode(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"access_token":"pqr","token_type":"bearer"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID: "id",
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:  "http://token",  //nolint:gosec // G101: Mock credentials for testing
						AuthURL:   "http://auth",   //nolint:gosec // G101: Mock credentials for testing
						DeviceURL: "http://device", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				tok, err := client.ExchangeDeviceCode(context.Background(), "dc123")
				require.NoError(t, err)
				require.Equal(t, "pqr", tok.AccessToken)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_Revoke(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, "")
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:         "http://token",      //nolint:gosec // G101: Mock credentials for testing
						AuthURL:          "http://auth",       //nolint:gosec // G101: Mock credentials for testing
						DeviceURL:        "http://device",     //nolint:gosec // G101: Mock credentials for testing
						RevocationURL:    "http://revoke",     //nolint:gosec // G101: Mock credentials for testing
						IntrospectionURL: "http://introspect", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				err := client.Revoke(context.Background(), "token123", "access_token")
				require.NoError(t, err)

				mockClient.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestClient_Introspect(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "success",
			test: func(t *testing.T) {
				mockClient := new(mocking.MockHTTPClient)
				resp := newResponse(200, `{"active":true,"scope":"read","client_id":"id"}`)
				mockClient.On("Do", mock.Anything).Return(resp, nil)

				client := &Client{ //nolint:gosec // G101: Mock credentials for testing
					ClientID:     "id",
					ClientSecret: "secret", //nolint:gosec // G101: Mock credentials for testing
					Endpoints: &Endpoints{ //nolint:gosec // G101: Mock credentials for testing
						TokenURL:         "http://token",      //nolint:gosec // G101: Mock credentials for testing
						AuthURL:          "http://auth",       //nolint:gosec // G101: Mock credentials for testing
						DeviceURL:        "http://device",     //nolint:gosec // G101: Mock credentials for testing
						RevocationURL:    "http://revoke",     //nolint:gosec // G101: Mock credentials for testing
						IntrospectionURL: "http://introspect", //nolint:gosec // G101: Mock credentials for testing
					},
					AuthMethod: AuthMethodClientSecretPost,
					HTTPClient: mockClient,
				}
				ir, err := client.Introspect(context.Background(), "token123", "")
				require.NoError(t, err)
				require.True(t, ir.Active)
				require.Equal(t, "read", ir.Scope)

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
					// nolint: gosec // no credentials
					Endpoints: &Endpoints{
						AuthURL:  "http://auth",
						TokenURL: "http://auth/token",
					},
				}
				u, err := client.AuthCodeURL("http://redirect", "state123", "challenge", pkce.MethodS256.String(), []string{"openid"})
				require.NoError(t, err)

				parsed, _ := url.Parse(u)
				q := parsed.Query()
				require.Equal(t, "id", q.Get(ClientIDKey))
				require.Equal(t, "http://redirect", q.Get(RedirectURIKey))
				require.Equal(t, "state123", q.Get(StateKey))
				require.Equal(t, "challenge", q.Get(CodeChallengeKey))
				require.Equal(t, pkce.MethodS256.String(), q.Get(CodeChallengeMethodKey))
				require.Equal(t, "openid", q.Get(ScopeKey))
			},
		},
		{
			name: "error if AuthURL is empty",
			test: func(t *testing.T) {
				client := &Client{
					// nolint: gosec // no credentials
					ClientID: "id",
					Endpoints: &Endpoints{
						AuthURL:  "",
						TokenURL: "https://url/token",
					},
				}
				u, err := client.AuthCodeURL("http://redirect", "state123", "challenge", pkce.MethodS256.String(), []string{"openid"})
				require.Error(t, err)
				require.Equal(t, "authorization endpoint is not set", err.Error())
				require.Empty(t, u)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
