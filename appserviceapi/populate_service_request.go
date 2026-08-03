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

// NewPopulateServiceRequest creates a new instance of PopulateServiceRequest.
func NewPopulateServiceRequest() *PopulateServiceRequest {
	return &PopulateServiceRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *PopulateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeCollectionOfObjectValuesFunc[*ServiceRelation](serviceRelationsKey, m.GetServiceRelations),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *PopulateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		serviceRelationsKey: internalSerialization.DeserializeCollectionOfObjectValuesFunc[*ServiceRelation](CreateServiceRelationFromDiscriminatorValue, m.setServiceRelations),
	}
}

// GetServiceRelations returns the service relations value.
func (m *PopulateServiceRequest) GetServiceRelations() ([]*ServiceRelation, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceRequest, []*ServiceRelation](m, serviceRelationsKey)
}

func (m *PopulateServiceRequest) setServiceRelations(val []*ServiceRelation) error {
	return store.DefaultBackedModelMutatorFunc(m, serviceRelationsKey, val)
}

// CreatePopulateServiceRequestFromDiscriminatorValue creates a new PopulateServiceRequest from a ParseNode.
func CreatePopulateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceRequest(), nil
}

// PopulateServiceResult represents the result details of populating a service.
type PopulateServiceResult struct {
	core.BaseModel
}

// NewPopulateServiceResult creates a new instance of PopulateServiceResult.
func NewPopulateServiceResult() *PopulateServiceResult {
	return &PopulateServiceResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *PopulateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *PopulateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc(m.setMessage),
	}
}

// GetStatus returns the status value.
func (m *PopulateServiceResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceResult, *string](m, statusKey)
}

func (m *PopulateServiceResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

// GetMessage returns the message value.
func (m *PopulateServiceResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*PopulateServiceResult, *string](m, messageKey)
}

func (m *PopulateServiceResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

// CreatePopulateServiceResultFromDiscriminatorValue creates a new PopulateServiceResult from a ParseNode.
func CreatePopulateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewPopulateServiceResult(), nil
}
