package tableapi

import (
	"reflect"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTablePostRequestConfigurationToConfiguration(t *testing.T) {
	t.Run("ValidConfiguration", func(t *testing.T) {
		rC := &TablePostRequestConfiguration{
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
		}

		config := rC.toConfiguration()

		// Validate each field of the configuration
		if !reflect.DeepEqual(config.Header, rC.Header) {
			t.Fatalf("Expected header %v, got: %v", rC.Header, config.Header)
		}

		if !reflect.DeepEqual(config.QueryParameters, rC.QueryParameters) {
			t.Fatalf("Expected query parameters %v, got: %v", rC.QueryParameters, config.QueryParameters)
		}

		if !reflect.DeepEqual(config.Data, rC.Data) {
			t.Fatalf("Expected data %v, got: %v", rC.Data, config.Data)
		}

		if !reflect.DeepEqual(config.ErrorMapping, rC.ErrorMapping) {
			t.Fatalf("Expected error mapping %v, got: %v", rC.ErrorMapping, config.ErrorMapping)
		}

		if !reflect.DeepEqual(config.Response, rC.response) {
			t.Fatalf("Expected response %v, got: %v", rC.response, config.Response)
		}
	})

	t.Run("NilQueryParameters", func(t *testing.T) {
		rC := &TablePostRequestConfiguration{
			Header: map[string]string{"Authorization": "Bearer token"},
			// Nil QueryParameters
			Data:         map[string]string{"key": "value"},
			ErrorMapping: core.ErrorMapping{"4XX": "error"},
			response:     &TableItemResponse{},
		}

		config := rC.toConfiguration()

		// Validate that nil QueryParameters are handled
		assert.Equal(t, (*TableRequestBuilderPostQueryParameters)(nil), config.QueryParameters)
	})
}
