package tableapi

import (
	"context"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestNewPageIteratorWithClient(t *testing.T) {
	// Mock client and current page
	client := &mockClient{}
	currentPage := TableCollectionResponse{
		Result:           []*TableEntry{&fakeEntry},
		FirstPageLink:    fakeFirstLink,
		LastPageLink:     fakeLastLink,
		NextPageLink:     fakeNextLink,
		PreviousPageLink: fakePrevLink,
	}

	pageIterator, err := NewPageIterator(context.Background(), currentPage, client)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if pageIterator == nil {
		t.Error("Expected PageIterator, but got nil")
	}

	assert.Equal(t, &expectedIterator, pageIterator)
}

func TestNewPageIteratorWithoutClient(t *testing.T) {
	// Mock current page
	currentPage := TableCollectionResponse{
		// Initialize with test data
	}

	pageIterator, err := NewPageIterator(context.Background(), currentPage, nil)

	assert.Equal(t, (*PageIterator)(nil), pageIterator)
	assert.Equal(t, ErrNilClient, err)
}

func TestNewPageIteratorNilCurrentPageWithClient(t *testing.T) {
	client := &mockClient{}
	pageIterator, err := NewPageIterator(context.Background(), nil, client)

	assert.Equal(t, (*PageIterator)(nil), pageIterator)
	assert.Equal(t, ErrNilResponse, err)
}

func TestPageIteratorNextWithLinkNoError(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeCollectionLinkWithLinks,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(context.Background(), currentPage, client)
	assert.Nil(t, err)

	page, err := pageIterator.next()
	assert.Nil(t, err)

	assert.Equal(t, expectedResult, page)
}

func TestPageIteratorEnumerateAll(t *testing.T) {
	pageIterator := PageIterator{
		currentPage: expectedResult,
	}

	enumCount := 0

	keepIterating := pageIterator.enumerate(func(item *TableEntry) bool {
		index := pageIterator.pauseIndex

		result := expectedResult.Result[index]

		assert.Equal(t, result, item)

		enumCount += 1

		return true
	})
	assert.Equal(t, true, keepIterating)
	assert.Equal(t, len(expectedResult.Result), enumCount)
}

func TestPageIteratorEnumerateOnce(t *testing.T) {
	pageIterator := PageIterator{
		currentPage: expectedResult,
	}

	enumCount := 0

	keepIterating := pageIterator.enumerate(func(item *TableEntry) bool {
		index := pageIterator.pauseIndex

		result := expectedResult.Result[index]

		assert.Equal(t, result, item)

		enumCount += 1

		return false
	})
	assert.Equal(t, false, keepIterating)
	assert.Equal(t, 1, enumCount)
}

func TestIterateWithNoCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	err := pageIterator.Iterate(nil)
	assert.ErrorIs(t, err, ErrNilCallback)
}

func TestPageIteratorIterateSinglePageWithoutNextLinkWithoutCurrentPageWithCallback(t *testing.T) {
	expectedIterator.currentPage = PageResult{}

	// Mock callback function
	callback := func(pageItem *TableEntry) bool {
		// Implement your callback logic for testing
		return true
	}

	err := expectedIterator.Iterate(callback)
	assert.Nil(t, err)
}

func TestPageIteratorIterateSinglePageWithoutNextLinkWithCurrentPageWithCallback(t *testing.T) {
	expectedIterator.currentPage.NextPageLink = ""

	// Mock callback function
	callback := func(pageItem *TableEntry) bool {
		// Implement your callback logic for testing
		return true
	}

	err := expectedIterator.Iterate(callback)
	assert.Nil(t, err)
}

func TestPageIteratorIterateSinglePageWithNextLinkWithCurrentPageWithCallback(t *testing.T) {
	expectedIterator.currentPage.NextPageLink = fakeCollectionLinkKey

	// Mock callback function
	callback := func(pageItem *TableEntry) bool {
		// Implement your callback logic for testing
		return true
	}

	err := expectedIterator.Iterate(callback)
	assert.Nil(t, err)
}

func TestPageIteratorIterateMultiplePagesWithCallback(t *testing.T) {
	// Mock PageIterator
	pageIterator := &PageIterator{
		currentPage: PageResult{
			NextPageLink: fakeCollectionLinkWithLinks,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	// Mock callback function
	callback := func(pageItem *TableEntry) bool {
		// Implement your callback logic for testing
		return true
	}

	err := pageIterator.Iterate(callback)
	assert.Nil(t, err)
}

func TestPageIteratorLast(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeNextLink,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	_, err := pageIterator.Last()

	assert.Error(t, err)
}

func TestPageFetchPageSendErr(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeCollectionLinkStatusFailed,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	_, err := pageIterator.fetchPage(fakeCollectionLinkStatusFailed)
	assert.Error(t, err)
}

func TestPageFetchPageEmptyUri(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeCollectionLinkStatusFailed,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	_, err := pageIterator.fetchPage("")
	assert.ErrorIs(t, err, ErrEmptyURI)
}

func TestPageIteratorFetchAndConvertPageWithLinkErrNilResponseBody(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeCollectionLinkWithLinksErr,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(context.Background(), currentPage, client)
	assert.Nil(t, err)

	page, err := pageIterator.fetchAndConvertPage(pageIterator.currentPage.NextPageLink)
	assert.ErrorIs(t, err, core.ErrNilResponseBody)

	assert.Equal(t, PageResult{}, page)
}

func TestPageIteratorFetchAndConvertPageWithoutLink(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeCollectionLinkNilResponse,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(context.Background(), currentPage, client)
	assert.Nil(t, err)

	page, err := pageIterator.fetchAndConvertPage(pageIterator.currentPage.NextPageLink)
	assert.ErrorIs(t, err, core.ErrNilResponse)

	assert.Equal(t, PageResult{}, page)
}
