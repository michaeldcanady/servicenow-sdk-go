package attachmentapi

import abstractions "github.com/microsoft/kiota-abstractions-go"

// AttachmentUploadRequestBuilderPostRequestConfiguration request configurations for attachment upload post request
type AttachmentUploadRequestBuilderPostRequestConfiguration struct {
	// Headers Request headers
	Headers *abstractions.RequestHeaders
	// Options Request options
	Options []abstractions.RequestOption
}
