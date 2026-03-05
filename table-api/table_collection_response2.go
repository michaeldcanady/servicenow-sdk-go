package tableapi

import (
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
)

// Deprecated: deprecated since v1.9.0. Please use [newInternal.ServiceNowCollectionResponse]
type TableCollectionResponse2[T Entry] struct {
	// Result is a slice of pointers to table entries.
	Result []*T
	// NextPageLink is the URL to the next page of results.
	NextPageLink string
	// PreviousPageLink is the URL to the previous page of results.
	PreviousPageLink string
	// FirstPageLink is the URL to the first page of results.
	FirstPageLink string
	// LastPageLink is the URL to the last page of results.
	LastPageLink string
}

// SetFirstPageLink sets the first page link
func (cR *TableCollectionResponse2[T]) SetFirstPageLink(link string) {
	cR.FirstPageLink = link
}

// SetPreviousPageLink sets the previous page link
func (cR *TableCollectionResponse2[T]) SetPreviousPageLink(link string) {
	cR.PreviousPageLink = link
}

// SetNextPageLink sets the next page link
func (cR *TableCollectionResponse2[T]) SetNextPageLink(link string) {
	cR.NextPageLink = link
}

// SetLastPageLink sets the last page link
func (cR *TableCollectionResponse2[T]) SetLastPageLink(link string) {
	cR.LastPageLink = link
}

// ParseHeaders parses the needed headers from the response.
func (cR *TableCollectionResponse2[T]) ParseHeaders(headers http.Header) {
	newInternal.ParseHTTPHeaders(cR, headers)
}

// ToPage converts a TableCollectionResponse2 to a PageResult
func (cR *TableCollectionResponse2[T]) ToPage() core.PageResult[T] {
	return core.PageResult[T]{
		Result:           cR.Result,
		NextPageLink:     cR.NextPageLink,
		PreviousPageLink: cR.PreviousPageLink,
		FirstPageLink:    cR.FirstPageLink,
		LastPageLink:     cR.LastPageLink,
	}
}
