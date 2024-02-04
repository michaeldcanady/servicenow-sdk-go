package tableapi

import (
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableGetRequestConfiguration2_toConfiguration(t *testing.T) {
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    &TableRequestBuilderGetQueryParameters{},
		data:     nil,
		response: &TableCollectionResponse2[TableEntry]{},
	}

	expected := &core.RequestConfiguration{
		Header:          nil,
		QueryParameters: &TableRequestBuilderGetQueryParameters{},
		Data:            nil,
		Response:        &TableCollectionResponse2[TableEntry]{},
	}

	assert.Equal(t, expected, config.toConfiguration())
}
