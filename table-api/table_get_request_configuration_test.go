package tableapi

import (
	"reflect"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestToConfiguration(t *testing.T) {
	t.Run("ValidConfiguration", func(t *testing.T) {
		rC := &TableGetRequestConfiguration{
			Header:          map[string]string{"Authorization": "Bearer token"},
			QueryParameters: &TableRequestBuilderGetQueryParameters{Limit: 10},
			Data:            map[string]interface{}{"key": "value"},
			ErrorMapping:    core.ErrorMapping{"4XX": "error"},
			response:        &TableCollectionResponse{},
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
		rC := &TableGetRequestConfiguration{
			Header:       map[string]string{"Authorization": "Bearer token"},
			Data:         map[string]interface{}{"key": "value"},
			ErrorMapping: core.ErrorMapping{"4XX": "error"},
			response:     &TableCollectionResponse{},
		}

		config := rC.toConfiguration()

		assert.Equal(t, (*TableRequestBuilderGetQueryParameters)(nil), config.QueryParameters)
	})
}
