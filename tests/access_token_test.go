package tests

import (
	"testing"
	"time"

	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func TestAccessToken_IsExpired(t *testing.T) {
	tests := []struct {
		name     string
		token    credentials.AccessToken
		expected bool
	}{
		{
			name: "ExpiredToken",
			token: credentials.AccessToken{
				ExpiresAt: time.Now().Add(-time.Minute), // An expired token
			},
			expected: true,
		},
		{
			name: "NonExpiredToken",
			token: credentials.AccessToken{
				ExpiresAt: time.Now().Add(time.Minute), // A non-expired token
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.token.IsExpired(); got != tt.expected {
				t.Errorf("AccessToken.IsExpired() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
