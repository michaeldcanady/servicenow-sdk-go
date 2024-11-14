package tableapi

import (
	"context"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// Deprecated: deprecated in v1.5.0. Use TablePageIterator[T] instead.
// PageIterator represents an iterator for paginated results from a table.
type PageIterator struct {
	ctx         context.Context
	currentPage PageResult
	client      core.Client
	pauseIndex  int
}

// Deprecated: deprecated in v1.5.0. Use NewTablePageIterator[T] instead.
// NewPageIterator creates a new PageIterator instance.
func NewPageIterator(ctx context.Context, currentPage interface{}, client core.Client) (*PageIterator, error) {
	if client == nil {
		return nil, ErrNilClient
	}

	page, err := convertToPage(currentPage)
	if err != nil {
		return nil, err
	}

	return &PageIterator{
		ctx:         ctx,
		currentPage: page,
		client:      client,
	}, nil
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (pI *PageIterator) Iterate(callback func(pageItem *TableEntry) bool) error {
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
			return nil
		}

		nextPage, err := pI.next()
		if err != nil {
			return err
		}

		pI.currentPage = nextPage
		pI.pauseIndex = 0
	}
}

// enumerate iterates through the items on the current page and invokes the callback.
func (pI *PageIterator) enumerate(callback func(item *TableEntry) bool) bool {
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
func (pI *PageIterator) next() (PageResult, error) {
	return pI.fetchAndConvertPage(pI.currentPage.NextPageLink)
}

// Last fetches the last page of results.
func (pI *PageIterator) Last() (PageResult, error) {
	return pI.fetchAndConvertPage(pI.currentPage.LastPageLink)
}

func (pI *PageIterator) fetchAndConvertPage(uri string) (PageResult, error) {
	var page PageResult

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
func (pI *PageIterator) fetchPage(uri string) (*TableCollectionResponse, error) {
	var collectionResp *TableCollectionResponse
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
	requestInformation := core.NewRequestInformation()
	requestInformation.Method = core.GET
	requestInformation.SetUri(nextLink)

	resp, err := pI.client.Send(pI.ctx, requestInformation, nil)
	if err != nil {
		return nil, err
	}

	err = core.ParseResponse(resp, &collectionResp)
	if err != nil {
		return nil, err
	}

	return collectionResp, nil
}
