package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViewString(t *testing.T) {
	tests := []struct {
		title    string
		input    View
		expected string
	}{
		{
			title:    "DESKTOP",
			input:    DESKTOP,
			expected: "desktop",
		},
		{
			title:    "MOBILE",
			input:    MOBILE,
			expected: "mobile",
		},
		{
			title:    "BOTH",
			input:    BOTH,
			expected: "both",
		},
		{
			title:    "DESKTOP",
			input:    ViewDesktop,
			expected: "desktop",
		},
		{
			title:    "MOBILE",
			input:    ViewMobile,
			expected: "mobile",
		},
		{
			title:    "BOTH",
			input:    ViewBoth,
			expected: "both",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := test.input
			assert.Equal(t, test.expected, string(result))
		})
	}
}
