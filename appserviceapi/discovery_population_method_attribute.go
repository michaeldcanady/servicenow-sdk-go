package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

var _ DiscoveryPopulationMethodAttribute = (*DiscoveryPopulationMethodAttributeModel)(nil)

// DiscoveryPopulationMethodAttribute represents an attribute used by a discovery population method to filter or refine service discovery.
type DiscoveryPopulationMethodAttribute interface {
	core.Model
	GetName() (*string, error)
	SetName(*string) error
	GetValue() (*string, error)
	SetValue(*string) error
}

// DiscoveryPopulationMethodAttributeModel is the backing-store-backed implementation of [DiscoveryPopulationMethodAttribute].
type DiscoveryPopulationMethodAttributeModel struct {
	*core.BaseModel
}

// NewDiscoveryPopulationMethodAttributeModel creates a new instance of [DiscoveryPopulationMethodAttributeModel].
func NewDiscoveryPopulationMethodAttributeModel() *DiscoveryPopulationMethodAttributeModel {
	return &DiscoveryPopulationMethodAttributeModel{
		BaseModel: core.NewBaseModel(),
	}
}

// CreateDiscoveryPopulationMethodAttributeModelFromDiscriminatorValue creates a new [DiscoveryPopulationMethodAttributeModel] from a ParseNode.
func CreateDiscoveryPopulationMethodAttributeModelFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewDiscoveryPopulationMethodAttributeModel(), nil
}

// GetFieldDeserializers implements [DiscoveryPopulationMethodAttribute].
func (d *DiscoveryPopulationMethodAttributeModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:  internalSerialization.DeserializeStringFunc(d.SetName),
		valueKey: internalSerialization.DeserializeStringFunc(d.SetValue),
	}
}

// Serialize implements [DiscoveryPopulationMethodAttribute].
func (d *DiscoveryPopulationMethodAttributeModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(d) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, d.GetName),
		internalSerialization.SerializeStringFunc(valueKey, d.GetValue),
	)
}

// GetName
func (d *DiscoveryPopulationMethodAttributeModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*DiscoveryPopulationMethodAttributeModel, *string](d, nameKey)
}

// GetValue
func (d *DiscoveryPopulationMethodAttributeModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*DiscoveryPopulationMethodAttributeModel, *string](d, valueKey)
}

// SetName
func (d *DiscoveryPopulationMethodAttributeModel) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(d, nameKey, val)
}

// SetValue
func (d *DiscoveryPopulationMethodAttributeModel) SetValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(d, valueKey, val)
}
