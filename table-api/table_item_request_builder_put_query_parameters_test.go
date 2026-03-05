package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableItemRequestBuilderPutQueryParameters(t *testing.T) {
	tests := []struct {
		name     string
		params   *TableItemRequestBuilderPutQueryParameters
		expected *TableItemRequestBuilderPutQueryParameters
	}{
		{
			name:   "DefaultValues",
			params: &TableItemRequestBuilderPutQueryParameters{},
			expected: &TableItemRequestBuilderPutQueryParameters{
				DisplayValue:         "",
				ExcludeReferenceLink: false,
				Fields:               nil,
				InputDisplayValue:    false,
				QueryNoDomain:        false,
				View:                 "",
			},
		},
		{
			name: "CustomValues",
			params: &TableItemRequestBuilderPutQueryParameters{
				DisplayValue:         "true",
				ExcludeReferenceLink: true,
				Fields:               []string{"field1", "field2"},
				InputDisplayValue:    true,
				QueryNoDomain:        true,
				View:                 "desktop",
			},
			expected: &TableItemRequestBuilderPutQueryParameters{
				DisplayValue:         "true",
				ExcludeReferenceLink: true,
				Fields:               []string{"field1", "field2"},
				InputDisplayValue:    true,
				QueryNoDomain:        true,
				View:                 "desktop",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected.DisplayValue, test.params.DisplayValue)
			assert.Equal(t, test.expected.ExcludeReferenceLink, test.params.ExcludeReferenceLink)
			assert.Equal(t, test.expected.Fields, test.params.Fields)
			assert.Equal(t, test.expected.InputDisplayValue, test.params.InputDisplayValue)
			assert.Equal(t, test.expected.QueryNoDomain, test.params.QueryNoDomain)
			assert.Equal(t, test.expected.View, test.params.View)
		})
	}
}
