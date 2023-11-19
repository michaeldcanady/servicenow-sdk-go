package tableapi

import (
	"reflect"
	"testing"
)

func TestTableItemRequestBuilderPutQueryParameters(t *testing.T) {
	t.Run("DefaultValues", func(t *testing.T) {
		params := &TableItemRequestBuilderPutQueryParameters{}

		// Validate default values
		if params.DisplayValue != "" {
			t.Fatalf("Expected empty DisplayValue, got: %v", params.DisplayValue)
		}

		if params.ExcludeReferenceLink != false {
			t.Fatalf("Expected ExcludeReferenceLink to be false, got: %v", params.ExcludeReferenceLink)
		}

		if len(params.Fields) != 0 {
			t.Fatalf("Expected empty Fields slice, got: %v", params.Fields)
		}

		if params.InputDisplayValue != false {
			t.Fatalf("Expected InputDisplayValue to be false, got: %v", params.InputDisplayValue)
		}

		if params.QueryNoDomain != false {
			t.Fatalf("Expected QueryNoDomain to be false, got: %v", params.QueryNoDomain)
		}

		if params.View != "" {
			t.Fatalf("Expected empty View, got: %v", params.View)
		}
	})

	t.Run("CustomValues", func(t *testing.T) {
		// Set custom values
		params := &TableItemRequestBuilderPutQueryParameters{
			DisplayValue:         "true",
			ExcludeReferenceLink: true,
			Fields:               []string{"field1", "field2"},
			InputDisplayValue:    true,
			QueryNoDomain:        true,
			View:                 "desktop",
		}

		// Validate custom values
		if params.DisplayValue != "true" {
			t.Fatalf("Expected DisplayValue to be 'true', got: %v", params.DisplayValue)
		}

		if params.ExcludeReferenceLink != true {
			t.Fatalf("Expected ExcludeReferenceLink to be true, got: %v", params.ExcludeReferenceLink)
		}

		expectedFields := []string{"field1", "field2"}
		if !reflect.DeepEqual(params.Fields, expectedFields) {
			t.Fatalf("Expected Fields %v, got: %v", expectedFields, params.Fields)
		}

		if params.InputDisplayValue != true {
			t.Fatalf("Expected InputDisplayValue to be true, got: %v", params.InputDisplayValue)
		}

		if params.QueryNoDomain != true {
			t.Fatalf("Expected QueryNoDomain to be true, got: %v", params.QueryNoDomain)
		}

		if params.View != "desktop" {
			t.Fatalf("Expected View to be 'desktop', got: %v", params.View)
		}
	})
}
