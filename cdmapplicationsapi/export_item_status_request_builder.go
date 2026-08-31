// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

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

const exportItemStatusURLTemplate = "{+baseurl}/api/sn_cdm/applications/deployables/exports/{export_id}/status"

// ExportItemStatusRequestBuilder provides operations to check export status.
type ExportItemStatusRequestBuilder struct {
	core.RequestBuilder
}

// NewExportItemStatusRequestBuilderInternal instantiates a new [ExportItemStatusRequestBuilder].
func NewExportItemStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ExportItemStatusRequestBuilder {
	return &ExportItemStatusRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, exportItemStatusURLTemplate, pathParameters),
	}
}

// NewExportItemStatusRequestBuilder instantiates a new [ExportItemStatusRequestBuilder].
func NewExportItemStatusRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ExportItemStatusRequestBuilder {
	return NewExportItemStatusRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Get gets the status of a specific export.
func (rB *ExportItemStatusRequestBuilder) Get(ctx context.Context, config *ExportItemStatusRequestBuilderGetRequestConfiguration) (ExportStatusResponse, error) {
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

	res, err := rB.GetRequestAdapter().Send(ctx, requestInfo, CreateExportStatusResponseFromDiscriminatorValue, core.DefaultErrorMapping())
	if err != nil {
		return nil, err
	}
	if conversion.IsNil(res) {
		return nil, snerrors.ErrNilResponse
	}
	typedRes, ok := res.(ExportStatusResponse)
	if !ok {
		return nil, fmt.Errorf("res is not %T", (*ExportStatusResponse)(nil))
	}
	return typedRes, nil
}

// ToGetRequestInformation builds the request information for the Get method.
func (rB *ExportItemStatusRequestBuilder) ToGetRequestInformation(_ context.Context, config *ExportItemStatusRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	if !conversion.IsNil(config) {
		if config.Headers != nil {
			requestInfo.Headers.AddAll(config.Headers)
		}
		if config.Options != nil {
			requestInfo.AddRequestOptions(config.Options)
		}
	}
	requestInfo.Headers.TryAdd(internalhttp.RequestHeaderAccept.String(), internalhttp.ContentTypeApplicationJSON.String())
	return requestInfo, nil
}
