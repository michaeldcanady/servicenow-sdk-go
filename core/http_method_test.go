package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpMethod_String(t *testing.T) {
	tests := []struct {
		name     string
		method   HttpMethod
		expected string
	}{
		{
			name:     "Test GET",
			method:   GET,
			expected: "GET",
		},
		{
			name:     "Test POST",
			method:   POST,
			expected: "POST",
		},
		{
			name:     "Test PATCH",
			method:   PATCH,
			expected: "PATCH",
		},
		{
			name:     "Test DELETE",
			method:   DELETE,
			expected: "DELETE",
		},
		{
			name:     "Test OPTIONS",
			method:   OPTIONS,
			expected: "OPTIONS",
		},
		{
			name:     "Test CONNECT",
			method:   CONNECT,
			expected: "CONNECT",
		},
		{
			name:     "Test PUT",
			method:   PUT,
			expected: "PUT",
		},
		{
			name:     "Test TRACE",
			method:   TRACE,
			expected: "TRACE",
		},
		{
			name:     "Test HEAD",
			method:   HEAD,
			expected: "HEAD",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.method.String())
		})
	}
}
