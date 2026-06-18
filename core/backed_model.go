package core

import (
	"github.com/microsoft/kiota-abstractions-go/store"
)

// BackedModel Represents a model backed by a BackingStore
type BackedModel interface {
	store.BackedModel
}
