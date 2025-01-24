package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplayValue2_String(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{
		{
			Title:    "Unknown",
			Input:    -1,
			Err:      nil,
			Expected: DisplayValue2Unknown,
		},
		{
			Title:    "True",
			Input:    0,
			Err:      nil,
			Expected: DisplayValue2True,
		},
		{
			Title:    "False",
			Input:    1,
			Err:      nil,
			Expected: DisplayValue2False,
		},
		{
			Title:    "All",
			Input:    2,
			Err:      nil,
			Expected: DisplayValue2All,
		},
		{
			Title:    "Invalid",
			Input:    3,
			Err:      nil,
			Expected: DisplayValue2Unknown,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			input, ok := test.Input.(int)
			if !ok {
				t.Error("test.Input is not int")
			}
			displayValue := DisplayValue2(input)
			// using strings since you can cast any int to this type
			assert.Equal(t, test.Expected.(DisplayValue2).String(), displayValue.String())
		})
	}
}

func TestParseDisplayValue2(t *testing.T) {
	tests := []struct {
		Title    string
		Input    interface{}
		Err      error
		Expected interface{}
	}{
		{
			Title:    "Unknown",
			Input:    "unknown",
			Err:      nil,
			Expected: DisplayValue2Unknown,
		},
		{
			Title:    "True",
			Input:    "true",
			Err:      nil,
			Expected: DisplayValue2True,
		},
		{
			Title:    "False",
			Input:    "false",
			Err:      nil,
			Expected: DisplayValue2False,
		},
		{
			Title:    "All",
			Input:    "all",
			Err:      nil,
			Expected: DisplayValue2All,
		},
		{
			Title:    "Invalid",
			Input:    "invalid",
			Err:      nil,
			Expected: DisplayValue2Unknown,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			input, ok := test.Input.(string)
			if !ok {
				t.Error("test.Input is not string")
			}

			displayValue := ParseDisplayValue2(input)
			assert.Equal(t, test.Expected, displayValue)
		})
	}
}
