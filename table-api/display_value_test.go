package tableapi

import (
	"testing"

	inttesting "github.com/michaeldcanady/servicenow-sdk-go/internal/testing"
	"github.com/stretchr/testify/assert"
)

func TestDisplayValue(t *testing.T) {
	testCases := []inttesting.InternalTest{
		{
			Title:     "TRUE",
			Input:     string(TRUE),
			Expected:  "true",
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "FALSE",
			Input:     string(FALSE),
			Expected:  "false",
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "ALL",
			Input:     string(ALL),
			Expected:  "all",
			ShouldErr: false,
			Error:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			result := tc.Input
			if result != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, result)
			}
		})
	}
}

func TestDisplayOption(t *testing.T) {
	testCases := []inttesting.InternalTest{
		{
			Title:     "DisplayTrue_String",
			Input:     DisplayTrue.String(),
			Expected:  "true",
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "DisplayTrue_Int",
			Input:     int(DisplayTrue),
			Expected:  1,
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "DisplayFalse_String",
			Input:     DisplayFalse.String(),
			Expected:  "false",
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "DisplayFalse_Int",
			Input:     int(DisplayFalse),
			Expected:  2,
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "DisplayAll_String",
			Input:     DisplayAll.String(),
			Expected:  "all",
			ShouldErr: false,
			Error:     nil,
		},
		{
			Title:     "DisplayAll_Int",
			Input:     int(DisplayAll),
			Expected:  3,
			ShouldErr: false,
			Error:     nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			result := tc.Input
			assert.Equal(t, tc.Expected, result)
		})
	}
}
