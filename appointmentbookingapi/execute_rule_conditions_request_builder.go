// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appointmentbookingapi

import (
	"context"
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	internalhttp "github.com/michaeldcanady/servicenow-sdk-go/v2/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appointmentExecuteRuleConditionsBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/execute_rule_conditions"

// ExecuteRuleConditionsRequestBuilder provides operations to manage execute rule conditions.
type ExecuteRuleConditionsRequestBuilder struct {
	core.RequestBuilder
}

// NewExecuteRuleConditionsRequestBuilderInternal instantiates a new [ExecuteRuleConditionsRequestBuilder].
func NewExecuteRuleConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExecuteRuleConditionsRequestBuilder {
	return &ExecuteRuleConditionsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentExecuteRuleConditionsBookingURLTemplate, pathParameters),
	}
}

// NewExecuteRuleConditionsRequestBuilder instantiates a new [ExecuteRuleConditionsRequestBuilder] with the provided base URL
// and request adapter.
func NewExecuteRuleConditionsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ExecuteRuleConditionsRequestBuilder {
	return NewExecuteRuleConditionsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to execute rule conditions.
func (rB *ExecuteRuleConditionsRequestBuilder) Post(ctx context.Context, body ExecuteRuleConditionsRequest, config *ExecuteRuleConditionsRequestBuilderPostRequestConfiguration) (ExecuteRuleConditionsResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExecuteRuleConditionsResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(ExecuteRuleConditionsResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*ExecuteRuleConditionsResponse)(nil))
	}

	return typedResp, nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *ExecuteRuleConditionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ExecuteRuleConditionsRequest, config *ExecuteRuleConditionsRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	// TODO: check for nil/empty template and path parameters
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return requestInfo, nil
}
