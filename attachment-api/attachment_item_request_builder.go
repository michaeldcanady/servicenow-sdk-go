package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentItemRequestBuilderURLTemplate = "{+baseurl}/attachment{/sysId}"
)

type AttachmentItemRequestBuilder struct {
	intCore.RequestBuilder
}

func NewAttachmentItemRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentItemRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		attachmentItemRequestBuilderURLTemplate,
		pathParameters,
	)
	return &AttachmentItemRequestBuilder{
		requestBuilder,
	}
}

func (rB *AttachmentItemRequestBuilder) Delete() error {}

func (rB *AttachmentItemRequestBuilder) Get() AttachmentItemResponse {}

func (rB *AttachmentItemRequestBuilder) File() *AttachmentItemFileRequestBuilder {}
