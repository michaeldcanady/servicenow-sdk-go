package tableapi

import (
	"errors"
	"fmt"

	"github.com/hetiansu5/urlquery"
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

func convertType[T any](val interface{}) (T, error) {
	v, ok := val.(T)
	if !ok {
		return v, fmt.Errorf("value (%v) cannot be converted to %T", val, v)
	}
	return v, nil
}
