package actsubapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	subscriptionsURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions"
)

// SubscriptionsRequestBuilder provides operations to manage subscriptions.
type SubscriptionsRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewSubscriptionsRequestBuilderInternal instantiates a new SubscriptionsRequestBuilder.
func NewSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriptionsRequestBuilder {
	return &SubscriptionsRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, subscriptionsURLTemplate, pathParameters),
	}
}

// BySubscriberId returns a SubscriptionItemRequestBuilder.
func (rB *SubscriptionsRequestBuilder) BySubscriberId(subscriberId string) *SubscriptionItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["subscriber_id"] = subscriberId
	return NewSubscriptionItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// ByObjectId returns a SubscriptionObjectRequestBuilder.
func (rB *SubscriptionsRequestBuilder) ByObjectId(subObjId string) *SubscriptionObjectRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sub_obj_id"] = subObjId
	return NewSubscriptionObjectRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// SubscriptionItemRequestBuilder provides operations to manage a specific subscription.
type SubscriptionItemRequestBuilder struct {
	newInternal.RequestBuilder
}

const subscriptionItemURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{subscriber_id}"

// NewSubscriptionItemRequestBuilderInternal instantiates a new SubscriptionItemRequestBuilder.
func NewSubscriptionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriptionItemRequestBuilder {
	return &SubscriptionItemRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, subscriptionItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a specific subscription.
func (rB *SubscriptionItemRequestBuilder) Get(ctx context.Context, config *SubscriptionsRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *SubscriptionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *SubscriptionsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !internal.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}

// SubscriptionObjectRequestBuilder provides operations to manage subscriptions for a specific object.
type SubscriptionObjectRequestBuilder struct {
	newInternal.RequestBuilder
}

const subscriptionObjectURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}"

// NewSubscriptionObjectRequestBuilderInternal instantiates a new SubscriptionObjectRequestBuilder.
func NewSubscriptionObjectRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriptionObjectRequestBuilder {
	return &SubscriptionObjectRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, subscriptionObjectURLTemplate, pathParameters),
	}
}

// IsSubscribed returns an IsSubscribedRequestBuilder.
func (rB *SubscriptionObjectRequestBuilder) IsSubscribed() *IsSubscribedRequestBuilder {
	return NewIsSubscribedRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Subscribe returns a SubscribeRequestBuilder.
func (rB *SubscriptionObjectRequestBuilder) Subscribe() *SubscribeRequestBuilder {
	return NewSubscribeRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// Unsubscribe returns an UnsubscribeRequestBuilder.
func (rB *SubscriptionObjectRequestBuilder) Unsubscribe() *UnsubscribeRequestBuilder {
	return NewUnsubscribeRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// IsSubscribedRequestBuilder provides operations to check if the current user is subscribed to an object.
type IsSubscribedRequestBuilder struct {
	newInternal.RequestBuilder
}

const isSubscribedURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}/isSubscribed"

// NewIsSubscribedRequestBuilderInternal instantiates a new IsSubscribedRequestBuilder.
func NewIsSubscribedRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *IsSubscribedRequestBuilder {
	return &IsSubscribedRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, isSubscribedURLTemplate, pathParameters),
	}
}

// Get sends a GET request to check if the current user is subscribed.
func (rB *IsSubscribedRequestBuilder) Get(ctx context.Context, config *IsSubscribedRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *IsSubscribedRequestBuilder) ToGetRequestInformation(ctx context.Context, config *IsSubscribedRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !internal.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}

// SubscribeRequestBuilder provides operations to subscribe to an object.
type SubscribeRequestBuilder struct {
	newInternal.RequestBuilder
}

const subscribeURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}/subscribe"

// NewSubscribeRequestBuilderInternal instantiates a new SubscribeRequestBuilder.
func NewSubscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscribeRequestBuilder {
	return &SubscribeRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, subscribeURLTemplate, pathParameters),
	}
}

// Post sends a POST request to subscribe to an object.
func (rB *SubscribeRequestBuilder) Post(ctx context.Context, body *ActivitySubscriptionModel, config *SubscribeRequestBuilderPostRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, newInternal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *SubscribeRequestBuilder) ToPostRequestInformation(ctx context.Context, body *ActivitySubscriptionModel, config *SubscribeRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	if !internal.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), newInternal.ContentTypeApplicationJSON, body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// UnsubscribeRequestBuilder provides operations to unsubscribe from an object.
type UnsubscribeRequestBuilder struct {
	newInternal.RequestBuilder
}

const unsubscribeURLTemplate = "{+baseurl}/api/now/v1/actsub/subscriptions/{sub_obj_id}/unsubscribe"

// NewUnsubscribeRequestBuilderInternal instantiates a new UnsubscribeRequestBuilder.
func NewUnsubscribeRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UnsubscribeRequestBuilder {
	return &UnsubscribeRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, unsubscribeURLTemplate, pathParameters),
	}
}

// Delete sends a DELETE request to unsubscribe from an object.
func (rB *UnsubscribeRequestBuilder) Delete(ctx context.Context, config *UnsubscribeRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := rB.ToDeleteRequestInformation(ctx, config)
	if err != nil {
		return err
	}

	err = rB.GetRequestAdapter().SendNoContent(ctx, requestInfo, nil)
	if err != nil {
		return err
	}

	return nil
}

// ToDeleteRequestInformation creates a RequestInformation object for a DELETE request.
func (rB *UnsubscribeRequestBuilder) ToDeleteRequestInformation(ctx context.Context, config *UnsubscribeRequestBuilderDeleteRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.DELETE, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(config) {
		if headers := config.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), newInternal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
