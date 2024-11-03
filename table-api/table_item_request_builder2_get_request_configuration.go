package tableapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableItemRequestBuilder2GetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableItemRequestBuilder2GetRequestConfiguration struct {
	// Request headers
	Headers *abstractions.RequestHeaders
	// Request options
	Options []abstractions.RequestOption
	// Request query parameters
	QueryParameters *TableItemRequestBuilder2GetQueryParameters
}
