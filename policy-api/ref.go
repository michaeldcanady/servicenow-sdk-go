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
	refLinkKey  = "link"
	refValueKey = "value"
)

// Refable interface for Ref model
type Refable interface {
	GetLink() (*string, error)
	SetLink(*string) error
	GetValue() (*string, error)
	SetValue(*string) error
	serialization.Parsable
	newInternal.Model
}

// Ref represents a reference with a link and a value.
type Ref struct {
	newInternal.Model
}

// NewRef creates a new instance of Ref.
func NewRef() *Ref {
	return &Ref{
		Model: newInternal.NewBaseModel(),
	}
}

// CreateRefFromDiscriminatorValue creates a new Ref from a ParseNode.
func CreateRefFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRef(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (r *Ref) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		refLinkKey:  internalSerialization.DeserializeStringFunc(r.SetLink),
		refValueKey: internalSerialization.DeserializeStringFunc(r.SetValue),
	}
}

// Serialize writes the objects properties to the current writer.
func (r *Ref) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(r) {
		return nil
	}

	{
		val, err := r.GetLink()
		if err != nil {
			return err
		}
		if val != nil {
			err = writer.WriteStringValue(refLinkKey, val)
			if err != nil {
				return err
			}
		}
	}
	{
		val, err := r.GetValue()
		if err != nil {
			return err
		}
		if val != nil {
			err = writer.WriteStringValue(refValueKey, val)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetLink returns the link.
func (r *Ref) GetLink() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), refLinkKey)
}

// SetLink sets the link.
func (r *Ref) SetLink(val *string) error {
	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), refLinkKey, val)
}

// GetValue returns the value.
func (r *Ref) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), refValueKey)
}

// SetValue sets the value.
func (r *Ref) SetValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(r.GetBackingStore(), refValueKey, val)
}
