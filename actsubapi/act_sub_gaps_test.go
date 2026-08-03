package actsubapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
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
	t.Run("FollowingsRequestBuilder ByFollower", func(t *testing.T) {
		var builder *FollowingsRequestBuilder
		assert.Nil(t, builder.ByFollower("user"))
		assert.Nil(t, (&FollowingsRequestBuilder{}).ByFollower("user"))
	})

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

	t.Run("UserStreamRequestBuilder ByProfileID", func(t *testing.T) {
		var builder *UserStreamRequestBuilder
		assert.Nil(t, builder.ByProfileID("profile"))
		assert.Nil(t, (&UserStreamRequestBuilder{}).ByProfileID("profile"))
	})

	t.Run("SubscribersRequestBuilder BySubObject", func(t *testing.T) {
		var builder *SubscribersRequestBuilder
		assert.Nil(t, builder.BySubObject("object"))
		assert.Nil(t, (&SubscribersRequestBuilder{}).BySubObject("object"))
	})

	t.Run("PreferencesRequestBuilder ByProfileID", func(t *testing.T) {
		var builder *PreferencesRequestBuilder
		assert.Nil(t, builder.ByProfileID("profile"))
		assert.Nil(t, (&PreferencesRequestBuilder{}).ByProfileID("profile"))
	})

	t.Run("SubscriptionsRequestBuilder BySubscriberID", func(t *testing.T) {
		var builder *SubscriptionsRequestBuilder
		assert.Nil(t, builder.BySubscriberID("subscriber"))
		assert.Nil(t, (&SubscriptionsRequestBuilder{}).BySubscriberID("subscriber"))
	})

	t.Run("SubscriptionsRequestBuilder ByObjectID", func(t *testing.T) {
		var builder *SubscriptionsRequestBuilder
		assert.Nil(t, builder.ByObjectID("object"))
		assert.Nil(t, (&SubscriptionsRequestBuilder{}).ByObjectID("object"))
	})

	t.Run("SubscriptionObjectRequestBuilder actions", func(t *testing.T) {
		var builder *SubscriptionObjectRequestBuilder
		assert.Nil(t, builder.IsSubscribed())
		assert.Nil(t, builder.Subscribe())
		assert.Nil(t, builder.Unsubscribe())

		empty := &SubscriptionObjectRequestBuilder{}
		assert.Nil(t, empty.IsSubscribed())
		assert.Nil(t, empty.Subscribe())
		assert.Nil(t, empty.Unsubscribe())
	})
}

