package attachmentapi

import (
	"context"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentUploadURLTemplate = "{+baseurl}/attachment/upload"
)

type AttachementUploadRequestBuilder interface {
	Post(ctx context.Context, contentType, tableName, tableSysId string) (AttachmentItemResponse, error)
}

type attachmentUploadRequestBuilder struct {
	intCore.Sendable2
}

func NewAttachmentUploadRequestBuilder(client intCore.ClientSendable, pathParameters map[string]string) AttachementUploadRequestBuilder {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		attachmentUploadURLTemplate,
		pathParameters,
	)
	return &attachmentUploadRequestBuilder{
		requestBuilder,
	}
}

func (rB *attachmentUploadRequestBuilder) Post(ctx context.Context, contentType, tableName, tableSysId string) (AttachmentItemResponse, error) {

}
