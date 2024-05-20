package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentFileRequestBuilderURLTemplate = "{+baseurl}/attachment/file{?encryption_context,file_name,table_name,table_sys_id}"
)

type AttachmentFileRequestBuilder struct {
	intCore.RequestBuilder
}

func NewAttachmentFileRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentFileRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		attachmentFileRequestBuilderURLTemplate,
		pathParameters,
	)
	return &AttachmentFileRequestBuilder{
		requestBuilder,
	}
}

func (rB *AttachmentFileRequestBuilder) Post(data []byte, params *AttachmentRequestBuilderFileQueryParameters) (*AttachmentItemResponse, error) {
	if params == nil {
		return nil, ErrNilParams
	}

	config := &AttachmentCollectionFileRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            data,
		response:        &AttachmentItemResponse{},
	}

	err := rB.SendPost3(config.toConfiguration())
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
