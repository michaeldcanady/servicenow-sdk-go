package tests

import (
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// MockPrompt is a mock implementation of the user prompt function for testing.
func MockPrompt() (string, string, error) {
	return "testuser", "testpassword", nil
}

func TestNewTokenCredential(t *testing.T) {
	// Test valid input.
	credential, err := credentials.NewTokenCredential("clientID", "clientSecret", "http://example.com", nil)
	if err != nil {
		t.Errorf("Expected no error, got '%v'", err)
	}
	if credential == nil {
		t.Error("Expected non-nil credential, got nil")
	}

	// Test empty client ID.
	_, err = credentials.NewTokenCredential("", "clientSecret", "http://example.com", nil)
	if err == nil {
		t.Error("Expected error for empty client Id, got nil")
	} else if !errors.Is(err, credentials.NewCredentialError("clientId is empty")) {
		t.Errorf("Expected 'clientId is empty' error, got '%v'", err)
	}

	// Test empty client secret.
	_, err = credentials.NewTokenCredential("clientID", "", "http://example.com", nil)
	if err == nil {
		t.Error("Expected error for empty client secret, got nil")
	} else if !errors.Is(err, credentials.NewCredentialError("clientSecret is empty")) {
		t.Errorf("Expected 'clientSecret is empty' error, got '%v'", err)
	}

	// Test empty base URL.
	_, err = credentials.NewTokenCredential("clientID", "clientSecret", "", nil)
	if err == nil {
		t.Error("Expected error for empty base URL, got nil")
	} else if !errors.Is(err, credentials.NewCredentialError("baseURL is empty")) {
		t.Errorf("Expected 'baseURL is empty' error, got '%v'", err)
	}
}

func TestTokenCredential_GetAuthentication(t *testing.T) {
	// Create a TokenCredential with a mock prompt function.
	credential, _ := credentials.NewTokenCredential("clientID", "clientSecret", "http://example.com", MockPrompt)

	// Test obtaining a new access token.
	authHeader, err := credential.GetAuthentication()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if authHeader == "" {
		t.Error("Expected non-empty auth header, got empty")
	}

	// Test refreshing an expired token (mocked).
	credential.Token = &credentials.AccessToken{
		AccessToken: "expired_token",
		ExpiresIn:   1, // 1 second expiration for testing.
	}
	authHeader, err = credential.GetAuthentication()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if authHeader == "" {
		t.Error("Expected non-empty auth header, got empty")
	}
}

func TestTokenCredential_GetAuthentication_NoPrompt(t *testing.T) {
	// Create a TokenCredential with no prompt function.
	credential, _ := credentials.NewTokenCredential("clientID", "clientSecret", "http://example.com", nil)

	// Test obtaining a new access token without a prompt function.
	_, err := credential.GetAuthentication()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestTokenCredential_GetAuthentication_RefreshError(t *testing.T) {
	// Create a TokenCredential with a mock prompt function.
	credential, _ := credentials.NewTokenCredential("clientID", "clientSecret", "http://example.com", MockPrompt)

	// Set an expired token.
	credential.Token = &credentials.AccessToken{
		AccessToken: "expired_token",
		ExpiresIn:   1, // 1 second expiration for testing.
	}

	// Test refreshing a token with an error (mocked).
	_, err := credential.GetAuthentication()
	if err == nil {
		t.Error("Expected error when refreshing token, got nil")
	}
}
