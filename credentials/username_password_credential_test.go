package credentials

import (
	"encoding/base64"
	"testing"
)

func TestNewUsernamePasswordCredential(t *testing.T) {
	username := "testuser"
	password := "testpassword"

	// Create a new UsernamePasswordCredential instance.
	credential := NewUsernamePasswordCredential(username, password)

	// Check if the credential fields are set correctly.
	if credential.Username != username {
		t.Errorf("Expected username '%s', got '%s'", username, credential.Username)
	}
	if credential.Password != password {
		t.Errorf("Expected password '%s', got '%s'", password, credential.Password)
	}
}

func TestUsernamePasswordCredential_GetAuthentication(t *testing.T) {
	username := "testuser"
	password := "testpassword"

	// Create a new UsernamePasswordCredential instance.
	credential := NewUsernamePasswordCredential(username, password)

	// Test the GetAuthentication method.
	authHeader, err := credential.GetAuthentication()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the generated authentication header is in the correct format.
	expectedAuthHeader := "Basic " + credential.BasicAuth(username, password)
	if authHeader != expectedAuthHeader {
		t.Errorf("Expected auth header '%s', got '%s'", expectedAuthHeader, authHeader)
	}
}

func TestUsernamePasswordCredential_BasicAuth(t *testing.T) {
	username := "testuser"
	password := "testpassword"

	// Create a new UsernamePasswordCredential instance.
	credential := NewUsernamePasswordCredential(username, password)

	// Test the BasicAuth method.
	authHeader := credential.BasicAuth(username, password)

	// Check if the generated Basic Authentication string is in the correct format.
	expectedAuth := username + ":" + password
	expectedEncodedAuth := base64.StdEncoding.EncodeToString([]byte(expectedAuth))
	if authHeader != expectedEncodedAuth {
		t.Errorf("Expected encoded auth '%s', got '%s'", expectedEncodedAuth, authHeader)
	}
}
