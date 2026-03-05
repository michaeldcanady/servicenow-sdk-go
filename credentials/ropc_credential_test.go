package credentials

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUsernamePasswordCredential(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		username     string
		password     string
		handler      func(w http.ResponseWriter, r *http.Request)
		refresh      bool
		expected     string
		expectedNext string
	}{
		{
			name:         "Confidential client",
			clientID:     "clientID",
			clientSecret: "clientSecret",
			username:     "user",
			password:     "pass",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = r.ParseForm()
				if r.Form.Get("grant_type") == "password" {
					assert.Equal(t, "clientID", r.Form.Get("client_id"))
					assert.Equal(t, "clientSecret", r.Form.Get("client_secret"))
					assert.Equal(t, "user", r.Form.Get("username"))
					assert.Equal(t, "pass", r.Form.Get("password"))

					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(map[string]interface{}{
						"access_token":  "token1",
						"refresh_token": "refresh1",
						"expires_in":    3600,
					})
				} else if r.Form.Get("grant_type") == "refresh_token" {
					assert.Equal(t, "refresh1", r.Form.Get("refresh_token"))
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(map[string]interface{}{
						"access_token": "token2",
						"expires_in":   3600,
					})
				}
			},
			refresh:      true,
			expected:     "token1",
			expectedNext: "token2",
		},
		{
			name:         "Public client",
			clientID:     "clientID",
			clientSecret: "",
			username:     "user",
			password:     "pass",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_ = r.ParseForm()
				assert.Equal(t, "password", r.Form.Get("grant_type"))
				assert.Equal(t, "clientID", r.Form.Get("client_id"))
				assert.Empty(t, r.Form.Get("client_secret"))

				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(map[string]interface{}{
					"access_token": "public-token",
					"expires_in":   3600,
				})
			},
			refresh:  false,
			expected: "public-token",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(test.handler))
			defer server.Close()

			authority := Authority(server.URL)
			cred, err := NewROPCCredential(test.clientID, test.clientSecret, test.username, test.password, authority, nil)
			assert.NoError(t, err)

			uri, _ := url.Parse(server.URL)
			token, err := cred.GetAuthorizationToken(context.Background(), uri, nil)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, token)

			if test.refresh {
				if cred.token != nil {
					cred.token.ExpiresAt = time.Now().Add(-time.Hour)
				}
				token, err = cred.GetAuthorizationToken(context.Background(), uri, nil)
				assert.NoError(t, err)
				assert.Equal(t, test.expectedNext, token)
			}
		})
	}
}
