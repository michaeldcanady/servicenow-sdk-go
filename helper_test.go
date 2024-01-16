package servicenowsdkgo

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestValidateURI(t *testing.T) {

	test := []internal.Test[string]{
		{
			Title:    "Valid",
			Input:    "https://instance.service-now.com/api",
			Expected: "https://instance.service-now.com/api",
			Error:    nil,
		},
		{
			Title:    "Valid",
			Input:    "https://instance.service-now.com",
			Expected: "https://instance.service-now.com/api",
			Error:    nil,
		},
		{
			Title:    "Valid",
			Input:    "instance.service-now.com",
			Expected: "https://instance.service-now.com/api",
			Error:    nil,
		},
		{
			Title:    "Valid",
			Input:    "instance",
			Expected: "https://instance.service-now.com/api",
			Error:    nil,
		},
		{
			Title:    "Valid",
			Input:    "https://instance.service-now.com:404",
			Expected: "https://instance.service-now.com:404/api",
			Error:    nil,
		},
		{
			Title:    "Valid",
			Input:    "https://www.instance.service-now.com:404",
			Expected: "https://instance.service-now.com:404/api",
			Error:    nil,
		},
		{
			Title:    "Missing Host",
			Input:    "https://",
			Expected: "https://instance.service-now.com:404/api",
			Error:    ErrMissingHost,
		},
		{
			Title:    "Incorrect path string",
			Input:    "https://instance.service-now.com/wrong",
			Expected: "",
			Error:    ErrInvalidURIPath,
		},
	}

	for _, tt := range test {

		output, err := validateURI(tt.Input.(string))

		assert.Equal(t, tt.Error, err)
		assert.Equal(t, tt.Expected, output)
	}
}
