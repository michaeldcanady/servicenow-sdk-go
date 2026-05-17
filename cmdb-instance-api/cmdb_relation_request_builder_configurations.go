package cmdbinstanceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CmdbRelationRequestBuilderPostRequestConfiguration ...
type CmdbRelationRequestBuilderPostRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    CmdbInstance
}

// CmdbRelationItemRequestBuilderDeleteRequestConfiguration ...
type CmdbRelationItemRequestBuilderDeleteRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
}
