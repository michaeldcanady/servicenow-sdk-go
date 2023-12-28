package tableapi

import "net/http"

// TableItemResponse2 [T TableEntry2] a single item response from the Table API.
type TableItemResponse2[T Entry] struct {
	Result *T
}

// ParseHeaders parses needed headers from Table Item Response.
func (iR *TableItemResponse2[T]) ParseHeaders(headers http.Header) {}
