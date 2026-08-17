package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// FindServiceResult represents the result details retrieved from the find_service endpoint.
type FindServiceResult struct {
	core.BaseModel
}

// NewFindServiceResult creates a new instance of FindServiceResult.
func NewFindServiceResult() *FindServiceResult {
	return &FindServiceResult{BaseModel: *core.NewBaseModel()}
}

// CreateFindServiceResultFromDiscriminatorValue creates a new FindServiceResult from a ParseNode.
func CreateFindServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewFindServiceResult(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *FindServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*Service](servicesKey, m.GetServices),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *FindServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		servicesKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*Service](CreateServiceFromDiscriminatorValue, m.SetServices),
	}
}

// GetServices returns the services value.
func (m *FindServiceResult) GetServices() ([]*Service, error) {
	return store.DefaultBackedModelAccessorFunc[*FindServiceResult, []*Service](m, servicesKey)
}

// SetServices
func (m *FindServiceResult) SetServices(val []*Service) error {
	return store.DefaultBackedModelMutatorFunc(m, servicesKey, val)
}
