package tableapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableRequestBuilderKiotaGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableItemRequestBuilder2DeleteRequestConfiguration struct {
	// Request headers
	Headers *abstractions.RequestHeaders
	// Request options
	Options []abstractions.RequestOption
	// Request query parameters
	QueryParameters *TableItemRequestBuilder2DeleteQueryParameters
}
