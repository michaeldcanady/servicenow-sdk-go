package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// AppointmentResponse represents the appointment response.
type AppointmentResponse = core.ServiceNowItemResponse[*AppointmentResultModel]

// CreateAppointmentResponseFromDiscriminatorValue is a factory for creating an AppointmentResponse.
func CreateAppointmentResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*AppointmentResultModel](CreateAppointmentResultFromDiscriminatorValue), nil
}

// AppointmentResult represents the result object in appointment response.
type AppointmentResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetData() (*string, error)
	SetData(*string) error
	GetMessage() (*string, error)
	SetMessage(*string) error
	GetReason() (*string, error)
	SetReason(*string) error
	GetSuccess() (*bool, error)
	SetSuccess(*bool) error
}

type AppointmentResultModel struct {
	core.BaseModel
}

func NewAppointmentResult() *AppointmentResultModel {
	return &AppointmentResultModel{BaseModel: *core.NewBaseModel()}
}

func CreateAppointmentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppointmentResult(), nil
}

func (m *AppointmentResultModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(dataKey, m.GetData),
		internalSerialization.SerializeStringFunc(messageKey, m.GetMessage),
		internalSerialization.SerializeStringFunc(reasonKey, m.GetReason),
		internalSerialization.SerializeBoolFunc(successKey, m.GetSuccess),
	)
}

func (m *AppointmentResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dataKey:    internalSerialization.DeserializeStringFunc(m.SetData),
		messageKey: internalSerialization.DeserializeStringFunc(m.SetMessage),
		reasonKey:  internalSerialization.DeserializeStringFunc(m.SetReason),
		successKey: internalSerialization.DeserializeBoolFunc(m.SetSuccess),
	}
}

func (m *AppointmentResultModel) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *string](m, dataKey)
}
func (m *AppointmentResultModel) SetData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}
func (m *AppointmentResultModel) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *string](m, messageKey)
}
func (m *AppointmentResultModel) SetMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}
func (m *AppointmentResultModel) GetReason() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *string](m, reasonKey)
}
func (m *AppointmentResultModel) SetReason(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, reasonKey, val)
}
func (m *AppointmentResultModel) GetSuccess() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *bool](m, successKey)
}
func (m *AppointmentResultModel) SetSuccess(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, successKey, val)
}
