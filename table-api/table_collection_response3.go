package tableapi

import (
	"net/http"
	"regexp"
	"strconv"
)

const (
	firstPageHeaderKey    = "first"
	prevPageHeaderKey     = "prev"
	nextPageHeaderKey     = "next"
	lastPageHeaderKey     = "last"
	paginationHeaderRegex = `<([^>]+)>;rel="([^"]+)"`
	countHeaderKey        = "X-Total-Count"
)

type TableCollectionResponse3[T TableRecord] interface {
	GetResults() []*T
	GetNextPageLink() string
	GetPreviousPageLink() string
	GetFirstPageLink() string
	GetLastPageLink() string
	GetCount() int
}

// TableCollectionResponse2 represents a collection of table entries.
type tableCollectionResponse3[T TableRecord] struct {
	// Result is a slice of pointers to table entries.
	Result []*T
	// nextPageLink is the URL to the next page of results.
	nextPageLink string
	// previousPageLink is the URL to the previous page of results.
	previousPageLink string
	// firstPageLink is the URL to the first page of results.
	firstPageLink string
	// lastPageLink is the URL to the last page of results.
	lastPageLink string
	// count is the count of total results. Is -1 if there was an error parsing.
	count int
}

// parsePaginationHeaders parses the pagination headers from the response
func (cR *tableCollectionResponse3[T]) parsePaginationHeaders(headers http.Header) {
	linkHeaderRegex := regexp.MustCompile(paginationHeaderRegex)

	links := make(map[string]string)

	hearderLinks := headers["Link"]

	for _, header := range hearderLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := match[2]

			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case firstPageHeaderKey:
				links[firstPageHeaderKey] = link
			case prevPageHeaderKey:
				links[prevPageHeaderKey] = link
			case nextPageHeaderKey:
				links[nextPageHeaderKey] = link
			case lastPageHeaderKey:
				links[lastPageHeaderKey] = link
			}
		}
	}

	cR.firstPageLink = links[firstPageHeaderKey]
	cR.previousPageLink = links[prevPageHeaderKey]
	cR.nextPageLink = links[nextPageHeaderKey]
	cR.lastPageLink = links[lastPageHeaderKey]
}

func (cR *tableCollectionResponse3[T]) parseCountHeader(headers http.Header) {
	count, err := strconv.Atoi(headers.Get(countHeaderKey))
	if err != nil {
		count = -1
	}
	cR.count = count
}

// ParseHeaders parses the needed headers from the response.
func (cR *tableCollectionResponse3[T]) ParseHeaders(headers http.Header) {
	cR.parsePaginationHeaders(headers)
	cR.parseCountHeader(headers)
}

func (cR *tableCollectionResponse3[T]) GetCount() int {
	return cR.count
}

func (cR *tableCollectionResponse3[T]) GetNextPageLink() string {
	return cR.nextPageLink
}

func (cR *tableCollectionResponse3[T]) GetPreviousPageLink() string {
	return cR.previousPageLink
}

func (cR *tableCollectionResponse3[T]) GetFirstPageLink() string {
	return cR.firstPageLink
}

func (cR *tableCollectionResponse3[T]) GetLastPageLink() string {
	return cR.lastPageLink
}

func (cR *tableCollectionResponse3[T]) GetResults() []*T {
	return cR.Result
}
