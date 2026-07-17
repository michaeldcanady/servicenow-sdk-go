package statsapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayValue_String(t *testing.T) {
	tests := []struct {
		name     string
		value    DisplayValue
		expected string
	}{
		{name: "Unknown", value: DisplayValueUnknown, expected: displayValueUnknown},
		{name: "True", value: DisplayValueTrue, expected: displayValueTrue},
		{name: "False", value: DisplayValueFalse, expected: displayValueFalse},
		{name: "All", value: DisplayValueAll, expected: displayValueAll},
		{name: "out of range falls back to unknown", value: DisplayValue(99), expected: displayValueUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.value.String())
		})
	}
}
