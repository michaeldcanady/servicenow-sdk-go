package tableapi

import "testing"

func TestDisplayValue(t *testing.T) {
	testCases := []struct {
		displayValue DisplayValue
		expected     string
	}{
		{TRUE, "true"},
		{FALSE, "false"},
		{ALL, "all"},
	}

	for _, tc := range testCases {
		t.Run("String() for DisplayValue", func(t *testing.T) {
			result := tc.displayValue
			if string(result) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
