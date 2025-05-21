package internal

import "github.com/microsoft/kiota-abstractions-go/store"

// BackedModel Represents a model backed by a BackingStore
type BackedModel interface {
	// GetBackingStore returns the backing store, if store is nil it instantiates a new store.
	GetBackingStore() (store.BackingStore, error)
}
