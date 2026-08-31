// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package documentsapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	versionStateURLTemplate = "{+baseurl}/api/now/v1/documents/versionstate/{version_sys_id}"
)

// VersionStateRequestBuilder provides operations to manage the versionstate endpoint.
type VersionStateRequestBuilder struct {
	*documentGetRequestBuilder
}

// NewVersionStateRequestBuilderInternal instantiates a new VersionStateRequestBuilder.
func NewVersionStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *VersionStateRequestBuilder {
	return &VersionStateRequestBuilder{
		newDocumentGetRequestBuilder(requestAdapter, versionStateURLTemplate, pathParameters),
	}
}

// Get retrieves the state of the specified document version.
func (rB *VersionStateRequestBuilder) Get(ctx context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*core.BaseServiceNowItemResponse[Document], error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.get(ctx, (*documentGetRequestConfiguration)(requestConfiguration))
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *VersionStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VersionStateRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if conversion.IsNil(rB) {
		return nil, snerrors.ErrNilRequestBuilder
	}
	return rB.toGetRequestInformation(ctx, (*documentGetRequestConfiguration)(requestConfiguration))
}
