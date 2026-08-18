package caseapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

var _ ActivitiesResult = (*ActivitiesResultModel)(nil)

// ActivitiesResult represents a single activity entry from the CSM case activity stream.
type ActivitiesResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	// GetSysID returns the sys_id of the activity entry.
	GetSysID() (*string, error)
	// GetType returns the type of activity (e.g. work_notes, comments).
	GetType() (*string, error)
	// GetValue returns the text value of the activity entry.
	GetValue() (*string, error)
	// GetUser returns the login name of the user who created the entry.
	GetUser() (*string, error)
	// GetSysCreatedOn returns the creation timestamp of the activity entry.
	GetSysCreatedOn() (*string, error)
	// GetFieldName returns the field name associated with the activity entry.
	GetFieldName() (*string, error)
}

// ActivitiesResultModel is the implementation of [ActivitiesResult].
type ActivitiesResultModel struct {
	core.BaseModel
}

// NewActivitiesResult creates a new instance of [ActivitiesResultModel].
func NewActivitiesResult() *ActivitiesResultModel {
	return &ActivitiesResultModel{BaseModel: *core.NewBaseModel()}
}

// CreateActivitiesResultFromDiscriminatorValue creates a new instance of [ActivitiesResult].
func CreateActivitiesResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewActivitiesResult(), nil
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

// GetSysID returns the sys_id of the activity entry.
func (m *ActivitiesResultModel) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysIDKey)
}

func (m *ActivitiesResultModel) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetType returns the type of activity.
func (m *ActivitiesResultModel) GetType() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, typeKey)
}

func (m *ActivitiesResultModel) setType(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, typeKey, val)
}

// GetValue returns the text value of the activity entry.
func (m *ActivitiesResultModel) GetValue() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, valueKey)
}

func (m *ActivitiesResultModel) setValue(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, valueKey, val)
}

// GetUser returns the login name of the user who created the entry.
func (m *ActivitiesResultModel) GetUser() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, userKey)
}

func (m *ActivitiesResultModel) setUser(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, userKey, val)
}

// GetSysCreatedOn returns the creation timestamp of the activity entry.
func (m *ActivitiesResultModel) GetSysCreatedOn() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, sysCreatedOnKey)
}

func (m *ActivitiesResultModel) setSysCreatedOn(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysCreatedOnKey, val)
}

// GetFieldName returns the field name associated with the activity entry.
func (m *ActivitiesResultModel) GetFieldName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ActivitiesResultModel, *string](m, fieldNameKey)
}

func (m *ActivitiesResultModel) setFieldName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, fieldNameKey, val)
}
