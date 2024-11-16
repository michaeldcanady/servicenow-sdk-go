package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableAttachmentRequestBuilder2GetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableAttachmentRequestBuilder2GetRequestConfiguration struct {
	//Headers Request headers
	Headers *abstractions.RequestHeaders
	//Options Request options
	Options []abstractions.RequestOption
	//QueryParameters Request query parameters
	QueryParameters *AttachmentRequestBuilder2GetQueryParameters
}
