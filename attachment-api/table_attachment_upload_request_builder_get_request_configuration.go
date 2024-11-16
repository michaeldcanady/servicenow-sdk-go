package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableAttachmentUploadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableAttachmentUploadRequestBuilderGetRequestConfiguration struct {
	//Headers Request headers
	Headers *abstractions.RequestHeaders
	//Options Request options
	Options []abstractions.RequestOption
}
