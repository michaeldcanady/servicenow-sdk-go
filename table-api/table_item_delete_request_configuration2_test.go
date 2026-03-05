package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemDeleteRequestConfiguration2_toConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *tableItemDeleteRequestConfiguration2[TableEntry]
		expected *core.RequestConfiguration
	}{
		{
			name: "Standard configuration",
			config: &tableItemDeleteRequestConfiguration2[TableEntry]{
				header:   nil,
				query:    &TableItemRequestBuilderDeleteQueryParameters{},
				data:     nil,
				response: nil,
			},
			expected: &core.RequestConfiguration{
				Header:          nil,
				QueryParameters: &TableItemRequestBuilderDeleteQueryParameters{},
				Data:            nil,
				Response:        (*TableItemResponse2[TableEntry])(nil),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.toConfiguration())
		})
	}
}
