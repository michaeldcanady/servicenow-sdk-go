// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package appserviceapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const csdmAppServiceItemURLTemplate = "{+baseurl}/api/now/cmdb/csdm/app_service/{sys_id}"

// CsdmAppServiceItemRequestBuilder provides operations for a specific CSDM application service.
type CsdmAppServiceItemRequestBuilder struct {
	core.RequestBuilder
}

// NewCsdmAppServiceItemRequestBuilderInternal instantiates a new [CsdmAppServiceItemRequestBuilder].
func NewCsdmAppServiceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *CsdmAppServiceItemRequestBuilder {
	return &CsdmAppServiceItemRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, csdmAppServiceItemURLTemplate, pathParameters),
	}
}

// PopulateService returns a [PopulateServiceRequestBuilder].
func (rB *CsdmAppServiceItemRequestBuilder) PopulateService() *PopulateServiceRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewPopulateServiceRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}

// ServiceDetails returns a [ServiceDetailsRequestBuilder].
func (rB *CsdmAppServiceItemRequestBuilder) ServiceDetails() *ServiceDetailsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}

	return NewServiceDetailsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
