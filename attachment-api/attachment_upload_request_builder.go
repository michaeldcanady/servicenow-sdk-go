package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentUploadRequestBuilderURLTemplate = "{+baseurl}/attachment/upload"
)

type AttachmentUploadRequestBuilder struct {
	intCore.RequestBuilder
}

func NewAttachmentUploadRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentUploadRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		attachmentUploadRequestBuilderURLTemplate,
		pathParameters,
	)
	return &AttachmentUploadRequestBuilder{
		requestBuilder,
	}
}

func (rB *AttachmentUploadRequestBuilder) Post() {}
