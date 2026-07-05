package actsubapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	subscriptionsURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions"
)

// SubscriptionsRequestBuilder provides operations to manage subscriptions.
type SubscriptionsRequestBuilder struct {
	core.RequestBuilder
}

// NewSubscriptionsRequestBuilderInternal instantiates a new SubscriptionsRequestBuilder.
func NewSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriptionsRequestBuilder {
	return &SubscriptionsRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, subscriptionsURLTemplate, pathParameters),
	}
}

// BySubscriberId returns a SubscriptionItemRequestBuilder.
func (rB *SubscriptionsRequestBuilder) BySubscriberId(subscriberId string) *SubscriptionItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["subscriber_id"] = subscriberId
	return NewSubscriptionItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// ByObjectId returns a SubscriptionObjectRequestBuilder.
func (rB *SubscriptionsRequestBuilder) ByObjectId(subObjId string) *SubscriptionObjectRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sub_obj_id"] = subObjId
	return NewSubscriptionObjectRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// SubscriptionItemRequestBuilder provides operations to manage a specific subscription.
type SubscriptionItemRequestBuilder struct {
	core.RequestBuilder
}

const subscriptionItemURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{subscriber_id}"

// NewSubscriptionItemRequestBuilderInternal instantiates a new SubscriptionItemRequestBuilder.
func NewSubscriptionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriptionItemRequestBuilder {
	return &SubscriptionItemRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, subscriptionItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a specific subscription.
func (rB *SubscriptionItemRequestBuilder) Get(ctx context.Context, config *SubscriptionsRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *SubscriptionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *SubscriptionsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// SubscriptionObjectRequestBuilder provides operations to manage subscriptions for a specific object.
type SubscriptionObjectRequestBuilder struct {
	core.RequestBuilder
}

const subscriptionObjectURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}"

// NewSubscriptionObjectRequestBuilderInternal instantiates a new SubscriptionObjectRequestBuilder.
func NewSubscriptionObjectRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriptionObjectRequestBuilder {
	return &SubscriptionObjectRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, subscriptionObjectURLTemplate, pathParameters),
	}
}

// IsSubscribed returns an IsSubscribedRequestBuilder.
func (rB *SubscriptionObjectRequestBuilder) IsSubscribed() *IsSubscribedRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewIsSubscribedRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Subscribe returns a SubscribeRequestBuilder.
func (rB *SubscriptionObjectRequestBuilder) Subscribe() *SubscribeRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewSubscribeRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Unsubscribe returns an UnsubscribeRequestBuilder.
func (rB *SubscriptionObjectRequestBuilder) Unsubscribe() *UnsubscribeRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewUnsubscribeRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// IsSubscribedRequestBuilder provides operations to check if the current user is subscribed to an object.
type IsSubscribedRequestBuilder struct {
	core.RequestBuilder
}

const isSubscribedURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}/isSubscribed"

// NewIsSubscribedRequestBuilderInternal instantiates a new IsSubscribedRequestBuilder.
func NewIsSubscribedRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *IsSubscribedRequestBuilder {
	return &IsSubscribedRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, isSubscribedURLTemplate, pathParameters),
	}
}

// Get sends a GET request to check if the current user is subscribed.
func (rB *IsSubscribedRequestBuilder) Get(ctx context.Context, config *IsSubscribedRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *IsSubscribedRequestBuilder) ToGetRequestInformation(ctx context.Context, config *IsSubscribedRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}

// SubscribeRequestBuilder provides operations to subscribe to an object.
type SubscribeRequestBuilder struct {
	core.RequestBuilder
}

const subscribeURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}/subscribe"

// NewSubscribeRequestBuilderInternal instantiates a new SubscribeRequestBuilder.
func NewSubscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscribeRequestBuilder {
	return &SubscribeRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, subscribeURLTemplate, pathParameters),
	}
}

// Post sends a POST request to subscribe to an object.
func (rB *SubscribeRequestBuilder) Post(ctx context.Context, body *ActivitySubscriptionModel, config *SubscribeRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *SubscribeRequestBuilder) ToPostRequestInformation(ctx context.Context, body *ActivitySubscriptionModel, config *SubscribeRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// UnsubscribeRequestBuilder provides operations to unsubscribe from an object.
type UnsubscribeRequestBuilder struct {
	core.RequestBuilder
}

const unsubscribeURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}/unsubscribe"

// NewUnsubscribeRequestBuilderInternal instantiates a new UnsubscribeRequestBuilder.
func NewUnsubscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UnsubscribeRequestBuilder {
	return &UnsubscribeRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, unsubscribeURLTemplate, pathParameters),
	}
}

// Delete sends a DELETE request to unsubscribe from an object.
func (rB *UnsubscribeRequestBuilder) Delete(ctx context.Context, config *UnsubscribeRequestBuilderDeleteRequestConfiguration) error {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	requestInfo, err := rB.ToDeleteRequestInformation(ctx, config)
	if err != nil {
		return err
	}

	err = rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, core.DefaultErrorMapping())
	if err != nil {
		return err
	}

	return nil
}

// ToDeleteRequestInformation creates a RequestInformation object for a DELETE request.
func (rB *UnsubscribeRequestBuilder) ToDeleteRequestInformation(ctx context.Context, config *UnsubscribeRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
