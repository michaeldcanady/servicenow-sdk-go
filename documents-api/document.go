package documentsapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

const (
	sysIDKey = "sys_id"
	nameKey  = "name"
	typeKey  = "type"
)

// Document represents a ServiceNow document record.
type Document interface {
	GetSysID() (*string, error)
	SetSysID(val *string) error
	GetName() (*string, error)
	SetName(val *string) error
	GetType() (*string, error)
	SetType(val *string) error

	serialization.Parsable
	kiotaStore.BackedModel
}

// DocumentModel implementation of Document
type DocumentModel struct {
	newInternal.BaseModel
}

// NewDocument creates a new instance of DocumentModel
func NewDocument() *DocumentModel {
	return &DocumentModel{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

// CreateDocumentFromDiscriminatorValue creates a new instance of Document.
func CreateDocumentFromDiscriminatorValue(parseNode serialization.ParseNode) (serialization.Parsable, error) {
	return NewDocument(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *DocumentModel) Serialize(writer serialization.SerializationWriter) error {
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey)(m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey)(m.GetName),
		internalSerialization.SerializeStringFunc(typeKey)(m.GetType),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *DocumentModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey: internalSerialization.DeserializeStringFunc()(m.SetSysID),
		nameKey:  internalSerialization.DeserializeStringFunc()(m.SetName),
		typeKey:  internalSerialization.DeserializeStringFunc()(m.SetType),
	}
}

// GetSysID ...
func (m *DocumentModel) GetSysID() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), sysIDKey)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// SetSysID ...
func (m *DocumentModel) SetSysID(val *string) error {
	return m.GetBackingStore().Set(sysIDKey, val)
}

// GetName ...
func (m *DocumentModel) GetName() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), nameKey)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// SetName ...
func (m *DocumentModel) SetName(val *string) error {
	return m.GetBackingStore().Set(nameKey, val)
}

// GetType ...
func (m *DocumentModel) GetType() (*string, error) {
	val, err := store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), typeKey)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// SetType ...
func (m *DocumentModel) SetType(val *string) error {
	return m.GetBackingStore().Set(typeKey, val)
}
