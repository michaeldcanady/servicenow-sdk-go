package attachmentapi

import (
	"os"
	"path/filepath"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

type AttachmentRequestBuilder struct {
	core.RequestBuilder
}

func NewAttachmentRequestBuilder(client core.Client, pathParameters map[string]string) *AttachmentRequestBuilder {
	requestBuilder := core.NewRequestBuilder(
		client,
		"{+baseurl}/attachment{/attachment}{?encryption_context,file_name,table_name,table_sys_id}",
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
	configuration := &AttachmentCollectionGetRequestConfiguration2{
		header:   nil,
		query:    params,
		data:     nil,
		response: &AttachmentCollectionResponse{},
	}

	err := rB.SendGet3(configuration)
	if err != nil {
		return nil, err
	}

	return configuration.response, nil
}

// File Uploads the provided file to ServiceNow.
func (rB *AttachmentRequestBuilder) File(filePath string, params *AttachmentRequestBuilderFileQueryParameters) (*AttachmentItemResponse, error) {
	if params == nil {
		return nil, ErrNilParams
	}

	cleanPath := filepath.Clean(filePath)

	_, err := os.Stat(cleanPath)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, err
	}

	config := &AttachmentCollectionFileRequestConfiguration2{
		header:   nil,
		query:    params,
		data:     data,
		response: &AttachmentItemResponse{},
	}

	err = rB.SendPost4(config)
	if err != nil {
		return nil, err
	}

	return config.response, nil
}
