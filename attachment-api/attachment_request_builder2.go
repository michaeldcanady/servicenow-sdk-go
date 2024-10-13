package attachmentapi

import (
	"context"

	intCore "github.com/michaeldcanady/servicenow-sdk-go/internal/core"
)

const (
	attachmentURLTemplate = "{+baseurl}/attachment{?sysparm_limit,sysparm_offset,sysparm_query}"
)

type AttachmentRequestBuilder2 interface {
	Get(ctx context.Context, params *AttachmentRequestBuilderGetQueryParameters) (*AttachmentCollectionResponse, error)
	ByID(id string) AttachmentItemRequestBuilder
	File() AttachmentFileRequestBuilder
}

type attachmentRequestBuilder2 struct {
	intCore.RequestBuilder2
}

func NewAttachmentRequestBuilder2(client intCore.ClientSendable, pathParameters map[string]string) AttachmentRequestBuilder2 {
	requestBuilder := intCore.NewRequestBuilder2(
		client,
		attachmentURLTemplate,
		pathParameters,
	)
	return &attachmentRequestBuilder2{
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
func (rB *attachmentRequestBuilder2) Get(ctx context.Context, params *AttachmentRequestBuilderGetQueryParameters) (*AttachmentCollectionResponse, error) {
	resp, err := rB.Send2(ctx, intCore.MethodGet, intCore.WithQueryParameters(params), intCore.WithResponse(&AttachmentCollectionResponse{}))
	if err != nil {
		return nil, err
	}

	return resp.(*AttachmentCollectionResponse), nil
}

func (rB *attachmentRequestBuilder2) ByID(id string) AttachmentItemRequestBuilder {
	params := rB.GetPathParameters()
	params["attachment"] = id

	return NewAttachmentItemRequestBuilder(rB.GetClient(), params)
}

func (rB *attachmentRequestBuilder2) File() AttachmentFileRequestBuilder {
	return NewAttachmentFileRequestBuilder(rB.GetClient(), rB.GetPathParameters())
}

// TODO: Create AttachmentUploadRequestBuilder: https://docs.servicenow.com/bundle/xanadu-api-reference/page/integrate/inbound-rest/concept/c_AttachmentAPI.html#title_attachment-POST-upload
func (rB *attachmentRequestBuilder2) Upload() AttachmentUploadRequestBuilder {
	return NewAttachmentUploadRequestBuilder(rB.GetClient(), rB.GetPathParameters())
}
