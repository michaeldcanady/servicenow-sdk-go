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

// Get sends an HTTP GET request using the specified query parameters and returns a AttachmentCollectionResponse.
//
// Parameters:
//   - params: An instance of AttachmentRequestBuilderGetQueryParameters to include in the GET request.
//
// Returns:
//   - *AttachmentCollectionResponse: The response data as a AttachmentCollectionResponse.
//   - error: An error if there was an issue with the request or response.
func (rB *AttachmentRequestBuilder) Get(params *AttachmentRequestBuilderGetQueryParameters) (*AttachmentCollectionResponse, error) {
	response, err := core.SendGet[*AttachmentCollectionResponse](&rB.RequestBuilder, params, nil)

	return (*response), err
}
