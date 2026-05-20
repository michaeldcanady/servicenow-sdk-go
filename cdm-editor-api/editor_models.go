package cdmeditorapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
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
	newInternal.BaseModel
}

func NewNodeResult() *NodeResultModel {
	return &NodeResultModel{BaseModel: *newInternal.NewBaseModel()}
}

func (m *NodeResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
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
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIdKey)
}
func (m *NodeResultModel) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), sysIdKey, val)
}
func (m *NodeResultModel) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
}
func (m *NodeResultModel) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), nameKey, val)
}
func (m *NodeResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
}
func (m *NodeResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), typeKey, val)
}
func (m *NodeResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), valueKey)
}
func (m *NodeResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), valueKey, val)
}
func (m *NodeResultModel) GetParent() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), parentKey)
}
func (m *NodeResultModel) setParent(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), parentKey, val)
}
func (m *NodeResultModel) GetCdmId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), cdmIdKey)
}
func (m *NodeResultModel) setCdmId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), cdmIdKey, val)
}

func CreateNodeResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewNodeResult(), nil
}

// ValidationResult represents validation status.
type ValidationResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetStatus() (*string, error)
	setStatus(*string) error
	GetErrors() (any, error)
	setErrors(any) error
	GetWarnings() (any, error)
	setWarnings(any) error
}

type ValidationResultModel struct {
	newInternal.BaseModel
}

func NewValidationResult() *ValidationResultModel {
	return &ValidationResultModel{BaseModel: *newInternal.NewBaseModel()}
}

func (m *ValidationResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeAnyFunc(errorsKey)(m.GetErrors),
		internalSerialization.SerializeAnyFunc(warningsKey)(m.GetWarnings),
	)
}

func (m *ValidationResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:   internalSerialization.DeserializeStringFunc()(m.setStatus),
		errorsKey:   internalSerialization.DeserializeAnyFunc()(m.setErrors),
		warningsKey: internalSerialization.DeserializeAnyFunc()(m.setWarnings),
	}
}

func (m *ValidationResultModel) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), statusKey)
}
func (m *ValidationResultModel) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), statusKey, val)
}
func (m *ValidationResultModel) GetErrors() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), errorsKey)
}
func (m *ValidationResultModel) setErrors(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), errorsKey, val)
}
func (m *ValidationResultModel) GetWarnings() (any, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, any](m.GetBackingStore(), warningsKey)
}
func (m *ValidationResultModel) setWarnings(val any) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), warningsKey, val)
}

func CreateValidationResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewValidationResult(), nil
}

// MessageResult represents a simple string result.
type MessageResult struct {
	newInternal.BaseModel
	Message *string
}

func NewMessageResult(message *string) *MessageResult {
	return &MessageResult{
		BaseModel: *newInternal.NewBaseModel(),
		Message:   message,
	}
}

func (m *MessageResult) Serialize(writer serialization.SerializationWriter) error { return nil }
func (m *MessageResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return nil
}

func CreateMessageResultFromDiscriminatorValue(node serialization.ParseNode) (serialization.Parsable, error) {
	val, err := node.GetStringValue()
	if err != nil {
		return nil, err
	}
	return NewMessageResult(val), nil
}
