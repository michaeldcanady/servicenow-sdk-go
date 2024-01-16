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

func TestPageIterator_Next(t *testing.T) {
	currentPage := TableCollectionResponse{
		NextPageLink: fakeCollectionLinkWithLinks,
	}

	client := &mockClient{}

	pageIterator, err := NewPageIterator(currentPage, client)
	assert.Nil(t, err)

	tests := []internal.Test[PageResult]{
		{
			Title:    "LinkNoError",
			Expected: expectedResult,
		},
	}

	for _, tt := range tests {
		page, err := pageIterator.next()

		assert.Equal(t, tt.Error, err)
		assert.Equal(t, tt.Expected, page)
	}
}

func TestPageIterator_Enumerate(t *testing.T) {
	pageIterator := PageIterator{
		currentPage: expectedResult,
	}

	enumCount := 0

	tests := []internal.Test[int]{
		{
			Title: "All",
			Input: func(item *TableEntry) bool {
				index := pageIterator.pauseIndex

				result := expectedResult.Result[index]

				assert.Equal(t, result, item)

				enumCount += 1

				return true
			},
			Expected: len(expectedResult.Result),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			enumCount = 0
			_ = pageIterator.enumerate(tt.Input.(func(item *TableEntry) bool))

			assert.Equal(t, tt.Expected, enumCount)
		})
	}
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

func TestPageIterator_Iterate(t *testing.T) {
	pageIterator := &PageIterator{
		currentPage: PageResult{
			// Initialize with test data
		},
		client:     &mockClient{},
		pauseIndex: 0,
	}

	tests := []internal.Test[any]{
		{
			Title: "NilCallback",
			Input: (func(pageItem *TableEntry) bool)(nil),
			Error: ErrNilCallback,
		},
		{
			Title: "SinglePageWithoutNextLinkWithoutCurrentPageWithCallback",
			Prepare: func() {
				pageIterator.currentPage = PageResult{}
			},
			Input: func(pageItem *TableEntry) bool {
				// Implement your callback logic for testing
				return true
			},
			Error: nil,
		},
		{
			Title: "SinglePageWithoutNextLinkWithCurrentPageWithCallback",
			Prepare: func() {
				pageIterator.currentPage.NextPageLink = ""
			},
			Input: func(pageItem *TableEntry) bool {
				// Implement your callback logic for testing
				return true
			},
			Error: nil,
		},
		{
			Title: "SinglePageWithNextLinkWithCurrentPageWithCallback",
			Prepare: func() {
				pageIterator.currentPage.NextPageLink = fakeCollectionLinkKey
			},
			Input: func(pageItem *TableEntry) bool {
				// Implement your callback logic for testing
				return true
			},
			Error: nil,
		},
		{
			Title: "MultiplePagesWithCallback",
			Prepare: func() {
				pageIterator.currentPage.NextPageLink = fakeCollectionLinkWithLinks
			},
			Input: func(pageItem *TableEntry) bool {
				// Implement your callback logic for testing
				return true
			},
			Error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Title, func(t *testing.T) {
			err := pageIterator.Iterate(tt.Input.(func(pageItem *TableEntry) bool))

			assert.Equal(t, tt.Error, err)
		})
	}
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
