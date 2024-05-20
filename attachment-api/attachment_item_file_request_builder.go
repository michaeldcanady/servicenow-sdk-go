package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentItemFileRequestBuilderURLTemplate = "{+baseurl}/attachment{/sysId}/file"
)

type AttachmentItemFileRequestBuilder struct {
	intCore.RequestBuilder
}

func NewAttachmentItemFileRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentItemFileRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		attachmentItemFileRequestBuilderURLTemplate,
		pathParameters,
	)
	return &AttachmentItemFileRequestBuilder{
		requestBuilder,
	}
}

func (rB *AttachmentItemFileRequestBuilder) Get() {}
