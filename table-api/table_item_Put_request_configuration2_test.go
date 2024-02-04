package tableapi

import (
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemPutRequestConfiguration2_toConfiguration(t *testing.T) {
	config := &tableItemPutRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    &TableItemRequestBuilderPutQueryParameters{},
		data:     nil,
		response: &TableItemResponse2[TableEntry]{},
	}

	expected := &core.RequestConfiguration{
		Header:          nil,
		QueryParameters: &TableItemRequestBuilderPutQueryParameters{},
		Data:            nil,
		Response:        &TableItemResponse2[TableEntry]{},
	}

	assert.Equal(t, expected, config.toConfiguration())
}
