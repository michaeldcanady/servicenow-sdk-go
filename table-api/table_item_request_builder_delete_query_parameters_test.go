package tableapi

import (
	"errors"
	"testing"

	"github.com/hetiansu5/urlquery"
	"github.com/stretchr/testify/assert"
)

func toQueryMap(source interface{}) (map[string]string, error) {
	if source == nil {
		return nil, errors.New("source or request is nil")
	}

	queryBytes, err := urlquery.Marshal(source)
	if err != nil {
		return nil, err
	}

	var queryMap map[string]string
	err = urlquery.Unmarshal(queryBytes, &queryMap)
	if err != nil {
		return nil, err
	}

	return queryMap, nil
}

func TestTableItemRequestBuilderDeleteQueryParameters(t *testing.T) {

	params := &TableItemRequestBuilderDeleteQueryParameters{
		QueryNoDomain: true,
	}

	queryMap, err := toQueryMap(params)
	if err != nil {
		t.Error(err)
	}

	expectedValue := map[string]string{"sysparm_query_no_domain": "1"}

	assert.Equal(t, expectedValue, queryMap)
}
