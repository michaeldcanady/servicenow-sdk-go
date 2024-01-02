package tableapi

import "net/http"

// TableItemResponse2[T] represents a T Entry single table record response.
type TableItemResponse2[T Entry] struct {
	Result *T
}

// ParseHeaders parses information from headers.
func (r *TableItemResponse2[T]) ParseHeaders(headers http.Header) {
	// no headers need parsing.
}
