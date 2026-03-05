package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemPutRequestConfiguration_toConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *TableItemPutRequestConfiguration
		expected *core.RequestConfiguration
	}{
		{
			name: "ValidConfiguration",
			config: &TableItemPutRequestConfiguration{
				Header: map[string]string{"Authorization": "Bearer token"},
				QueryParameters: &TableItemRequestBuilderPutQueryParameters{
					DisplayValue:         "true",
					ExcludeReferenceLink: true,
					Fields:               []string{"field1", "field2"},
					InputDisplayValue:    true,
					QueryNoDomain:        true,
					View:                 "desktop",
				},
				Data:         map[string]interface{}{"key": "value"},
				ErrorMapping: core.ErrorMapping{"4XX": "error"},
				response:     &TableItemResponse{},
			},
			expected: &core.RequestConfiguration{
				Header: map[string]string{"Authorization": "Bearer token"},
				QueryParameters: &TableItemRequestBuilderPutQueryParameters{
					DisplayValue:         "true",
					ExcludeReferenceLink: true,
					Fields:               []string{"field1", "field2"},
					InputDisplayValue:    true,
					QueryNoDomain:        true,
					View:                 "desktop",
				},
				Data:         map[string]interface{}{"key": "value"},
				ErrorMapping: core.ErrorMapping{"4XX": "error"},
				Response:     &TableItemResponse{},
			},
		},
		{
			name: "ValidConfigurationWithNilParams",
			config: &TableItemPutRequestConfiguration{
				Header:          nil,
				QueryParameters: nil,
				Data:            nil,
				ErrorMapping:    nil,
				response:        nil,
			},
			expected: &core.RequestConfiguration{
				Header:          nil,
				QueryParameters: (*TableItemRequestBuilderPutQueryParameters)(nil),
				Data:            nil,
				ErrorMapping:    nil,
				Response:        (*TableItemResponse)(nil),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.toConfiguration())
		})
	}
}
