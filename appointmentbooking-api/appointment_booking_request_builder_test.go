package appointmentbookingapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
)

func TestAppointmentBookingRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAppointmentBookingRequestBuilder(adapter, map[string]string{"baseurl": "https://example.com"})

	t.Run("Calendar Get", func(t *testing.T) {
		calendarBuilder := builder.Calendar()
		requestInfo, err := calendarBuilder.ToGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/calendar{?catalog_id,location,opened_for}", requestInfo.UrlTemplate)
	})

	t.Run("Configuration Get", func(t *testing.T) {
		configBuilder := builder.Configuration()
		requestInfo, err := configBuilder.ToGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/configuration{?catalog_id}", requestInfo.UrlTemplate)
	})
}

func TestAppointmentRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewAppointmentBookingRequestBuilder(adapter, map[string]string{"baseurl": "https://example.com"}).Appointment()

	body := NewAppointmentRequest()
	requestInfo, err := builder.ToPostRequestInformation(context.Background(), body, nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/appointment", requestInfo.UrlTemplate)
}
