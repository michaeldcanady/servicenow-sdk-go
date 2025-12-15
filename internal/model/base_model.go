package model

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/microsoft/kiota-abstractions-go/store"
)

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
	if internal.IsNil(bM) {
		return nil
	}

	if internal.IsNil(factory) {
		return errors.New("factory is nil")
	}

	bM.backingStoreFactory = factory

	// TODO: invalidate existing store
	// TODO: transfer data from one store to the new
	// TODO: replace existing store with new

	return nil
}

// GetBackingStore retrieves the backing store for the model.
func (bM *BaseModel) GetBackingStore() store.BackingStore {
	if internal.IsNil(bM) {
		return nil
	}

	if internal.IsNil(bM.backingStore) {
		bM.backingStore = bM.backingStoreFactory()
	}

	return bM.backingStore
}
