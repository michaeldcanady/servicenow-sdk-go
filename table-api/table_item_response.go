package tableapi

import "net/http"

type TableItemResponse struct {
	Result *TableEntry
}

func (iR *TableItemResponse) ParseHeaders(headers http.Header) {}
