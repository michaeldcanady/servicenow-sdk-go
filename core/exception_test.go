package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestException(t *testing.T) {
	ex := Exception{
		Message: "Sample message",
		Detail:  "Sample detail",
	}

	assert.Equal(t, "Sample message", ex.Message)
	assert.Equal(t, "Sample detail", ex.Detail)
}
