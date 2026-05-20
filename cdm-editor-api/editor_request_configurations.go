package cdmeditorapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NodesRequestBuilderGetQueryParameters represents query parameters for GET /nodes.
type NodesRequestBuilderGetQueryParameters struct {
	SysId    *string `url:"sys_id,omitempty"`
	ParentId *string `url:"parent_id,omitempty"`
	Type     *string `url:"type,omitempty"`
}

// ValidationRequestBuilderGetQueryParameters represents query parameters for GET /validation.
type ValidationRequestBuilderGetQueryParameters struct {
	CdmId *string `url:"cdm_id,omitempty"`
}

// Request Configurations
type NodesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[NodesRequestBuilderGetQueryParameters]
type NodesRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type NodeItemRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type NodeItemRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type ValidationRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ValidationRequestBuilderGetQueryParameters]
