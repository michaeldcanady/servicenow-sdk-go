package actsubapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestContextsRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewContextsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := internal.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestFacetsInstanceRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewFacetsInstanceRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "activity_context": "ctx1", "context_instance": "inst1"}, adapter)

	mockRes := internal.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestPreferencesRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewPreferencesRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewActivitySubscriptionModel(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestPreferenceItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewPreferenceItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "profileId": "prof1"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestSubscriptionItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewSubscriptionItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "subscriber_id": "sub1"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestIsSubscribedRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewIsSubscribedRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sub_obj_id": "obj1"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestSubscribeRequestBuilder_Post(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewSubscribeRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sub_obj_id": "obj1"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Post(context.Background(), NewActivitySubscriptionModel(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestUnsubscribeRequestBuilder_Delete(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewUnsubscribeRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sub_obj_id": "obj1"}, adapter)

	adapter.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := builder.Delete(context.Background(), nil)

	assert.NoError(t, err)
}

func TestFollowingItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewFollowingItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "follower": "user1"}, adapter)

	mockRes := internal.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestSubObjectsRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewSubObjectsRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	mockRes := internal.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestSubscriberItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewSubscriberItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "sub_obj_id": "obj1"}, adapter)

	mockRes := internal.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestUserStreamItemRequestBuilder_Get(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewUserStreamItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "profileId": "prof1"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Get(context.Background(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestUserStreamItemRequestBuilder_Put(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewUserStreamItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "profileId": "prof1"}, adapter)

	mockRes := internal.NewBaseServiceNowItemResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue)
	adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(mockRes, nil)

	resp, err := builder.Put(context.Background(), NewActivitySubscriptionModel(), nil)

	assert.NoError(t, err)
	assert.Equal(t, mockRes, resp)
}

func TestActSubRequestBuilder_Hierarchy_Extra(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewActSubRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

	assert.NotNil(t, builder.Preferences().ByProfileId("prof1"))
	assert.NotNil(t, builder.Subscribers().BySubObject("obj1"))
}

