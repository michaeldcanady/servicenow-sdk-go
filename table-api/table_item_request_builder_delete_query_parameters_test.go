package tableapi

import (
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableItemRequestBuilderDeleteQueryParameters(t *testing.T) {
	params := &TableItemRequestBuilderDeleteQueryParameters{
		QueryNoDomain: true,
	}

	queryMap, err := core.ToQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	expectedValue := map[string]string{"sysparm_query_no_domain": "true"}

	assert.Equal(t, expectedValue, queryMap)
}
