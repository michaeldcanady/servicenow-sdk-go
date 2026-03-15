package credentials

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestClientCredentialsCredential(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		scopes       []string
		responder    httpmock.Responder
		expected     string
		wantErr      bool
	}{
		{
			name:         "Successful token acquisition",
			clientID:     "clientID",
			clientSecret: "clientSecret",
			scopes:       []string{"scope1", "scope2"},
			responder: httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]interface{}{
				"access_token": "token123",
				"expires_in":   3600,
				"token_type":   "Bearer",
			}),
			expected: "token123",
			wantErr:  false,
		},
		{
			name:         "Token request failed",
			clientID:     "clientID",
			clientSecret: "clientSecret",
			responder:    httpmock.NewStringResponder(http.StatusBadRequest, `{"error":"invalid_request"}`),
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authority := NewInstanceAuthority("dev12345")
			httpmock.RegisterResponder("POST", authority.TokenURL(), tt.responder)

			cred, err := NewClientCredentialsCredential(tt.clientID, tt.clientSecret, authority, nil, tt.scopes)
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

func TestNewClientCredentialsAuthenticationProvider(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	authority := NewInstanceAuthority("dev12345")
	httpmock.RegisterResponder("POST", authority.TokenURL(),
		httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]interface{}{
			"access_token": "token123",
			"expires_in":   3600,
		}))

	provider, err := NewClientCredentialsAuthenticationProvider("clientID", "clientSecret", authority, nil, nil)
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
