package tableapi

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

const (
	headerLinkKey  = "Link"
	headerCountKey = "X-Total-Count"
)

var linkHeaderRegex = regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)

// TableCollectionResponse2 represents a collection of table entries.
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
	// Count is the count of records returned by the query.
	Count int
}

// parsePaginationHeaders parses the pagination headers from the response
// parsePaginationHeaders parses the pagination headers from the response.
func (cR *TableCollectionResponse2[T]) parsePaginationHeaders(headers http.Header) {
	headerLinks, ok := headers[headerLinkKey]
	if !ok {
		return // Handle missing link headers
	}

	for _, header := range headerLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := match[2]

			// Map rel attribute to corresponding field
			switch rel {
			case "first":
				cR.FirstPageLink = link
			case "prev":
				cR.PreviousPageLink = link
			case "next":
				cR.NextPageLink = link
			case "last":
				cR.LastPageLink = link
			}
		}
	}
}

func (cR *TableCollectionResponse2[T]) parseCountHeaders(headers http.Header) {
	hearderLinks, ok := headers[headerCountKey]
	if !ok {
		return // Handle missing count headers
	}
	i, err := strconv.Atoi(hearderLinks[0])
	if err != nil {
		cR.Count = -1
		return
	}
	cR.Count = i
}

// ParseHeaders parses the needed headers from the response.
func (cR *TableCollectionResponse2[T]) ParseHeaders(headers http.Header) {
	cR.parsePaginationHeaders(headers)
	cR.parseCountHeaders(headers)
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
