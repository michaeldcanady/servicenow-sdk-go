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

func TestUsernamePasswordCredential_Confidential(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	}))
	defer server.Close()

	authority := Authority(server.URL)
	cred, err := NewROPCCredential("clientID", "clientSecret", "user", "pass", authority, nil)
	assert.NoError(t, err)

	uri, _ := url.Parse(server.URL)
	token, err := cred.GetAuthorizationToken(context.Background(), uri, nil)
	assert.NoError(t, err)
	assert.Equal(t, "token1", token)

	// Refresh
	if cred.token != nil {
		cred.token.ExpiresAt = time.Now().Add(-time.Hour)
	}
	token, err = cred.GetAuthorizationToken(context.Background(), uri, nil)
	assert.NoError(t, err)
	assert.Equal(t, "token2", token)
}

func TestUsernamePasswordCredential_Public(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		assert.Equal(t, "password", r.Form.Get("grant_type"))
		assert.Equal(t, "clientID", r.Form.Get("client_id"))
		assert.Empty(t, r.Form.Get("client_secret"))

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "public-token",
			"expires_in":   3600,
		})
	}))
	defer server.Close()

	authority := Authority(server.URL)
	cred, err := NewROPCCredential("clientID", "", "user", "pass", authority, nil)
	assert.NoError(t, err)

	uri, _ := url.Parse(server.URL)
	token, err := cred.GetAuthorizationToken(context.Background(), uri, nil)
	assert.NoError(t, err)
	assert.Equal(t, "public-token", token)
}

func rURL(s string) *url.URL {
	u, _ := url.Parse(s)
	return u
}
