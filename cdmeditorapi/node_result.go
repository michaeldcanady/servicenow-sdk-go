package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// NodeResult represents a node in the configuration tree.
type NodeResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetSysId() (*string, error)
	setSysId(*string) error
	GetName() (*string, error)
	setName(*string) error
	GetType() (*string, error)
	setType(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	GetParent() (*string, error)
	setParent(*string) error
	GetCdmId() (*string, error)
	setCdmId(*string) error
}

type NodeResultModel struct {
	core.BaseModel
}

func NewNodeResult() *NodeResultModel {
	return &NodeResultModel{BaseModel: *core.NewBaseModel()}
}

func (m *NodeResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
		internalSerialization.SerializeStringFunc(parentKey)(m.GetParent),
		internalSerialization.SerializeStringFunc(cdmIdKey)(m.GetCdmId),
	)
}

func (m *NodeResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:  internalSerialization.DeserializeStringFunc()(m.setSysId),
		nameKey:   internalSerialization.DeserializeStringFunc()(m.setName),
		typeKey:   internalSerialization.DeserializeStringFunc()(m.setType),
		valueKey:  internalSerialization.DeserializeStringFunc()(m.setValue),
		parentKey: internalSerialization.DeserializeStringFunc()(m.setParent),
		cdmIdKey:  internalSerialization.DeserializeStringFunc()(m.setCdmId),
	}
}

func (m *NodeResultModel) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, sysIdKey)
}
func (m *NodeResultModel) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *NodeResultModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, nameKey)
}
func (m *NodeResultModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}
func (m *NodeResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, typeKey)
}
func (m *NodeResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
func (m *NodeResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, valueKey)
}
func (m *NodeResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}
func (m *NodeResultModel) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, parentKey)
}
func (m *NodeResultModel) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, parentKey, val)
}
func (m *NodeResultModel) GetCdmId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*NodeResultModel, *string](m, cdmIdKey)
}
func (m *NodeResultModel) setCdmId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, cdmIdKey, val)
}

func CreateNodeResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeResult(), nil
}
