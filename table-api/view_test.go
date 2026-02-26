package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView2_String(t *testing.T) {
	tests := []struct {
		name     string
		value    View
		expected string
	}{
		{
			name:     "Unknown value",
			value:    ViewUnknown,
			expected: "unknown",
		},
		{
			name:     "Desktop value",
			value:    ViewDesktop,
			expected: "desktop",
		},
		{
			name:     "Mobile value",
			value:    ViewMobile,
			expected: "mobile",
		},
		{
			name:     "Both value",
			value:    ViewBoth,
			expected: "both",
		},
		{
			name:     "Invalid value falls back to Unknown",
			value:    View(42),
			expected: ViewUnknown.String(), // fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}
