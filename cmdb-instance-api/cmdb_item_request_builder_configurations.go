package cmdbinstanceapi

import (
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// CmdbItemRequestBuilderGetRequestConfiguration ...
type CmdbItemRequestBuilderGetRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
}

// CmdbItemRequestBuilderPutRequestConfiguration ...
type CmdbItemRequestBuilderPutRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    CmdbInstance
}

// CmdbItemRequestBuilderPatchRequestConfiguration ...
type CmdbItemRequestBuilderPatchRequestConfiguration struct {
	Headers *abstractions.RequestHeaders
	Options []abstractions.RequestOption
	Data    CmdbInstance
}
