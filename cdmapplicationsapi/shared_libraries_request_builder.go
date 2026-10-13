// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const sharedLibrariesURLTemplate = "{+baseurl}/api/sn_cdm/applications/shared_libraries"

// SharedLibrariesRequestBuilder provides operations to manage shared libraries.
type SharedLibrariesRequestBuilder struct {
	core.RequestBuilder
}

// NewSharedLibrariesRequestBuilderInternal instantiates a new [SharedLibrariesRequestBuilder].
func NewSharedLibrariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *SharedLibrariesRequestBuilder {
	return &SharedLibrariesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, sharedLibrariesURLTemplate, pathParameters),
	}
}

// Components returns a [SharedLibrariesComponentsRequestBuilder].
func (rB *SharedLibrariesRequestBuilder) Components() *SharedLibrariesComponentsRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewSharedLibrariesComponentsRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
