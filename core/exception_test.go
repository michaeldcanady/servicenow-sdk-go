package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestException(t *testing.T) {
	tests := []struct {
		name     string
		ex       Exception
		expected Exception
	}{
		{
			name: "Standard exception",
			ex: Exception{
				Message: "Sample message",
				Detail:  "Sample detail",
			},
			expected: Exception{
				Message: "Sample message",
				Detail:  "Sample detail",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected.Message, test.ex.Message)
			assert.Equal(t, test.expected.Detail, test.ex.Detail)
		})
	}
}
