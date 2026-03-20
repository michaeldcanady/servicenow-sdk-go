package model

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceNowCollectionResponseFromDiscriminatorValue[T] Creates a new ServiceNowCollectionResponse[T] based on the kind.
func ServiceNowCollectionResponseFromDiscriminatorValue[T serialization.Parsable](factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewBaseServiceNowCollectionResponse[T](factory), nil
	}
}

// ServiceNowItemResponseFromDiscriminatorValue[T] Creates a new ServiceNowItemResponse[T] based on the kind.
func ServiceNowItemResponseFromDiscriminatorValue[T serialization.Parsable](factory serialization.ParsableFactory) serialization.ParsableFactory {
	return func(parseNode serialization.ParseNode) (serialization.Parsable, error) {
		return NewBaseServiceNowItemResponse[T](factory), nil
	}
}
