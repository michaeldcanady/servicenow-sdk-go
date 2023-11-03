package tableapi

import (
	"testing"
)

func TestOperatorString(t *testing.T) {
	tests := []struct {
		op         Operator
		expected   string
		shouldFail bool
	}{
		{Is, "=", false},
		{IsNot, "!=", false},
		{GreaterThan, ">", false},
		{GreaterOrEqual, ">=", false},
		{LessThan, "<", false},
		{LessOrEqual, "<=", false},
		{Contains, "CONTAINS", false},
		{NotContains, "!CONTAINS", false},
		{StartsWith, "STARTSWITH", false},
		{EndsWith, "ENDSWITH", false},
		{Between, "BETWEEN", false},
		{IsSame, "SAMEAS", false},
		{IsDifferent, "NSAMEAS", false},
		{IsEmpty, "ISEMPTY", false},
		{100, "", true}, // Invalid Operator value
	}

	for _, test := range tests {
		result := test.op.String()
		if result != test.expected {
			t.Errorf("Expected %s for Operator %d, but got %s", test.expected, test.op, result)
		}
	}
}

func TestIsValidOperator(t *testing.T) {
	tests := []struct {
		op         Operator
		isValid    bool
		shouldFail bool
	}{
		{Is, true, false},
		{IsNot, true, false},
		{GreaterThan, true, false},
		{GreaterOrEqual, true, false},
		{LessThan, true, false},
		{LessOrEqual, true, false},
		{Contains, true, false},
		{NotContains, true, false},
		{StartsWith, true, false},
		{EndsWith, true, false},
		{Between, true, false},
		{IsSame, true, false},
		{IsDifferent, true, false},
		{IsEmpty, true, false},
		{100, false, true}, // Invalid Operator value
	}

	for _, test := range tests {
		result := IsValidOperator(test.op)
		if result != test.isValid {
			t.Errorf("Expected IsValidOperator(%d) to be %t, but got %t", test.op, test.isValid, result)
		}
	}
}
