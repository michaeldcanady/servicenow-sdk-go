package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableAttachmentFileRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableAttachmentFileRequestBuilderPostRequestConfiguration struct {
	// Headers Request headers
	Headers *abstractions.RequestHeaders
	// Options Request options
	Options []abstractions.RequestOption
	// QueryParameters request query parameters
	QueryParameters *AttachmentFileRequestBuilderPostQueryParameters
}
