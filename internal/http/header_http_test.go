package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPHeader(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "HTTPHeaderUnknown",
			test: func(t *testing.T) {
				assert.Equal(t, "unknown", HTTPHeaderUnknown.String())
			},
		},
		{
			name: "HTTPHeaderContentType",
			test: func(t *testing.T) {
				assert.Equal(t, "content-type", HTTPHeaderContentType.String())
			},
		},
		{
			name: "Unknown",
			test: func(t *testing.T) {
				assert.Equal(t, "unknown", HTTPHeader(-9999).String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
