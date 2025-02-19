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
			input:    View2Desktop,
			expected: "desktop",
		},
		{
			title:    "MOBILE",
			input:    View2Mobile,
			expected: "mobile",
		},
		{
			title:    "BOTH",
			input:    View2Both,
			expected: "both",
		},
		{
			title:    "unknown",
			input:    View2(-999),
			expected: "unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := test.input.String()
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestParseView2(t *testing.T) {
	tests := []struct {
		title    string
		input    string
		expected View2
	}{
		{
			title:    "successful",
			input:    "both",
			expected: View2Both,
		},
		{
			title:    "unknown",
			input:    "help",
			expected: View2Unknown,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result := ParseView2(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
