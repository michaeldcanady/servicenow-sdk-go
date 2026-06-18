package actsubapi

import (
	"context"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	userStreamURLTemplate = "{+baseurl}/api/now/v1/actsub/userstream"
)

// UserStreamRequestBuilder provides operations to manage user activity streams.
type UserStreamRequestBuilder struct {
	internal.RequestBuilder
}

// NewUserStreamRequestBuilderInternal instantiates a new UserStreamRequestBuilder.
func NewUserStreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserStreamRequestBuilder {
	return &UserStreamRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, userStreamURLTemplate, pathParameters),
	}
}

// ByProfileId returns a UserStreamItemRequestBuilder.
func (rB *UserStreamRequestBuilder) ByProfileId(profileId string) *UserStreamItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["profileId"] = profileId
	return NewUserStreamItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// UserStreamItemRequestBuilder provides operations to manage a specific user activity stream.
type UserStreamItemRequestBuilder struct {
	internal.RequestBuilder
}

const userStreamItemURLTemplate = "{+baseurl}/api/now/v1/actsub/userstream/{profileId}"

// NewUserStreamItemRequestBuilderInternal instantiates a new UserStreamItemRequestBuilder.
func NewUserStreamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserStreamItemRequestBuilder {
	return &UserStreamItemRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, userStreamItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a specific user activity stream.
func (rB *UserStreamItemRequestBuilder) Get(ctx context.Context, config *UserStreamRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *UserStreamItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *UserStreamRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return kiotaRequestInfo.RequestInformation, nil
}

// Put sends a PUT request to update a specific user activity stream.
func (rB *UserStreamItemRequestBuilder) Put(ctx context.Context, body *ActivitySubscriptionModel, config *UserStreamRequestBuilderPutRequestConfiguration) (*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := internal.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *UserStreamItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body *ActivitySubscriptionModel, config *UserStreamRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
		if err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}
