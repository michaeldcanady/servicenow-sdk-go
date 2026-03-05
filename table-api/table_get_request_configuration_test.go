package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestToConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		config   *TableGetRequestConfiguration
		expected *core.RequestConfiguration
	}{
		{
			name: "ValidConfiguration",
			config: &TableGetRequestConfiguration{
				Header:          map[string]string{"Authorization": "Bearer token"},
				QueryParameters: &TableRequestBuilderGetQueryParameters{Limit: 10},
				Data:            map[string]interface{}{"key": "value"},
				ErrorMapping:    core.ErrorMapping{"4XX": "error"},
				response:        &TableCollectionResponse{},
			},
			expected: &core.RequestConfiguration{
				Header:          map[string]string{"Authorization": "Bearer token"},
				QueryParameters: &TableRequestBuilderGetQueryParameters{Limit: 10},
				Data:            map[string]interface{}{"key": "value"},
				ErrorMapping:    core.ErrorMapping{"4XX": "error"},
				Response:        &TableCollectionResponse{},
			},
		},
		{
			name: "NilQueryParameters",
			config: &TableGetRequestConfiguration{
				Header:       map[string]string{"Authorization": "Bearer token"},
				Data:         map[string]interface{}{"key": "value"},
				ErrorMapping: core.ErrorMapping{"4XX": "error"},
				response:     &TableCollectionResponse{},
			},
			expected: &core.RequestConfiguration{
				Header:          map[string]string{"Authorization": "Bearer token"},
				QueryParameters: (*TableRequestBuilderGetQueryParameters)(nil),
				Data:            map[string]interface{}{"key": "value"},
				ErrorMapping:    core.ErrorMapping{"4XX": "error"},
				Response:        &TableCollectionResponse{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.config.toConfiguration())
		})
	}
}
