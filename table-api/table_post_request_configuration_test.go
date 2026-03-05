package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTablePostRequestConfigurationToConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *TablePostRequestConfiguration
		expected *core.RequestConfiguration
	}{
		{
			name: "ValidConfiguration",
			config: &TablePostRequestConfiguration{
				Header: map[string]string{"Authorization": "Bearer token"},
				QueryParameters: &TableRequestBuilderPostQueryParameters{
					DisplayValue:      "true",
					Fields:            []string{"field1", "field2"},
					InputDisplayValue: true,
					View:              "desktop",
				},
				Data:         map[string]string{"key": "value"},
				ErrorMapping: core.ErrorMapping{"4XX": "error"},
				response:     &TableItemResponse{},
			},
			expected: &core.RequestConfiguration{
				Header: map[string]string{"Authorization": "Bearer token"},
				QueryParameters: &TableRequestBuilderPostQueryParameters{
					DisplayValue:      "true",
					Fields:            []string{"field1", "field2"},
					InputDisplayValue: true,
					View:              "desktop",
				},
				Data:         map[string]string{"key": "value"},
				ErrorMapping: core.ErrorMapping{"4XX": "error"},
				Response:     &TableItemResponse{},
			},
		},
		{
			name: "NilQueryParameters",
			config: &TablePostRequestConfiguration{
				Header:       map[string]string{"Authorization": "Bearer token"},
				Data:         map[string]string{"key": "value"},
				ErrorMapping: core.ErrorMapping{"4XX": "error"},
				response:     &TableItemResponse{},
			},
			expected: &core.RequestConfiguration{
				Header:          map[string]string{"Authorization": "Bearer token"},
				QueryParameters: (*TableRequestBuilderPostQueryParameters)(nil),
				Data:            map[string]string{"key": "value"},
				ErrorMapping:    core.ErrorMapping{"4XX": "error"},
				Response:        &TableItemResponse{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.toConfiguration())
		})
	}
}
