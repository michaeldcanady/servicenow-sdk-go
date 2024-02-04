package tableapi

import (
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestTableRequestBuilderGetQueryParameters(t *testing.T) {
	params := &TableRequestBuilderGetQueryParameters{
		Limit: 1,
	}

	queryMap, err := core.ToQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	expectedValue := map[string]string{"sysparm_limit": "1"}

	assert.Equal(t, expectedValue, queryMap)
}
