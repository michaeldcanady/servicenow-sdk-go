package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceDetailsResult represents the result details of modifying service details.
type ServiceDetailsResult struct {
	core.BaseModel
}

func NewServiceDetailsResult() *ServiceDetailsResult {
	return &ServiceDetailsResult{BaseModel: *core.NewBaseModel()}
}

func (m *ServiceDetailsResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *ServiceDetailsResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *ServiceDetailsResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsResult, *string](m, statusKey)
}

func (m *ServiceDetailsResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

func (m *ServiceDetailsResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsResult, *string](m, messageKey)
}

func (m *ServiceDetailsResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

func CreateServiceDetailsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceDetailsResult(), nil
}
