package appointmentbookingapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
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
	core.BaseModel
}

func NewCalendarResponse() *CalendarResponseModel {
	return &CalendarResponseModel{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateCalendarResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCalendarResponse(), nil
}

func (m *CalendarResponseModel) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
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
	return store.DefaultBackedModelAccessorFunc[*CalendarResponseModel, *string](m, rangeEndKey)
}

func (m *CalendarResponseModel) setRangeEnd(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, rangeEndKey, val)
}

func (m *CalendarResponseModel) GetRangeStart() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CalendarResponseModel, *string](m, rangeStartKey)
}

func (m *CalendarResponseModel) setRangeStart(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, rangeStartKey, val)
}
