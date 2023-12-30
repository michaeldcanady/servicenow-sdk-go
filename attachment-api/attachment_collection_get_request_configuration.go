package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

// Deprecated: deprecated since v{version}. Use `AttachmentCollectionGetRequestConfiguration2` instead.
type AttachmentCollectionGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *AttachmentRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *AttachmentCollectionResponse //nolint:unused
}

//nolint:unused
func (rC *AttachmentCollectionGetRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
