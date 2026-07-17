package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// NodeCreateRequest represents the body for creating a node.
type NodeCreateRequest interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetName() (*string, error)
	setName(*string) error
	GetType() (*string, error)
	setType(*string) error
	GetParentId() (*string, error)
	setParentId(*string) error
	GetCdmId() (*string, error)
	setCdmId(*string) error
}

type NodeCreateRequestModel struct {
	core.BaseModel
}

func NewNodeCreateRequest() *NodeCreateRequestModel {
	return &NodeCreateRequestModel{BaseModel: *core.NewBaseModel()}
}

func (m *NodeCreateRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(parentIdKey, m.GetParentId),
		internalSerialization.SerializeStringFunc(cdmIdKey, m.GetCdmId),
	)
}

func (m *NodeCreateRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		typeKey:     internalSerialization.DeserializeStringFunc(m.setType),
		parentIdKey: internalSerialization.DeserializeStringFunc(m.setParentId),
		cdmIdKey:    internalSerialization.DeserializeStringFunc(m.setCdmId),
	}
}

func (m *NodeCreateRequestModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, nameKey)
}
func (m *NodeCreateRequestModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *NodeCreateRequestModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, typeKey)
}
func (m *NodeCreateRequestModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
func (m *NodeCreateRequestModel) GetParentId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, parentIdKey)
}
func (m *NodeCreateRequestModel) setParentId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentIdKey, val)
}
func (m *NodeCreateRequestModel) GetCdmId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeCreateRequestModel, *string](m, cdmIdKey)
}
func (m *NodeCreateRequestModel) setCdmId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmIdKey, val)
}

func CreateNodeCreateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeCreateRequest(), nil
}
