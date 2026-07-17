package statsapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	// statsURLTemplate is the url template for Service-Now's Stats API.
	//
	// sysparm_group_by and sysparm_having are intentionally excluded: when set, the platform
	// returns an array of StatsResult under "result" instead of a single object, which this
	// request builder does not yet model (see StatsResult).
	statsURLTemplate = "{+baseurl}/api/now/stats{/table}{?sysparm_count,sysparm_sum_fields,sysparm_avg_fields,sysparm_min_fields,sysparm_max_fields,sysparm_query,sysparm_display_value}"
)

// StatsRequestBuilder provides operations to retrieve aggregate statistics for a Service-Now table.
type StatsRequestBuilder struct {
	core.RequestBuilder
}

// NewStatsRequestBuilderInternal instantiates a new StatsRequestBuilder with the provided path parameters and request adapter.
func NewStatsRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *StatsRequestBuilder {
	return &StatsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, statsURLTemplate, pathParameters),
	}
}

// NewStatsRequestBuilder instantiates a new StatsRequestBuilder with a raw URL.
func NewStatsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *StatsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[internal.RawURLKey] = rawURL
	return NewStatsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get sends an HTTP GET request and returns the aggregate statistics for the table.
func (rB *StatsRequestBuilder) Get(ctx context.Context, requestConfiguration *StatsRequestBuilderGetRequestConfiguration) (core.ServiceNowItemResponse[*StatsResultModel], error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if conversion.IsNil(rB.GetRequestAdapter()) {
		return nil, snerrors.ErrNilRequestAdapter
	}

	if conversion.IsNil(requestConfiguration) {
		requestConfiguration = &StatsRequestBuilderGetRequestConfiguration{}
	}

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := core.DefaultErrorMapping()
	resp, err := rB.GetRequestAdapter().Send(ctx, requestInfo, core.ServiceNowItemResponseFromDiscriminatorValue[*StatsResultModel](CreateStatsResultFromDiscriminatorValue), errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := resp.(core.ServiceNowItemResponse[*StatsResultModel])
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*core.ServiceNowItemResponse[*StatsResultModel])(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation converts provided parameters into request information
func (rB *StatsRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *StatsRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &internal.KiotaRequestInformation{RequestInformation: requestInfo}

	internal.ConfigureRequestInformation(kiotaRequestInfo, requestConfiguration)

	kiotaRequestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return kiotaRequestInfo.RequestInformation, nil
}
