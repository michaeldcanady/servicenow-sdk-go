package cmdbinstanceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CmdbClassRequestBuilderGetQueryParameters ...
type CmdbClassRequestBuilderGetQueryParameters struct {
	Query  *string `uriparametername:"sysparm_query"`
	Limit  *int    `uriparametername:"sysparm_limit"`
	Offset *int    `uriparametername:"sysparm_offset"`
}

// CmdbClassRequestBuilderGetRequestConfiguration ...
type CmdbClassRequestBuilderGetRequestConfiguration struct {
	Headers         *abstractions.RequestHeaders
	Options         []abstractions.RequestOption
	QueryParameters *CmdbClassRequestBuilderGetQueryParameters
}

// CmdbClassRequestBuilderPostRequestConfiguration ...
type CmdbClassRequestBuilderPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    CmdbInstance
}
