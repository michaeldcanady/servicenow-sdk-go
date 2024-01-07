package tableapi

import (
	"net/http"
	"regexp"
)

// Deprecated: deprecated since v{version}. Use `TableCollectionResponse3` instead.
//
// TableCollectionResponse2[T] represents a collection of T table entries.
type TableCollectionResponse2[T Entry] struct {
	Result           []*T
	NextPageLink     string
	PreviousPageLink string
	FirstPageLink    string
	LastPageLink     string
}

// parsePaginationHeaders parses the pagination headers from the response
func (r *TableCollectionResponse2[T]) parsePaginationHeaders(headers http.Header) {
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
func (r *TableCollectionResponse2[T]) ParseHeaders(headers http.Header) {
	r.parsePaginationHeaders(headers)
}
