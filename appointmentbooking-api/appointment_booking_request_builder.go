package appointmentbookingapi

import (
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// AppointmentBookingRequestBuilder builds and executes requests for the 'api/sn_apptmnt_booking/v1/appointment' path.
type AppointmentBookingRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewAppointmentBookingRequestBuilder creates a new instance of the AppointmentBookingRequestBuilder.
func NewAppointmentBookingRequestBuilder(requestAdapter abstractions.RequestAdapter, pathParameters map[string]string) *AppointmentBookingRequestBuilder {
	return &AppointmentBookingRequestBuilder{
		RequestBuilder: newInternal.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment", pathParameters),
	}
}

// Appointment returns a RequestBuilder for the '/appointment' path.
func (rB *AppointmentBookingRequestBuilder) Appointment() *AppointmentRequestBuilder {
	return NewAppointmentRequestBuilder(rB.GetPathParameters(), rB.GetRequestAdapter())
}

// Availability returns a RequestBuilder for the '/availability' path.
func (rB *AppointmentBookingRequestBuilder) Availability() *AvailabilityRequestBuilder {
	return NewAvailabilityRequestBuilder(rB.GetPathParameters(), rB.GetRequestAdapter())
}

// Calendar returns a RequestBuilder for the '/calendar' path.
func (rB *AppointmentBookingRequestBuilder) Calendar() *CalendarRequestBuilder {
	return NewCalendarRequestBuilder(rB.GetPathParameters(), rB.GetRequestAdapter())
}

// Configuration returns a RequestBuilder for the '/configuration' path.
func (rB *AppointmentBookingRequestBuilder) Configuration() *ConfigurationRequestBuilder {
	return NewConfigurationRequestBuilder(rB.GetPathParameters(), rB.GetRequestAdapter())
}

// ExecuteRuleConditions returns a RequestBuilder for the '/execute_rule_conditions' path.
func (rB *AppointmentBookingRequestBuilder) ExecuteRuleConditions() *ExecuteRuleConditionsRequestBuilder {
	return NewExecuteRuleConditionsRequestBuilder(rB.GetPathParameters(), rB.GetRequestAdapter())
}

// UserWindow returns a RequestBuilder for the '/userwindow' path.
func (rB *AppointmentBookingRequestBuilder) UserWindow() *UserWindowRequestBuilder {
	return NewUserWindowRequestBuilder(rB.GetPathParameters(), rB.GetRequestAdapter())
}
