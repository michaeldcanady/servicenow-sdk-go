package core

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/yosida95/uritemplate/v3"
)

// normalizeVarNames normalizes variable names for URI template expansion.
func normalizeVarNames(varNames []string) map[string]string {
	normalizedNames := make(map[string]string)
	for _, varName := range varNames {
		normalizedNames[strings.ToLower(varName)] = varName
	}
	return normalizedNames
}

func getOriginalParameterName(key string, normalizedNames map[string]string) string {
	lowercaseKey := strings.ToLower(key)
	if paramName, ok := normalizedNames[lowercaseKey]; ok {
		return paramName
	}
	return key
}

// addParameterWithOriginalName adds the URI template parameter to the template using the right casing, because of Go conventions,
// casing might have changed for the generated property.
func addParameterWithOriginalName(key string, value string, normalizedNames map[string]string, values uritemplate.Values) {
	paramName := getOriginalParameterName(key, normalizedNames)
	values.Set(paramName, uritemplate.String(value))
}

func addParametersWithOrignialNames(params map[string]string, normalizedNames map[string]string) uritemplate.Values {
	var values uritemplate.Values

	for key, value := range params {
		addParameterWithOriginalName(key, value, normalizedNames, values)
	}

	return values
}

func IsPointer(value interface{}) bool {

	if value == nil {
		return false
	}

	valueKind := reflect.ValueOf(value).Kind()

	return valueKind == reflect.Ptr
}

func FromJson[T interface{}](response *http.Response) (*T, error) {

	if !IsPointer(response) {
		return nil, errors.New("response is nil or not a pointer")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var value T
	if err := json.Unmarshal(body, &value); err != nil {
		return nil, err
	}

	return &value, nil
}
