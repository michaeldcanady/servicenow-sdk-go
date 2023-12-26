package tableapi

import (
	"net/http"
	"regexp"
)

// TableCollectionResponse2[T] represents a collection of T table entries.
type TableCollectionResponse2[T TableEntry2] struct {
	Result           []*T
	NextPageLink     string
	PreviousPageLink string
	FirstPageLink    string
	LastPageLink     string
}

// parsePaginationHeaders parses the pagination headers from the response
func (cR *TableCollectionResponse2[T]) parsePaginationHeaders(headers http.Header) {
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

	cR.FirstPageLink = links["firstPageLink"]
	cR.PreviousPageLink = links["previousPageLink"]
	cR.NextPageLink = links["nextPageLink"]
	cR.LastPageLink = links["lastPageLink"]
}

// ParseHeaders parses the needed headers from the response.
func (cR *TableCollectionResponse2[T]) ParseHeaders(headers http.Header) {
	cR.parsePaginationHeaders(headers)
}
