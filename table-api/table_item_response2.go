package tableapi

import "net/http"

// TableItemResponse2 [T TableEntry2] a single item response from the Table API.
type TableItemResponse2[T TableEntry2] struct {
	Result *T
}

func (iR *TableItemResponse2[T]) ParseHeaders(headers http.Header) {}
