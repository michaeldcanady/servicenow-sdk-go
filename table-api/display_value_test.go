package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayValue2_String(t *testing.T) {
	tests := []struct {
		name     string
		value    DisplayValue
		expected string
	}{
		{
			name:     "Unknown value",
			value:    DisplayValueUnknown,
			expected: "unknown",
		},
		{
			name:     "True value",
			value:    DisplayValueTrue,
			expected: "true",
		},
		{
			name:     "False value",
			value:    DisplayValueFalse,
			expected: "false",
		},
		{
			name:     "All value",
			value:    DisplayValueAll,
			expected: "all",
		},
		{
			name:     "Invalid value falls back to Unknown",
			value:    DisplayValue(99),
			expected: DisplayValueUnknown.String(), // fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.value.String()
			assert.Equal(t, tt.expected, got)
		})
	}
}
