package cdmapplicationsapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedLibraryComponentApplication represents an application associated with shared libraries.
type SharedLibraryComponentApplication struct {
	core.BaseModel
}

// NewSharedLibraryComponentApplication instantiates a new SharedLibraryComponentApplication.
func NewSharedLibraryComponentApplication() *SharedLibraryComponentApplication {
	return &SharedLibraryComponentApplication{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the object's properties to the given writer.
func (m *SharedLibraryComponentApplication) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(versionKey, m.GetVersion),
		internalSerialization.SerializeStringFunc(descriptionKey, m.GetDescription),
		internalSerialization.SerializeStringFunc(appNameKey, m.GetAppName),
	)
}

// GetFieldDeserializers returns the deserializers for this object's fields.
func (m *SharedLibraryComponentApplication) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:       internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:        internalSerialization.DeserializeStringFunc(m.setName),
		versionKey:     internalSerialization.DeserializeStringFunc(m.setVersion),
		descriptionKey: internalSerialization.DeserializeStringFunc(m.setDescription),
		appNameKey:     internalSerialization.DeserializeStringFunc(m.setAppName),
	}
}

// GetSysID returns the sys id.
func (m *SharedLibraryComponentApplication) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedLibraryComponentApplication, *string](m, sysIDKey)
}
func (m *SharedLibraryComponentApplication) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName returns the name.
func (m *SharedLibraryComponentApplication) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedLibraryComponentApplication, *string](m, nameKey)
}
func (m *SharedLibraryComponentApplication) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetVersion returns the version.
func (m *SharedLibraryComponentApplication) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedLibraryComponentApplication, *string](m, versionKey)
}
func (m *SharedLibraryComponentApplication) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, versionKey, val)
}

// GetDescription returns the description.
func (m *SharedLibraryComponentApplication) GetDescription() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedLibraryComponentApplication, *string](m, descriptionKey)
}
func (m *SharedLibraryComponentApplication) setDescription(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, descriptionKey, val)
}

// GetAppName returns the app name.
func (m *SharedLibraryComponentApplication) GetAppName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*SharedLibraryComponentApplication, *string](m, appNameKey)
}
func (m *SharedLibraryComponentApplication) setAppName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, appNameKey, val)
}

// CreateSharedLibraryComponentApplicationFromDiscriminatorValue creates a new SharedLibraryComponentApplication from a ParseNode.
func CreateSharedLibraryComponentApplicationFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewSharedLibraryComponentApplication(), nil
}
