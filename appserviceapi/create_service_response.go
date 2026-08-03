package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateServiceResponse represents the response containing the created application service details.
type CreateServiceResponse interface {
	core.ServiceNowItemResponse[*CreateServiceResult]
}

// CreateCreateServiceResponseFromDiscriminatorValue creates a new CreateServiceResponse from a ParseNode.
func CreateCreateServiceResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*CreateServiceResult](CreateCreateServiceResultFromDiscriminatorValue), nil
}

// FindServiceResult represents the result details retrieved from the find_service endpoint.
type FindServiceResult struct {
	core.BaseModel
}

// NewFindServiceResult creates a new instance of FindServiceResult.
func NewFindServiceResult() *FindServiceResult {
	return &FindServiceResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *FindServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(numberKey, m.GetNumber),
		internalSerialization.SerializeStringFunc(environmentKey, m.GetEnvironment),
		internalSerialization.SerializeStringFunc(versionKey, m.GetVersion),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *FindServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:       internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:        internalSerialization.DeserializeStringFunc(m.setName),
		numberKey:      internalSerialization.DeserializeStringFunc(m.setNumber),
		environmentKey: internalSerialization.DeserializeStringFunc(m.setEnvironment),
		versionKey:     internalSerialization.DeserializeStringFunc(m.setVersion),
	}
}

// GetSysID returns the sys id value.
func (m *FindServiceResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, sysIDKey)
}

func (m *FindServiceResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName returns the name value.
func (m *FindServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, nameKey)
}

func (m *FindServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetNumber returns the number value.
func (m *FindServiceResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, numberKey)
}

func (m *FindServiceResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

// GetEnvironment returns the environment value.
func (m *FindServiceResult) GetEnvironment() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, environmentKey)
}

func (m *FindServiceResult) setEnvironment(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, environmentKey, val)
}

// GetVersion returns the version value.
func (m *FindServiceResult) GetVersion() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, *string](m, versionKey)
}

func (m *FindServiceResult) setVersion(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, versionKey, val)
}

// CreateFindServiceResultFromDiscriminatorValue creates a new FindServiceResult from a ParseNode.
func CreateFindServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFindServiceResult(), nil
}
