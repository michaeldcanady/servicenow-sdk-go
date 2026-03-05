package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenCredential(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		baseURL      string
		expectedErr  error
	}{
		{
			name:         "Valid input",
			clientID:     "clientID",
			clientSecret: "clientSecret",
			baseURL:      "http://example.com",
			expectedErr:  nil,
		},
		{
			name:         "Empty client ID",
			clientID:     "",
			clientSecret: "clientSecret",
			baseURL:      "http://example.com",
			expectedErr:  EmptyClientID,
		},
		{
			name:         "Empty client secret",
			clientID:     "clientID",
			clientSecret: "",
			baseURL:      "http://example.com",
			expectedErr:  EmptyClientSecret,
		},
		{
			name:         "Empty base URL",
			clientID:     "clientID",
			clientSecret: "clientSecret",
			baseURL:      "",
			expectedErr:  EmptyBaseURL,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			credential, err := NewTokenCredential(test.clientID, test.clientSecret, test.baseURL, nil)
			if test.expectedErr != nil {
				assert.ErrorIs(t, err, test.expectedErr)
				assert.Nil(t, credential)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, credential)
			}
		})
	}
}
