package credentials

import (
	"context"
	"errors"
	"net/url"
	"testing"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// stubTokenProvider satisfies authentication.AccessTokenProvider without implementing
// Preparable, so it also exercises the Initialize no-op branch.
type stubTokenProvider struct {
	token    string
	err      error
	calls    int
	lastURI  *url.URL
	lastCtxt map[string]any
}

func (s *stubTokenProvider) GetAuthorizationToken(_ context.Context, uri *url.URL, additional map[string]any) (string, error) {
	s.calls++
	s.lastURI = uri
	s.lastCtxt = additional

	return s.token, s.err
}

func (s *stubTokenProvider) GetAllowedHostsValidator() *authentication.AllowedHostsValidator {
	return nil
}

// preparableStubTokenProvider is a stubTokenProvider that also implements Preparable.
type preparableStubTokenProvider struct {
	stubTokenProvider
	initializedWith string
	initializeCalls int
}

func (p *preparableStubTokenProvider) Initialize(baseURL string) {
	p.initializedWith = baseURL
	p.initializeCalls++
}

// newTestRequest builds a RequestInformation with a resolvable URI.
func newTestRequest() *abstractions.RequestInformation {
	request := abstractions.NewRequestInformation()
	request.UrlTemplate = "{+baseurl}/api/now/table/incident"
	request.PathParameters = map[string]string{"baseurl": "https://dev12345.service-now.com"}

	return request
}

func TestNewBearerTokenAuthenticationProvider(t *testing.T) {
	tests := []struct {
		name          string
		tokenProvider authentication.AccessTokenProvider
	}{
		{
			name:          "wraps the supplied token provider",
			tokenProvider: &stubTokenProvider{token: "abc"},
		},
		{
			name:          "accepts a nil token provider",
			tokenProvider: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := NewBearerTokenAuthenticationProvider(test.tokenProvider)

			require.NotNil(t, provider)
			require.NotNil(t, provider.BaseBearerTokenAuthenticationProvider)
			assert.Equal(t, test.tokenProvider, provider.tokenProvider)
		})
	}
}

func TestBearerTokenAuthenticationProvider_Initialize(t *testing.T) {
	t.Run("preparable token provider is initialized", func(t *testing.T) {
		tokenProvider := &preparableStubTokenProvider{}

		NewBearerTokenAuthenticationProvider(tokenProvider).Initialize("https://dev12345.service-now.com")

		assert.Equal(t, 1, tokenProvider.initializeCalls)
		assert.Equal(t, "https://dev12345.service-now.com", tokenProvider.initializedWith)
	})

	t.Run("BaseAccessTokenProvider records the base URL and derives a host validator", func(t *testing.T) {
		tokenProvider := newBaseAccessTokenProvider(nil)

		NewBearerTokenAuthenticationProvider(tokenProvider).Initialize("https://dev12345.service-now.com")

		assert.Equal(t, "https://dev12345.service-now.com", tokenProvider.baseURL)
		require.NotNil(t, tokenProvider.GetAllowedHostsValidator())
		assert.True(t, tokenProvider.GetAllowedHostsValidator().GetAllowedHosts()["dev12345.service-now.com"])
	})

	t.Run("non-preparable token provider is a no-op", func(t *testing.T) {
		provider := NewBearerTokenAuthenticationProvider(&stubTokenProvider{})

		assert.NotPanics(t, func() { provider.Initialize("https://dev12345.service-now.com") })
	})

	t.Run("nil token provider is a no-op", func(t *testing.T) {
		provider := NewBearerTokenAuthenticationProvider(nil)

		assert.NotPanics(t, func() { provider.Initialize("https://dev12345.service-now.com") })
	})
}

func TestBearerTokenAuthenticationProvider_AuthenticateRequest(t *testing.T) {
	tests := []struct {
		name          string
		tokenProvider *stubTokenProvider
		request       func() *abstractions.RequestInformation
		additional    map[string]any
		expectedErr   string
		expectedCalls int
		expectedAuth  []string
	}{
		{
			name:          "token is applied as a bearer header",
			tokenProvider: &stubTokenProvider{token: "token123"},
			request:       newTestRequest,
			expectedCalls: 1,
			expectedAuth:  []string{"Bearer token123"},
		},
		{
			name:          "empty token leaves the header unset",
			tokenProvider: &stubTokenProvider{token: ""},
			request:       newTestRequest,
			expectedCalls: 1,
			expectedAuth:  []string{},
		},
		{
			name:          "token provider error propagates",
			tokenProvider: &stubTokenProvider{err: errors.New("token acquisition failed")},
			request:       newTestRequest,
			expectedErr:   "token acquisition failed",
			expectedCalls: 1,
		},
		{
			name:          "existing authorization header short-circuits the token provider",
			tokenProvider: &stubTokenProvider{token: "token123"},
			request: func() *abstractions.RequestInformation {
				request := newTestRequest()
				request.Headers.Add("Authorization", "Bearer preexisting")

				return request
			},
			expectedCalls: 0,
			expectedAuth:  []string{"Bearer preexisting"},
		},
		{
			name:          "claims in the additional context force a token refresh",
			tokenProvider: &stubTokenProvider{token: "token123"},
			request: func() *abstractions.RequestInformation {
				request := newTestRequest()
				request.Headers.Add("Authorization", "Bearer stale")

				return request
			},
			additional:    map[string]any{"claims": "eyJhY2Nlc3MifQ=="},
			expectedCalls: 1,
			expectedAuth:  []string{"Bearer token123"},
		},
		{
			name:          "unresolvable request URI fails",
			tokenProvider: &stubTokenProvider{token: "token123"},
			request: func() *abstractions.RequestInformation {
				return abstractions.NewRequestInformation()
			},
			expectedErr:   "uri",
			expectedCalls: 0,
		},
		{
			name:          "nil request fails",
			tokenProvider: &stubTokenProvider{token: "token123"},
			request:       func() *abstractions.RequestInformation { return nil },
			expectedErr:   "request is nil",
			expectedCalls: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider := NewBearerTokenAuthenticationProvider(test.tokenProvider)
			request := test.request()

			err := provider.AuthenticateRequest(context.Background(), request, test.additional)

			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectedAuth, request.Headers.Get("Authorization"))
			}

			assert.Equal(t, test.expectedCalls, test.tokenProvider.calls)
		})
	}

	t.Run("nil token provider fails", func(t *testing.T) {
		provider := NewBearerTokenAuthenticationProvider(nil)

		err := provider.AuthenticateRequest(context.Background(), newTestRequest(), nil)

		require.ErrorContains(t, err, "needs to be initialized with an access token provider")
	})

	t.Run("request URI is forwarded to the token provider", func(t *testing.T) {
		tokenProvider := &stubTokenProvider{token: "token123"}
		provider := NewBearerTokenAuthenticationProvider(tokenProvider)

		require.NoError(t, provider.AuthenticateRequest(context.Background(), newTestRequest(), nil))

		require.NotNil(t, tokenProvider.lastURI)
		assert.Equal(t, "dev12345.service-now.com", tokenProvider.lastURI.Host)
		assert.Equal(t, "/api/now/table/incident", tokenProvider.lastURI.Path)
	})
}
