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

const appointmentAvailabilityBookingURLTemplate = "{+baseurl}/api/sn_apptmnt_booking/v1/appointment/availability"

// AvailabilityRequestBuilder provides operations to manage availability.
type AvailabilityRequestBuilder struct {
	core.RequestBuilder
}

// NewAvailabilityRequestBuilderInternal instantiates a new AvailabilityRequestBuilder.
func NewAvailabilityRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AvailabilityRequestBuilder {
	return &AvailabilityRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appointmentAvailabilityBookingURLTemplate, pathParameters),
	}
}

// NewAvailabilityRequestBuilder instantiates a new [AvailabilityRequestBuilder] with the provided base URL
// and request adapter.
func NewAvailabilityRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AvailabilityRequestBuilder {
	return NewAvailabilityRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to get availability.
func (rB *AvailabilityRequestBuilder) Post(ctx context.Context, body AvailabilityRequest, config *AvailabilityRequestBuilderPostRequestConfiguration) (AvailabilityResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateAvailabilityResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}

	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}

	typedResp, ok := res.(AvailabilityResponse)
	if !ok {
		return nil, fmt.Errorf("resp is not %T", (*AvailabilityResponse)(nil))
	}

	return typedResp, nil
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *AvailabilityRequestBuilder) ToPostRequestInformation(ctx context.Context, body AvailabilityRequest, config *AvailabilityRequestBuilderPostRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil, snerrors.ErrNilRequestBuilder
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.POST, rB.GetURLTemplate(), rB.GetPathParameters())
	abstractions.ConfigureRequestInformation(requestInfo, config)

	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())

	err := requestInfo.SetContentFromParsable(ctx, rB.GetRequestAdapter(), internalhttp.ContentTypeApplicationJSON.String(), body)
	if err != nil {
		return nil, err
	}

	return requestInfo, nil
}
