package core

import (
	"net/url"
)

// PageIterator represents an iterator object that can be used to get subsequent pages of a collection.
type PageIterator[T any, C CollectionResponse[T]] struct {
	currentPage PageResult[T]
	client      Client
	pauseIndex  int
}

// NewPageIterator creates a new PageIterator instance.
func NewPageIterator[T any, C CollectionResponse[T]](currentPage CollectionResponse[T], client Client) (*PageIterator[T, C], error) {
	if isNil(client) {
		return nil, ErrNilClient
	}

	page, err := convertToPage(currentPage)
	if err != nil {
		return nil, err
	}

	return &PageIterator[T, C]{
		currentPage: page,
		client:      client,
	}, nil
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (pI *PageIterator[T, C]) Iterate(callback func(pageItem *T) bool) error {
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
func (pI *PageIterator[T, C]) enumerate(callback func(item *T) bool) bool {
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
func (pI *PageIterator[T, C]) next() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.NextPageLink)
}

// Last fetches the last page of results.
func (pI *PageIterator[T, C]) Last() (PageResult[T], error) {
	return pI.fetchAndConvertPage(pI.currentPage.LastPageLink)
}

// fetchAndConvertPage fetches next page and converts it
func (pI *PageIterator[T, C]) fetchAndConvertPage(uri string) (PageResult[T], error) {
	var page PageResult[T]

	resp, err := pI.fetchPage(uri)
	if err != nil {
		return page, err
	}

	page, err = convertToPage[T]((*resp))
	if err != nil {
		return page, err
	}

	return page, nil
}

// fetchPage fetches the specified uri page of results.
func (pI *PageIterator[T, C]) fetchPage(uri string) (*C, error) {
	var collectionResp C
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

	return &collectionResp, nil
}
