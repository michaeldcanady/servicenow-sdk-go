package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	kiotaStore "github.com/microsoft/kiota-abstractions-go/store"
)

// CalendarResponse represents the calendar response.
type CalendarResponse interface {
	serialization.Parsable
	kiotaStore.BackedModel
	GetRangeEnd() (*string, error)
	setRangeEnd(*string) error
	GetRangeStart() (*string, error)
	setRangeStart(*string) error
}

// CalendarResponseModel implementation of CalendarResponse
type CalendarResponseModel struct {
	newInternal.BaseModel
}

func NewCalendarResponse() *CalendarResponseModel {
	return &CalendarResponseModel{
		BaseModel: *newInternal.NewBaseModel(),
	}
}

func (m *CalendarResponseModel) Serialize(writer serialization.SerializationWriter) error {
	if internal.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(rangeEndKey)(m.GetRangeEnd),
		internalSerialization.SerializeStringFunc(rangeStartKey)(m.GetRangeStart),
	)
}

func (m *CalendarResponseModel) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		rangeEndKey:   internalSerialization.DeserializeStringFunc()(m.setRangeEnd),
		rangeStartKey: internalSerialization.DeserializeStringFunc()(m.setRangeStart),
	}
}

func (m *CalendarResponseModel) GetRangeEnd() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), rangeEndKey)
}

func (m *CalendarResponseModel) setRangeEnd(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), rangeEndKey, val)
}

func (m *CalendarResponseModel) GetRangeStart() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[kiotaStore.BackingStore, *string](m.GetBackingStore(), rangeStartKey)
}

func (m *CalendarResponseModel) setRangeStart(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m.GetBackingStore(), rangeStartKey, val)
}

func CreateCalendarResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCalendarResponse(), nil
}
