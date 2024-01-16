package core

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/yosida95/uritemplate/v3"
)

// ToQueryMap converts a struct to query parameter map
func ToQueryMap(source interface{}) (map[string]string, error) {
	return internal.ToQueryMap(source)
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

// IsPointer checks if the provided type is a pointer
func IsPointer(value interface{}) bool {
	if isNil(value) {
		return false
	}

	valueKind := reflect.ValueOf(value).Kind()

	return valueKind == reflect.Ptr
}

// Deprecated: deprecated since v{version}. Removed from public API.
//
// FromJson[T] marshalls the provided response into v
func FromJson[T any](response *http.Response, v *T) error { //nolint:stylecheck

	err := internal.FromJSON[T](response, v)

	return err
}

// Deprecated: deprecated since v{version}. Removed from public API.
//
// ParseResponse[T] parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response, value *T) error {
	return internal.ParseResponse(response, value)
}

// Deprecated: deprecated in v1.4.0. Please use SendGet2.
func sendGet[T Response](requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping, value *T) error {
	requestInfo, err := requestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

// SendGet2 sends a GET request using the given request builder and configuration.
// It returns an error if the request or the response parsing fails.
// The request builder is used to create the request information, which includes
// the URL, the headers, and the query parameters.
// The configuration specifies how to handle errors and how to parse the response.
func SendGet2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToGetRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

// Deprecated: deprecated in v1.4.0. Use `SendPost2` instead.
func sendPost[T Response](requestBuilder *RequestBuilder, data interface{}, params interface{}, errorMapping ErrorMapping, value *T) error {
	requestInfo, err := requestBuilder.ToPostRequestInformation2(data, params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

// SendPost2 sends a POST request using the given request builder and configuration.
// It returns an error if the request or the response parsing fails.
// The request builder is used to create the request information, which includes
// the URL, the headers, the body, and the query parameters.
// The configuration specifies how to handle errors and how to parse the response.
func SendPost2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToPostRequestInformation3(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

// Deprecated: deprecated in v1.4.0. Use `sendDelete2` instead.
func sendDelete(requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping) error {
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

func sendDelete2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToDeleteRequestInformation2(config)
	if err != nil {
		return err
	}

	_, err = requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: deprecated in v1.4.0. Use `sendPut2` instead.
func sendPut[T Response](requestBuilder *RequestBuilder, data map[string]string, params interface{}, errorMapping ErrorMapping, value *T) error {
	requestInfo, err := requestBuilder.ToPutRequestInformation(data, params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

// SendDelete2 sends a DELETE request using the given request builder and configuration.
// It returns an error if the request or the response fails.
// The request builder is used to create the request information, which includes
// the URL, the headers, and the query parameters.
// The configuration specifies how to handle errors.
func sendPut2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToPutRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

// isNil checks if a value is nil or a nil interface
func isNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// convertToPage converts a response into a PageResult.
func convertToPage[T any](response CollectionResponse[T]) (PageResult[T], error) {
	var page PageResult[T]

	if isNil(response) {
		return page, ErrNilResponse
	}

	return response.ToPage(), nil
}
