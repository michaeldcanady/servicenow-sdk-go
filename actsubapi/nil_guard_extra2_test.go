package actsubapi

import (
	"context"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Nil-builder guard coverage for UserStreamItemRequestBuilder, PreferencesRequestBuilder,
// PreferenceItemRequestBuilder, and the Subscriptions family of builders — none of which
// previously had a dedicated _test.go file.

func TestUserStreamItemRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *UserStreamItemRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestUserStreamItemRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *UserStreamItemRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestUserStreamItemRequestBuilder_Put_NilBuilder(t *testing.T) {
	var builder *UserStreamItemRequestBuilder

	resp, err := builder.Put(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestUserStreamItemRequestBuilder_ToPutRequestInformation_NilBuilder(t *testing.T) {
	var builder *UserStreamItemRequestBuilder

	requestInfo, err := builder.ToPutRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestPreferencesRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *PreferencesRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestPreferencesRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *PreferencesRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestPreferenceItemRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *PreferenceItemRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestPreferenceItemRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *PreferenceItemRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestSubscriptionItemRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *SubscriptionItemRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestSubscriptionItemRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *SubscriptionItemRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestIsSubscribedRequestBuilder_Get_NilBuilder(t *testing.T) {
	var builder *IsSubscribedRequestBuilder

	resp, err := builder.Get(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestIsSubscribedRequestBuilder_ToGetRequestInformation_NilBuilder(t *testing.T) {
	var builder *IsSubscribedRequestBuilder

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestSubscribeRequestBuilder_Post_NilBuilder(t *testing.T) {
	var builder *SubscribeRequestBuilder

	resp, err := builder.Post(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, resp)
}

func TestSubscribeRequestBuilder_ToPostRequestInformation_NilBuilder(t *testing.T) {
	var builder *SubscribeRequestBuilder

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}

func TestUnsubscribeRequestBuilder_Delete_NilBuilder(t *testing.T) {
	var builder *UnsubscribeRequestBuilder

	err := builder.Delete(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
}

func TestUnsubscribeRequestBuilder_ToDeleteRequestInformation_NilBuilder(t *testing.T) {
	var builder *UnsubscribeRequestBuilder

	requestInfo, err := builder.ToDeleteRequestInformation(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
	assert.Nil(t, requestInfo)
}
