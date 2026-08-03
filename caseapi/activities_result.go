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

	GetSysID() (*string, error)
	setSysID(*string) error
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

// ActivitiesResultModel implementation of ActivitiesResult.
type ActivitiesResultModel struct {
	core.BaseModel
}

// NewActivitiesResult creates a new instance of ActivitiesResultModel.
func NewActivitiesResult() *ActivitiesResultModel {
	return &ActivitiesResultModel{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *ActivitiesResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(typeKey, m.GetType),
		internalSerialization.SerializeStringFunc(valueKey, m.GetValue),
		internalSerialization.SerializeStringFunc(userKey, m.GetUser),
		internalSerialization.SerializeStringFunc(sysCreatedOnKey, m.GetSysCreatedOn),
		internalSerialization.SerializeStringFunc(fieldNameKey, m.GetFieldName),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ActivitiesResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:        internalSerialization.DeserializeStringFunc(m.setSysID),
		typeKey:         internalSerialization.DeserializeStringFunc(m.setType),
		valueKey:        internalSerialization.DeserializeStringFunc(m.setValue),
		userKey:         internalSerialization.DeserializeStringFunc(m.setUser),
		sysCreatedOnKey: internalSerialization.DeserializeStringFunc(m.setSysCreatedOn),
		fieldNameKey:    internalSerialization.DeserializeStringFunc(m.setFieldName),
	}
}

// GetSysID ...
func (m *ActivitiesResultModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysIDKey)
}
func (m *ActivitiesResultModel) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetType ...
func (m *ActivitiesResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, typeKey)
}
func (m *ActivitiesResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetValue ...
func (m *ActivitiesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, valueKey)
}
func (m *ActivitiesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// GetUser ...
func (m *ActivitiesResultModel) GetUser() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, userKey)
}
func (m *ActivitiesResultModel) setUser(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userKey, val)
}

// GetSysCreatedOn ...
func (m *ActivitiesResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysCreatedOnKey)
}
func (m *ActivitiesResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// GetFieldName ...
func (m *ActivitiesResultModel) GetFieldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, fieldNameKey)
}
func (m *ActivitiesResultModel) setFieldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, fieldNameKey, val)
}

// CreateActivitiesResultFromDiscriminatorValue creates a new instance of ActivitiesResult.
func CreateActivitiesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResult(), nil
}
