package tableapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableRequestBuilderKiotaGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableItemRequestBuilder2PutRequestConfiguration struct {
	// Headers Request headers
	Headers *abstractions.RequestHeaders
	// Options Request options
	Options []abstractions.RequestOption
	// QueryParameters Request query parameters
	QueryParameters *TableItemRequestBuilder2PutQueryParameters
}
