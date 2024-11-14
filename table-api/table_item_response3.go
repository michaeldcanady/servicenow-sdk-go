package tableapi

import (
	"net/http"

	intCore "github.com/RecoLabs/servicenow-sdk-go/internal/core"
)

// parseTableItemResponse3 parsable to create TableItemResponse3
//
//nolint:unused
func parseTableItemResponse3(resp *http.Response) (intCore.Response, error) {
	var itemResp *tableItemResponse3[*TableRecordImpl]

	err := intCore.ParseResponse(resp, &itemResp)
	if err != nil {
		return nil, err
	}

	return itemResp, nil
}

type TableItemResponse3[T TableRecord] interface {
	GetResult() T
}

// tableItemResponse3[T] represents a T Entry single table record response.
type tableItemResponse3[T TableRecord] struct {
	Result T
}

// ParseHeaders parses information from headers.
func (r *tableItemResponse3[T]) ParseHeaders(headers http.Header) {
	// no headers need parsing.
}

func (r *tableItemResponse3[T]) GetResult() T {
	return r.Result
}
