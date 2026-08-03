package appointmentbookingapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Nil-builder guard coverage for the request builders in this package that
// did not otherwise have a dedicated test exercising the guard.

func TestAppointmentRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *AppointmentRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestAppointmentRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *AppointmentRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestAvailabilityRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *AvailabilityRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestAvailabilityRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *AvailabilityRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestCalendarRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *CalendarRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestCalendarRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *CalendarRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestConfigurationRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *ConfigurationRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestConfigurationRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *ConfigurationRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestExecuteRuleConditionsRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *ExecuteRuleConditionsRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestExecuteRuleConditionsRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *ExecuteRuleConditionsRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestUserWindowRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *UserWindowRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestUserWindowRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *UserWindowRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestAppointmentRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewAppointmentRequestBuilder(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestAvailabilityRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewAvailabilityRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestCalendarRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewCalendarRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestConfigurationRequestBuilder_Get_NilRequestAdapter(t *testing.T) {
	builder := NewConfigurationRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestExecuteRuleConditionsRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewExecuteRuleConditionsRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestUserWindowRequestBuilder_Post_NilRequestAdapter(t *testing.T) {
	builder := NewUserWindowRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}
