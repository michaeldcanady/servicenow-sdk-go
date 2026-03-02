package tableapi

import "net/http"

// Deprecated: deprecated since v1.9.0. Please use [newInternal.ServiceNowItemResponse]
type TableItemResponse2[T Entry] struct {
	Result *T
}

// ParseHeaders parses information from headers.
func (r *TableItemResponse2[T]) ParseHeaders(headers http.Header) {
	// no headers need parsing.
}
