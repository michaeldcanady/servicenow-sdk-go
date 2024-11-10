package tableapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableCollectionResponse2ParsePaginationHeaders(t *testing.T) {
	headers := http.Header{}
	headers.Add("Link", `<http://example.com/first>;rel="first",<http://example.com/prev>;rel="prev",<http://example.com/next>;rel="next",<http://example.com/last>;rel="last"`)

	cR := &TableCollectionResponse2[Entry]{}
	cR.ParseHeaders(headers)

	assert.Equal(t, "http://example.com/first", cR.FirstPageLink)
	assert.Equal(t, "http://example.com/prev", cR.PreviousPageLink)
	assert.Equal(t, "http://example.com/next", cR.NextPageLink)
	assert.Equal(t, "http://example.com/last", cR.LastPageLink)
}

func TestTableCollectionResponse2ToPage(t *testing.T) {
	cR := &TableCollectionResponse2[Entry]{
		Result:           []*Entry{},
		NextPageLink:     "http://example.com/next",
		PreviousPageLink: "http://example.com/prev",
		FirstPageLink:    "http://example.com/first",
		LastPageLink:     "http://example.com/last",
	}

	page := cR.ToPage()

	assert.Equal(t, cR.Result, page.Result)
	assert.Equal(t, cR.NextPageLink, page.NextPageLink)
	assert.Equal(t, cR.PreviousPageLink, page.PreviousPageLink)
	assert.Equal(t, cR.FirstPageLink, page.FirstPageLink)
	assert.Equal(t, cR.LastPageLink, page.LastPageLink)
}
