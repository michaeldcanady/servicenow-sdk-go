package appointmentbookingapi // nolint:dupl // shares field-count shape with UserTimeFormatModel by coincidence, not copy-paste; distinct API concept, not worth sacrificing named accessors for

import (
	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalSerialization "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/serialization"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/store"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarItemResponse is the response envelope for a single CalendarResponse.
type CalendarItemResponse = core.ServiceNowItemResponse[*CalendarResponse]

// CreateAvailabilityResponseFromDiscriminatorValue is a factory for creating an AvailabilityResponse.
func CreateCalendarItemResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return core.NewBaseServiceNowItemResponse[*CalendarResponse](CreateCalendarResponseFromDiscriminatorValue), nil
}

// CalendarResponse implementation of CalendarResponse
type CalendarResponse struct {
	core.BaseModel
}

// NewCalendarResponse creates a new instance of CalendarResponse.
func NewCalendarResponse() *CalendarResponse {
	return &CalendarResponse{
		BaseModel: *core.NewBaseModel(),
	}
}

// CreateCalendarResponseFromDiscriminatorValue creates a new CalendarResponse from a ParseNode.
func CreateCalendarResponseFromDiscriminatorValue(_ serialization.ParseNode) (serialization.Parsable, error) {
	return NewCalendarResponse(), nil
}

// Serialize writes the objects properties to the current writer.
func (m *CalendarResponse) Serialize(writer serialization.SerializationWriter) error {
	if conversion.IsNil(m) {
		return nil
	}
	return internalSerialization.Serialize(writer,
		internalSerialization.SerializeStringFunc(rangeEndKey, m.GetRangeEnd),
		internalSerialization.SerializeStringFunc(rangeStartKey, m.GetRangeStart),
	)
}

// GetFieldDeserializers returns the deserialization information for this object.
func (m *CalendarResponse) GetFieldDeserializers() map[string]func(serialization.ParseNode) error {
	return map[string]func(serialization.ParseNode) error{
		rangeEndKey:   internalSerialization.DeserializeStringFunc(m.SetRangeEnd),
		rangeStartKey: internalSerialization.DeserializeStringFunc(m.SetRangeStart),
	}
}

// TODO: date
// GetRangeEnd returns the range end value.
func (m *CalendarResponse) GetRangeEnd() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CalendarResponse, *string](m, rangeEndKey)
}

// SetRangeEnd sets the range end value.
func (m *CalendarResponse) SetRangeEnd(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, rangeEndKey, val)
}

// TODO: date
// GetRangeStart returns the range start value.
func (m *CalendarResponse) GetRangeStart() (*string, error) {
	return store.DefaultBackedModelAccessorFunc[*CalendarResponse, *string](m, rangeStartKey)
}

// SetRangeStart sets the range start value.
func (m *CalendarResponse) SetRangeStart(val *string) error {
	return store.DefaultBackedModelMutatorFunc(m, rangeStartKey, val)
}
