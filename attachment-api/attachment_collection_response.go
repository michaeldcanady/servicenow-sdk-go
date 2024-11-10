package attachmentapi

import "net/http"

type AttachmentCollectionResponse struct {
	Result []*Attachment
}

func (cR *AttachmentCollectionResponse) ParseHeaders(headers http.Header) {
	//No headers to parse but needed for Response
}
