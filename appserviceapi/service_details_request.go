package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceDetailsRequest represents the request body for modifying basic details of a CSDM service.
type ServiceDetailsRequest struct {
	core.BaseModel
}

// NewServiceDetailsRequest creates a new instance of [ServiceDetailsRequest].
func NewServiceDetailsRequest() *ServiceDetailsRequest {
	return &ServiceDetailsRequest{BaseModel: *core.NewBaseModel()}
}

// CreateServiceDetailsRequestFromDiscriminatorValue creates a new [ServiceDetailsRequest] from a [serialization.ParseNode].
func CreateServiceDetailsRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceDetailsRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *ServiceDetailsRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*BasicDetails](basicDetailsKey, m.GetBasicDetails),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ServiceDetailsRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		basicDetailsKey: internalSerialization.DeserializeObjectValueFunc[*BasicDetails](CreateBasicDetailsFromDiscriminatorValue, m.SetBasicDetails),
	}
}

// GetBasicDetails returns the basic details value.
func (m *ServiceDetailsRequest) GetBasicDetails() (*BasicDetails, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsRequest, *BasicDetails](m, basicDetailsKey)
}

func (m *ServiceDetailsRequest) SetBasicDetails(val *BasicDetails) error {
	return store.DefaultBackedModelMutatorFunc(m, basicDetailsKey, val)
}

// GetRelationships
func (m *ServiceDetailsRequest) GetRelationships() (*ServiceRelationship, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsRequest, *ServiceRelationship](m, relationshipsKey)
}

// SetRelationships
func (m *ServiceDetailsRequest) SetRelationships(val *ServiceRelationship) error {
	return store.DefaultBackedModelMutatorFunc(m, relationshipsKey, val)
}
