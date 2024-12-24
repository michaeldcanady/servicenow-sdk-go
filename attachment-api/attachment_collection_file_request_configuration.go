package attachmentapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{unreleased}.
//
// AttachmentCollectionFileRequestConfiguration ...
type AttachmentCollectionFileRequestConfiguration struct {
	Header          interface{}
	QueryParameters *AttachmentRequestBuilderFileQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping //nolint: staticcheck
	response        *AttachmentItemResponse
}

func (rC *AttachmentCollectionFileRequestConfiguration) toConfiguration() *core.RequestConfiguration { //nolint: staticcheck
	return &core.RequestConfiguration{ //nolint: staticcheck
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
