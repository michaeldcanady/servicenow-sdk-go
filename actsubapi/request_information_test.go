package actsubapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
)

func TestActivitiesRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActivitiesRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestContextsRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewContextsRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestFacetsInstanceRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewFacetsInstanceRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestFollowingItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewFollowingItemRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestPreferenceItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewPreferenceItemRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestPreferencesRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewPreferencesRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), NewActivitySubscriptionModel(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "POST", requestInfo.Method.String())
}

func TestSubObjectsRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewSubObjectsRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestSubscriberItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewSubscriberItemRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestSubscriptionItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewSubscriptionItemRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestIsSubscribedRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewIsSubscribedRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestSubscribeRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewSubscribeRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToPostRequestInformation(context.Background(), NewActivitySubscriptionModel(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "POST", requestInfo.Method.String())
}

func TestUnsubscribeRequestBuilder_ToDeleteRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewUnsubscribeRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToDeleteRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "DELETE", requestInfo.Method.String())
}

func TestUserStreamItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewUserStreamItemRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "GET", requestInfo.Method.String())
}

func TestUserStreamItemRequestBuilder_ToPutRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewUserStreamItemRequestBuilderInternal(nil, adapter)

	requestInfo, err := builder.ToPutRequestInformation(context.Background(), NewActivitySubscriptionModel(), nil)
	assert.NoError(t, err)
	assert.NotNil(t, requestInfo)
	assert.Equal(t, "PUT", requestInfo.Method.String())
}
