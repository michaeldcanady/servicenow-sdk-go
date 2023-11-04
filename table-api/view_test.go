package tableapi

import "testing"

func TestViewString(t *testing.T) {
	tests := []struct {
		view     View
		expected string
	}{
		{DESKTOP, "desktop"},
		{MOBILE, "mobile"},
		{BOTH, "both"},
	}

	for _, test := range tests {
		t.Run(string(test.view), func(t *testing.T) {
			result := test.view
			if string(result) != test.expected {
				t.Errorf("Expected %s, got: %s", test.expected, result)
			}
		})
	}
}
