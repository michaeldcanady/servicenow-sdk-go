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

const updateDynamicNumberOfLevelsURLTemplate = "{+baseurl}/api/now/cmdb/app_service/updateDynamicNumberOfLevels"

// UpdateDynamicNumberOfLevelsRequestBuilder provides operations to update the number of levels of a dynamic application service.
type UpdateDynamicNumberOfLevelsRequestBuilder struct {
	*servicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse]
}

// NewUpdateDynamicNumberOfLevelsRequestBuilderInternal instantiates a new UpdateDynamicNumberOfLevelsRequestBuilder.
func NewUpdateDynamicNumberOfLevelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UpdateDynamicNumberOfLevelsRequestBuilder {
	return &UpdateDynamicNumberOfLevelsRequestBuilder{
		newServicePostRequestBuilder[*AppServiceActionRequest, AppServiceActionResponse](requestAdapter, updateDynamicNumberOfLevelsURLTemplate, pathParameters, CreateAppServiceActionResponseFromDiscriminatorValue),
	}
}

// NewUpdateDynamicNumberOfLevelsRequestBuilder instantiates a new [UpdateDynamicNumberOfLevelsRequestBuilder].
func NewUpdateDynamicNumberOfLevelsRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *UpdateDynamicNumberOfLevelsRequestBuilder {
	return NewUpdateDynamicNumberOfLevelsRequestBuilderInternal(map[string]string{internal.RawURLKey: rawURL}, requestAdapter)
}

// Post sends a POST request to update the number of levels of a dynamic application service.
func (rB *UpdateDynamicNumberOfLevelsRequestBuilder) Post(ctx context.Context, body *AppServiceActionRequest, config *UpdateDynamicNumberOfLevelsRequestConfiguration) (AppServiceActionResponse, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.post(ctx, body, config)
}

// ToPostRequestInformation creates a RequestInformation object for a POST request.
func (rB *UpdateDynamicNumberOfLevelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body *AppServiceActionRequest, config *UpdateDynamicNumberOfLevelsRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toPostRequestInformation(ctx, body, config)
}
