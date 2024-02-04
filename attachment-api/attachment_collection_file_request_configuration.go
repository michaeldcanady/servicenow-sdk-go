package attachmentapi

import "github.com/RecoLabs/servicenow-sdk-go/core"

type AttachmentCollectionFileRequestConfiguration struct {
	Header          interface{}
	QueryParameters *AttachmentRequestBuilderFileQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *AttachmentItemResponse
}

func (rC *AttachmentCollectionFileRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
