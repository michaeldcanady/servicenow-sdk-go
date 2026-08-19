package oauth2

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/oauth2/pkce"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// newResponse builds a *http.Response with the given status and body for mocking HTTPClient.Do.
func newResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func testEndpoints() *Endpoints {
	return &Endpoints{
		TokenURL:         "http://token",
		AuthURL:          "http://auth",
		DeviceURL:        "http://device",
		RevocationURL:    "http://revoke",
		IntrospectionURL: "http://introspect",
	}
}

// --------------------
// Options / constructor
// --------------------

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
		want func(t *testing.T, c *Client)
	}{
		{
			name: "no options - defaults",
			opts: nil,
			want: func(t *testing.T, c *Client) {
				require.Equal(t, "id", c.ClientID)
				require.Empty(t, c.ClientSecret)
				require.Nil(t, c.HTTPClient)
				require.Nil(t, c.ErrorParser)
				require.Equal(t, AuthMethodClientSecretPost, c.AuthMethod)
			},
		},
		{
			name: "WithClientSecret sets secret",
			opts: []Option{WithClientSecret("shh")},
			want: func(t *testing.T, c *Client) {
				require.Equal(t, "shh", c.ClientSecret)
			},
		},
		{
			name: "WithAuthMethod sets method",
			opts: []Option{WithAuthMethod(AuthMethodClientSecretBasic)},
			want: func(t *testing.T, c *Client) {
				require.Equal(t, AuthMethodClientSecretBasic, c.AuthMethod)
			},
		},
		{
			name: "WithHTTPClient sets client",
			opts: []Option{WithHTTPClient(new(mocking.MockHTTPClient))},
			want: func(t *testing.T, c *Client) {
				require.NotNil(t, c.HTTPClient)
			},
		},
		{
			name: "WithErrorParser sets parser",
			opts: []Option{WithErrorParser(func(int, []byte) error { return errors.New("parsed") })},
			want: func(t *testing.T, c *Client) {
				require.NotNil(t, c.ErrorParser)
				require.EqualError(t, c.ErrorParser(400, nil), "parsed")
			},
		},
		{
			name: "multiple options combine",
			opts: []Option{WithClientSecret("s"), WithAuthMethod(AuthMethodNone)},
			want: func(t *testing.T, c *Client) {
				require.Equal(t, "s", c.ClientSecret)
				require.Equal(t, AuthMethodNone, c.AuthMethod)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient("id", testEndpoints(), tt.opts...)
			require.NotNil(t, c)
			require.Equal(t, testEndpoints(), c.Endpoints)
			tt.want(t, c)
		})
	}
}

// --------------------
// httpClient()
// --------------------

func TestClient_httpClient(t *testing.T) {
	tests := []struct {
		name   string
		client *Client
		want   HTTPClient
	}{
		{
			name:   "custom client returned when set",
			client: &Client{HTTPClient: new(mocking.MockHTTPClient)},
		},
		{
			name:   "nil client falls back to default",
			client: &Client{},
			want:   defaultHTTPClient,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.client.httpClient()
			require.NotNil(t, got)
			if tt.want != nil {
				require.Same(t, tt.want, got)
			} else {
				require.Same(t, tt.client.HTTPClient, got)
			}
		})
	}
}

// --------------------
// Exchange (+ doRequest)
// --------------------

