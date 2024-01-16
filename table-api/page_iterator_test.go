package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewPageIterator(t *testing.T) {
	client := &mockClient{}
	currentPage := TableCollectionResponse{
		Result:           []*TableEntry{&fakeEntry},
		FirstPageLink:    fakeFirstLink,
		LastPageLink:     fakeLastLink,
		NextPageLink:     fakeNextLink,
		PreviousPageLink: fakePrevLink,
	}

	tests := []internal.Test[*PageIterator]{
		{
			Title:    "WithClient",
			Input:    []interface{}{currentPage, client},
			Expected: &expectedIterator,
		},
		{
			Title:    "WithoutClient",
			Input:    []interface{}{currentPage, (*mockClient)(nil)},
			Expected: nil,
			Error:    ErrNilClient,
		},
		{
			Title:    "NilCurrentPageWithClient",
			Input:    []interface{}{nil, client},
			Expected: nil,
			Error:    ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			input := tt.Input.([]interface{})

			pageIterator, err := NewPageIterator(input[0], input[1].(core.Client))

			assert.Equal(t, tt.Error, err)
			assert.Equal(t, tt.Expected, pageIterator)
		})
	}
}

func TestPageIteratorNextWithLinkNoError(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeCollectionLinkWithLinks,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(currentPage, client)
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

func TestPageIterator_Last(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeNextLink,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	tests := []internal.Test[any]{
		{
			Title: "Valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			_, err := pageIterator.Last()
			assert.Error(t, err)
		})
	}
}

func TestPage_FetchPage(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			LastPageLink: fakeCollectionLinkStatusFailed,
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	tests := []internal.Test[any]{
		{
			Title: "EmptyUri",
			Input: "",
			Error: ErrEmptyURI,
		},
		{
			Title: "SendErr",
			Input: fakeCollectionLinkStatusFailed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			_, err := pageIterator.fetchPage(tt.Input.(string))

			assert.Error(t, err)
		})
	}
}

func TestPageIterator_FetchAndConvertPage(t *testing.T) {
	client := &mockClient{}

	pageIterator, err := NewPageIterator(TableCollectionResponse{}, client)
	assert.Nil(t, err)

	tests := []internal.Test[PageResult]{
		{
			Title:    "WithoutLink",
			Input:    fakeCollectionLinkNilResponse,
			Expected: PageResult{},
			Error:    core.ErrNilResponse,
		},
		{
			Title:    "LinkErrNilResponseBody",
			Input:    fakeCollectionLinkWithLinksErr,
			Expected: PageResult{},
			Error:    core.ErrNilResponseBody,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			page, err := pageIterator.fetchAndConvertPage(tt.Input.(string))
			assert.Equal(t, tt.Error, err)

			assert.Equal(t, tt.Expected, page)
		})
	}
}
