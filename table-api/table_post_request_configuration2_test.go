package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTablePostRequestConfiguration2_toConfiguration(t *testing.T) {
	config := &tablePostRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    &TableRequestBuilderPostQueryParameters{},
		data:     map[string]string{"key1": "value1"},
		response: &TableItemResponse2[TableEntry]{},
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
