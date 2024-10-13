package attachmentapi

import (
	"context"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentItemFileURLTemplate = "{+baseurl}/attachment{/attachment}/file"
)

type AttachmentItemFileRequestBuilder interface {
	Get(ctx context.Context, filePath string) error
}

type attachmentItemFileRequestBuilder struct {
	intCore.Sendable
}

func NewAttachmentItemFileRequestBuilder(client intCore.ClientSendable, pathParameters map[string]string) AttachmentItemFileRequestBuilder {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		attachmentItemFileURLTemplate,
		pathParameters,
	)
	return &attachmentItemFileRequestBuilder{
		requestBuilder,
	}
}

func (rB *attachmentItemFileRequestBuilder) Get(ctx context.Context, filePath string) error {
	// TODO: Finish implementation
}
