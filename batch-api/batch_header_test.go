package batchapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBatchHeaderMarshal(t *testing.T) {
	header := batchHeader{
		Name:  "test",
		Value: "testy",
	}

	expectedJSON := "{\"name\":\"test\",\"value\":\"testy\"}"

	headerJSON, err := json.Marshal(header)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJSON, string(headerJSON))
}
