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

func NewServiceDetailsRequest() *ServiceDetailsRequest {
	return &ServiceDetailsRequest{BaseModel: *core.NewBaseModel()}
}

func (m *ServiceDetailsRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeObjectValueFunc[*BasicDetails](basicDetailsKey, m.GetBasicDetails),
	)
}

func (m *ServiceDetailsRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		basicDetailsKey: internalSerialization.DeserializeObjectValueFunc[*BasicDetails](CreateBasicDetailsFromDiscriminatorValue, m.setBasicDetails),
	}
}

func (m *ServiceDetailsRequest) GetBasicDetails() (*BasicDetails, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsRequest, *BasicDetails](m, basicDetailsKey)
}

func (m *ServiceDetailsRequest) setBasicDetails(val *BasicDetails) error {
	return store.DefaultBackedModelMutatorFunc(m, basicDetailsKey, val)
}

func CreateServiceDetailsRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceDetailsRequest(), nil
}
