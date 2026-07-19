package appointmentbookingapi // nolint:dupl // shares field-count shape with UserTimeFormatModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarResponse implementation of CalendarResponse
type CalendarResponse struct {
	core.BaseModel
}

func NewCalendarResponse() *CalendarResponse {
	return &CalendarResponse{
		BaseModel: *core.NewBaseModel(),
	}
}

func CreateCalendarResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCalendarResponse(), nil
}

func (m *CalendarResponse) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(rangeEndKey, m.GetRangeEnd),
		internalSerialization.SerializeStringFunc(rangeStartKey, m.GetRangeStart),
	)
}

func (m *CalendarResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		rangeEndKey:   internalSerialization.DeserializeStringFunc(m.SetRangeEnd),
		rangeStartKey: internalSerialization.DeserializeStringFunc(m.SetRangeStart),
	}
}

func (m *CalendarResponse) GetRangeEnd() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CalendarResponse, *string](m, rangeEndKey)
}

func (m *CalendarResponse) SetRangeEnd(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, rangeEndKey, val)
}

func (m *CalendarResponse) GetRangeStart() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CalendarResponse, *string](m, rangeStartKey)
}

func (m *CalendarResponse) SetRangeStart(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, rangeStartKey, val)
}
