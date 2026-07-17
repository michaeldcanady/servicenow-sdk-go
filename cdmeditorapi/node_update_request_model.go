package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// NodeUpdateRequest represents the body for updating a node.
type NodeUpdateRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetName() (*string, error)
	setName(*string) error
	GetValue() (*string, error)
	setValue(*string) error
}

type NodeUpdateRequestModel struct {
	core.BaseModel
}

func NewNodeUpdateRequest() *NodeUpdateRequestModel {
	return &NodeUpdateRequestModel{BaseModel: *core.NewBaseModel()}
}

func (m *NodeUpdateRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
	)
}

func (m *NodeUpdateRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:  internalSerialization.DeserializeStringFunc(m.setName),
		valueKey: internalSerialization.DeserializeStringFunc(m.setValue),
	}
}

func (m *NodeUpdateRequestModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeUpdateRequestModel, *string](m, nameKey)
}
func (m *NodeUpdateRequestModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *NodeUpdateRequestModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeUpdateRequestModel, *string](m, valueKey)
}
func (m *NodeUpdateRequestModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

func CreateNodeUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeUpdateRequest(), nil
}
