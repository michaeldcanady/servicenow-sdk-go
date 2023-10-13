package tableapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableRequestBuilderGetQueryParameters(t *testing.T) {

	params := &TableRequestBuilderGetQueryParameters{
		Limit: 1,
	}

	queryMap, err := toQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	expectedValue := map[string]string{"sysparm_limit": "1"}

	assert.Equal(t, expectedValue, queryMap)
}
