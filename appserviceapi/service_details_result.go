package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceDetailsResult represents the result details of modifying service details.
type ServiceDetailsResult struct {
	core.BaseModel
}

// NewServiceDetailsResult creates a new instance of ServiceDetailsResult.
func NewServiceDetailsResult() *ServiceDetailsResult {
	return &ServiceDetailsResult{BaseModel: *core.NewBaseModel()}
}

// Serialize writes the objects properties to the current writer.
func (m *ServiceDetailsResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(statusKey, m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *ServiceDetailsResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		statusKey:  internalSerialization.DeserializeStringFunc(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc(m.setMessage),
	}
}

// GetStatus returns the status value.
func (m *ServiceDetailsResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsResult, *string](m, statusKey)
}

func (m *ServiceDetailsResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

// GetMessage returns the message value.
func (m *ServiceDetailsResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*ServiceDetailsResult, *string](m, messageKey)
}

func (m *ServiceDetailsResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

// CreateServiceDetailsResultFromDiscriminatorValue creates a new ServiceDetailsResult from a ParseNode.
func CreateServiceDetailsResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewServiceDetailsResult(), nil
}
