package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableRequestBuilderPostQueryParametersDefaults(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Default values",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := &TableRequestBuilderPostQueryParameters{}

			assert.Equal(t, DisplayValue(""), params.DisplayValue)
			assert.Equal(t, false, params.ExcludeReferenceLink)
			assert.Equal(t, 0, len(params.Fields))
			assert.Equal(t, View(""), params.View)
		})
	}
}

func TestTableRequestBuilderPostQueryParamtersDefaults(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Default values",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := &TableRequestBuilderPostQueryParamters{}

			assert.Equal(t, DisplayValue(""), params.DisplayValue)
			assert.Equal(t, false, params.ExcludeReferenceLink)
			assert.Equal(t, 0, len(params.Fields))
			assert.Equal(t, View(""), params.View)
		})
	}
}
