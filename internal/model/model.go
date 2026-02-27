package model

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoft/kiota-abstractions-go/store"
)

type Model interface {
	serialization.Parsable
	store.BackedModel
}
