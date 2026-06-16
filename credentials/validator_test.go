package credentials

import (
	"net/url"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		name         string
		allowedHosts []string
		url          string
		expected     bool
	}{
		{
			name:         "Exact match",
			allowedHosts: []string{"example.com"},
			url:          "https://example.com/api",
			expected:     true,
		},
		{
			name:         "Match with port",
			allowedHosts: []string{"example.com"},
			url:          "https://example.com:443/api",
			expected:     true,
		},
		{
			name:         "Empty allowed hosts",
			allowedHosts: []string{},
			url:          "https://example.com/api",
			expected:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			v, err := authentication.NewAllowedHostsValidatorErrorCheck(test.allowedHosts)
			assert.NoError(t, err)
			u, err := url.Parse(test.url)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, v.IsUrlHostValid(u))
		})
	}
}
