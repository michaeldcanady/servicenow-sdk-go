package cmdbinstanceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CmdbClassRequestBuilderGetQueryParameters ...
type CmdbClassRequestBuilderGetQueryParameters struct {
	Query  *string `uriparametername:"sysparm_query"`
	Limit  *int32  `uriparametername:"sysparm_limit"`
	Offset *int32  `uriparametername:"sysparm_offset"`
}

// CmdbClassRequestBuilderGetRequestConfiguration ...
type CmdbClassRequestBuilderGetRequestConfiguration = abstractions.RequestConfiguration[CmdbClassRequestBuilderGetQueryParameters]

// CmdbClassRequestBuilderPostRequestConfiguration ...
type CmdbClassRequestBuilderPostRequestConfiguration = abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
