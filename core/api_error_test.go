package core

import (
	"testing"
)

func TestApiError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      *ApiError
		expected string
	}{
		{"WithMessage", &ApiError{Message: "failed"}, "hello"}, // Wait, I should use real values
		{"Default", &ApiError{}, "error status code received from the API"},
	}
	// Fixing the expected values
	tests[0].expected = "failed"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("got %s, expected %s", tt.err.Error(), tt.expected)
			}
		})
	}
}
