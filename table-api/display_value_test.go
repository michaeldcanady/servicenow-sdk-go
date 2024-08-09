package tableapi

import "testing"

func TestDisplayValue(t *testing.T) {
	testCases := []struct {
		title    string
		input    DisplayValue
		expected string
	}{
		{
			title:    "TRUE",
			input:    TRUE,
			expected: "true",
		},
		{
			title:    "FALSE",
			input:    FALSE,
			expected: "false",
		},
		{
			title:    "ALL",
			input:    ALL,
			expected: "all",
		},
		{
			title:    "DisplayValueTrue",
			input:    DisplayValueTrue,
			expected: "true",
		},
		{
			title:    "DisplayValueFalse",
			input:    DisplayValueFalse,
			expected: "false",
		},
		{
			title:    "DisplayValueAll",
			input:    DisplayValueAll,
			expected: "all",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.title, func(t *testing.T) {
			result := tc.input
			if string(result) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
