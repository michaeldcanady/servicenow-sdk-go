package appointmentbookingapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appointmentBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment"

// AppointmentBookingRequestBuilder builds and executes requests for the 'api/sn_apptmnt_booking/v1/appointment' path.
type AppointmentBookingRequestBuilder struct {
	core.RequestBuilder
}

// NewAppointmentBookingRequestBuilderInternal creates a new instance of the [AppointmentBookingRequestBuilder].
func NewAppointmentBookingRequestBuilderInternal(requestAdapter abstractions.RequestAdapter, pathParameters map[string]string) *AppointmentBookingRequestBuilder {
	return &AppointmentBookingRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentBookingURLTemplate, pathParameters),
	}
}

// NewAppointmentBookingRequestBuilder instantiates a new [AppointmentBookingRequestBuilder] with the provided base URL
// and request adapter.
func NewAppointmentBookingRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AppointmentBookingRequestBuilder {
	return NewAppointmentBookingRequestBuilderInternal(requestAdapter, map[string]string{internal.RawURLKey: rawURL})
}

// Appointment returns a [AppointmentRequestBuilder].
func (rB *AppointmentBookingRequestBuilder) Appointment() *AppointmentRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewAppointmentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Availability returns a [AvailabilityRequestBuilder].
func (rB *AppointmentBookingRequestBuilder) Availability() *AvailabilityRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewAvailabilityRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Calendar returns a [CalendarRequestBuilder].
func (rB *AppointmentBookingRequestBuilder) Calendar() *CalendarRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewCalendarRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Configuration returns a [ConfigurationRequestBuilder].
func (rB *AppointmentBookingRequestBuilder) Configuration() *ConfigurationRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewConfigurationRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ExecuteRuleConditions returns a [ExecuteRuleConditionsRequestBuilder].
func (rB *AppointmentBookingRequestBuilder) ExecuteRuleConditions() *ExecuteRuleConditionsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewExecuteRuleConditionsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
