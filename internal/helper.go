package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

func FromJSON[T any](response *http.Response, value *T) error {
	if response == nil {
		return ErrNilResponse
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close() //nolint:errcheck

	if len(body) == 0 {
		return ErrNilResponseBody
	}

	if err := json.Unmarshal(body, value); err != nil {
		return err
	}

	return nil
}

// ParseResponse[T] parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response, value *T) error {
	err := FromJSON[T](response, value)
	if err != nil {
		return err
	}

	(*value).ParseHeaders(response.Header)

	return nil
}

// ToQueryMap converts a struct to query parameter map
func ToQueryMap(source interface{}) (map[string]string, error) {
	if source == nil {
		return nil, ErrNilSource
	}

	queryValues, err := query.Values(source)
	if err != nil {
		return nil, err
	}

	queryParams := map[string]string{}

	for key, values := range queryValues {
		queryParams[key] = strings.Join(values, ",")
	}

	return queryParams, nil
}
