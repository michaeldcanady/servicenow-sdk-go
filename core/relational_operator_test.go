package core

import "testing"

func TestRelationalOperators(t *testing.T) {
	testCases := []struct {
		operator RelationalOperator
		expected string
	}{
		{Null, ""},
		{Is, "="},
		{IsNot, "!="},
		{GreaterThan, ">"},
		{GreaterOrEqual, ">="},
		{LessThan, "<"},
		{LessOrEqual, "<="},
		{Contains, "CONTAINS"},
		{NotContains, "!CONTAINS"},
		{StartsWith, "STARTSWITH"},
		{EndsWith, "ENDSWITH"},
		{Between, "BETWEEN"},
		{IsSame, "SAMEAS"},
		{IsDifferent, "NSAMEAS"},
		{IsEmpty, "ISEMPTY"},
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
