package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestDisplayValue(t *testing.T) {
	tests := []internal.Test[string]{
		{
			Title:    "True",
			Input:    TRUE,
			Expected: "true",
		},
		{
			Title:    "False",
			Input:    FALSE,
			Expected: "false",
		},
		{
			Title:    "All",
			Input:    ALL,
			Expected: "all",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			value := string(tt.Input.(DisplayValue))
			assert.Equal(t, tt.Expected, value)
		})
	}
}
