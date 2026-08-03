package actsubapi

import (
	"context"
	"maps"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	preferencesURLTemplate = "{+baseurl}/api/now/v1/actsub/preferences"
)

// PreferencesRequestBuilder provides operations to manage preferences.
type PreferencesRequestBuilder struct {
	core.RequestBuilder
}

var _ core.ItemPostRequestBuilder[*ActivitySubscription, abstractions.DefaultQueryParameters, abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]] = (*PreferencesRequestBuilder)(nil)

// NewPreferencesRequestBuilderInternal instantiates a new PreferencesRequestBuilder.
func NewPreferencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PreferencesRequestBuilder {
	return &PreferencesRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, preferencesURLTemplate, pathParameters),
	}
}

// NewPreferencesRequestBuilder instantiates a new [PreferencesRequestBuilder] with a raw URL and request adapter.
func NewPreferencesRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *PreferencesRequestBuilder {
	return NewPreferencesRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to create preferences.
func (rB *PreferencesRequestBuilder) Post(ctx context.Context, body *ActivitySubscription, config *PreferencesRequestBuilderPostRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscription], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToPostRequestInformation(ctx, body, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscription](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscription]), nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *PreferencesRequestBuilder) ToPostRequestInformation(ctx context.Context, body *ActivitySubscription, config *PreferencesRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	if !conversion.IsNil(body) {
		err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
		if err != nil {
			return nil, err
		}
	}

	return requestInfo, nil
}

// ByProfileID returns a PreferenceItemRequestBuilder.
func (rB *PreferencesRequestBuilder) ByProfileID(profileID string) *PreferenceItemRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	pathParameters := maps.Clone(rB.GetPathParameters())
	pathParameters["profileId"] = profileID
	return NewPreferenceItemRequestBuilderInternal(pathParameters, rB.GetRequestAdapter())
}

// PreferenceItemRequestBuilder provides operations to manage preferences for a specific profile.
type PreferenceItemRequestBuilder struct {
	core.RequestBuilder
}

const preferenceItemURLTemplate = "{+baseurl}/api/now/v1/actsub/preferences/{profileId}"

// NewPreferenceItemRequestBuilderInternal instantiates a new PreferenceItemRequestBuilder.
func NewPreferenceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *PreferenceItemRequestBuilder {
	return &PreferenceItemRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, preferenceItemURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve preferences.
func (rB *PreferenceItemRequestBuilder) Get(ctx context.Context, config *PreferencesRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscription], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*ActivitySubscription](CreateActivitySubscriptionModelFromDiscriminatorValue), core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	return res.(*core.BaseServiceNowItemResponse[*ActivitySubscription]), nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *PreferenceItemRequestBuilder) ToGetRequestInformation(_ context.Context, config *PreferencesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
