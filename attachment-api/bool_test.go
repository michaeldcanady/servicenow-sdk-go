package attachmentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_UnmarshalJSON_Success(t *testing.T) {

	rawJSON := []byte(`"true"`)

	boolVal := Bool(false)

	err := (&boolVal).UnmarshalJSON(rawJSON)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, Bool(true), boolVal)
}

func TestBool_UnmarshalJSON_Failed(t *testing.T) {
	rawJSON := []byte(`"w"`)

	boolVal := Bool(false)

	err := (&boolVal).UnmarshalJSON(rawJSON)
	assert.Error(t, err)
}
