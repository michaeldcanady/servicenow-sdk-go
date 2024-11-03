package tableapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableRequestBuilder2GetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableRequestBuilder2GetRequestConfiguration struct {
	// Request headers
	Headers *abstractions.RequestHeaders
	// Request options
	Options []abstractions.RequestOption
	// Request query parameters
	QueryParameters *TableRequestBuilder2GetQueryParameters
}
