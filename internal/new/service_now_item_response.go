package internal

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type ServiceNowItemResponse[T serialization.Parsable] interface {
	GetResult() (T, error)
	serialization.Parsable
	store.BackedModel
}
