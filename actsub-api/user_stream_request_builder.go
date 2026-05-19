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
	userStreamURLTemplate = "{+baseurl}/api/now/v1/actsub/userstream"
)

// UserStreamRequestBuilder provides operations to manage user activity streams.
type UserStreamRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewUserStreamRequestBuilderInternal instantiates a new UserStreamRequestBuilder.
func NewUserStreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserStreamRequestBuilder {
	return &UserStreamRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, userStreamURLTemplate, pathParameters),
	}
}

// ByProfileId returns a UserStreamItemRequestBuilder.
func (rB *UserStreamRequestBuilder) ByProfileId(profileId string) *UserStreamItemRequestBuilder {
	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["profileId"] = profileId
	return NewUserStreamItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// UserStreamItemRequestBuilder provides operations to manage a specific user activity stream.
type UserStreamItemRequestBuilder struct {
	newInternal.RequestBuilder
}

const userStreamItemURLTemplate = "{+baseurl}/api/now/v1/actsub/userstream/{profileId}"

// NewUserStreamItemRequestBuilderInternal instantiates a new UserStreamItemRequestBuilder.
func NewUserStreamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserStreamItemRequestBuilder {
	return &UserStreamItemRequestBuilder{
		newInternal.NewBaseRequestBuilder(requestAdapter, userStreamItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a specific user activity stream.
func (rB *UserStreamItemRequestBuilder) Get(ctx context.Context, config *UserStreamRequestBuilderGetRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
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
func (rB *UserStreamItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *UserStreamRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
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

// Put sends a PUT request to update a specific user activity stream.
func (rB *UserStreamItemRequestBuilder) Put(ctx context.Context, body *ActivitySubscriptionModel, config *UserStreamRequestBuilderPutRequestConfiguration) (*newInternal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
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

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *UserStreamItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body *ActivitySubscriptionModel, config *UserStreamRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
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
