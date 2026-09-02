// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"context"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const convertToDynamicServiceURLTemplate = "{+baseurl}/api/now/cmdb/app_service/convertToDynamicService"

// ConvertToDynamicServiceRequestBuilder provides operations to convert an application service to a dynamic service.
type ConvertToDynamicServiceRequestBuilder struct {
	*servicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse]
}

// NewConvertToDynamicServiceRequestBuilderInternal instantiates a new ConvertToDynamicServiceRequestBuilder.
func NewConvertToDynamicServiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *ConvertToDynamicServiceRequestBuilder {
	return &ConvertToDynamicServiceRequestBuilder{
		newServicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse](requestAdapter, convertToDynamicServiceURLTemplate, pathParameters, CreateAppServiceActionResponseFromDiscriminatorValue),
	}
}

// NewConvertToDynamicServiceRequestBuilder instantiates a new [ConvertToDynamicServiceRequestBuilder].
func NewConvertToDynamicServiceRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *ConvertToDynamicServiceRequestBuilder {
	return NewConvertToDynamicServiceRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to convert an application service to a dynamic service.
func (rB *ConvertToDynamicServiceRequestBuilder) Post(ctx context.Context, body *AppServiceActionRequest, config *ConvertToDynamicServiceRequestConfiguration) (AppServiceActionResponse, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *ConvertToDynamicServiceRequestBuilder) ToPostRequestInformation(ctx context.Context, body *AppServiceActionRequest, config *ConvertToDynamicServiceRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
