package attachmentapi

import (
	"context"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentItemURLTemplate = "{+baseurl}/attachment{/attachment}{?encryption_context,file_name,table_name,table_sys_id}"
)

type AttachmentItemRequestBuilder interface {
	Get(ctx context.Context) (*AttachmentItemResponse, error)
	File() AttachmentItemFileRequestBuilder
}

type attachmentItemRequestBuilder struct {
	intCore.RequestBuilder2
}

func NewAttachmentItemRequestBuilder(client intCore.ClientSendable, pathParameters map[string]string) AttachmentItemRequestBuilder {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		attachmentItemURLTemplate,
		pathParameters,
	)
	return &attachmentItemRequestBuilder{
		requestBuilder,
	}
}

// Get sends an HTTP GET request using the specified query parameters and returns a AttachmentCollectionResponse.
//
// Parameters:
//   - params: An instance of AttachmentRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *AttachmentCollectionResponse: The response data as a AttachmentCollectionResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *attachmentItemRequestBuilder) Get(ctx context.Context) (*AttachmentItemResponse, error) {
	resp, err := rB.Send2(ctx, intCore.MethodGet, intCore.WithResponse(&AttachmentItemResponse{}))
	if err != nil {
		return nil, err
	}

	return resp.(*AttachmentItemResponse), nil
}

// File ...
func (rB *attachmentItemRequestBuilder) File() AttachmentItemFileRequestBuilder {
	return NewAttachmentItemFileRequestBuilder(rB.GetClient(), rB.GetPathParameters())
}