// TestNavigationMethods_Clone guards the path-parameter clone in each navigation method: the key
// a child adds must not leak back into its parent.
func TestNavigationMethods_Clone(t *testing.T) {
	adapter := mocking.NewMockRequestAdapter()

	tests := []struct {
		name    string
		build   func() (parentParams map[string]string, child core.RequestBuilder)
		addsKey string
	}{
		{
			name: "SubscriptionsRequestBuilder BySubscriberID",
			build: func() (map[string]string, core.RequestBuilder) {
				parent := NewSubscriptionsRequestBuilderInternal(gapsTestPathParameters(), adapter)

				return parent.GetPathParameters(), parent.BySubscriberID("subscriber-1")
			},
			addsKey: "subscriber_id",
		},
		{
			name: "SubscriptionsRequestBuilder ByObjectID",
			build: func() (map[string]string, core.RequestBuilder) {
				parent := NewSubscriptionsRequestBuilderInternal(gapsTestPathParameters(), adapter)

				return parent.GetPathParameters(), parent.ByObjectID("object-1")
			},
			addsKey: "sub_obj_id",
		},
		{
			// Note the key is camelCase here, unlike the snake_case keys the other
			// navigation methods use. It matches this builder's URL template, which is
			// what actually matters for substitution.
			name: "UserStreamRequestBuilder ByProfileID",
			build: func() (map[string]string, core.RequestBuilder) {
				parent := NewUserStreamRequestBuilderInternal(gapsTestPathParameters(), adapter)

				return parent.GetPathParameters(), parent.ByProfileID("profile-1")
			},
			addsKey: "profileId",
		},
		{
			name: "PreferencesRequestBuilder ByProfileID",
			build: func() (map[string]string, core.RequestBuilder) {
				parent := NewPreferencesRequestBuilderInternal(gapsTestPathParameters(), adapter)

				return parent.GetPathParameters(), parent.ByProfileID("profile-1")
			},
			addsKey: "profileId",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			parentParams, child := test.build()

			require.NotNil(t, child)
			assert.Contains(t, child.GetPathParameters(), test.addsKey)
			assert.NotContains(t, parentParams, test.addsKey)

			// The key only matters if the URL template actually substitutes it, so resolve
			// the URI and confirm the value lands in the path.
			requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(
				abstractions.GET, child.GetURLTemplate(), child.GetPathParameters(),
			)
			uri, err := requestInfo.GetUri()
			require.NoError(t, err)
			assert.Contains(t, uri.String(), child.GetPathParameters()[test.addsKey],
				"the path parameter must resolve into the URL")
		})
	}
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

	return []verbCall{
		{
			name: "ActivitiesRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewActivitiesRequestBuilderInternal(params, a).Get(ctx, nil)
			},
		},
		{
			name: "ContextsRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewContextsRequestBuilderInternal(params, a).Get(ctx, nil)
			},
		},
		{
			name: "SubObjectsRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewSubObjectsRequestBuilderInternal(params, a).Get(ctx, nil)
			},
		},
		{
			name: "FollowingItemRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewFollowingsRequestBuilderInternal(params, a).ByFollower("user").Get(ctx, nil)
			},
		},
		{
			name: "FacetsInstanceRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewFacetsRequestBuilderInternal(params, a).ByContext("c").ByInstance("i").Get(ctx, nil)
			},
		},
		{
			name: "SubscriberItemRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewSubscribersRequestBuilderInternal(params, a).BySubObject("o").Get(ctx, nil)
			},
		},
		{
			name: "SubscriptionItemRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewSubscriptionsRequestBuilderInternal(params, a).BySubscriberID("s").Get(ctx, nil)
			},
		},
		{
			name: "IsSubscribedRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewSubscriptionsRequestBuilderInternal(params, a).ByObjectID("o").IsSubscribed().Get(ctx, nil)
			},
		},
		{
			name: "UserStreamItemRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewUserStreamRequestBuilderInternal(params, a).ByProfileID("p").Get(ctx, nil)
			},
		},
		{
			name: "PreferenceItemRequestBuilder Get",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewPreferencesRequestBuilderInternal(params, a).ByProfileID("p").Get(ctx, nil)
			},
		},
		{
			name: "PreferencesRequestBuilder Post",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewPreferencesRequestBuilderInternal(params, a).Post(ctx, NewActivitySubscription(), nil)
			},
		},
		{
			name: "SubscribeRequestBuilder Post",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewSubscriptionsRequestBuilderInternal(params, a).ByObjectID("o").Subscribe().
					Post(ctx, NewActivitySubscription(), nil)
			},
		},
		{
			name: "UserStreamItemRequestBuilder Put",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewUserStreamRequestBuilderInternal(params, a).ByProfileID("p").
					Put(ctx, NewActivitySubscription(), nil)
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
// back a nil response and a nil error rather than attempting a type assertion.
func TestVerbs_NilResponse(t *testing.T) {
	for _, verb := range verbCalls() {
		t.Run(verb.name, func(t *testing.T) {
			response, err := verb.call(newGapsAdapter(nil, nil))

			require.NoError(t, err)
			assert.Nil(t, response)
		})
	}
}

// TestUnsubscribeRequestBuilder_DeleteError covers the adapter-error branch of the one verb that
// reports only an error. Its happy path is covered in request_builders_extra_test.go.
func TestUnsubscribeRequestBuilder_DeleteError(t *testing.T) {
	builder := NewSubscriptionsRequestBuilderInternal(gapsTestPathParameters(), newGapsAdapter(nil, errSend)).
		ByObjectID("o").Unsubscribe()

	require.ErrorIs(t, builder.Delete(context.Background(), nil), errSend)
}

// TestBodyVerbs_SerializationFailure covers the body-write error branch of the verbs that carry a
// request body — the only way their request-information builders can fail past the nil guard.
func TestBodyVerbs_SerializationFailure(t *testing.T) {
	ctx := context.Background()
	params := gapsTestPathParameters()

	newFailingAdapter := func() *mocking.MockRequestAdapter {
		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(errSend)
		writer.On("Close").Return(nil)

		factory := mocking.NewMockSerializationWriterFactory()
		factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

		adapter := mocking.NewMockRequestAdapter()
		adapter.On("GetSerializationWriterFactory").Return(factory)

		return adapter
	}

	tests := []struct {
		name string
		call func(adapter *mocking.MockRequestAdapter) (any, error)
	}{
		{
			name: "PreferencesRequestBuilder Post",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewPreferencesRequestBuilderInternal(params, a).Post(ctx, NewActivitySubscription(), nil)
			},
		},
		{
			name: "SubscribeRequestBuilder Post",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewSubscriptionsRequestBuilderInternal(params, a).ByObjectID("o").Subscribe().
					Post(ctx, NewActivitySubscription(), nil)
			},
		},
		{
			name: "UserStreamItemRequestBuilder Put",
			call: func(a *mocking.MockRequestAdapter) (any, error) {
				return NewUserStreamRequestBuilderInternal(params, a).ByProfileID("p").
					Put(ctx, NewActivitySubscription(), nil)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			adapter := newFailingAdapter()

			response, err := test.call(adapter)

			require.ErrorIs(t, err, errSend)
			assert.Nil(t, response)
			adapter.AssertNotCalled(t, "Send")
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
