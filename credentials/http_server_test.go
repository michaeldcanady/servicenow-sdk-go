package credentials

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPServer_StartStop(t *testing.T) {
	// Create a new HTTP server for testing with a random port.
	server := NewHTTPServer(":0")

	// Start the server in a goroutine.
	go server.Start()

	// Stop the server.
	server.Stop()
}

func TestOAuthRedirectHandler(t *testing.T) {
	// Create a test HTTP request.
	req := httptest.NewRequest("GET", "/oauth_redirect.do", nil)
	w := httptest.NewRecorder()

	// Call the OAuth redirect handler function.
	OauthRedirectHandler(w, req)

	// Check the response status code (200 OK is expected).
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Check the response body (expected message).
	expectedMessage := "OAuth2 token obtained successfully!"
	actualMessage := w.Body.String()
	if actualMessage != expectedMessage {
		t.Errorf("Expected response body '%s', got '%s'", expectedMessage, actualMessage)
	}
}
