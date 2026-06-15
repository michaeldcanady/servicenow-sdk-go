package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
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
	internal.BaseModel
}

func NewNodeCreateRequest() *NodeCreateRequestModel {
	return &NodeCreateRequestModel{BaseModel: *internal.NewBaseModel()}
}

func (m *NodeCreateRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(parentIdKey)(m.GetParentId),
		internalSerialization.SerializeStringFunc(cdmIdKey)(m.GetCdmId),
	)
}

func (m *NodeCreateRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc()(m.setName),
		typeKey:     internalSerialization.DeserializeStringFunc()(m.setType),
		parentIdKey: internalSerialization.DeserializeStringFunc()(m.setParentId),
		cdmIdKey:    internalSerialization.DeserializeStringFunc()(m.setCdmId),
	}
}

func (m *NodeCreateRequestModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *NodeCreateRequestModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *NodeCreateRequestModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}
func (m *NodeCreateRequestModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}
func (m *NodeCreateRequestModel) GetParentId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), parentIdKey)
}
func (m *NodeCreateRequestModel) setParentId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), parentIdKey, val)
}
func (m *NodeCreateRequestModel) GetCdmId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), cdmIdKey)
}
func (m *NodeCreateRequestModel) setCdmId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdmIdKey, val)
}

func CreateNodeCreateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeCreateRequest(), nil
}

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
	internal.BaseModel
}

func NewNodeUpdateRequest() *NodeUpdateRequestModel {
	return &NodeUpdateRequestModel{BaseModel: *internal.NewBaseModel()}
}

func (m *NodeUpdateRequestModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
	)
}

func (m *NodeUpdateRequestModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:  internalSerialization.DeserializeStringFunc()(m.setName),
		valueKey: internalSerialization.DeserializeStringFunc()(m.setValue),
	}
}

func (m *NodeUpdateRequestModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *NodeUpdateRequestModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *NodeUpdateRequestModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *NodeUpdateRequestModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}

func CreateNodeUpdateRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeUpdateRequest(), nil
}
