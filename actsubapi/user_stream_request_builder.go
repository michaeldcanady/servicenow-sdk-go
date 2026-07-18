package actsubapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	userStreamURLTemplate = "{+baseurl}/api/now/v1/actsub/userstream"
)

// UserStreamRequestBuilder provides operations to manage user activity streams.
type UserStreamRequestBuilder struct {
	core.RequestBuilder
}

// NewUserStreamRequestBuilderInternal instantiates a new UserStreamRequestBuilder.
func NewUserStreamRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserStreamRequestBuilder {
	return &UserStreamRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, userStreamURLTemplate, pathParameters),
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
	core.RequestBuilder
}

const userStreamItemURLTemplate = "{+baseurl}/api/now/v1/actsub/userstream/{profileId}"

// NewUserStreamItemRequestBuilderInternal instantiates a new UserStreamItemRequestBuilder.
func NewUserStreamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UserStreamItemRequestBuilder {
	return &UserStreamItemRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, userStreamItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve a specific user activity stream.
func (rB *UserStreamItemRequestBuilder) Get(ctx context.Context, config *UserStreamRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *UserStreamItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *UserStreamRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}

// Put sends a PUT request to update a specific user activity stream.
func (rB *UserStreamItemRequestBuilder) Put(ctx context.Context, body *ActivitySubscriptionModel, config *UserStreamRequestBuilderPutRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo, err := rB.ToPutRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToPutRequestInformation creates a RequestInformation object for a PUT request.
func (rB *UserStreamItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body *ActivitySubscriptionModel, config *UserStreamRequestBuilderPutRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.PUT, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
		if err != nil {
			return nil, err
		}
	}

	return kiotaRequestInfo.RequestInformation, nil
}
