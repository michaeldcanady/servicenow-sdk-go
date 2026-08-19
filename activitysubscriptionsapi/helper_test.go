package activitysubscriptionsapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNilOrEmpty(t *testing.T) {
	type payload struct {
		Name string
	}

	var (
		nilPointer   *string
		nilMap       map[string]string
		nilSlice     []string
		nilInterface interface{}
		emptyString  = ""
		filledString = "value"
	)

	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{name: "untyped nil", value: nil, expected: true},
		{name: "typed nil pointer", value: nilPointer, expected: true},
		{name: "nil map", value: nilMap, expected: true},
		{name: "nil slice", value: nilSlice, expected: true},
		{name: "nil interface", value: nilInterface, expected: true},
		{name: "empty string", value: "", expected: true},
		{name: "empty slice", value: []string{}, expected: true},
		{name: "empty map", value: map[string]string{}, expected: true},
		{name: "empty array", value: [0]string{}, expected: true},
		{name: "zero int", value: 0, expected: true},
		{name: "false bool", value: false, expected: true},
		{name: "zero struct", value: payload{}, expected: true},
		{name: "pointer to empty string", value: &emptyString, expected: true},
		{name: "non-empty string", value: "value", expected: false},
		{name: "pointer to non-empty string", value: &filledString, expected: false},
		{name: "non-empty slice", value: []string{"a"}, expected: false},
		{name: "non-empty map", value: map[string]string{"a": "b"}, expected: false},
		{name: "non-zero array", value: [1]string{"a"}, expected: false},
		{name: "non-zero int", value: 1, expected: false},
		{name: "true bool", value: true, expected: false},
		{name: "populated struct", value: payload{Name: "a"}, expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, isNilOrEmpty(test.value))
		})
	}
}
