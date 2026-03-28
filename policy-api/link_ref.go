package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	linkKey  = "link"
	valueKey = "value"
)

// LinkRef defines the interface for the LinkRef model.
type LinkRef interface {

	// GetLink gets the link property value.
	GetLink() (*string, error)
	// SetLink sets the link property value.
	SetLink(*string) error

	// GetValue gets the value property value.
	GetValue() (*string, error)
	// SetValue sets the value property value.
	SetValue(*string) error

	serialization.Parsable
	kiotaStore.BackedModel
}

// LinkRefModel is the concrete implementation of the LinkRef interface.
type LinkRefModel struct {
	newInternal.Model
}

// NewLinkRef creates a new instance of LinkRefModel with a backing store.
func NewLinkRef() *LinkRefModel {
	return &LinkRefModel{
		newInternal.NewBaseModel(),
	}
}

// CreateLinkRefFromDiscriminatorValue is a factory function for creating LinkRefModel instances during deserialization.
func CreateLinkRefFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewLinkRef(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *LinkRefModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}

	return internalSerialization.Serialize(writer,

		internalSerialization.SerializeStringFunc(linkKey)(m.GetLink),

		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *LinkRefModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{

		linkKey: internalSerialization.DeserializeStringFunc()(m.SetLink),

		valueKey: internalSerialization.DeserializeStringFunc()(m.SetValue),
	}
}

// GetLink returns the link property value.
func (m *LinkRefModel) GetLink() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, linkKey)
}

// SetLink sets the link property value.
func (m *LinkRefModel) SetLink(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, linkKey, val)
}

// GetValue returns the value property value.
func (m *LinkRefModel) GetValue() (*string, error) {
	if internal.IsNil(m) {
		return nil, nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](backingStore, valueKey)
}

// SetValue sets the value property value.
func (m *LinkRefModel) SetValue(val *string) error {
	if internal.IsNil(m) {
		return nil
	}

	backingStore := m.GetBackingStore()
	return store.DefaultBackedModelMutatorFunc(backingStore, valueKey, val)
}
