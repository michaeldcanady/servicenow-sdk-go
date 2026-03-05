package tableapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableCollectionResponse2ParsePaginationHeaders(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected struct {
			first    string
			prev     string
			next     string
			last     string
		}
	}{
		{
			name: "All links present",
			headers: func() http.Header {
				h := http.Header{}
				h.Add("Link", `<http://example.com/first>;rel="first",<http://example.com/prev>;rel="prev",<http://example.com/next>;rel="next",<http://example.com/last>;rel="last"`)
				return h
			}(),
			expected: struct {
				first    string
				prev     string
				next     string
				last     string
			}{
				first: "http://example.com/first",
				prev:  "http://example.com/prev",
				next:  "http://example.com/next",
				last:  "http://example.com/last",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cR := &TableCollectionResponse2[Entry]{}
			cR.ParseHeaders(test.headers)

			assert.Equal(t, test.expected.first, cR.FirstPageLink)
			assert.Equal(t, test.expected.prev, cR.PreviousPageLink)
			assert.Equal(t, test.expected.next, cR.NextPageLink)
			assert.Equal(t, test.expected.last, cR.LastPageLink)
		})
	}
}

func TestTableCollectionResponse2ToPage(t *testing.T) {
	tests := []struct {
		name string
		cR   *TableCollectionResponse2[Entry]
	}{
		{
			name: "Standard response",
			cR: &TableCollectionResponse2[Entry]{
				Result:           []*Entry{},
				NextPageLink:     "http://example.com/next",
				PreviousPageLink: "http://example.com/prev",
				FirstPageLink:    "http://example.com/first",
				LastPageLink:     "http://example.com/last",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			page := test.cR.ToPage()

			assert.Equal(t, test.cR.Result, page.Result)
			assert.Equal(t, test.cR.NextPageLink, page.NextPageLink)
			assert.Equal(t, test.cR.PreviousPageLink, page.PreviousPageLink)
			assert.Equal(t, test.cR.FirstPageLink, page.FirstPageLink)
			assert.Equal(t, test.cR.LastPageLink, page.LastPageLink)
		})
	}
}
