package appointmentbookingapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// errSerialize is a stand-in for an error surfaced while writing a request body.
var errSerialize = errors.New("serialize failed")

// newSerializingAdapter returns an adapter that serializes a request body successfully,
// which every POST builder must get past before it reaches Send.
func newSerializingAdapter() *mocking.MockRequestAdapter {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())

	return adapter
}

// newFailingSerializationAdapter returns an adapter whose writer fails on write. A body-write
// failure is the only way ToPostRequestInformation can fail once past its nil guard.
func newFailingSerializationAdapter() *mocking.MockRequestAdapter {
	writer := mocking.NewMockSerializationWriter()
	writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(errSerialize)
	writer.On("Close").Return(nil)

	factory := mocking.NewMockSerializationWriterFactory()
	factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(factory, nil)

	return adapter
}

// postInvocation names one POST builder and how to invoke it with a valid body. The happy
// paths for these live in appointment_booking_request_builder_test.go; this table drives the
// failure and empty-response outcomes, which are identical across all four builders.
type postInvocation struct {
	name string
	post func(adapter *mocking.MockRequestAdapter) (any, error)
	// postNilBuilder invokes the same method on a nil builder.
	postNilBuilder func() (any, error)
}

func postInvocations() []postInvocation {
	ctx := context.Background()

	return []postInvocation{
		{
			name: "AvailabilityRequestBuilder",
			post: func(adapter *mocking.MockRequestAdapter) (any, error) {
				builder := NewAvailabilityRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter)

				return builder.Post(ctx, NewAvailabilityRequest(), nil)
			},
			postNilBuilder: func() (any, error) {
				var builder *AvailabilityRequestBuilder

				return builder.Post(ctx, NewAvailabilityRequest(), nil)
			},
		},
		{
			name: "AppointmentRequestBuilder",
			post: func(adapter *mocking.MockRequestAdapter) (any, error) {
				builder := NewAppointmentRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter)

				return builder.Post(ctx, NewAppointmentRequest(), nil)
			},
			postNilBuilder: func() (any, error) {
				var builder *AppointmentRequestBuilder

				return builder.Post(ctx, NewAppointmentRequest(), nil)
			},
		},
		{
			name: "ExecuteRuleConditionsRequestBuilder",
			post: func(adapter *mocking.MockRequestAdapter) (any, error) {
				builder := NewExecuteRuleConditionsRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter)

				return builder.Post(ctx, NewExecuteRuleConditionsRequest(), nil)
			},
			postNilBuilder: func() (any, error) {
				var builder *ExecuteRuleConditionsRequestBuilder

				return builder.Post(ctx, NewExecuteRuleConditionsRequest(), nil)
			},
		},
		{
			name: "UserWindowRequestBuilder",
			post: func(adapter *mocking.MockRequestAdapter) (any, error) {
				builder := NewUserWindowRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter)

				return builder.Post(ctx, NewAvailabilityRequest(), nil)
			},
			postNilBuilder: func() (any, error) {
				var builder *UserWindowRequestBuilder

				return builder.Post(ctx, NewAvailabilityRequest(), nil)
			},
		},
	}
}

// TestPostRequestBuilders_AdapterErrorPropagates covers the Send-error branch of every POST builder.
func TestPostRequestBuilders_AdapterErrorPropagates(t *testing.T) {
	for _, invocation := range postInvocations() {
		t.Run(invocation.name, func(t *testing.T) {
			adapter := newSerializingAdapter()
			adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errNetwork)

			response, err := invocation.post(adapter)

			require.ErrorIs(t, err, errNetwork)
			assert.Nil(t, response)
		})
	}
}

// TestPostRequestBuilders_NilResponse covers the "adapter returned no response" branch, where
// the builders return a nil response and a nil error rather than attempting a type assertion.
func TestPostRequestBuilders_NilResponse(t *testing.T) {
	for _, invocation := range postInvocations() {
		t.Run(invocation.name, func(t *testing.T) {
			adapter := newSerializingAdapter()
			adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

			response, err := invocation.post(adapter)

			require.NoError(t, err)
			assert.Nil(t, response)
		})
	}
}

// TestPostRequestBuilders_SerializationFailure covers each builder's propagation of a
// ToPostRequestInformation error, which only a body-write failure can trigger.
func TestPostRequestBuilders_SerializationFailure(t *testing.T) {
	for _, invocation := range postInvocations() {
		t.Run(invocation.name, func(t *testing.T) {
			adapter := newFailingSerializationAdapter()

			response, err := invocation.post(adapter)

			require.ErrorIs(t, err, errSerialize)
			assert.Nil(t, response)
			adapter.AssertNotCalled(t, "Send")
		})
	}
}

// TestPostRequestBuilders_NilBuilder covers the nil-guard prologue of every POST builder.
func TestPostRequestBuilders_NilBuilder(t *testing.T) {
	for _, invocation := range postInvocations() {
		t.Run(invocation.name, func(t *testing.T) {
			response, err := invocation.postNilBuilder()

			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, response)
		})
	}
}

// TestUserWindowRequestBuilder_PostNilBody covers the one POST builder that tolerates a nil
// body: it skips serialization entirely rather than casting nil to a Parsable.
func TestUserWindowRequestBuilder_PostNilBody(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	builder := NewUserWindowRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter)

	response, err := builder.Post(context.Background(), nil, nil)

	require.NoError(t, err)
	assert.Nil(t, response)
	adapter.AssertNotCalled(t, "GetSerializationWriterFactory")
}
