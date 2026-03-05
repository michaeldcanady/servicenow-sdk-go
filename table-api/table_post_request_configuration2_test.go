package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTablePostRequestConfiguration2_toConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *tablePostRequestConfiguration2[TableEntry]
		expected *core.RequestConfiguration
	}{
		{
			name: "Standard configuration",
			config: &tablePostRequestConfiguration2[TableEntry]{
				header:   nil,
				query:    &TableRequestBuilderPostQueryParameters{},
				data:     map[string]string{"key1": "value1"},
				response: &TableItemResponse2[TableEntry]{},
			},
			expected: &core.RequestConfiguration{
				Header:          nil,
				QueryParameters: &TableRequestBuilderPostQueryParameters{},
				Data:            map[string]string{"key1": "value1"},
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
