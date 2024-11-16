package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableAttachmentItemFileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableAttachmentItemFileRequestBuilderGetRequestConfiguration struct {
	//Headers Request headers
	Headers *abstractions.RequestHeaders
	//Options Request options
	Options []abstractions.RequestOption
}
