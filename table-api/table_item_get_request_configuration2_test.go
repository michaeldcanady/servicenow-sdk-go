package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemGetRequestConfiguration2_toConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *tableItemGetRequestConfiguration2[TableEntry]
		expected *core.RequestConfiguration
	}{
		{
			name: "Standard configuration",
			config: &tableItemGetRequestConfiguration2[TableEntry]{
				header:   nil,
				query:    &TableItemRequestBuilderGetQueryParameters{},
				data:     nil,
				response: &TableItemResponse2[TableEntry]{},
			},
			expected: &core.RequestConfiguration{
				Header:          nil,
				QueryParameters: &TableItemRequestBuilderGetQueryParameters{},
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
