package core

import (
	"context"
	"encoding/json"
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
		return nil, ErrNilSource
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

func FromJson[T any](response *http.Response, v *T) error { //nolint:stylecheck
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

// ParseResponse[T] parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response, value *T) error {
	err := FromJson(response, &value)
	if err != nil {
		return err
	}

	(*value).ParseHeaders(response.Header)

	return nil
}

// Deprecated: deprecated in v1.4.0. Please use SendGet2.
func sendGet[T Response](requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping, value *T) error {
	requestInfo, err := requestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(context.Background(), requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

func SendGet2(ctx context.Context, requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToGetRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(ctx, requestInfo, config.ErrorMapping)
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

	response, err := requestBuilder.Client.Send(context.Background(), requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

func SendPost2(ctx context.Context, requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToPostRequestInformation3(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(ctx, requestInfo, config.ErrorMapping)
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

	_, err = requestBuilder.Client.Send(context.Background(), requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return nil
}

func sendDelete2(ctx context.Context, requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToDeleteRequestInformation2(config)
	if err != nil {
		return err
	}

	_, err = requestBuilder.Client.Send(ctx, requestInfo, config.ErrorMapping)
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

	response, err := requestBuilder.Client.Send(context.Background(), requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

func sendPut2(ctx context.Context, requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToPutRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(ctx, requestInfo, config.ErrorMapping)
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
