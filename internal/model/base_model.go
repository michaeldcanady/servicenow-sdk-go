package model

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
	"github.com/microsoft/kiota-abstractions-go/serialization"
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

func (bM *BaseModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}

func (bM *BaseModel) Serialize(_ serialization.SerializationWriter) error {
	return nil
}

// SetBackingStoreFactory sets the store.BackingStoreFactory for the model.
func (bM *BaseModel) SetBackingStoreFactory(factory store.BackingStoreFactory) error {
	if utils.IsNil(bM) {
		return nil
	}

	if utils.IsNil(factory) {
		return errors.New("factory is nil")
	}

	bM.backingStoreFactory = factory

	return nil
}

// GetBackingStore retrieves the backing store for the model.
func (bM *BaseModel) GetBackingStore() store.BackingStore {
	if utils.IsNil(bM) {
		return nil
	}

	if utils.IsNil(bM.backingStore) {
		bM.backingStore = bM.backingStoreFactory()
	}

	return bM.backingStore
}
