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

// AppointmentResultModel represents the appointment result model.
type AppointmentResultModel struct {
	core.BaseModel
}

// NewAppointmentResult creates a new instance of AppointmentResultModel.
func NewAppointmentResult() *AppointmentResultModel {
	return &AppointmentResultModel{BaseModel: *core.NewBaseModel()}
}

// CreateAppointmentResultFromDiscriminatorValue creates a new AppointmentResult from a ParseNode.
func CreateAppointmentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppointmentResult(), nil
}

// Serialize writes the objects properties to the current writer.
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

// GetFieldDeserializers returns the deserialization information for this object.
func (m *AppointmentResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dataKey:    internalSerialization.DeserializeStringFunc(m.SetData),
		messageKey: internalSerialization.DeserializeStringFunc(m.SetMessage),
		reasonKey:  internalSerialization.DeserializeStringFunc(m.SetReason),
		successKey: internalSerialization.DeserializeBoolFunc(m.SetSuccess),
	}
}

// GetData returns the data value.
func (m *AppointmentResultModel) GetData() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *string](m, dataKey)
}

// SetData sets the data value.
func (m *AppointmentResultModel) SetData(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, dataKey, val)
}

// GetMessage returns the message value.
func (m *AppointmentResultModel) GetMessage() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *string](m, messageKey)
}

// SetMessage sets the message value.
func (m *AppointmentResultModel) SetMessage(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, messageKey, val)
}

// GetReason returns the reason value.
func (m *AppointmentResultModel) GetReason() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *string](m, reasonKey)
}

// SetReason sets the reason value.
func (m *AppointmentResultModel) SetReason(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, reasonKey, val)
}

// GetSuccess returns the success value.
func (m *AppointmentResultModel) GetSuccess() (*bool, error) {
	return store.DefaultBackedModelAccessorFunc[*AppointmentResultModel, *bool](m, successKey)
}

// SetSuccess sets the success value.
func (m *AppointmentResultModel) SetSuccess(val *bool) error {
	return store.DefaultBackedModelMutatorFunc(m, successKey, val)
}
