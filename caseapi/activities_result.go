package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// ActivitiesResult represents case activities.
type ActivitiesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetSysId() (*string, error)
	setSysId(*string) error
	GetType() (*string, error)
	setType(*string) error
	GetValue() (*string, error)
	setValue(*string) error
	GetUser() (*string, error)
	setUser(*string) error
	GetSysCreatedOn() (*string, error)
	setSysCreatedOn(*string) error
	GetFieldName() (*string, error)
	setFieldName(*string) error
}

type ActivitiesResultModel struct {
	core.BaseModel
}

func NewActivitiesResult() *ActivitiesResultModel {
	return &ActivitiesResultModel{BaseModel: *core.NewBaseModel()}
}

func (m *ActivitiesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
		internalSerialization.SerializeStringFunc(valueKey)(m.GetValue),
		internalSerialization.SerializeStringFunc(userKey)(m.GetUser),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey)(m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(fieldNameKey)(m.GetFieldName),
	)
}

func (m *ActivitiesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:        internalSerialization.DeserializeStringFunc()(m.setSysId),
		typeKey:         internalSerialization.DeserializeStringFunc()(m.setType),
		valueKey:        internalSerialization.DeserializeStringFunc()(m.setValue),
		userKey:         internalSerialization.DeserializeStringFunc()(m.setUser),
		sysCreatedOnKey: internalSerialization.DeserializeStringFunc()(m.setSysCreatedOn),
		fieldNameKey:    internalSerialization.DeserializeStringFunc()(m.setFieldName),
	}
}

func (m *ActivitiesResultModel) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysIdKey)
}
func (m *ActivitiesResultModel) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}
func (m *ActivitiesResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, typeKey)
}
func (m *ActivitiesResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}
func (m *ActivitiesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, valueKey)
}
func (m *ActivitiesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}
func (m *ActivitiesResultModel) GetUser() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, userKey)
}
func (m *ActivitiesResultModel) setUser(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userKey, val)
}
func (m *ActivitiesResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysCreatedOnKey)
}
func (m *ActivitiesResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}
func (m *ActivitiesResultModel) GetFieldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, fieldNameKey)
}
func (m *ActivitiesResultModel) setFieldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, fieldNameKey, val)
}

func CreateActivitiesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResult(), nil
}
