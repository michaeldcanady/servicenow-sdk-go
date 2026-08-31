// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAppointmentBookingRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewAppointmentBookingRequestBuilderInternal(adapter, map[string]string{"baseurl": "https://example.com"})

	t.Run("Calendar Get", func(t *testing.T) {
		calendarBuilder := builder.Calendar()
		requestInfo, err := calendarBuilder.ToGetRequestInformation(context.Background(), nil)
		require.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/calendar{?catalog_id,location,opened_for}", requestInfo.UrlTemplate)
	})

	t.Run("Configuration Get", func(t *testing.T) {
		configBuilder := builder.Configuration()
		requestInfo, err := configBuilder.ToGetRequestInformation(context.Background(), nil)
		require.NoError(t, err)
		assert.NotNil(t, requestInfo)
		assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/configuration{?catalog_id}", requestInfo.UrlTemplate)
	})

	t.Run("Availability", func(t *testing.T) {
		assert.NotNil(t, builder.Availability())
	})

	t.Run("ExecuteRuleConditions", func(t *testing.T) {
		assert.NotNil(t, builder.ExecuteRuleConditions())
	})

	t.Run("UserWindow", func(t *testing.T) {
		assert.NotNil(t, builder.UserWindow())
	})
}

func TestAppointmentRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewAppointmentRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*AppointmentResultModel](CreateAppointmentResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewAppointmentRequest(), nil)

	require.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestAvailabilityRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewAvailabilityRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*AvailabilityResultModel](CreateAvailabilityResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewAvailabilityRequest(), nil)

	require.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestExecuteRuleConditionsRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewExecuteRuleConditionsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*ExecuteRuleConditionsResult](CreateExecuteRuleConditionsResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewExecuteRuleConditionsRequest(), nil)

	require.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestUserWindowRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewUserWindowRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := core.NewBaseServiceNowItemResponse[*UserWindowResult](CreateUserWindowResultFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewUserWindowRequest(), nil)

	require.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestUserWindowRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewAppointmentBookingRequestBuilderInternal(adapter, map[string]string{"baseurl": "https://example.com"}).UserWindow()

	body := NewUserWindowRequest()
	requestInfo, err := builder.ToPostRequestInformation(context.Background(), body, nil)
	require.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/userwindow", requestInfo.UrlTemplate)
}

func TestAppointmentRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewAppointmentBookingRequestBuilderInternal(adapter, map[string]string{"baseurl": "https://example.com"}).Appointment()

	body := NewAppointmentRequest()
	requestInfo, err := builder.ToPostRequestInformation(context.Background(), body, nil)
	require.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/appointment", requestInfo.UrlTemplate)
}
