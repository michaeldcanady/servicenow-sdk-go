package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableGetRequestConfiguration2_toConfiguration(t *testing.T) {
	config := &tableGetRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    &TableRequestBuilderGetQueryParameters{},
		data:     nil,
		response: &TableCollectionResponse2[TableEntry]{},
	}

	// Call the toConfiguration method
	coreConfig := config.toConfiguration()

	// Assert that the fields in the returned core.RequestConfiguration match the original config
	assert.Equal(t, config.header, coreConfig.Header)
	assert.Equal(t, config.query, coreConfig.QueryParameters)
	assert.Equal(t, config.data, coreConfig.Data)
	assert.Equal(t, config.mapping, coreConfig.ErrorMapping)
	assert.Equal(t, config.response, coreConfig.Response)
}
