package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemRequestBuilderDeleteQueryParameters(t *testing.T) {
	tests := []struct {
		name     string
		params   *TableItemRequestBuilderDeleteQueryParameters
		expected map[string]string
	}{
		{
			name: "QueryNoDomain true",
			params: &TableItemRequestBuilderDeleteQueryParameters{
				QueryNoDomain: true,
			},
			expected: map[string]string{"sysparm_query_no_domain": "true"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			queryMap, err := core.ToQueryMap(test.params)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, queryMap)
		})
	}
}
