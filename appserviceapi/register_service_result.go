package appserviceapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegisterServiceResult represents the result details of a registered CSDM service.
type RegisterServiceResult struct {
	core.BaseModel
}

func NewRegisterServiceResult() *RegisterServiceResult {
	return &RegisterServiceResult{BaseModel: *core.NewBaseModel()}
}

func (m *RegisterServiceResult) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(sysIdKey)(m.GetSysId),
		internalSerialization.SerializeStringFunc(numberKey)(m.GetNumber),
		internalSerialization.SerializeStringFunc(statusKey)(m.GetStatus),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
	)
}

func (m *RegisterServiceResult) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		sysIdKey:   internalSerialization.DeserializeStringFunc()(m.setSysId),
		numberKey:  internalSerialization.DeserializeStringFunc()(m.setNumber),
		statusKey:  internalSerialization.DeserializeStringFunc()(m.setStatus),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
	}
}

func (m *RegisterServiceResult) GetSysId() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceResult, *string](m, sysIdKey)
}

func (m *RegisterServiceResult) setSysId(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, sysIdKey, val)
}

func (m *RegisterServiceResult) GetNumber() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceResult, *string](m, numberKey)
}

func (m *RegisterServiceResult) setNumber(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, numberKey, val)
}

func (m *RegisterServiceResult) GetStatus() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceResult, *string](m, statusKey)
}

func (m *RegisterServiceResult) setStatus(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, statusKey, val)
}

func (m *RegisterServiceResult) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*RegisterServiceResult, *string](m, messageKey)
}

func (m *RegisterServiceResult) setMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

func CreateRegisterServiceResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewRegisterServiceResult(), nil
}
