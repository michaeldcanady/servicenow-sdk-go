package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{unreleased}.
//
// AttachmentCollectionGetRequestConfiguration ...
type AttachmentCollectionGetRequestConfiguration struct {
	Header          interface{}
	QueryParameters *AttachmentRequestBuilderGetQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping //nolint: staticcheck
	response        *AttachmentCollectionResponse
}

func (rC *AttachmentCollectionGetRequestConfiguration) toConfiguration() *core.RequestConfiguration { //nolint: staticcheck
	return &core.RequestConfiguration{ //nolint: staticcheck
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
