package attachmentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_UnmarshalJSON_Success(t *testing.T) {

	rawJson := []byte(`"true"`)

	boolVal := Bool(false)

	err := (&boolVal).UnmarshalJSON(rawJson)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, Bool(true), boolVal)
}

func TestBool_UnmarshalJSON_Failed(t *testing.T) {
	rawJson := []byte(`"w"`)

	boolVal := Bool(false)

	err := (&boolVal).UnmarshalJSON(rawJson)
	assert.Error(t, err)
}
