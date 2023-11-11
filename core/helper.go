package core

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/hetiansu5/urlquery"
	"github.com/yosida95/uritemplate/v3"
)

// ToQueryMap converts a struct to query parameter map
func ToQueryMap(source interface{}) (map[string]string, error) {
	if source == nil {
		return nil, errors.New("source is nil")
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

func IsPointer(value interface{}) bool {

	if value == nil {
		return false
	}

	valueKind := reflect.ValueOf(value).Kind()

	return valueKind == reflect.Ptr
}

func FromJson[T any](response *http.Response) (*T, error) {

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

type Response interface {
	ParseHeaders(headers http.Header)
}

func ParseHeaders(response Response, headers http.Header) {
	response.ParseHeaders(headers)
}

// ParseResponse parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response) (*T, error) {

	responseObject, err := FromJson[T](response)
	if err != nil {
		return nil, err
	}

	ParseHeaders(any(*responseObject).(Response), response.Header)

	return responseObject, nil
}

func SendGet[T Response](requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping) (*T, error) {
	requestInfo, err := requestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return nil, err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	return ParseResponse[T](response)
}

func SendPost[T Response](requestBuilder *RequestBuilder, data map[string]string, params interface{}, errorMapping ErrorMapping) (*T, error) {
	requestInfo, err := requestBuilder.ToPostRequestInformation(data, params)
	if err != nil {
		return nil, err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	return ParseResponse[T](response)
}

func SendDelete(requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping) error {
	requestInfo, err := requestBuilder.ToDeleteRequestInformation(params)
	if err != nil {
		return err
	}

	_, err = requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return nil
}

func SendPut[T Response](requestBuilder *RequestBuilder, data map[string]string, params interface{}, errorMapping ErrorMapping) (*T, error) {
	requestInfo, err := requestBuilder.ToPutRequestInformation(data, params)
	if err != nil {
		return nil, err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	return ParseResponse[T](response)
}
