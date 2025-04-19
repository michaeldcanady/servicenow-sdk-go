package batchapi

import (
	"fmt"
	"strings"

	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// throwErrors returns error is provided req is an error
func throwErrors(req ServicedRequest, typeName string) error {
	code, err := req.GetStatusCode()
	if err != nil {
		return err
	}

	if code != nil && *code < 400 {
		return nil
	}

	body, err := req.GetErrorMessage()
	if err != nil {
		return err
	}

	headers, err := req.GetHeaders()
	if err != nil {
		return err
	}

	contentType := getHTTPHeader(headers, newInternal.HTTPHeaderContentType, "")

	return newInternal.ThrowErrors(typeName, *code, contentType, []byte(*body))
}

// serializeContent serializes the provided content using the provided ParsableFactory
func serializeContent[T serialization.Parsable](contentType string, content []byte, constructor serialization.ParsableFactory) (T, error) {
	var res T

	parseNodeFactory := serialization.DefaultParseNodeFactoryInstance
	parseNode, err := parseNodeFactory.GetRootParseNode(contentType, content)
	if err != nil {
		return res, err
	}

	result, err := parseNode.GetObjectValue(constructor)
	if err != nil {
		return res, err
	}

	typedResult, ok := result.(T)
	if !ok {
		return res, fmt.Errorf("result is not %T", res)
	}
	return typedResult, nil
}

// getHTTPHeader gets the requested header's value from a slice of RestRequestHeader
func getHTTPHeader(headers []RestRequestHeader, httpHeader newInternal.HTTPHeader, defaultValue string) string { // nolint: unparam // httpHeader currently on receives one value; however, can be used for any header in the future.
	for _, header := range headers {
		name, err := header.GetName()
		if err != nil {
			continue // Skip this header if there's an error
		}

		if strings.EqualFold(*name, httpHeader.String()) {
			value, err := header.GetValue()
			if err == nil {
				return *value // Return the value if no error
			}
		}
	}
	return defaultValue // Return the provided default value if no match is found
}
