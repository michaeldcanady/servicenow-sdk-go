package attachmentapi

import "net/http"

// Deprecated: deprecated since v1.8.0.
//
// AttachmentItemResponse ...
type AttachmentItemResponse struct {
	Result *Attachment
}

func (cR *AttachmentItemResponse) ParseHeaders(headers http.Header) {
	//No headers to parse but needed for Response
}
