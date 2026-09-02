// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

package cdmeditorapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NodesRequestBuilderGetQueryParameters represents query parameters for GET /nodes.
type NodesRequestBuilderGetQueryParameters struct {
	SysID    *string `uriparametername:"sys_id"`
	ParentID *string `uriparametername:"parent_id"`
	Type     *string `uriparametername:"type"`
}

// ValidationRequestBuilderGetQueryParameters represents query parameters for GET /validation.
type ValidationRequestBuilderGetQueryParameters struct {
	CdmID *string `uriparametername:"cdm_id"`
}

// NodesRequestBuilderGetRequestConfiguration represents the configuration for a Get request.
type NodesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[NodesRequestBuilderGetQueryParameters]

// NodesRequestBuilderPostRequestConfiguration represents the POST request configuration for the Nodes resource.
type NodesRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// NodeItemRequestBuilderPutRequestConfiguration represents the PUT request configuration for the Node Item resource.
type NodeItemRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// NodeItemRequestBuilderDeleteRequestConfiguration represents the DELETE request configuration for the Node Item resource.
type NodeItemRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]

// ValidationRequestBuilderGetRequestConfiguration represents the GET request configuration for the Validation resource.
type ValidationRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ValidationRequestBuilderGetQueryParameters]
