package core

import (
	"net/url"
)

type PageIterator[T any, C CollectionResponse[T]] interface {
	fetchAndConvertPage(uri string) (PageResult[T], error)
	fetchPage(uri string) (*C, error)
	Current() *PageResult[T]
	SetCurret(page PageResult[T])
}

// BasePageIterator represents an iterator object that can be used to get subsequent pages of a collection.
type BasePageIterator[T any, C CollectionResponse[T]] struct {
	currentPage PageResult[T]
	client      Client
}

// NewPageIterator creates a new PageIterator instance.
func NewPageIterator[T any, C CollectionResponse[T]](currentPage CollectionResponse[T], client Client) (*BasePageIterator[T, C], error) {
	if isNil(client) {
		return nil, ErrNilClient
	}

	page, err := convertToPage(currentPage)
	if err != nil {
		return nil, err
	}

	return &BasePageIterator[T, C]{
		currentPage: page,
		client:      client,
	}, nil
}

// fetchAndConvertPage fetches next page and converts it
func (pI *BasePageIterator[T, C]) fetchAndConvertPage(uri string) (PageResult[T], error) {
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
func (pI *BasePageIterator[T, C]) fetchPage(uri string) (*C, error) {
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

func (pI *BasePageIterator[T, C]) Current() *PageResult[T] {
	return &pI.currentPage
}

func (pI *BasePageIterator[T, C]) SetCurret(page PageResult[T]) {
	pI.currentPage = page
}
