package core

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageResult represents a page object built from a graph response object
type PageResult[T serialization.Parsable] struct {
	NextLink  *string
	PrevLink  *string
	FirstLink *string
	LastLink  *string
	Result    []T
}

// convertToPage converts a ServiceNowCollectionResponse to a PageResult
func convertToPage[T serialization.Parsable](response ServiceNowCollectionResponse[T]) (PageResult[T], error) {
	var page PageResult[T]
	var err error

	if response == nil {
		return page, errors.New("response cannot be nil")
	}

	page.Result, err = response.GetResult()
	if err != nil {
		return page, err
	}

	page.NextLink, err = response.GetNextLink()
	if err != nil {
		return page, err
	}

	page.PrevLink, err = response.GetPreviousLink()
	if err != nil {
		return page, err
	}

	page.FirstLink, err = response.GetFirstLink()
	if err != nil {
		return page, err
	}

	page.LastLink, err = response.GetLastLink()
	if err != nil {
		return page, err
	}

	return page, nil
}
