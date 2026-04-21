package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// ResolvedInputs defines the interface for the ResolvedInputs model.
type ResolvedInputs interface {
	serialization.Parsable
	kiotaStore.BackedModel
}

// ResolvedInputsModel is the concrete implementation of the ResolvedInputs interface.
type ResolvedInputsModel struct {
	newInternal.Model
}

// NewResolvedInputs creates a new instance of ResolvedInputsModel with a backing store.
func NewResolvedInputs() *ResolvedInputsModel {
	return &ResolvedInputsModel{
		newInternal.NewBaseModel(),
	}
}

// CreateResolvedInputsFromDiscriminatorValue is a factory function for creating ResolvedInputsModel instances during deserialization.
func CreateResolvedInputsFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewResolvedInputs(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ResolvedInputsModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ResolvedInputsModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{}
}
