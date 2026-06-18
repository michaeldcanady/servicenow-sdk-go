package appointmentbookingapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// AppointmentBookingRequestBuilder builds and executes requests for the 'api/sn_apptmnt_booking/v1/appointment' path.
type AppointmentBookingRequestBuilder struct {
	core.RequestBuilder
}

// NewAppointmentBookingRequestBuilder creates a new instance of the AppointmentBookingRequestBuilder.
func NewAppointmentBookingRequestBuilder(requestAdapter abstractions.RequestAdapter, pathParameters map[string]string) *AppointmentBookingRequestBuilder {
	return &AppointmentBookingRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment", pathParameters),
	}
}

// Appointment returns a RequestBuilder for the '/appointment' path.
func (rB *AppointmentBookingRequestBuilder) Appointment() *AppointmentRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewAppointmentRequestBuilder(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Availability returns a RequestBuilder for the '/availability' path.
func (rB *AppointmentBookingRequestBuilder) Availability() *AvailabilityRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewAvailabilityRequestBuilder(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Calendar returns a RequestBuilder for the '/calendar' path.
func (rB *AppointmentBookingRequestBuilder) Calendar() *CalendarRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewCalendarRequestBuilder(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Configuration returns a RequestBuilder for the '/configuration' path.
func (rB *AppointmentBookingRequestBuilder) Configuration() *ConfigurationRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewConfigurationRequestBuilder(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ExecuteRuleConditions returns a RequestBuilder for the '/execute_rule_conditions' path.
func (rB *AppointmentBookingRequestBuilder) ExecuteRuleConditions() *ExecuteRuleConditionsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewExecuteRuleConditionsRequestBuilder(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// UserWindow returns a RequestBuilder for the '/userwindow' path.
func (rB *AppointmentBookingRequestBuilder) UserWindow() *UserWindowRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUserWindowRequestBuilder(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
