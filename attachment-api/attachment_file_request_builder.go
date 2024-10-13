package attachmentapi

import (
	"context"
	"errors"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentFileURLTemplate = "{+baseurl}/attachment/file{?encryption_context,file_name,table_name,table_sys_id}"
)

type AttachmentFileRequestBuilder interface {
	// TODO: Create type for headers and query parameters
	Post(ctx context.Context, filePath string, headers *AttachmentFileRequestBuilderPostHeaders, params *AttachmentFileRequestBuilderPostQueryParameter) (*AttachmentItemResponse, error)
}

type attachmentFileRequestBuilder struct {
	intCore.Sendable2
}

func NewAttachmentFileRequestBuilder(client intCore.ClientSendable, pathParameters map[string]string) AttachmentFileRequestBuilder {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		attachmentURLTemplate,
		pathParameters,
	)
	return &attachmentFileRequestBuilder{
		requestBuilder,
	}
}

func (rB *attachmentFileRequestBuilder) Post(ctx context.Context, filePath string, headers *AttachmentFileRequestBuilderPostHeaders, params *AttachmentFileRequestBuilderPostQueryParameter) (*AttachmentItemResponse, error) {
	if intCore.IsNil(params) {
		return nil, errors.New("params is nil")
	}

	if params.FileName == "" {
		// TODO: get file name from filePath
	}

	if params.TableName == "" {
		return nil, errors.New("params.TableName is nil")
	}

	if params.TableSysID == "" {
		return nil, errors.New("params.TableSysID is nil")
	}

	if intCore.IsNil(headers) {
		// TODO: builder headers
	}

	resp, err := rB.Send2(ctx, intCore.MethodPost, intCore.WithHeader(headers), intCore.WithQueryParameters(params), intCore.WithResponse(&AttachmentItemResponse{}))
	if err != nil {
		return nil, err
	}

	return resp.(*AttachmentItemResponse), nil
}
