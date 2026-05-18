package cmdbinstanceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CmdbClassRequestBuilderGetQueryParameters ...
type CmdbClassRequestBuilderGetQueryParameters struct {
	Query  *string `url:"sysparm_query,omitempty"`
	Limit  *int    `url:"sysparm_limit,omitempty"`
	Offset *int    `url:"sysparm_offset,omitempty"`
}

// CmdbClassRequestBuilderGetRequestConfiguration ...
type CmdbClassRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CmdbClassRequestBuilderGetQueryParameters]

// CmdbClassRequestBuilderPostRequestConfiguration ...
type CmdbClassRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
