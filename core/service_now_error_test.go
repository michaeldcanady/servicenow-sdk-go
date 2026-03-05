package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceNowError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      ServiceNowError
		expected string
	}{
		{
			name: "Standard error",
			err: ServiceNowError{
				Exception: Exception{
					Message: "Some error",
					Detail:  "Error details",
				},
				Status: "500 Internal Server Error",
			},
			expected: "Some error: Error details",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.err.Error())
		})
	}
}
