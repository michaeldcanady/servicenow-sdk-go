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
	preferencesURLTemplate = "{+baseurl}/api/now/v1/actsub/preferences"
)

// PreferencesRequestBuilder provides operations to manage preferences.
type PreferencesRequestBuilder struct {
	internal.RequestBuilder
}

// NewPreferencesRequestBuilderInternal instantiates a new PreferencesRequestBuilder.
func NewPreferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PreferencesRequestBuilder {
	return &PreferencesRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, preferencesURLTemplate, pathParameters),
	}
}

// Post sends a POST request to create preferences.
func (rB *PreferencesRequestBuilder) Post(ctx context.Context, body *ActivitySubscriptionModel, config *PreferencesRequestBuilderPostRequestConfiguration) (*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), internal.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *PreferencesRequestBuilder) ToPostRequestInformation(ctx context.Context, body *ActivitySubscriptionModel, config *PreferencesRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	if !conversion.IsNil(body) {
		err := kiotaRequestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internal.ContentTypeApplicationJSON, body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// ByProfileId returns a PreferenceItemRequestBuilder.
func (rB *PreferencesRequestBuilder) ByProfileId(profileId string) *PreferenceItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["profileId"] = profileId
	return NewPreferenceItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// PreferenceItemRequestBuilder provides operations to manage preferences for a specific profile.
type PreferenceItemRequestBuilder struct {
	internal.RequestBuilder
}

const preferenceItemURLTemplate = "{+baseurl}/api/now/v1/actsub/preferences/{profileId}"

// NewPreferenceItemRequestBuilderInternal instantiates a new PreferenceItemRequestBuilder.
func NewPreferenceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PreferenceItemRequestBuilder {
	return &PreferenceItemRequestBuilder{
		internal.NewBaseRequestBuilder(requestAdapter, preferenceItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve preferences.
func (rB *PreferenceItemRequestBuilder) Get(ctx context.Context, config *PreferencesRequestBuilderGetRequestConfiguration) (*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, internal.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue), internal.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res.(*internal.BaseServiceNowItemResponse[*ActivitySubscriptionModel]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *PreferenceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, config *PreferencesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, config)

	kiotaRequestInfo.Headers.TryAdd(internalHttp.RequestHeaderAccept.String(), internal.ContentTypeApplicationJSON)

	return requestInfo, nil
}
