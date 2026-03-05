package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemPutRequestConfiguration2_toConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *tableItemPutRequestConfiguration2[TableEntry]
		expected *core.RequestConfiguration
	}{
		{
			name: "Standard configuration",
			config: &tableItemPutRequestConfiguration2[TableEntry]{
				header:   nil,
				query:    &TableItemRequestBuilderPutQueryParameters{},
				data:     nil,
				response: &TableItemResponse2[TableEntry]{},
			},
			expected: &core.RequestConfiguration{
				Header:          nil,
				QueryParameters: &TableItemRequestBuilderPutQueryParameters{},
				Data:            nil,
				Response:        &TableItemResponse2[TableEntry]{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.toConfiguration())
		})
	}
}
