package attachmentapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime_UnmarshalJSON_Success(t *testing.T) {

	rawJson := []byte(`2006-01-02 15:04:05`)

	actual := &Time{}

	err := actual.UnmarshalJSON(rawJson)
	if err != nil {
		t.Error(err)
	}

	expectedRaw, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	if err != nil {
		t.Error(err)
	}

	expected := Time(expectedRaw)

	assert.Equal(t, &expected, actual)
}

func TestTime_UnmarshallJSON_Failed(t *testing.T) {
	rawJson := []byte(`2006-01-02 15:04:053`)

	actual := &Time{}

	err := actual.UnmarshalJSON(rawJson)
	if err != nil {
		assert.Error(t, err)
	}
}
