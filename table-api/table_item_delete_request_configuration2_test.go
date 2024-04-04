package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemDeleteRequestConfiguration2_toConfiguration(t *testing.T) {
	config := &tableItemDeleteRequestConfiguration2[TableEntry]{
		header:   nil,
		query:    &TableItemRequestBuilderDeleteQueryParameters{},
		data:     nil,
		response: nil,
	}

	expected := &core.RequestConfiguration{
		Header:          nil,
		QueryParameters: &TableItemRequestBuilderDeleteQueryParameters{},
		Data:            nil,
		Response:        (*TableItemResponse2[TableEntry])(nil),
	}

	assert.Equal(t, expected, config.toConfiguration())
}
