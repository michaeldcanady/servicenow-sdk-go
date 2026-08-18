package actsubapi

import (
	"context"
	"errors"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	activitiesURLTemplate = "{+baseurl}/api/now/v1/actsub/activities{?before,context,context_instance,end_date,facets,last,record_id,start_date,stFrom}"
)

// ActivitiesRequestBuilder provides operations to manage activities.
type ActivitiesRequestBuilder struct {
	core.RequestBuilder
}

// NewActivitiesRequestBuilderInternal instantiates a new ActivitiesRequestBuilder.
func NewActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ActivitiesRequestBuilder {
	return &ActivitiesRequestBuilder{
		core.NewBaseRequestBuilder(requestAdapter, activitiesURLTemplate, pathParameters),
	}
}

// NewActivitiesRequestBuilder instantiates a new [ActivitiesRequestBuilder] with a raw URL and request adapter.
func NewActivitiesRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ActivitiesRequestBuilder {
	return NewActivitiesRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get sends a GET request to retrieve activities.
func (rB *ActivitiesRequestBuilder) Get(ctx context.Context, config *ActivitiesRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[*ActivitySubscription], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	if conversion.IsNil(config) || isNilOrEmpty(config.QueryParameters) {
		return nil, ErrNilOrEmptyQueryParameters
	}

	if isNilOrEmpty(config.QueryParameters.Context) {
		return nil, ErrNilContextQueryParameter
	}

	if isNilOrEmpty(config.QueryParameters.ContextInstance) {
		return nil, ErrNilContextInstanceQueryParameter
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, config)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateActivitySubscriptionItemResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, nil
	}

	typedRes, ok := res.(*ActivitySubscriptionItemResponse)
	if !ok {
		// TODO: standardize error
		return nil, errors.New("unexpected type")
	}

	return typedRes, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request to retrieve activities.
func (rB *ActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, config *ActivitiesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
