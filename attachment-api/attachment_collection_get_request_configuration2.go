package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

type AttachmentCollectionGetRequestConfiguration2 struct {
	header   interface{}
	query    *AttachmentRequestBuilderGetQueryParameters
	data     interface{}
	mapping  core.ErrorMapping
	response *AttachmentCollectionResponse
}

func (rC *AttachmentCollectionGetRequestConfiguration2) Header() interface{} {
	return rC.header
}

func (rC *AttachmentCollectionGetRequestConfiguration2) Query() interface{} {
	return rC.query
}

func (rC *AttachmentCollectionGetRequestConfiguration2) Data() interface{} {
	return rC.data
}

func (rC *AttachmentCollectionGetRequestConfiguration2) Mapping() core.ErrorMapping {
	return rC.mapping
}

func (rC *AttachmentCollectionGetRequestConfiguration2) Response() core.Response {
	return rC.response
}
