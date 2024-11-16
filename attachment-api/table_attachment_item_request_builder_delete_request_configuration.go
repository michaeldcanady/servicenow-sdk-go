package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// TableAttachmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TableAttachmentItemRequestBuilderDeleteRequestConfiguration struct {
	//Headers Request headers
	Headers *abstractions.RequestHeaders
	//Options Request options
	Options []abstractions.RequestOption
}
