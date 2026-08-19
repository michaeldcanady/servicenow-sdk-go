package activitysubscriptionsapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const gapsTestBaseURL = "https://example.service-now.com"

// errSend is a stand-in for a transport-level failure from the request adapter.
var errSend = errors.New("send failed")

func gapsTestPathParameters() map[string]string {
	return map[string]string{"baseurl": gapsTestBaseURL}
}

// newGapsAdapter returns an adapter that can serialize a body and whose Send/SendNoContent
// return the supplied outcome.
func newGapsAdapter(response any, err error) *mocking.MockRequestAdapter {
	adapter := mocking.NewMockRequestAdapter()
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(response, err)
	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(err)

	return adapter
}

// TestNavigationMethods_NilReceivers covers the nil-receiver guard on every navigation method in
// the package. Each returns a nil child rather than panicking, so a broken chain surfaces at the
// call site instead of mid-traversal.
func TestNavigationMethods_NilReceivers(t *testing.T) {
	t.Run("FacetsRequestBuilder ByContext", func(t *testing.T) {
		var builder *FacetsRequestBuilder
		assert.Nil(t, builder.ByContext("context"))
		assert.Nil(t, (&FacetsRequestBuilder{}).ByContext("context"))
	})

	t.Run("FacetsContextRequestBuilder ByInstance", func(t *testing.T) {
		var builder *FacetsContextRequestBuilder
		assert.Nil(t, builder.ByInstance("instance"))
		assert.Nil(t, (&FacetsContextRequestBuilder{}).ByInstance("instance"))
	})
}

// verbCall names one verb and how to invoke it, so the shared adapter-error and empty-response
// outcomes can be driven across every verb in the package.
type verbCall struct {
	name string
	call func(adapter *mocking.MockRequestAdapter) (any, error)
}

func verbCalls() []verbCall {
	ctx := context.Background()
	params := gapsTestPathParameters()

	activitiesConfig := &ActivitiesRequestBuilderGetRequestConfiguration{
		QueryParameters: &ActivitiesRequestBuilderGetQueryParameters{
			Context:         strPtr("ctx1"),
			ContextInstance: strPtr("inst1"),
		},
	}

	return []verbCall{
		{
			name: "ActivitiesRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewActivitiesRequestBuilderInternal(params, a).Get(ctx, activitiesConfig)
			},
		},

		{
			name: "FacetsInstanceRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewFacetsRequestBuilderInternal(params, a).ByContext("c").ByInstance("i").Get(ctx, nil)
			},
		},
	}
}

// TestVerbs_AdapterErrorPropagates covers the Send-error branch of every verb that returns a
// response value.
func TestVerbs_AdapterErrorPropagates(t *testing.T) {
	for _, verb := range verbCalls() {
		t.Run(verb.name, func(t *testing.T) {
			response, err := verb.call(newGapsAdapter(nil, errSend))

			require.ErrorIs(t, err, errSend)
			assert.Nil(t, response)
		})
	}
}

// TestVerbs_NilResponse covers the branch where the adapter reports no response: the verbs hand
// back a nil response and snerrors.ErrNilResponse rather than attempting a type assertion.
func TestVerbs_NilResponse(t *testing.T) {
	for _, verb := range verbCalls() {
		t.Run(verb.name, func(t *testing.T) {
			response, err := verb.call(newGapsAdapter(nil, nil))

			require.ErrorIs(t, err, snerrors.ErrNilResponse)
			assert.Nil(t, response)
		})
	}
}

// TestModels_SerializeNilReceiver covers the nil-receiver branch of each model's Serialize.
// These models are read-only in practice — Serialize writes nothing either way — so the only
// behaviour worth pinning down is that a nil receiver is tolerated rather than panicking.
func TestModels_SerializeNilReceiver(t *testing.T) {
	writer := mocking.NewMockSerializationWriter()

	t.Run("ActivitySubscriptionModel", func(t *testing.T) {
		var model *ActivitySubscription
		require.NoError(t, model.Serialize(writer))
		require.NoError(t, NewActivitySubscription().Serialize(writer))
	})

	t.Run("Activity", func(t *testing.T) {
		var model *Activity
		require.NoError(t, model.Serialize(writer))
		require.NoError(t, NewActivity().Serialize(writer))
	})

	t.Run("Field", func(t *testing.T) {
		var model *Field
		require.NoError(t, model.Serialize(writer))
		require.NoError(t, NewField().Serialize(writer))
	})

	writer.AssertNotCalled(t, "WriteStringValue", mock.Anything, mock.Anything)
}
