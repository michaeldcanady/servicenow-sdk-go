package model

import (
	"errors"

	"github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageResult represents a page object built from a graph response object
type PageResult[T serialization.Parsable] struct {
	nextLink  *string
	prevLink  *string
	firstLink *string
	lastLink  *string
	value     []T
}

func convertToPage[T serialization.Parsable](response ServiceNowCollectionResponse[T]) (PageResult[T], error) {
	if response == nil {
		return PageResult[T]{}, errors.New("response cannot be nil")
	}

	results, err := response.GetResult()
	if err != nil {
		return PageResult[T]{}, err
	}

	nextLink, err := response.GetNextLink()
	if err != nil {
		return PageResult[T]{}, err
	}

	prevLink, err := response.GetPreviousLink()
	if err != nil {
		return PageResult[T]{}, err
	}

	firstLink, err := response.GetFirstLink()
	if err != nil {
		return PageResult[T]{}, err
	}

	lastLink, err := response.GetLastLink()
	if err != nil {
		return PageResult[T]{}, err
	}

	return PageResult[T]{
		nextLink:  nextLink,
		prevLink:  prevLink,
		firstLink: firstLink,
		lastLink:  lastLink,
		value:     results,
	}, nil
}
