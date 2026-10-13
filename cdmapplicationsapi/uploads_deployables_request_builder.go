// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmapplicationsapi

import (
	"maps"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const uploadsDeployablesURLTemplate = "{+baseurl}/api/sn_cdm/applications/uploads/deployables"

// UploadsDeployablesRequestBuilder provides operations to manage deployables uploads.
type UploadsDeployablesRequestBuilder struct {
	core.RequestBuilder
}

// NewUploadsDeployablesRequestBuilderInternal instantiates a new [UploadsDeployablesRequestBuilder].
func NewUploadsDeployablesRequestBuilderInternal(pathParameters map[string]string, requestAdapter abstractions.RequestAdapter) *UploadsDeployablesRequestBuilder {
	return &UploadsDeployablesRequestBuilder{
		RequestBuilder: core.NewBaseRequestBuilder(requestAdapter, uploadsDeployablesURLTemplate, pathParameters),
	}
}

// File returns a [UploadsDeployablesFileRequestBuilder].
func (rB *UploadsDeployablesRequestBuilder) File() *UploadsDeployablesFileRequestBuilder {
	if conversion.IsNil(rB) || conversion.IsNil(rB.RequestBuilder) {
		return nil
	}
	return NewUploadsDeployablesFileRequestBuilderInternal(maps.Clone(rB.GetPathParameters()), rB.GetRequestAdapter())
}