func TestClient_Exchange(t *testing.T) {
	tests := []struct {
		name        string
		client      func(mockClient *mocking.MockHTTPClient) *Client
		params      url.Values
		mockSetup   func(mockClient *mocking.MockHTTPClient)
		wantErr     bool
		wantErrIs   string // substring to assert with ErrorContains
		wantToken   string
		wantAuthHdr string // expected Authorization header captured on the outgoing request, "" = not asserted
	}{
		{
			name: "nil endpoints fails validation",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", HTTPClient: mockClient}
			},
			params:    url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			wantErr:   true,
			wantErrIs: "endpoints are not set",
		},
		{
			name: "empty client id fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			params:    url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			wantErr:   true,
			wantErrIs: "client id is empty",
		},
		{
			name: "client_secret_basic missing client secret fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodClientSecretBasic, HTTPClient: mockClient}
			},
			params:    url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			wantErr:   true,
			wantErrIs: "client_secret_basic requires client_id and client_secret",
		},
		{
			name: "unsupported auth method fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethod(99), HTTPClient: mockClient}
			},
			params:    url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			wantErr:   true,
			wantErrIs: "unsupported auth method",
		},
		{
			name: "invalid token URL fails building request",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{
					ClientID:   "id",
					Endpoints:  &Endpoints{TokenURL: "http://token\n"},
					AuthMethod: AuthMethodNone,
					HTTPClient: mockClient,
				}
			},
			params:  url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			wantErr: true,
		},
		{
			name: "http client Do error propagates",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(nil, errors.New("boom"))
			},
			wantErr:   true,
			wantErrIs: "token request failed",
		},
		{
			name: "non-2xx with error parser returning error",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{
					ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient,
					ErrorParser: func(int, []byte) error { return errors.New("parsed error") },
				}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "parsed error",
		},
		{
			name: "non-2xx with error parser returning nil falls back to TokenError",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{
					ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient,
					ErrorParser: func(int, []byte) error { return nil },
				}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "invalid_request",
		},
		{
			name: "non-2xx without error parser returns TokenError",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request","error_description":"bad"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "invalid_request",
		},
		{
			name: "2xx with invalid json body fails to parse",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `not-json`), nil)
			},
			wantErr:   true,
			wantErrIs: "failed to parse token response",
		},
		{
			name: "client_secret_post success sets client_secret param",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", ClientSecret: "sec", Endpoints: testEndpoints(), AuthMethod: AuthMethodClientSecretPost, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"access_token":"abc","token_type":"bearer","expires_in":3600}`), nil)
			},
			wantToken: "abc",
		},
		{
			name: "client_secret_basic success sets Authorization header",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", ClientSecret: "sec", Endpoints: testEndpoints(), AuthMethod: AuthMethodClientSecretBasic, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"access_token":"basic-tok","token_type":"bearer"}`), nil)
			},
			wantToken:   "basic-tok",
			wantAuthHdr: "Basic aWQ6c2Vj",
		},
		{
			name: "auth_method_none success omits client secret",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"access_token":"none-tok","token_type":"bearer"}`), nil)
			},
			wantToken: "none-tok",
		},
		{
			name: "client id filled in from client when params missing it",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient}
			},
			params: url.Values{GrantTypeKey: []string{GrantTypeClientCreds}, ClientIDKey: []string{"already-set"}},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"access_token":"tok","token_type":"bearer"}`), nil)
			},
			wantToken: "tok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocking.MockHTTPClient)
			var captured *http.Request
			if tt.mockSetup != nil {
				tt.mockSetup(mockClient)
			}
			mockClient.Test(t)
			// wrap capture on top of any existing expectation by re-registering with Run.
			if tt.wantAuthHdr != "" {
				mockClient.ExpectedCalls = nil
				mockClient.On("Do", mock.Anything).Run(func(args mock.Arguments) {
					captured = args.Get(0).(*http.Request)
				}).Return(newResponse(200, `{"access_token":"basic-tok","token_type":"bearer"}`), nil)
			}

			c := tt.client(mockClient)
			tok, err := c.Exchange(context.Background(), tt.params)

			if tt.wantErr {
				require.Error(t, err)
				if tt.wantErrIs != "" {
					require.ErrorContains(t, err, tt.wantErrIs)
				}
				require.Nil(t, tok)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, tok)
			require.Equal(t, tt.wantToken, tok.AccessToken)
			if tt.wantAuthHdr != "" {
				require.NotNil(t, captured)
				require.Equal(t, tt.wantAuthHdr, captured.Header.Get(AuthorizationKey))
			}
		})
	}
}

// --------------------
// Grant-type convenience wrappers
// --------------------

