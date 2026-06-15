package actsubapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	subscribersURLTemplate = "{+baseurl}/api/now/v1/actsub/subscribers"
)

// SubscribersRequestBuilder provides operations to manage subscribers.
type SubscribersRequestBuilder struct {
	internal.RequestBuilder
}

// NewSubscribersRequestBuilderInternal instantiates a new SubscribersRequestBuilder.
func NewSubscribersRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscribersRequestBuilder {
	return &SubscribersRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, subscribersURLTemplate, pathParameters),
	}
}

// BySubObject returns a SubscriberItemRequestBuilder.
func (rB *SubscribersRequestBuilder) BySubObject(subObject string) *SubscriberItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["sub_object"] = subObject
	return NewSubscriberItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// SubscriberItemRequestBuilder provides operations to manage subscribers for a specific object.
type SubscriberItemRequestBuilder struct {
	internal.RequestBuilder
}

const subscriberItemURLTemplate = "{+baseurl}/api/now/v1/actsub/subscribers/{sub_object}"

// NewSubscriberItemRequestBuilderInternal instantiates a new SubscriberItemRequestBuilder.
func NewSubscriberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SubscriberItemRequestBuilder {
	return &SubscriberItemRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, subscriberItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve subscribers.
func (rB *SubscriberItemRequestBuilder) Get(ctx context.Context, config *SubscribersRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowCollectionResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), nil)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowCollectionResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *SubscriberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *SubscribersRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !conversion.IsNil(config) {
		if headers := config.Headers; !conversion.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := config.Options; !conversion.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
		if queryParameters := config.QueryParameters; !conversion.IsNil(queryParameters) {
			kiotaRequestInfo.AddQueryParameters(queryParameters)
		}
	}
	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
