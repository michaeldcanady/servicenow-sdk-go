package tableapi

import (
	"net/http"
	"regexp"

	"github.com/RecoLabs/servicenow-sdk-go/core"
)

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
