package tableapi

import (
	"reflect"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
)

func TestTableItemPutRequestConfiguration_toConfiguration(t *testing.T) {
	t.Run("ValidConfiguration", func(t *testing.T) {
		// Create a TableItemPutRequestConfiguration with valid parameters
		rC := &TableItemPutRequestConfiguration{
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
		}

		// Call the toConfiguration method
		config := rC.toConfiguration()

		// Check if the converted configuration is as expected
		expectedConfig := &core.RequestConfiguration{
			Header:          rC.Header,
			QueryParameters: rC.QueryParameters,
			Data:            rC.Data,
			ErrorMapping:    rC.ErrorMapping,
			Response:        rC.response,
		}

		if !reflect.DeepEqual(config, expectedConfig) {
			t.Fatalf("Expected configuration %v, got: %v", expectedConfig, config)
		}
	})

	t.Run("ValidConfigurationWithNilParams", func(t *testing.T) {
		// ... Similar to the previous test but with nil parameters ...
	})

	t.Run("InvalidConfiguration", func(t *testing.T) {
		// ... Test when the configuration is invalid (e.g., missing required parameters) ...
	})
}
