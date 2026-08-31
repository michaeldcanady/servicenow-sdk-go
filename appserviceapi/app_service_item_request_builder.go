// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const appServiceItemURLTemplate = "{+baseurl}/api/now/cmdb/app_service/{sys_id}"

// AppServiceItemRequestBuilder provides operations for a specific application service.
type AppServiceItemRequestBuilder struct {
	core.RequestBuilder
}

// NewAppServiceItemRequestBuilderInternal instantiates a new [AppServiceItemRequestBuilder].
func NewAppServiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *AppServiceItemRequestBuilder {
	return &AppServiceItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, appServiceItemURLTemplate, pathParameters),
	}
}

// GetContent returns a [GetContentRequestBuilder].
func (rB *AppServiceItemRequestBuilder) GetContent() *GetContentRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewGetContentRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
