package core

import "testing"

func TestLogicalOperators(t *testing.T) {
	testCases := []struct {
		operator LogicalOperator
		expected string
	}{
		{And, "^"},
		{Or, "^OR"},
	}

	for _, tc := range testCases {
		t.Run(string(tc.operator), func(t *testing.T) {
			result := string(tc.operator)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
