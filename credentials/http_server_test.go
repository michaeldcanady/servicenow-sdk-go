package credentials

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPServer_StartStop(t *testing.T) {
	tests := []struct {
		name string
		addr string
	}{
		{
			name: "Start and stop",
			addr: ":0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := NewHTTPServer(test.addr)
			go server.Start()
			server.Stop()
		})
	}
}

func TestOAuthRedirectHandler(t *testing.T) {
	tests := []struct {
		name            string
		req             *http.Request
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:            "Successful redirect",
			req:             httptest.NewRequest("GET", "/oauth_redirect.do", nil),
			expectedStatus:  http.StatusOK,
			expectedMessage: "OAuth2 token obtained successfully!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			OauthRedirectHandler(w, test.req)

			resp := w.Result()
			assert.Equal(t, test.expectedStatus, resp.StatusCode)
			assert.Equal(t, test.expectedMessage, w.Body.String())
		})
	}
}
