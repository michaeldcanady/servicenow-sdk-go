package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestHeader_String(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "RequestHeaderAuthorization",
			test: func(t *testing.T) {
				assert.Equal(t, "authorization", RequestHeaderAuthorization.String())
			},
		},
		{
			name: "RequestHeaderUnknown",
			test: func(t *testing.T) {
				assert.Equal(t, "unknown", RequestHeaderUnknown.String())
			},
		},
		{
			name: "RequestHeaderAccept",
			test: func(t *testing.T) {
				assert.Equal(t, "accept", RequestHeaderAccept.String())
			},
		},
		{
			name: "unknown RequestHeader",
			test: func(t *testing.T) {
				assert.Equal(t, "unknown", RequestHeader(-2).String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
