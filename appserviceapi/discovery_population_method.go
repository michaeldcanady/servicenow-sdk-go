package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

var _ DiscoveryPopulationMethod = (*DiscoveryPopulationMethodModel)(nil)

// DiscoveryPopulationMethod represents a discovery-based population method that uses entry points and attributes to discover service relationships.
type DiscoveryPopulationMethod interface {
	PopulationMethod
	GetEntryPointID() (*string, error)
	SetEntryPointID(*string) error
	GetAttributes() ([]DiscoveryPopulationMethodAttribute, error)
	SetAttributes([]DiscoveryPopulationMethodAttribute) error
	AddAttribute(DiscoveryPopulationMethodAttribute) error
}

// DiscoveryPopulationMethodModel is the backing-store-backed implementation of [DiscoveryPopulationMethod].
type DiscoveryPopulationMethodModel struct {
	PopulationMethod
}

// NewDiscoveryPopulationMethod creates a new instance of [DiscoveryPopulationMethodModel].
func NewDiscoveryPopulationMethod() *DiscoveryPopulationMethodModel {
	return &DiscoveryPopulationMethodModel{
		PopulationMethod: NewPopulationMethod(),
	}
}

// CreateDiscoveryPopulationMethodFromDiscriminatorValue creates a new [DiscoveryPopulationMethodModel] from a ParseNode.
func CreateDiscoveryPopulationMethodFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewDiscoveryPopulationMethod(), nil
}

// GetFieldDeserializers
func (d *DiscoveryPopulationMethodModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return extendmap(map[string]func(serialization.ParseNode) error{
		attributesKey:   internalSerialization.DeserializeCollectionOfObjectValuesFunc[DiscoveryPopulationMethodAttribute](CreateDiscoveryPopulationMethodAttributeModelFromDiscriminatorValue, d.SetAttributes),
		entryPointIDKey: internalSerialization.DeserializeStringFunc(d.SetEntryPointID),
	}, d.PopulationMethod.GetFieldDeserializers())
}

// Serialize implements
func (d *DiscoveryPopulationMethodModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(d) {
		return nil
	}

	if err := d.PopulationMethod.Serialize(writer); err != nil {
		return err
	}

	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[DiscoveryPopulationMethodAttribute](attributesKey, d.GetAttributes),
		internalSerialization.SerializeStringFunc(entryPointIDKey, d.GetEntryPointID),
	)
}

// AddAttribute
func (d *DiscoveryPopulationMethodModel) AddAttribute(val DiscoveryPopulationMethodAttribute) error {
	attrs, err := d.GetAttributes()
	if err != nil {
		return err
	}
	return d.SetAttributes(append(attrs, val))
}

// GetAttributes
func (d *DiscoveryPopulationMethodModel) GetAttributes() ([]DiscoveryPopulationMethodAttribute, error) {
	return store.DefaultBackedModelAccessorFunc[*DiscoveryPopulationMethodModel, []DiscoveryPopulationMethodAttribute](d, attributesKey)
}

// SetAttributes
func (d *DiscoveryPopulationMethodModel) SetAttributes(val []DiscoveryPopulationMethodAttribute) error {
	return store.DefaultBackedModelMutatorFunc(d, attributesKey, val)
}

// GetEntryPointID
func (d *DiscoveryPopulationMethodModel) GetEntryPointID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*DiscoveryPopulationMethodModel, *string](d, entryPointIDKey)
}

// SetEntryPointID
func (d *DiscoveryPopulationMethodModel) SetEntryPointID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(d, entryPointIDKey, val)
}
