package core

import (
	"fmt"
	"testing"
)

func TestOrderByString(t *testing.T) {
	testCases := []struct {
		order    OrderBy
		expected string
	}{
		{OrderBy{Unset, "fieldName"}, ""},
		{OrderBy{Asc, "fieldName"}, "^ORDERBYfieldName"},
		{OrderBy{Desc, "otherField"}, "^ORDERBYDESCotherField"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("String() with Direction %s and Field %s", tc.order.Direction, tc.order.Field), func(t *testing.T) {
			result := tc.order.String()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestNewOrderBy(t *testing.T) {
	oB := NewOrderBy()
	if oB.Direction != Unset || oB.Field != "" {
		t.Error("Expected Direction and Field to be unset and empty, got", oB)
	}
}
