package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// Deprecated: deprecated since v{version}. Use `AttachmentCollectionFileRequestConfiguration2` instead.
type AttachmentCollectionFileRequestConfiguration struct {
	Header          interface{}
	QueryParameters *AttachmentRequestBuilderFileQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *AttachmentItemResponse //nolint:unused
}

//nolint:unused
func (rC *AttachmentCollectionFileRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
