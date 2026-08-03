package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CreateServiceRequest represents the request body for creating an application service.
type CreateServiceRequest struct {
	core.BaseModel
}

// NewCreateServiceRequest creates a new instance of CreateServiceRequest.
func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *CreateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CreateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc(m.setComments),
	}
}

// GetName returns the name value.
func (m *CreateServiceRequest) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceRequest, *string](m, nameKey)
}

func (m *CreateServiceRequest) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetComments returns the comments value.
func (m *CreateServiceRequest) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceRequest, *string](m, commentsKey)
}

func (m *CreateServiceRequest) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

// CreateCreateServiceRequestFromDiscriminatorValue creates a new CreateServiceRequest from a ParseNode.
func CreateCreateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceRequest(), nil
}

// CreateServiceResult represents the result details of a created application service.
type CreateServiceResult struct {
	core.BaseModel
}

// NewCreateServiceResult creates a new instance of CreateServiceResult.
func NewCreateServiceResult() *CreateServiceResult {
	return &CreateServiceResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *CreateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIDKey, m.GetSysID),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CreateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIDKey:    internalSerialization.DeserializeStringFunc(m.setSysID),
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc(m.setComments),
	}
}

// GetSysID returns the sys id value.
func (m *CreateServiceResult) GetSysID() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, sysIDKey)
}

func (m *CreateServiceResult) setSysID(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIDKey, val)
}

// GetName returns the name value.
func (m *CreateServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, nameKey)
}

func (m *CreateServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

// GetComments returns the comments value.
func (m *CreateServiceResult) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, commentsKey)
}

func (m *CreateServiceResult) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

// CreateCreateServiceResultFromDiscriminatorValue creates a new CreateServiceResult from a ParseNode.
func CreateCreateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceResult(), nil
}
