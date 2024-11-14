package core

import (
	"net/http"
	"net/url"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
)

type Parsable[T any] func(*http.Response) (CollectionResponse[T], error)

// PageIterator2[T]
type PageIterator2[T any] struct {
	currentPage     PageResult[T]
	client          Client
	pauseIndex      int
	constructorFunc Parsable[T]
}

// NewPageIterator2[T] creates a new PageIterator instance.
func NewPageIterator2[T any](currentPage CollectionResponse[T], client Client, constructorFunc Parsable[T]) (*PageIterator2[T], error) {
	if internal.IsNil(client) {
		return nil, ErrNilClient
	}

	page, err := convertToPage[T](currentPage)
	if err != nil {
		return nil, err
	}

	return &PageIterator2[T]{
		currentPage:     page,
		client:          client,
		constructorFunc: constructorFunc,
	}, nil
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (pI *PageIterator2[T]) Iterate(callback func(pageItem *T) bool, reversed bool) error {
	if callback == nil {
		return ErrNilCallback
	}

	for {
		keepIterating := pI.enumerate(callback)

		if !keepIterating {
			// Callback returned false, stop iterating through pages.
			return nil
		}

		nextPage, err := pI.nextPage(reversed)
		if err != nil {
			return err
		}

		if len(nextPage.Result) == 0 {
			return nil
		}

		pI.currentPage = nextPage
		pI.pauseIndex = 0
	}
}

// enumerate iterates through the items on the current page and invokes the callback.
func (pI *PageIterator2[T]) enumerate(callback func(item *T) bool) bool {
	keepIterating := true

	pageItems := pI.currentPage.Result
	if pageItems == nil {
		return false
	}

	for i := pI.pauseIndex; i < len(pageItems); i++ {
		keepIterating = callback(pageItems[i])

		if !keepIterating {
			break
		}

		pI.pauseIndex = i + 1
	}
	return keepIterating
}

// nextPage fetches the next (or previous) page of results.
// If reversed is true, it fetches the previous page; otherwise, it fetches the next page.
// Returns an empty PageResult if there is no next (or previous) link.
func (pI *PageIterator2[T]) nextPage(reversed bool) (PageResult[T], error) {
	nextLink := pI.currentPage.NextPageLink
	if reversed {
		nextLink = pI.currentPage.PreviousPageLink
	}

	if nextLink == "" {
		return PageResult[T]{}, nil
	}

	return pI.fetchAndConvertPage(nextLink)
}

// Next fetches the Next page of results.
func (pI *PageIterator2[T]) Next() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.NextPageLink)
}

// Last fetches the last page of results.
func (pI *PageIterator2[T]) Last() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.LastPageLink)
}

// First fetches the first page of results.
func (pI *PageIterator2[T]) First() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.FirstPageLink)
}

// Previous fetches the previous page of results.
func (pI *PageIterator2[T]) Previous() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.PreviousPageLink)
}

// fetchAndConvertPage fetches provided page and converts it
func (pI *PageIterator2[T]) fetchAndConvertPage(uri string) (PageResult[T], error) {
	var page PageResult[T]

	resp, err := pI.fetchPage(uri)
	if err != nil {
		return page, err
	}

	page, err = convertToPage(resp)
	if err != nil {
		return page, err
	}

	return page, nil
}

// fetchPage fetches the specified uri page of results.
func (pI *PageIterator2[T]) fetchPage(uri string) (CollectionResponse[T], error) {
	if uri == "" {
		return nil, ErrEmptyURI
	}

	// parse provided link
	nextLink, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}

	// build request information
	requestInformation := NewRequestInformation()
	requestInformation.Method = GET
	requestInformation.SetUri(nextLink)

	resp, err := pI.client.Send(requestInformation, nil)
	if err != nil {
		return nil, err
	}

	collectionResp, err := pI.constructorFunc(resp)
	if err != nil {
		return nil, err
	}

	return collectionResp, nil
}