func TestClient_ExchangeWrappers(t *testing.T) {
	tests := []struct {
		name          string
		call          func(c *Client) (*Token, error)
		wantGrantType string
		wantParams    map[string]string
		absentParams  []string
	}{
		{
			name: "client credentials with scopes",
			call: func(c *Client) (*Token, error) {
				return c.ExchangeClientCredentials(context.Background(), []string{"a", "b"})
			},
			wantGrantType: GrantTypeClientCreds,
			wantParams:    map[string]string{ScopeKey: "a b"},
		},
		{
			name:          "client credentials without scopes",
			call:          func(c *Client) (*Token, error) { return c.ExchangeClientCredentials(context.Background(), nil) },
			wantGrantType: GrantTypeClientCreds,
			absentParams:  []string{ScopeKey},
		},
		{
			name:          "refresh token",
			call:          func(c *Client) (*Token, error) { return c.ExchangeRefreshToken(context.Background(), "refresh123") },
			wantGrantType: GrantTypeRefreshToken,
			wantParams:    map[string]string{RefreshTokenKey: "refresh123"},
		},
		{
			name: "password with scopes",
			call: func(c *Client) (*Token, error) {
				return c.ExchangePassword(context.Background(), "user", "pass", []string{"scope"})
			},
			wantGrantType: GrantTypePassword,
			wantParams:    map[string]string{UsernameKey: "user", PasswordKey: "pass", ScopeKey: "scope"},
		},
		{
			name:          "password without scopes",
			call:          func(c *Client) (*Token, error) { return c.ExchangePassword(context.Background(), "user", "pass", nil) },
			wantGrantType: GrantTypePassword,
			absentParams:  []string{ScopeKey},
		},
		{
			name: "authorization code with all fields",
			call: func(c *Client) (*Token, error) {
				return c.ExchangeCode(context.Background(), "code123", "http://redirect", "verifier", "state1")
			},
			wantGrantType: GrantTypeAuthCode,
			wantParams: map[string]string{
				CodeKey: "code123", RedirectURIKey: "http://redirect", CodeVerifierKey: "verifier", StateKey: "state1",
			},
		},
		{
			name: "authorization code omits empty optional fields",
			call: func(c *Client) (*Token, error) {
				return c.ExchangeCode(context.Background(), "code123", "", "", "")
			},
			wantGrantType: GrantTypeAuthCode,
			wantParams:    map[string]string{CodeKey: "code123"},
			absentParams:  []string{RedirectURIKey, StateKey, CodeVerifierKey},
		},
		{
			name:          "jwt bearer",
			call:          func(c *Client) (*Token, error) { return c.ExchangeJWT(context.Background(), "assertion") },
			wantGrantType: GrantTypeJWTBearer,
			wantParams:    map[string]string{AssertionKey: "assertion"},
		},
		{
			name:          "device code",
			call:          func(c *Client) (*Token, error) { return c.ExchangeDeviceCode(context.Background(), "dc123") },
			wantGrantType: GrantTypeDeviceCode,
			wantParams:    map[string]string{DeviceCodeKey: "dc123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocking.MockHTTPClient)
			var captured *http.Request
			mockClient.On("Do", mock.Anything).Run(func(args mock.Arguments) {
				captured = args.Get(0).(*http.Request)
			}).Return(newResponse(200, `{"access_token":"tok","token_type":"bearer"}`), nil)

			c := &Client{ClientID: "id", Endpoints: testEndpoints(), AuthMethod: AuthMethodNone, HTTPClient: mockClient}
			tok, err := tt.call(c)
			require.NoError(t, err)
			require.Equal(t, "tok", tok.AccessToken)

			require.NotNil(t, captured)
			body, err := io.ReadAll(captured.Body)
			require.NoError(t, err)
			values, err := url.ParseQuery(string(body))
			require.NoError(t, err)

			require.Equal(t, tt.wantGrantType, values.Get(GrantTypeKey))
			for k, v := range tt.wantParams {
				require.Equal(t, v, values.Get(k), "param %q", k)
			}
			for _, k := range tt.absentParams {
				require.False(t, values.Has(k), "expected param %q to be absent", k)
			}

			mockClient.AssertExpectations(t)
		})
	}
}

