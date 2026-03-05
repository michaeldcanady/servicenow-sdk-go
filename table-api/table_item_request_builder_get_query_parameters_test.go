package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableItemRequestBuilderGetQueryParametersDefaults(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Default values",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := &TableItemRequestBuilderGetQueryParameters{}

			assert.Equal(t, DisplayValue(""), params.DisplayValue)
			assert.Equal(t, false, params.ExcludeReferenceLink)
			assert.Equal(t, 0, len(params.Fields))
			assert.Equal(t, false, params.QueryNoDomain)
			assert.Equal(t, View(""), params.View)
		})
	}
}
