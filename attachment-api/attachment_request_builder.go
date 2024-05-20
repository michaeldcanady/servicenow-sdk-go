package attachmentapi

import (
	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

const (
	attachmentRequestBuilderURLTemplate = "{+baseurl}/attachment{/attachment}{?encryption_context,file_name,table_name,table_sys_id}"
)

type AttachmentRequestBuilder struct {
	core.RequestBuilder
}

func NewAttachmentRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		attachmentRequestBuilderURLTemplate,
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
	configuration := &AttachmentCollectionGetRequestConfiguration{
		Header:          nil,
		QueryParameters: params,
		Data:            nil,
		response:        &AttachmentCollectionResponse{},
	}

	err := rB.SendGet2(configuration.toConfiguration())
	if err != nil {
		return nil, err
	}

	return configuration.response, nil
}

// File ...
func (rB *AttachmentRequestBuilder) File(filePath string, params *AttachmentRequestBuilderFileQueryParameters) *AttachmentFileRequestBuilder {
}
