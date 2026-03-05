package attachmentapi

import (
	"testing"
)

func TestAttachmentItemResponse_ParseHeaders(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Nil headers",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var attachmentItemResponse AttachmentItemResponse
			attachmentItemResponse.ParseHeaders(nil)
		})
	}
}
