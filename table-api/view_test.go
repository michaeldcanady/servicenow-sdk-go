package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView2_String(t *testing.T) {
	tests := []struct {
		name     string
		value    View2
		expected string
	}{
		{
			name:     "Unknown value",
			value:    View2Unknown,
			expected: "unknown",
		},
		{
			name:     "Desktop value",
			value:    View2Desktop,
			expected: "desktop",
		},
		{
			name:     "Mobile value",
			value:    View2Mobile,
			expected: "mobile",
		},
		{
			name:     "Both value",
			value:    View2Both,
			expected: "both",
		},
		{
			name:     "Invalid value falls back to Unknown",
			value:    View2(42),
			expected: View2Unknown.String(), // fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}
