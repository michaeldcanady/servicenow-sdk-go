package tableapi

import (
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

// PageIterator represents an iterator for paginated results from a table.
type PageIterator struct {
	currentPage PageResult
	client      core.Client
	pauseIndex  int
}

// defaultCallback is the default function to determine whether to continue iterating.
func defaultCallback(pageItem *TableEntry) bool {
	return true
}

// NewPageIterator creates a new PageIterator instance.
func NewPageIterator(currentPage interface{}, client core.Client) (*PageIterator, error) {
	if client == nil {
		return nil, ErrNilClient
	}

	page, err := convertToPage(currentPage)
	if err != nil {
		return nil, err
	}

	return &PageIterator{
		currentPage: page,
		client:      client,
	}, nil
}

// Iterate iterates through pages and invokes the provided callback for each page item.
func (p *PageIterator) Iterate(callback func(pageItem *TableEntry) bool) error {

	if callback == nil {
		callback = defaultCallback
	}

	for {
		keepIterating := p.enumerate(callback)

		if !keepIterating {
			// Callback returned false, stop iterating through pages.
			return nil
		}

		if p.currentPage.NextPageLink == "" {
			return nil
		}

		nextPage, err := p.next()
		if err != nil {
			return err
		}

		p.currentPage = nextPage
		p.pauseIndex = 0
	}
}

// enumerate iterates through the items on the current page and invokes the callback.
func (p *PageIterator) enumerate(callback func(item *TableEntry) bool) bool {
	keepIterating := true

	pageItems := p.currentPage.Result
	if pageItems == nil {
		return false
	}

	for i := p.pauseIndex; i < len(pageItems); i++ {
		keepIterating = callback(pageItems[i])

		if !keepIterating {
			break
		}

		p.pauseIndex = i + 1
	}
	return keepIterating
}

// next fetches the next page of results.
func (pI *PageIterator) next() (PageResult, error) {
	var page PageResult

	resp, err := pI.fetchNextPage()
	if err != nil {
		return page, err
	}

	page, err = convertToPage(resp)
	if err != nil {
		return page, err
	}

	return page, nil
}

// Last fethces the last page of results.
func (pI *PageIterator) Last() (PageResult, error) {
	var page PageResult

	resp, err := pI.fetchLastPage()
	if err != nil {
		return page, err
	}

	page, err = convertToPage(resp)
	if err != nil {
		return page, err
	}

	return page, nil
}

// fetchNextPage fetches the next page of results.
func (pI *PageIterator) fetchNextPage() (*TableCollectionResponse, error) {
	return pI.fetchPage(pI.currentPage.NextPageLink)
}

// fetchLastPage fetches the last page of results.
func (pI *PageIterator) fetchLastPage() (*TableCollectionResponse, error) {
	return pI.fetchPage(pI.currentPage.LastPageLink)
}

// fetchPage fetches the specified uri page of results.
func (pI *PageIterator) fetchPage(uri string) (*TableCollectionResponse, error) {
	var collectionResp TableCollectionResponse
	var err error

	if uri == "" {
		return nil, nil
	}

	parsedLink, err := url.Parse(uri)
	if err != nil {
		return &collectionResp, ErrParsing
	}

	requestInformation := core.NewRequestInformation()
	requestInformation.Method = core.GET
	requestInformation.SetUri(parsedLink)

	resp, err := pI.client.Send(requestInformation, nil)

	if err != nil {
		return nil, err
	}

	err = core.FromJson(resp, &collectionResp)
	if err != nil {
		return nil, nil
	}

	return &collectionResp, nil
}
