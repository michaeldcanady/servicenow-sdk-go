package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)


// AppointmentResponse represents the appointment response.
type AppointmentResponse = newInternal.ServiceNowItemResponse[*AppointmentResultModel]

// CreateAppointmentResponseFromDiscriminatorValue is a factory for creating an AppointmentResponse.
func CreateAppointmentResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return newInternal.NewBaseServiceNowItemResponse[*AppointmentResultModel](CreateAppointmentResultFromDiscriminatorValue), nil
}

// AppointmentResult represents the result object in appointment response.
type AppointmentResult interface {
	serialization.Parsable
	kiotaStore.BackedModel

	GetData() (*string, error)
	setData(*string) error
	GetMessage() (*string, error)
	setMessage(*string) error
	GetReason() (*string, error)
	setReason(*string) error
	GetSuccess() (*bool, error)
	setSuccess(*bool) error
}

type AppointmentResultModel struct {
	newInternal.BaseModel
}

func NewAppointmentResult() *AppointmentResultModel {
	return &AppointmentResultModel{BaseModel: *newInternal.NewBaseModel()}
}

func CreateAppointmentResultFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewAppointmentResult(), nil
}

func (m *AppointmentResultModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(dataKey)(m.GetData),
		internalSerialization.SerializeStringFunc(messageKey)(m.GetMessage),
		internalSerialization.SerializeStringFunc(reasonKey)(m.GetReason),
		internalSerialization.SerializeBoolFunc(successKey)(m.GetSuccess),
	)
}

func (m *AppointmentResultModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		dataKey:    internalSerialization.DeserializeStringFunc()(m.setData),
		messageKey: internalSerialization.DeserializeStringFunc()(m.setMessage),
		reasonKey:  internalSerialization.DeserializeStringFunc()(m.setReason),
		successKey: internalSerialization.DeserializeBoolFunc()(m.setSuccess),
	}
}

func (m *AppointmentResultModel) GetData() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), dataKey) }
func (m *AppointmentResultModel) setData(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), dataKey, val) }
func (m *AppointmentResultModel) GetMessage() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), messageKey) }
func (m *AppointmentResultModel) setMessage(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), messageKey, val) }
func (m *AppointmentResultModel) GetReason() (*string, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), reasonKey) }
func (m *AppointmentResultModel) setReason(val *string) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), reasonKey, val) }
func (m *AppointmentResultModel) GetSuccess() (*bool, error) { return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *bool](m.GetBackingStore(), successKey) }
func (m *AppointmentResultModel) setSuccess(val *bool) error { return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), successKey, val) }
