package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceNowError_Error(t *testing.T) {
	e := ServiceNowError{
		Exception: Exception{
			Message: "Some error",
			Detail:  "Error details",
		},
		Status: "500 Internal Server Error",
	}

	expectedErrorMessage := "Some error: Error details"
	assert.Equal(t, expectedErrorMessage, e.Error())
}
