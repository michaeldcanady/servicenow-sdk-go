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

// NewRegisterServiceRequest creates a new instance of RegisterServiceRequest.
func NewRegisterServiceRequest() *RegisterServiceRequest {
	return &RegisterServiceRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *RegisterServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*BasicDetails](basicDetailsKey, m.GetBasicDetails),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *RegisterServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		basicDetailsKey: internalSerialization.DeserializeObjectValueFunc[*BasicDetails](CreateBasicDetailsFromDiscriminatorValue, m.setBasicDetails),
	}
}

// GetBasicDetails returns the basic details value.
func (m *RegisterServiceRequest) GetBasicDetails() (*BasicDetails, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceRequest, *BasicDetails](m, basicDetailsKey)
}

func (m *RegisterServiceRequest) setBasicDetails(val *BasicDetails) error {
	return store.DefaultBackedModelMutatorFunc(m, basicDetailsKey, val)
}

// CreateRegisterServiceRequestFromDiscriminatorValue creates a new RegisterServiceRequest from a ParseNode.
func CreateRegisterServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRegisterServiceRequest(), nil
}
