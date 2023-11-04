package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

type AttachmentRequestBuilder struct {
	core.RequestBuilder
}

func NewAttachmentRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/attachment{/attachment}",
		pathParameters,
	)
	return &AttachmentRequestBuilder{
		*requestBuilder,
	}
}

func (rB *AttachmentRequestBuilder) Get() {

}
