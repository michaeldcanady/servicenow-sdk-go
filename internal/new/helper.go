package internal

import (
	"fmt"
	"reflect"
	"strconv"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// IsNil checks if a value is nil or a nil interface
func IsNil(a interface{}) bool {
	defer func() { _ = recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

// ThrowErrors
func ThrowErrors(typeName string, statusCode int64, contentType string, content []byte) error {
	var errorCtor serialization.ParsableFactory
	statusAsString := strconv.Itoa(int(statusCode))

	errorMappings, err := GetErrorRegistryInstance().Get(typeName)
	if err != nil {
		return err
	}

	if factory, ok := errorMappings[statusAsString]; ok {
		errorCtor = factory
	} else if factory, ok := errorMappings["XXX"]; ok {
		errorCtor = factory
	} else if factory, ok := errorMappings["4XX"]; statusCode >= 400 && statusCode < 500 && ok {
		errorCtor = factory
	} else if factory, ok := errorMappings["5XX"]; statusCode >= 500 && statusCode < 600 && ok {
		errorCtor = factory
	}

	if errorCtor == nil {
		return &abstractions.ApiError{
			Message: fmt.Sprintf("The server returned an unexpected status code and no error factory is registered for this code: %s", statusAsString),
		}
	}
	rootNode, err := serialization.DefaultParseNodeFactoryInstance.GetRootParseNode(contentType, content)
	if err != nil {
		return err
	}
	if rootNode == nil {
		return &abstractions.ApiError{
			Message: fmt.Sprintf("The server returned an unexpected status code with no response body: %s", statusAsString),
		}
	}

	errParsableValue, err := rootNode.GetObjectValue(errorCtor)
	if err != nil {
		return err
	}

	errValue, ok := errParsableValue.(error)
	if !ok {
		return fmt.Errorf("%T is not error", errParsableValue)
	}

	return errValue
}

// ToPointer
func ToPointer[T any](value T) *T {
	return &value
}
