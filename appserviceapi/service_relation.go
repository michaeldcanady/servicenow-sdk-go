package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceRelation represents a relationship between components inside Populate request.
type ServiceRelation struct {
	core.BaseModel
}

func NewServiceRelation() *ServiceRelation {
	return &ServiceRelation{BaseModel: *core.NewBaseModel()}
}

func (m *ServiceRelation) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(parentKey, m.GetParent),
		internalSerialization.SerializeStringFunc(childKey, m.GetChild),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
	)
}

func (m *ServiceRelation) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		parentKey: internalSerialization.DeserializeStringFunc(m.setParent),
		childKey:  internalSerialization.DeserializeStringFunc(m.setChild),
		typeKey:   internalSerialization.DeserializeStringFunc(m.setType),
	}
}

func (m *ServiceRelation) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelation, *string](m, parentKey)
}

func (m *ServiceRelation) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentKey, val)
}

func (m *ServiceRelation) GetChild() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelation, *string](m, childKey)
}

func (m *ServiceRelation) setChild(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, childKey, val)
}

func (m *ServiceRelation) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceRelation, *string](m, typeKey)
}

func (m *ServiceRelation) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

func CreateServiceRelationFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceRelation(), nil
}
