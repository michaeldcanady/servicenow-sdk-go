package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayValue2_String(t *testing.T) {
	tests := []struct {
		name     string
		value    DisplayValue2
		expected string
	}{
		{
			name:     "Unknown value",
			value:    DisplayValue2Unknown,
			expected: "unknown",
		},
		{
			name:     "True value",
			value:    DisplayValue2True,
			expected: "true",
		},
		{
			name:     "False value",
			value:    DisplayValue2False,
			expected: "false",
		},
		{
			name:     "All value",
			value:    DisplayValue2All,
			expected: "all",
		},
		{
			name:     "Invalid value falls back to Unknown",
			value:    DisplayValue2(99),
			expected: DisplayValue2Unknown.String(), // fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}
