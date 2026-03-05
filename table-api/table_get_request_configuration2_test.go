package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableGetRequestConfiguration2_toConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *tableGetRequestConfiguration2[TableEntry]
		expected *core.RequestConfiguration
	}{
		{
			name: "Standard configuration",
			config: &tableGetRequestConfiguration2[TableEntry]{
				header:   nil,
				query:    &TableRequestBuilderGetQueryParameters{},
				data:     nil,
				response: &TableCollectionResponse2[TableEntry]{},
			},
			expected: &core.RequestConfiguration{
				Header:          nil,
				QueryParameters: &TableRequestBuilderGetQueryParameters{},
				Data:            nil,
				Response:        &TableCollectionResponse2[TableEntry]{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.toConfiguration())
		})
	}
}
