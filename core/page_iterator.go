package core

import (
	"net/url"
)

// PageIterator
type PageIterator[T any] struct {
	currentPage PageResult[T]
	client      Client
	pauseIndex  int
}

// NewPageIterator creates a new PageIterator instance.
func NewPageIterator[T any](currentPage CollectionResponse[T], client Client) (*PageIterator[T], error) {
	if client == nil {
		return nil, ErrNilClient
	}

	page, err := convertToPage(currentPage)
	if err != nil {
		return nil, err
	}

	return &PageIterator[T]{
		currentPage: page,
		client:      client,
	}, nil
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (pI *PageIterator[T]) Iterate(callback func(pageItem *T) bool) error {
	if callback == nil {
		return ErrNilCallback
	}

	for {
		keepIterating := pI.enumerate(callback)

		if !keepIterating {
			// Callback returned false, stop iterating through pages.
			return nil
		}

		if pI.currentPage.NextPageLink == "" {
			// NextPageLink is empty, stop iterating through pages.
			return nil
		}

		// TODO: Add option for reverse
		nextPage, err := pI.next()
		if err != nil {
			return err
		}

		pI.currentPage = nextPage
		pI.pauseIndex = 0
	}
}

// enumerate iterates through the items on the current page and invokes the callback.
func (pI *PageIterator[T]) enumerate(callback func(item *T) bool) bool {
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

// next fetches the next page of results.
func (pI *PageIterator[T]) next() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.NextPageLink)
}

// Last fetches the last page of results.
func (pI *PageIterator[T]) Last() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.LastPageLink)
}

func (pI *PageIterator[T]) fetchAndConvertPage(uri string) (PageResult[T], error) {
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
func (pI *PageIterator[T]) fetchPage(uri string) (CollectionResponse[T], error) {
	var collectionResp CollectionResponse[T]
	var err error

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

	err = ParseResponse(resp, &collectionResp)
	if err != nil {
		return nil, err
	}

	return collectionResp, nil
}
