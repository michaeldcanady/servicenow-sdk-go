package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/mock"
	"github.com/yosida95/uritemplate/v3"
)

func toQueryMapFromStruct(source interface{}) (map[string]string, error) {
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

// ToQueryMap converts a struct to query parameter map
func ToQueryMap(source interface{}) (map[string]string, error) {
	var err error

	if IsNil(source) {
		return nil, ErrNilSource
	}

	queryParams := make(map[string]string)

	sourceType := reflect.TypeOf(source)

	if sourceType.Kind() != reflect.Map {
		source, err = toQueryMapFromStruct(source)
		if err != nil {
			return nil, err
		}
	}

	sourceValue := reflect.ValueOf(source)
	for _, key := range sourceValue.MapKeys() {
		strKey := fmt.Sprintf("%v", key.Interface())
		strValue := fmt.Sprintf("%v", sourceValue.MapIndex(key).Interface())
		queryParams[strKey] = strValue
	}
	return queryParams, nil
}

// normalizeVarNames normalizes variable names for URI template expansion.
func normalizeVarNames(varNames []string) map[string]string {
	normalizedNames := make(map[string]string)
	for _, varName := range varNames {
		normalizedNames[strings.ToLower(varName)] = varName
	}
	return normalizedNames
}

func addParametersWithOriginalNames(params map[string]string, normalizedNames map[string]string, values uritemplate.Values) uritemplate.Values {
	if values == nil {
		values = uritemplate.Values{}
	}

	for key, value := range params {
		values.Set(getKeyWithOriginalName(key, normalizedNames), uritemplate.String(value))
	}
	return values
}

func getKeyWithOriginalName(key string, normalizedNames map[string]string) string {
	originalName, exists := normalizedNames[key]
	if exists {
		return originalName
	}
	return key
}

func FromJSON[T any](response *http.Response, v *T) error {
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

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

// IsNil checks if a value is nil or a nil interface
func IsNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// ParseResponse[T] parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response, value *T) error {
	var err error

	if IsNil(response) {
		return ErrNilResponse
	}

	switch contentType := response.Header.Get(contentTypeHeader); contentType {
	case jsonContentType:
		err = FromJSON(response, &value)
	default:
		err = fmt.Errorf("unsupported content type: %s", contentType)
	}
	if err != nil {
		return err
	}

	(*value).ParseHeaders(response.Header)

	return nil
}

func ResetCalls(calls ...*mock.Call) {
	for _, call := range calls {
		if !IsNil(call) {
			call.Unset()
		}
	}
}
