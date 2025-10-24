package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v1.8.0.
//
// AttachmentCollectionFileRequestConfiguration ...
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
