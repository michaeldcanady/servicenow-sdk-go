package tableapi

import (
	"net/http"
	"regexp"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// TableCollectionResponse3[T] represents a collection of T table entries.
type TableCollectionResponse3[T Entry] struct {
	// Result the list of page items from the current page.
	Result []*T
	// NextPageLink the URI for the next page.
	NextPageLink string
	// PreviousPageLink the URI for the previous page.
	PreviousPageLink string
	// FirstPageLink the URI for the first page.
	FirstPageLink string
	// LastPageLink the URI for the last page.
	LastPageLink string
}

// parsePaginationHeaders parses the pagination headers from the response
func (r TableCollectionResponse3[T]) parsePaginationHeaders(headers http.Header) {
	linkHeaderRegex := regexp.MustCompile(`<([^>]+)>;rel="([^"]+)"`)

	links := make(map[string]string)

	hearderLinks := headers["Link"]

	for _, header := range hearderLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := match[2]

			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case "first":
				links["firstPageLink"] = link
			case "prev":
				links["previousPageLink"] = link
			case "next":
				links["nextPageLink"] = link
			case "last":
				links["lastPageLink"] = link
			}
		}
	}

	r.FirstPageLink = links["firstPageLink"]
	r.PreviousPageLink = links["previousPageLink"]
	r.NextPageLink = links["nextPageLink"]
	r.LastPageLink = links["lastPageLink"]
}

// ParseHeaders parses the needed headers from the response.
func (r TableCollectionResponse3[T]) ParseHeaders(headers http.Header) {
	r.parsePaginationHeaders(headers)
}

// ToPage converts r to `core.PageResult`
func (r TableCollectionResponse3[T]) ToPage() core.PageResult[T] {
	return core.PageResult[T]{
		Result:           r.Result,
		NextPageLink:     r.NextPageLink,
		PreviousPageLink: r.PreviousPageLink,
		LastPageLink:     r.LastPageLink,
		FirstPageLink:    r.FirstPageLink,
	}
}
