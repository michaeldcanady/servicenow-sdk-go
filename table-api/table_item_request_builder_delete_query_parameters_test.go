package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/core"
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

	expectedValue := map[string]string{"sysparm_query_no_domain": "1"}

	assert.Equal(t, expectedValue, queryMap)
}
