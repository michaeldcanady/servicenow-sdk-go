package attachmentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolUnmarshalJSON_Success(t *testing.T) {
	rawJSON := []byte(`"true"`)

	boolVal := Bool(false)

	err := (&boolVal).UnmarshalJSON(rawJSON)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, Bool(true), boolVal)
}

func TestBoolUnmarshalJSON_Failed(t *testing.T) {
	rawJSON := []byte(`"w"`)

	boolVal := Bool(false)

	err := (&boolVal).UnmarshalJSON(rawJSON)
	assert.Error(t, err)
}
