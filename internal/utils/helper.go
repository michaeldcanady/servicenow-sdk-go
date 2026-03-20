package utils

import (
	"reflect"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// IsNil checks if a value is nil or a nil interface.
func IsNil(a interface{}) bool {
	if a == nil {
		return true
	}
	v := reflect.ValueOf(a)
	k := v.Kind()
	if k == reflect.Chan || k == reflect.Func || k == reflect.Map || k == reflect.Pointer || k == reflect.UnsafePointer || k == reflect.Interface || k == reflect.Slice {
		return v.IsNil()
	}
	return false
}

// ToPointer Converts provided value to pointer.
func ToPointer[T any](value T) *T {
	return &value
}

// IsPointer
func IsPointer(value any) bool {
	return reflect.ValueOf(value).Kind() == reflect.Pointer
}

func ConfigureRequestInformation[T any](request *KiotaRequestInformation, config *abstractions.RequestConfiguration[T]) {
	if request == nil {
		return
	}
	if config == nil {
		return
	}
	if headers := config.Headers; !internal.IsNil(headers) {
		request.Headers.AddAll(headers)
	}
	if options := config.Options; !internal.IsNil(options) {
		request.AddRequestOptions(options)
	}
	if queryParams := config.QueryParameters; !internal.IsNil(queryParams) {
		request.AddQueryParameters(queryParams)
	}
}

// ModelSetter represents a function that sets a value.
type ModelSetter[T any] func(val T) error

// Mutator represents a function that transforms a value.
type Mutator[T, S any] func(input T) (S, error)
