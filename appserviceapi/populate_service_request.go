package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// PopulateServiceRequest represents the request body for populating a CSDM service.
type PopulateServiceRequest struct {
	core.BaseModel
}

func NewPopulateServiceRequest() *PopulateServiceRequest {
	return &PopulateServiceRequest{BaseModel: *core.NewBaseModel()}
}

func (m *PopulateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*ServiceRelation](serviceRelationsKey)(m.GetServiceRelations),
	)
}

func (m *PopulateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		serviceRelationsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*ServiceRelation](CreateServiceRelationFromDiscriminatorValue)(m.setServiceRelations),
	}
}

func (m *PopulateServiceRequest) GetServiceRelations() ([]*ServiceRelation, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceRequest, []*ServiceRelation](m, serviceRelationsKey)
}

func (m *PopulateServiceRequest) setServiceRelations(val []*ServiceRelation) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceRelationsKey, val)
}

func CreatePopulateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceRequest(), nil
}

// PopulateServiceResult represents the result details of populating a service.
type PopulateServiceResult struct {
	core.BaseModel
}

func NewPopulateServiceResult() *PopulateServiceResult {
	return &PopulateServiceResult{BaseModel: *core.NewBaseModel()}
}

func (m *PopulateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *PopulateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *PopulateServiceResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceResult, *string](m, statusKey)
}

func (m *PopulateServiceResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

func (m *PopulateServiceResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceResult, *string](m, messageKey)
}

func (m *PopulateServiceResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

func CreatePopulateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceResult(), nil
}
