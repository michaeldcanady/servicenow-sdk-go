package credentials

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationCodeURL(t *testing.T) {
	clientID := "test-client-id"
	redirectURI := "http://localhost:8080/callback"
	authority := NewInstanceAuthority("dev12345")
	state := "random-state"
	scopes := []string{"user_account", "glide_api"}

	authURL, err := GetAuthorizationCodeURL(clientID, redirectURI, authority, state, scopes)
	assert.NoError(t, err)

	u, err := url.Parse(authURL)
	assert.NoError(t, err)

	assert.Equal(t, "https", u.Scheme)
	assert.Equal(t, "dev12345.service-now.com", u.Host)
	assert.Equal(t, "/oauth_auth.do", u.Path)

	q := u.Query()
	assert.Equal(t, "code", q.Get("response_type"))
	assert.Equal(t, clientID, q.Get("client_id"))
	assert.Equal(t, redirectURI, q.Get("redirect_uri"))
	assert.Equal(t, state, q.Get("state"))
	assert.Equal(t, strings.Join(scopes, " "), q.Get("scope"))
}

func TestAuthorizationCodeCredential(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	authority := NewInstanceAuthority("dev12345")
	clientID := "test-client-id"
	clientSecret := "test-client-secret"
	redirectURI := "http://localhost:8080/callback"
	code := "test-auth-code"

	tests := []struct {
		name      string
		responder httpmock.Responder
		expected  string
		wantErr   bool
	}{
		{
			name: "Successful token acquisition",
			responder: httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]interface{}{
				"access_token": "token123",
				"expires_in":   3600,
				"token_type":   "Bearer",
				"refresh_token": "refresh123",
			}),
			expected: "token123",
			wantErr:  false,
		},
		{
			name:      "Token request failed",
			responder: httpmock.NewStringResponder(http.StatusBadRequest, `{"error":"invalid_grant"}`),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", authority.TokenURL(), tt.responder)

			cred, err := NewAuthorizationCodeCredential(clientID, clientSecret, code, redirectURI, authority, nil)
			assert.NoError(t, err)

			token, err := cred.GetAuthorizationToken(context.Background(), nil, nil)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, token)
			}
		})
	}
}

func TestNewAuthorizationCodeAuthenticationProvider(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	authority := NewInstanceAuthority("dev12345")
	httpmock.RegisterResponder("POST", authority.TokenURL(),
		httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]interface{}{
			"access_token": "token123",
			"expires_in":   3600,
		}))

	provider, err := NewAuthorizationCodeAuthenticationProvider("clientID", "clientSecret", "code", "http://localhost", authority, nil)
	assert.NoError(t, err)
	assert.NotNil(t, provider)

	request := abstractions.NewRequestInformation()
	u, _ := url.Parse("https://dev12345.service-now.com/api/now/table/incident")
	request.SetUri(*u)

	err = provider.AuthenticateRequest(context.Background(), request, nil)
	assert.NoError(t, err)

	authHeader := request.Headers.Get("Authorization")
	assert.NotEmpty(t, authHeader)
	assert.Equal(t, "Bearer token123", authHeader[0])
}
