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

// fetchAndConvertPage retrieves a page from a given URI and converts it into a PageResult.
// The function first fetches the page using the fetchPage method. If an error occurs during fetching, it returns immediately.
// If the page is fetched successfully, it is then converted into a PageResult using the convertToPage function.
// Note: The error from convertToPage is currently ignored as no test case has been found to cause an error.
// The function returns the PageResult and any error encountered.
//
// Parameters:
//   - uri: A string representing the URI of the page to be fetched and converted.
//
// Returns:
//   - A PageResult representing the fetched and converted page.
//   - An error that will be non-nil if an error occurred during the fetch operation.
func (pI *PageIterator) fetchAndConvertPage(uri string) (PageResult, error) {
	var page PageResult

	resp, err := pI.fetchPage(uri)
	if err != nil {
		return page, err
	}

	page, _ = convertToPage(resp) //Can't find test case to cause error

	return page, nil
}

// fetchPage fetches the specified uri page of results.
func (pI *PageIterator) fetchPage(uri string) (*TableCollectionResponse, error) {
	var collectionResp *TableCollectionResponse
	var err error

	if uri == "" {
		return nil, ErrEmptyURI
	}

	nextLink, err := url.ParseRequestURI(pI.currentPage.NextPageLink)
	if err != nil {
		return nil, err
	}

	requestInformation := core.NewRequestInformation()
	requestInformation.Method = core.GET
	requestInformation.SetUri(nextLink)

	resp, err := pI.client.Send(requestInformation, nil)
	if err != nil {
		return nil, err
	}

	err = core.ParseResponse(resp, &collectionResp)
	if err != nil {
		return nil, err
	}

	return collectionResp, nil
}
