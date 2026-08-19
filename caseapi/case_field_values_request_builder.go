package caseapi

import (
	"context"
	"fmt"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	fieldValuesURLTemplate     = "{+baseurl}/api/sn_customerservice/v1/case/field_values/{field_name}{?sysparm_dependent_value,sysparm_limit,sysparm_offset,sysparm_reference_field_columns,sysparm_query,sysparm_ref_qual_input}"
	itemFieldValuesURLTemplate = "{+baseurl}/api/sn_customerservice/v1/case/{id}/field_values/{field_name}"
)

// CaseFieldValuesRequestBuilder provides operations to manage field values.
type CaseFieldValuesRequestBuilder struct {
	core.RequestBuilder
}

// NewCaseFieldValuesRequestBuilderInternal instantiates a new CaseFieldValuesRequestBuilder for the case-level field_values endpoint.
func NewCaseFieldValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseFieldValuesRequestBuilder {
	return &CaseFieldValuesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, fieldValuesURLTemplate, pathParameters),
	}
}

// NewItemFieldValuesRequestBuilderInternal instantiates a new CaseFieldValuesRequestBuilder for a single case's field_values endpoint.
func NewItemFieldValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CaseFieldValuesRequestBuilder {
	return &CaseFieldValuesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, itemFieldValuesURLTemplate, pathParameters),
	}
}

// Get sends a GET request to retrieve field values.
func (rB *CaseFieldValuesRequestBuilder) Get(ctx context.Context, config *CaseFieldValuesRequestBuilderGetRequestConfiguration) (FieldValuesResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateFieldValuesResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(FieldValuesResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*FieldValuesResponse)(nil))
	}

	return typedResp, nil
}

// ToGetRequestInformation creates a RequestInformation object for a GET request.
func (rB *CaseFieldValuesRequestBuilder) ToGetRequestInformation(_ context.Context, config *CaseFieldValuesRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	return requestInfo, nil
}
