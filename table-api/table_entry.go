package tableapi

import (
	"errors"

	"github.com/hetiansu5/urlquery"
)

type TableEntry map[string]interface{}

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
