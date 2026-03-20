package policyapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
	model.Model
}

// Ref represents a reference with a link and a value.
type Ref struct {
	model.Model
}

// NewRef creates a new instance of Ref.
func NewRef() *Ref {
	return &Ref{
		Model: model.NewBaseModel(),
	}
}

// CreateRefFromDiscriminatorValue creates a new Ref from a ParseNode.
func CreateRefFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRef(), nil
}

// GetFieldDeserializers returns the deserialization information for this object.
func (r *Ref) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		refLinkKey:  internalSerialization.DeserializeStringFunc()(r.SetLink),
		refValueKey: internalSerialization.DeserializeStringFunc()(r.SetValue),
	}
}

// Serialize writes the objects properties to the current writer.
func (r *Ref) Serialize(writer serialization.SerializationWriter) error {
	if utils.IsNil(r) {
		return nil
	}

	return kiota.Serialize(writer,
		kiota.SerializeStringFunc(refLinkKey)(r.GetLink),
		kiota.SerializeStringFunc(refValueKey)(r.GetValue),
	)
}

// GetLink returns the link.
func (r *Ref) GetLink() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), refLinkKey)
}

// SetLink sets the link.
func (r *Ref) SetLink(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(r.GetBackingStore(), refLinkKey, val)
}

// GetValue returns the value.
func (r *Ref) GetValue() (*string, error) {
	return kiota.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](r.GetBackingStore(), refValueKey)
}

// SetValue sets the value.
func (r *Ref) SetValue(val *string) error {
	return kiota.DefaultBackedModelMutatorFunc(r.GetBackingStore(), refValueKey, val)
}
