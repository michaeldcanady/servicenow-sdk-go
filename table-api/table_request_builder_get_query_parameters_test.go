package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableRequestBuilderGetQueryParameters(t *testing.T) {
	tests := []struct {
		name     string
		params   *TableRequestBuilderGetQueryParameters
		expected map[string]string
	}{
		{
			name: "Limit 1",
			params: &TableRequestBuilderGetQueryParameters{
				Limit: 1,
			},
			expected: map[string]string{"sysparm_limit": "1"},
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
