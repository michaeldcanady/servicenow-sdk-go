package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegisterServiceRequest represents the request body for registering a CSDM service.
type RegisterServiceRequest struct {
	core.BaseModel
}

// NewRegisterServiceRequest creates a new instance of [RegisterServiceRequest].
func NewRegisterServiceRequest() *RegisterServiceRequest {
	return &RegisterServiceRequest{BaseModel: *core.NewBaseModel()}
}

// CreateRegisterServiceRequestFromDiscriminatorValue creates a new [RegisterServiceRequest] from a [serialization.ParseNode].
func CreateRegisterServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRegisterServiceRequest(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *RegisterServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*BasicDetails](basicDetailsKey, m.GetBasicDetails),
		internalSerialization.SerializeObjectValueFunc[*ServiceRelationship](relationshipsKey, m.GetRelationships),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *RegisterServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		basicDetailsKey:  internalSerialization.DeserializeObjectValueFunc[*BasicDetails](CreateBasicDetailsFromDiscriminatorValue, m.SetBasicDetails),
		relationshipsKey: internalSerialization.DeserializeObjectValueFunc[*ServiceRelationship](CreateServiceRelationshipFromDiscriminatorValue, m.SetRelationships),
	}
}

// GetBasicDetails returns the basic details value.
func (m *RegisterServiceRequest) GetBasicDetails() (*BasicDetails, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceRequest, *BasicDetails](m, basicDetailsKey)
}

// SetBasicDetails
func (m *RegisterServiceRequest) SetBasicDetails(val *BasicDetails) error {
	return store.DefaultBackedModelMutatorFunc(m, basicDetailsKey, val)
}

// GetRelationships
func (m *RegisterServiceRequest) GetRelationships() (*ServiceRelationship, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceRequest, *ServiceRelationship](m, relationshipsKey)
}

// SetRelationships
func (m *RegisterServiceRequest) SetRelationships(val *ServiceRelationship) error {
	return store.DefaultBackedModelMutatorFunc(m, relationshipsKey, val)
}
