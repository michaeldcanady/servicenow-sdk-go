package tableapi

import (
	"errors"
	"net/url"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
)

var (
	ErrNilClient         = errors.New("client can't be nil")
	ErrNilResponse       = errors.New("response can't be nil")
	ErrNilResult         = errors.New("result property missing in response object")
	ErrWrongResponseType = errors.New("incorrect Response Type")
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
func (p *PageIterator) next() (PageResult, error) {
	var page PageResult

	resp, err := p.fetchNextPage()
	if err != nil {
		return page, err
	}

	page, err = convertToPage(resp)
	if err != nil {
		return page, err
	}

	return page, nil
}

// convertToPage converts a response into a PageResult.
func convertToPage(response interface{}) (PageResult, error) {
	var page PageResult

	if response == nil {
		return page, ErrNilResponse
	}

	collectionRep, ok := response.(TableCollectionResponse)
	if !ok {
		return page, ErrWrongResponseType
	}

	page.Result = collectionRep.Result
	page.FirstPageLink = collectionRep.FirstPageLink
	page.LastPageLink = collectionRep.LastPageLink
	page.NextPageLink = collectionRep.NextPageLink
	page.PreviousPageLink = collectionRep.PreviousPageLink

	return page, nil
}

// fetchNextPage fetches the next page of results.
func (pI *PageIterator) fetchNextPage() (*TableCollectionResponse, error) {
	var collectionResp *TableCollectionResponse
	var err error

	if pI.currentPage.NextPageLink == "" {
		return nil, nil
	}

	nextLink, err := url.Parse(pI.currentPage.NextPageLink)
	if err != nil {
		return collectionResp, errors.New("parsing nextLink url failed")
	}

	requestInformation := core.NewRequestInformation()
	requestInformation.Method = core.GET
	requestInformation.SetUri(nextLink)

	resp, err := pI.client.Send(requestInformation, nil)
	if err != nil {
		return nil, err
	}

	collectionResp, err = core.FromJson[TableCollectionResponse](resp)
	if err != nil {
		return nil, nil
	}

	return collectionResp, nil
}