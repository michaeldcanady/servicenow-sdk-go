package internal

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	"github.com/microsoft/kiota-abstractions-go/store"
)

// BackingStoreFactorySetter represents a struct with a settable backing store factory.
type BackingStoreFactorySetter interface {
	SetBackingStoreFactory(store.BackingStoreFactory) error
}

type BaseModel struct {
	// backingStoreFactory factory to create backingStore
	backingStoreFactory store.BackingStoreFactory
	// backingStore the store backing the model
	backingStore store.BackingStore
}

// NewBaseModel instantiates a new BaseModel.
func NewBaseModel() *BaseModel {
	return &BaseModel{
		backingStoreFactory: store.NewInMemoryBackingStore,
		backingStore:        nil,
	}
}

// SetBackingStoreFactory sets the store.BackingStoreFactory for the model.
func (bM *BaseModel) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	if conversion.IsNil(bM) {
		return nil
	}

	if conversion.IsNil(factory) {
		return errors.New("factory is nil")
	}

	bM.backingStoreFactory = factory

	return nil
}

// GetBackingStore retrieves the backing store for the model.
func (bM *BaseModel) GetBackingStore() store.BackingStore {
	if conversion.IsNil(bM) {
		return nil
	}

	if conversion.IsNil(bM.backingStore) {
		bM.backingStore = bM.backingStoreFactory()
	}

	return bM.backingStore
}
