package cmdbinstanceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	sysIDKey     = "sys_id"
	nameKey      = "name"
	classNameKey = "className"
)

// CmdbInstance represents a ServiceNow CMDB record.
type CmdbInstance interface {
	GetSysID() (*string, error)
	SetSysID(val *string) error
	GetName() (*string, error)
	SetName(val *string) error
	GetClassName() (*string, error)
	SetClassName(val *string) error

	serialization.Parsable
	kiotaStore.BackedModel
}

// CmdbInstanceModel implementation of CmdbInstance
type CmdbInstanceModel struct {
	core.BaseModel
}

// NewCmdbInstance creates a new instance of CmdbInstanceModel
func NewCmdbInstance() *CmdbInstanceModel {
	return &CmdbInstanceModel{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateCmdbInstanceFromDiscriminatorValue creates a new instance of CmdbInstance.
func CreateCmdbInstanceFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewCmdbInstance(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *CmdbInstanceModel) Serialize(writer serialization.SerializationWriter) error {
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey)(m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(classNameKey)(m.GetClassName),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CmdbInstanceModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:     internalSerialization.DeserializeStringFunc()(m.SetSysID),
		nameKey:      internalSerialization.DeserializeStringFunc()(m.SetName),
		classNameKey: internalSerialization.DeserializeStringFunc()(m.SetClassName),
	}
}

// GetSysID ...
func (m *CmdbInstanceModel) GetSysID() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIDKey)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// SetSysID ...
func (m *CmdbInstanceModel) SetSysID(val *string) error {
	return m.GetBackingStore().Set(sysIDKey, val)
}

// GetName ...
func (m *CmdbInstanceModel) GetName() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// SetName ...
func (m *CmdbInstanceModel) SetName(val *string) error {
	return m.GetBackingStore().Set(nameKey, val)
}

// GetClassName ...
func (m *CmdbInstanceModel) GetClassName() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), classNameKey)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// SetClassName ...
func (m *CmdbInstanceModel) SetClassName(val *string) error {
	return m.GetBackingStore().Set(classNameKey, val)
}
