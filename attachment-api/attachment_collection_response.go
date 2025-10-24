package attachmentapi

import "net/http"

// Deprecated: deprecated since v1.8.0.
//
// AttachmentCollectionResponse ...
type AttachmentCollectionResponse struct {
	Result []*Attachment
}

func (cR *AttachmentCollectionResponse) ParseHeaders(headers http.Header) {
	//No headers to parse but needed for Response
}
