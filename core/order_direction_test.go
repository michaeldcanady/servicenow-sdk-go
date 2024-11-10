package core

import "testing"

func TestOrderDirections(t *testing.T) {
	testCases := []struct {
		direction OrderDirection
		expected  string
	}{
		{Unset, ""},
		{Asc, "^ORDERBY"},
		{Desc, "^ORDERBYDESC"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.direction), func(t *testing.T) {
			result := string(tc.direction)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
