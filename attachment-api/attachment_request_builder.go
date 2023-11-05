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
	requestInfo, err := rB.RequestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return nil, err
	}

	errorMapping := core.ErrorMapping{"4XX": "hi"}

	response, err := rB.RequestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return nil, err
	}

	value, err := core.FromJson[AttachmentCollectionResponse](response)
	if err != nil {
		return nil, err
	}
	//value.parsePaginationHeaders(response.Header)

	return value, nil
}
