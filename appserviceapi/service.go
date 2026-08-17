package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// Service represents the result details retrieved from the find_service endpoint.
type Service struct {
	core.BaseModel
}

// NewService creates a new instance of Service.
func NewService() *Service {
	return &Service{BaseModel: *core.NewBaseModel()}
}

// CreateServiceFromDiscriminatorValue creates a new Service from a ParseNode.
func CreateServiceFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewService(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *Service) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*ServiceRelationship](relationshipsKey, m.GetRelationships),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(environmentKey, m.GetEnvironment),
		internalSerialization.SerializeStringFunc(versionKey, m.GetVersion),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *Service) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:         internalSerialization.DeserializeStringFunc(m.SetSysID),
		nameKey:          internalSerialization.DeserializeStringFunc(m.SetName),
		numberKey:        internalSerialization.DeserializeStringFunc(m.SetNumber),
		relationshipsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*ServiceRelationship](CreateServiceRelationshipFromDiscriminatorValue, m.SetRelationships),
		environmentKey:   internalSerialization.DeserializeStringFunc(m.SetEnvironment),
		versionKey:       internalSerialization.DeserializeStringFunc(m.SetVersion),
	}
}


// GetEnvironment returns the environment value.
func (m *Service) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Service, *string](m, environmentKey)
}

// SetEnvironment
func (m *Service) SetEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, environmentKey, val)
}

// GetName returns the name value.
func (m *Service) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Service, *string](m, nameKey)
}

// SetName
func (m *Service) SetName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetNumber returns the number value.
func (m *Service) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Service, *string](m, numberKey)
}

// SetNumber
func (m *Service) SetNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// GetRelationships
func (m *Service) GetRelationships() ([]*ServiceRelationship, error) {
	return store.DefaultBackedModelAccessorFunc[*Service, []*ServiceRelationship](m, relationshipsKey)
}

// SetRelationships
func (m *Service) SetRelationships(val []*ServiceRelationship) error {
	return store.DefaultBackedModelMutatorFunc(m, relationshipsKey, val)
}

// GetSysID returns the sys id value.
func (m *Service) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Service, *string](m, sysIDKey)
}

// SetSysID
func (m *Service) SetSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetVersion returns the version value.
func (m *Service) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*Service, *string](m, versionKey)
}

// SetVersion
func (m *Service) SetVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, versionKey, val)
}
