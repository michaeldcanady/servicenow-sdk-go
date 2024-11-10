package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView2String(t *testing.T) {
	tests := []struct {
		title    string
		input    View2
		expected string
	}{
		{
			title:    "DESKTOP",
			input:    ViewDesktop2,
			expected: "desktop",
		},
		{
			title:    "MOBILE",
			input:    ViewMobile2,
			expected: "mobile",
		},
		{
			title:    "BOTH",
			input:    ViewBoth2,
			expected: "both",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := test.input.String()
			assert.Equal(t, test.expected, result)
		})
	}
}
