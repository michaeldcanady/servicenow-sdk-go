package internal

import (
	"fmt"
	"reflect"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceNowResponseFromDiscriminatorValue[T] Creates a new ServiceNowCollectionResponse[T] or ServiceNowItemResponse[T] based on the kind.
func ServiceNowResponseFromDiscriminatorValue[T serialization.Parsable](factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		rawValue, err := parseNode.GetRawValue()
		if err != nil {
			return nil, err
		}

		val := reflect.ValueOf(rawValue)

		kind := val.Kind()
		if kind == reflect.Pointer {
			val = val.Elem()
		}
		kind = val.Kind()

		switch kind {
		case reflect.Slice:
			return NewBaseServiceNowCollectionResponse[T](factory), nil
		case reflect.Struct:
			return NewBaseServiceNowItemResponse[T](factory), nil
		}
		return nil, fmt.Errorf("unsupported type: %s", kind)
	}
}
