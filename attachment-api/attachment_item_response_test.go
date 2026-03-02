package attachmentapi

import (
	"testing"
)

func TestAttachmentItemResponse_ParseHeaders(t *testing.T) {
	var attachmentItemResponse AttachmentItemResponse
	attachmentItemResponse.ParseHeaders(nil)
}
