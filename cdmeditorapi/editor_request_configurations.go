package cdmeditorapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// NodesRequestBuilderGetQueryParameters represents query parameters for GET /nodes.
type NodesRequestBuilderGetQueryParameters struct {
	SysId    *string `uriparametername:"sys_id"`
	ParentId *string `uriparametername:"parent_id"`
	Type     *string `uriparametername:"type"`
}

// ValidationRequestBuilderGetQueryParameters represents query parameters for GET /validation.
type ValidationRequestBuilderGetQueryParameters struct {
	CdmId *string `uriparametername:"cdm_id"`
}

// Request Configurations
type NodesRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[NodesRequestBuilderGetQueryParameters]
type NodesRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type NodeItemRequestBuilderPutRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type NodeItemRequestBuilderDeleteRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
type ValidationRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[ValidationRequestBuilderGetQueryParameters]
