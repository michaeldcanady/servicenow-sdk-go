package attachmentapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

type AttachmentCollectionGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *AttachmentRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *AttachmentCollectionResponse
}

func (rC *AttachmentCollectionGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
