package internal

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/store"
)

type Model interface {
	BackingStoreFactorySetter
	store.BackedModel
}

type BackingStoreFactorySetter interface {
	SetBackingStoreFactory(store.BackingStoreFactory) error
}

type BaseModel struct {
	// backingStoreFactory factory to create backingStore
	backingStoreFactory store.BackingStoreFactory
	// backingStore the store backing the model
	backingStore store.BackingStore
}

func NewBaseModel() *BaseModel {
	return &BaseModel{
		backingStoreFactory: store.NewInMemoryBackingStore,
		backingStore:        nil,
	}
}

func (bM *BaseModel) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	if IsNil(bM) {
		return nil
	}

	if IsNil(factory) {
		return errors.New("factory is nil")
	}

	bM.backingStoreFactory = factory

	return nil
}

// GetBackingStore retrieves the backing store for the model.
func (bM *BaseModel) GetBackingStore() store.BackingStore {
	if IsNil(bM) {
		return nil
	}

	if IsNil(bM.backingStore) {
		bM.backingStore = bM.backingStoreFactory()
	}

	return bM.backingStore
}
