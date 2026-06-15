package internal

import (
	"reflect"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

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
	if headers := config.Headers; !conversion.IsNil(headers) {
		request.Headers.AddAll(headers)
	}
	if options := config.Options; !conversion.IsNil(options) {
		request.AddRequestOptions(options)
	}
	if queryParams := config.QueryParameters; !conversion.IsNil(queryParams) {
		request.AddQueryParameters(queryParams)
	}
}
