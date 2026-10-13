// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const sharedLibrariesComponentsURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_libraries/components"

// SharedLibrariesComponentsRequestBuilder provides operations to manage shared library components.
type SharedLibrariesComponentsRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesComponentsRequestBuilderInternal instantiates a new [SharedLibrariesComponentsRequestBuilder].
func NewSharedLibrariesComponentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesComponentsRequestBuilder {
	return &SharedLibrariesComponentsRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedLibrariesComponentsURLTemplate, pathParameters),
	}
}

// Applications returns a [SharedLibrariesComponentsApplicationsRequestBuilder].
func (rB *SharedLibrariesComponentsRequestBuilder) Applications() *SharedLibrariesComponentsApplicationsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewSharedLibrariesComponentsApplicationsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
