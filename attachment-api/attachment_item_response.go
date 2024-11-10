package attachmentapi

import "net/http"

type AttachmentItemResponse struct {
	Result *Attachment
}

func (cR *AttachmentItemResponse) ParseHeaders(headers http.Header) {
	//No headers to parse but needed for Response
}
