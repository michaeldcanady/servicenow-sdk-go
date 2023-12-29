package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

type AttachmentCollectionFileRequestConfiguration2 struct {
	header   interface{}
	query    *AttachmentRequestBuilderFileQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *AttachmentItemResponse
}

func (rC *AttachmentCollectionFileRequestConfiguration2) Header() interface{} {
	return rC.header
}

func (rC *AttachmentCollectionFileRequestConfiguration2) Query() interface{} {
	return rC.query
}

func (rC *AttachmentCollectionFileRequestConfiguration2) Data() interface{} {
	return rC.data
}

func (rC *AttachmentCollectionFileRequestConfiguration2) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *AttachmentCollectionFileRequestConfiguration2) Response() core.Response {
	return rC.response
}
