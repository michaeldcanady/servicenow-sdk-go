package caseapi

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

const caseActivitiesURLTemplate = "{+baseurl}/api/sn_customerservice/v1/case/{id}/activities{?sysparm_activity_type,sysparm_limit,sysparm_offset}"

// CaseActivitiesRequestBuilder provides operations to manage case activities.
type CaseActivitiesRequestBuilder struct {
	core.RequestBuilder
}

// NewCaseActivitiesRequestBuilderInternal instantiates a new [CaseActivitiesRequestBuilder].
func NewCaseActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseActivitiesRequestBuilder {
	return &CaseActivitiesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, caseActivitiesURLTemplate, pathParameters),
	}
}

// NewCaseActivitiesRequestBuilder instantiates a new [FindServiceRequestBuilder].
func NewCaseActivitiesRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *CaseActivitiesRequestBuilder {
	return NewCaseActivitiesRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get sends a GET request to retrieve case activities.
func (rB *CaseActivitiesRequestBuilder) Get(ctx context.Context, config *CaseActivitiesRequestBuilderGetRequestConfiguration) (ActivitiesResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateActivitiesResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(ActivitiesResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*ActivitiesResponse)(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a [abstractions.RequestInformation] object for a GET request.
func (rB *CaseActivitiesRequestBuilder) ToGetRequestInformation(_ context.Context, config *CaseActivitiesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
