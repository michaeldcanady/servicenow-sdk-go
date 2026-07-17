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

func NewCreateServiceRequest() *CreateServiceRequest {
	return &CreateServiceRequest{BaseModel: *core.NewBaseModel()}
}

func (m *CreateServiceRequest) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
	)
}

func (m *CreateServiceRequest) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc(m.setComments),
	}
}

func (m *CreateServiceRequest) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceRequest, *string](m, nameKey)
}

func (m *CreateServiceRequest) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

func (m *CreateServiceRequest) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceRequest, *string](m, commentsKey)
}

func (m *CreateServiceRequest) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

func CreateCreateServiceRequestFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceRequest(), nil
}

// CreateServiceResult represents the result details of a created application service.
type CreateServiceResult struct {
	core.BaseModel
}

func NewCreateServiceResult() *CreateServiceResult {
	return &CreateServiceResult{BaseModel: *core.NewBaseModel()}
}

func (m *CreateServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey, m.GetSysId),
		internalSerialization.SerializeStringFunc(nameKey, m.GetName),
		internalSerialization.SerializeStringFunc(commentsKey, m.GetComments),
	)
}

func (m *CreateServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:    internalSerialization.DeserializeStringFunc(m.setSysId),
		nameKey:     internalSerialization.DeserializeStringFunc(m.setName),
		commentsKey: internalSerialization.DeserializeStringFunc(m.setComments),
	}
}

func (m *CreateServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, sysIdKey)
}

func (m *CreateServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}

func (m *CreateServiceResult) GetName() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, nameKey)
}

func (m *CreateServiceResult) setName(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, nameKey, val)
}

func (m *CreateServiceResult) GetComments() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CreateServiceResult, *string](m, commentsKey)
}

func (m *CreateServiceResult) setComments(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, commentsKey, val)
}

func CreateCreateServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCreateServiceResult(), nil
}
