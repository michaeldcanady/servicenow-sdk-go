package internal

import (
	"regexp"
	"strings"

	abstractions "github.com/microsoft/kiota-abstractions-go"
)

const (
	linkHeaderKey = "Link"
)

var (
	linkHeaderRegex = regexp.MustCompile(`<([^>]+)>\s*;\s*(?:[^,;]+\s*;\s*)*rel="?([^";, ]+)"?`)
)

// ParseHeaders parses Link headers from ResponseHeaders into a ServiceNowCollectionResponse.
func ParseHeaders(collection Paginatable, headers *abstractions.ResponseHeaders) {
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
			case firstKey:
				_ = collection.SetFirstLink(&link)
			case previousKey:
				_ = collection.SetPreviousLink(&link)
			case nextKey:
				_ = collection.SetNextLink(&link)
			case lastKey:
				_ = collection.SetLastLink(&link)
			}
		}
	}
}

// Paginatable represents an object that can have pagination links set.
type Paginatable interface {
	SetFirstLink(*string) error
	SetPreviousLink(*string) error
	SetNextLink(*string) error
	SetLastLink(*string) error
}
