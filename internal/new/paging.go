package internal

import (
	"net/http"
	"regexp"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

const (
	linkHeaderKey = "Link"
)

var (
	linkHeaderRegex = regexp.MustCompile(`<([^>]+)>\s*;\s*(?:[^,;]+\s*;\s*)*rel="?([^";, ]+)"?`)
)

// ParseHeaders parses Link headers from ResponseHeaders into a ServiceNowCollectionResponse.
func ParseHeaders[T serialization.Parsable](collection ServiceNowCollectionResponse[T], headers *abstractions.ResponseHeaders) {
	if IsNil(collection) || IsNil(headers) {
		return
	}

	headerLinks := headers.Get(linkHeaderKey)
	for _, header := range headerLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := strings.ToLower(match[2])

			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case "first":
				_ = collection.SetFirstLink(&link)
			case "prev":
				_ = collection.SetPreviousLink(&link)
			case "next":
				_ = collection.SetNextLink(&link)
			case "last":
				_ = collection.SetLastLink(&link)
			}
		}
	}
}

// Paginatable represents an object that can have pagination links set.
type Paginatable interface {
	SetFirstPageLink(string)
	SetPreviousPageLink(string)
	SetNextPageLink(string)
	SetLastPageLink(string)
}

// ParseHTTPHeaders parses Link headers from http.Header into a Paginatable object.
func ParseHTTPHeaders(paginatable Paginatable, headers http.Header) {
	if IsNil(paginatable) || IsNil(headers) {
		return
	}

	headerLinks := headers[linkHeaderKey]
	for _, header := range headerLinks {
		linkMatches := linkHeaderRegex.FindAllStringSubmatch(header, -1)

		for _, match := range linkMatches {
			link := match[1]
			rel := strings.ToLower(match[2])

			// Determine the type of link based on the 'rel' attribute
			switch rel {
			case "first":
				paginatable.SetFirstPageLink(link)
			case "prev":
				paginatable.SetPreviousPageLink(link)
			case "next":
				paginatable.SetNextPageLink(link)
			case "last":
				paginatable.SetLastPageLink(link)
			}
		}
	}
}