// --------------------
// RequestDeviceAuthorization
// --------------------

func TestClient_RequestDeviceAuthorization(t *testing.T) {
	tests := []struct {
		name      string
		client    func(mockClient *mocking.MockHTTPClient) *Client
		scopes    []string
		mockSetup func(mockClient *mocking.MockHTTPClient)
		wantErr   bool
		wantErrIs string
		wantDAR   *DeviceAuthorizationResponse
	}{
		{
			name: "missing device endpoint fails validation",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: &Endpoints{TokenURL: "http://token"}, HTTPClient: mockClient}
			},
			wantErr:   true,
			wantErrIs: "device authorization endpoint is not set",
		},
		{
			name: "invalid device URL fails building request",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: &Endpoints{TokenURL: "http://token", DeviceURL: "http://device\n"}, HTTPClient: mockClient}
			},
			wantErr: true,
		},
		{
			name: "http client Do error propagates",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(nil, errors.New("boom"))
			},
			wantErr: true,
		},
		{
			name: "non-2xx with error parser",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{
					ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient,
					ErrorParser: func(int, []byte) error { return errors.New("parsed") },
				}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "parsed",
		},
		{
			name: "non-2xx without error parser returns TokenError",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "invalid_request",
		},
		{
			name: "invalid json body fails to parse",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `not-json`), nil)
			},
			wantErr:   true,
			wantErrIs: "failed to parse device authorization response",
		},
		{
			name:   "success with scopes",
			scopes: []string{"scope"},
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"device_code":"dc","user_code":"uc","verification_uri":"http://verify","expires_in":300,"interval":5}`), nil)
			},
			wantDAR: &DeviceAuthorizationResponse{DeviceCode: "dc", UserCode: "uc", VerificationURI: "http://verify", ExpiresIn: 300, Interval: 5},
		},
		{
			name: "success without scopes",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"device_code":"dc2","user_code":"uc2","verification_uri":"http://verify"}`), nil)
			},
			wantDAR: &DeviceAuthorizationResponse{DeviceCode: "dc2", UserCode: "uc2", VerificationURI: "http://verify"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocking.MockHTTPClient)
			if tt.mockSetup != nil {
				tt.mockSetup(mockClient)
			}
			c := tt.client(mockClient)
			dar, err := c.RequestDeviceAuthorization(context.Background(), tt.scopes)

			if tt.wantErr {
				require.Error(t, err)
				if tt.wantErrIs != "" {
					require.ErrorContains(t, err, tt.wantErrIs)
				}
				require.Nil(t, dar)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.wantDAR, dar)
		})
	}
}

// --------------------
// Revoke
// --------------------

func TestClient_Revoke(t *testing.T) {
	tests := []struct {
		name      string
		client    func(mockClient *mocking.MockHTTPClient) *Client
		mockSetup func(mockClient *mocking.MockHTTPClient)
		wantErr   bool
		wantErrIs string
	}{
		{
			name: "nil endpoints fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{HTTPClient: mockClient}
			},
			wantErr:   true,
			wantErrIs: "revocation endpoint is not set",
		},
		{
			name: "empty revocation URL fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{Endpoints: &Endpoints{TokenURL: "http://token"}, HTTPClient: mockClient}
			},
			wantErr:   true,
			wantErrIs: "revocation endpoint is not set",
		},
		{
			name: "invalid revocation URL fails building request",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{Endpoints: &Endpoints{RevocationURL: "http://revoke\n"}, HTTPClient: mockClient}
			},
			wantErr: true,
		},
		{
			name: "http client Do error propagates",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(nil, errors.New("boom"))
			},
			wantErr: true,
		},
		{
			name: "non-2xx with error parser returning error",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{
					ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient,
					ErrorParser: func(int, []byte) error { return errors.New("parsed") },
				}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "parsed",
		},
		{
			name: "non-2xx without error parser returns TokenError",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "invalid_request",
		},
		{
			name: "success with token type hint",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", ClientSecret: "secret", Endpoints: testEndpoints(), AuthMethod: AuthMethodClientSecretPost, HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, ""), nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocking.MockHTTPClient)
			if tt.mockSetup != nil {
				tt.mockSetup(mockClient)
			}
			c := tt.client(mockClient)
			err := c.Revoke(context.Background(), "token123", "access_token")

			if tt.wantErr {
				require.Error(t, err)
				if tt.wantErrIs != "" {
					require.ErrorContains(t, err, tt.wantErrIs)
				}
				return
			}
			require.NoError(t, err)
		})
	}
}

// --------------------
// Introspect
// --------------------

func TestClient_Introspect(t *testing.T) {
	tests := []struct {
		name          string
		tokenTypeHint string
		client        func(mockClient *mocking.MockHTTPClient) *Client
		mockSetup     func(mockClient *mocking.MockHTTPClient)
		wantErr       bool
		wantErrIs     string
		wantActive    bool
		wantScope     string
	}{
		{
			name: "nil endpoints fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{HTTPClient: mockClient}
			},
			wantErr:   true,
			wantErrIs: "introspection endpoint is not set",
		},
		{
			name: "empty introspection URL fails",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{Endpoints: &Endpoints{TokenURL: "http://token"}, HTTPClient: mockClient}
			},
			wantErr:   true,
			wantErrIs: "introspection endpoint is not set",
		},
		{
			name: "invalid introspection URL fails building request",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{Endpoints: &Endpoints{IntrospectionURL: "http://introspect\n"}, HTTPClient: mockClient}
			},
			wantErr: true,
		},
		{
			name: "http client Do error propagates",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(nil, errors.New("boom"))
			},
			wantErr: true,
		},
		{
			name: "non-2xx with error parser returning error",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{
					ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient,
					ErrorParser: func(int, []byte) error { return errors.New("parsed") },
				}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "parsed",
		},
		{
			name: "non-2xx without error parser returns TokenError",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(400, `{"error":"invalid_request"}`), nil)
			},
			wantErr:   true,
			wantErrIs: "invalid_request",
		},
		{
			name: "invalid json body fails to parse",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", Endpoints: testEndpoints(), HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `not-json`), nil)
			},
			wantErr:   true,
			wantErrIs: "failed to parse introspection response",
		},
		{
			name:          "success with token type hint",
			tokenTypeHint: "access_token",
			client: func(mockClient *mocking.MockHTTPClient) *Client {
				return &Client{ClientID: "id", ClientSecret: "secret", Endpoints: testEndpoints(), AuthMethod: AuthMethodClientSecretPost, HTTPClient: mockClient}
			},
			mockSetup: func(mockClient *mocking.MockHTTPClient) {
				mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"active":true,"scope":"read","client_id":"id"}`), nil)
			},
			wantActive: true,
			wantScope:  "read",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(mocking.MockHTTPClient)
			if tt.mockSetup != nil {
				tt.mockSetup(mockClient)
			}
			c := tt.client(mockClient)
			ir, err := c.Introspect(context.Background(), "token123", tt.tokenTypeHint)

			if tt.wantErr {
				require.Error(t, err)
				if tt.wantErrIs != "" {
					require.ErrorContains(t, err, tt.wantErrIs)
				}
				require.Nil(t, ir)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.wantActive, ir.Active)
			require.Equal(t, tt.wantScope, ir.Scope)
			require.NotNil(t, ir.Raw)
		})
	}
}

// --------------------
// newAuthenticatedRequest
// --------------------

func TestClient_newAuthenticatedRequest(t *testing.T) {
	tests := []struct {
		name        string
		client      *Client
		url         string
		params      url.Values
		wantErr     bool
		wantAuthHdr string
	}{
		{
			name:        "client_secret_basic sets Authorization header",
			client:      &Client{ClientID: "id", ClientSecret: "sec", AuthMethod: AuthMethodClientSecretBasic},
			url:         "http://x",
			params:      url.Values{},
			wantAuthHdr: "Basic aWQ6c2Vj",
		},
		{
			name:   "client_secret_basic missing creds sets no header and no error",
			client: &Client{AuthMethod: AuthMethodClientSecretBasic},
			url:    "http://x",
			params: url.Values{},
		},
		{
			name:   "client_secret_post sets client_secret param",
			client: &Client{ClientSecret: "sec", AuthMethod: AuthMethodClientSecretPost},
			url:    "http://x",
			params: url.Values{},
		},
		{
			name:   "auth_method_none no header",
			client: &Client{AuthMethod: AuthMethodNone},
			url:    "http://x",
			params: url.Values{},
		},
		{
			name:   "unsupported auth method silently produces no header",
			client: &Client{AuthMethod: AuthMethod(99)},
			url:    "http://x",
			params: url.Values{},
		},
		{
			name:    "invalid url fails",
			client:  &Client{AuthMethod: AuthMethodNone},
			url:     "http://x\n",
			params:  url.Values{},
			wantErr: true,
		},
		{
			name:   "client id already present is not overwritten",
			client: &Client{ClientID: "id", AuthMethod: AuthMethodNone},
			url:    "http://x",
			params: url.Values{ClientIDKey: []string{"already-there"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := tt.client.newAuthenticatedRequest(context.Background(), http.MethodPost, tt.url, tt.params)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, req)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, req)
			require.Equal(t, tt.wantAuthHdr, req.Header.Get(AuthorizationKey))
		})
	}
}

// --------------------
// AuthCodeURL
// --------------------

func TestClient_AuthCodeURL(t *testing.T) {
	tests := []struct {
		name    string
		client  *Client
		wantErr string // substring; empty means no error expected
		check   func(t *testing.T, u string)
	}{
		{
			name: "builds correct URL with all fields",
			client: &Client{
				ClientID: "id",
				Endpoints: &Endpoints{
					AuthURL:  "http://auth",
					TokenURL: "http://auth/token",
				},
			},
			check: func(t *testing.T, u string) {
				parsed, err := url.Parse(u)
				require.NoError(t, err)
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
			name: "includes trimmed client secret when set",
			client: &Client{
				ClientID:     "id",
				ClientSecret: "  sec  ",
				Endpoints: &Endpoints{
					AuthURL:  "http://auth",
					TokenURL: "http://auth/token",
				},
			},
			check: func(t *testing.T, u string) {
				parsed, err := url.Parse(u)
				require.NoError(t, err)
				require.Equal(t, "sec", parsed.Query().Get(ClientSecretKey))
			},
		},
		{
			name: "omits code_challenge_method when code challenge method empty",
			client: &Client{
				ClientID: "id",
				Endpoints: &Endpoints{
					AuthURL:  "http://auth",
					TokenURL: "http://auth/token",
				},
			},
			check: func(t *testing.T, u string) {
				parsed, err := url.Parse(u)
				require.NoError(t, err)
				require.False(t, parsed.Query().Has(CodeChallengeMethodKey))
			},
		},
		{
			name: "error if AuthURL is empty",
			client: &Client{
				ClientID: "id",
				Endpoints: &Endpoints{
					AuthURL:  "",
					TokenURL: "https://url/token",
				},
			},
			wantErr: "authorization endpoint is not set",
		},
		{
			name: "error if endpoints validation fails on missing token URL",
			client: &Client{
				ClientID:  "id",
				Endpoints: &Endpoints{AuthURL: "http://auth"},
			},
			wantErr: "token endpoint is not set",
		},
		{
			name: "error if auth URL is unparsable",
			client: &Client{
				ClientID: "id",
				Endpoints: &Endpoints{
					AuthURL:  "://bad",
					TokenURL: "http://token",
				},
			},
			wantErr: "invalid auth URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u string
			var err error
			if tt.name == "omits code_challenge_method when code challenge method empty" {
				u, err = tt.client.AuthCodeURL("", "", "", "", nil)
			} else {
				u, err = tt.client.AuthCodeURL("http://redirect", "state123", "challenge", pkce.MethodS256.String(), []string{"openid"})
			}

			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				require.Empty(t, u)
				return
			}
			require.NoError(t, err)
			tt.check(t, u)
		})
	}
}
