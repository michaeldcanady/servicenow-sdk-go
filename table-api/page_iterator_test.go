package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/stretchr/testify/assert"
)

func TestNewPageIterator(t *testing.T) {
	tests := []struct {
		name        string
		currentPage interface{}
		client      core.Client
		expectedErr error
	}{
		{
			name: "Valid with client",
			currentPage: TableCollectionResponse{
				Result:           []*TableEntry{&fakeEntry},
				FirstPageLink:    fakeFirstLink,
				LastPageLink:     fakeLastLink,
				NextPageLink:     fakeNextLink,
				PreviousPageLink: fakePrevLink,
			},
			client:      &mockClient{},
			expectedErr: nil,
		},
		{
			name: "Without client",
			currentPage: TableCollectionResponse{
				Result: []*TableEntry{&fakeEntry},
			},
			client:      nil,
			expectedErr: ErrNilClient,
		},
		{
			name:        "Nil response with client",
			currentPage: nil,
			client:      &mockClient{},
			expectedErr: ErrNilResponse,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iterator, err := NewPageIterator(test.currentPage, test.client)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr, err)
				assert.Nil(t, iterator)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, iterator)
				assert.Equal(t, &expectedIterator, iterator)
			}
		})
	}
}

func TestPageIteratorNext(t *testing.T) {
	tests := []struct {
		name           string
		currentPage    TableCollectionResponse
		client         core.Client
		expectedResult PageResult
		expectedErr    error
	}{
		{
			name: "With link no error",
			currentPage: TableCollectionResponse{
				NextPageLink: fakeCollectionLinkWithLinks,
			},
			client:         &mockClient{},
			expectedResult: expectedResult,
			expectedErr:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iterator, err := NewPageIterator(test.currentPage, test.client)
			assert.NoError(t, err)

			page, err := iterator.next()
			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.expectedResult, page)
		})
	}
}

func TestPageIteratorEnumerate(t *testing.T) {
	tests := []struct {
		name           string
		iterator       PageIterator
		stopAfter      int
		expectedIter   int
		expectedStatus bool
	}{
		{
			name: "Enumerate all",
			iterator: PageIterator{
				currentPage: expectedResult,
			},
			stopAfter:      -1,
			expectedIter:   len(expectedResult.Result),
			expectedStatus: true,
		},
		{
			name: "Enumerate once",
			iterator: PageIterator{
				currentPage: expectedResult,
			},
			stopAfter:      1,
			expectedIter:   1,
			expectedStatus: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			count := 0
			keepIterating := test.iterator.enumerate(func(item *TableEntry) bool {
				count++
				if test.stopAfter > 0 && count >= test.stopAfter {
					return false
				}
				return true
			})
			assert.Equal(t, test.expectedStatus, keepIterating)
			assert.Equal(t, test.expectedIter, count)
		})
	}
}

func TestPageIteratorIterate(t *testing.T) {
	tests := []struct {
		name        string
		iterator    *PageIterator
		callback    func(pageItem *TableEntry) bool
		expectedErr error
	}{
		{
			name: "No callback",
			iterator: &PageIterator{
				currentPage: PageResult{},
				client:      &mockClient{},
			},
			callback:    nil,
			expectedErr: ErrNilCallback,
		},
		{
			name: "Single page without next link",
			iterator: &PageIterator{
				currentPage: PageResult{},
				client:      &mockClient{},
			},
			callback: func(item *TableEntry) bool {
				return true
			},
			expectedErr: nil,
		},
		{
			name: "Multiple pages with callback",
			iterator: &PageIterator{
				currentPage: PageResult{
					NextPageLink: fakeCollectionLinkWithLinks,
				},
				client: &mockClient{},
			},
			callback: func(item *TableEntry) bool {
				return true
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.iterator.Iterate(test.callback)
			if test.expectedErr != nil {
				assert.ErrorIs(t, err, test.expectedErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPageIteratorLast(t *testing.T) {
	tests := []struct {
		name        string
		iterator    *PageIterator
		expectedErr bool
	}{
		{
			name: "Returns error",
			iterator: &PageIterator{
				currentPage: PageResult{
					LastPageLink: fakeNextLink,
				},
				client: &mockClient{},
			},
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.iterator.Last()
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPageIteratorFetchPage(t *testing.T) {
	tests := []struct {
		name        string
		iterator    *PageIterator
		uri         string
		expectedErr error
	}{
		{
			name: "Send error",
			iterator: &PageIterator{
				currentPage: PageResult{},
				client:      &mockClient{},
			},
			uri:         fakeCollectionLinkStatusFailed,
			expectedErr: assert.AnError, // any error
		},
		{
			name: "Empty URI",
			iterator: &PageIterator{
				currentPage: PageResult{},
				client:      &mockClient{},
			},
			uri:         "",
			expectedErr: ErrEmptyURI,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.iterator.fetchPage(test.uri)
			if test.expectedErr != nil {
				if test.expectedErr == assert.AnError {
					assert.Error(t, err)
				} else {
					assert.ErrorIs(t, err, test.expectedErr)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPageIteratorFetchAndConvertPage(t *testing.T) {
	tests := []struct {
		name        string
		nextLink    string
		expectedErr error
	}{
		{
			name:        "Nil response body error",
			nextLink:    fakeCollectionLinkWithLinksErr,
			expectedErr: core.ErrNilResponseBody,
		},
		{
			name:        "Nil response error",
			nextLink:    fakeCollectionLinkNilResponse,
			expectedErr: core.ErrNilResponse,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iterator, err := NewPageIterator(TableCollectionResponse{NextPageLink: test.nextLink}, &mockClient{})
			assert.NoError(t, err)

			page, err := iterator.fetchAndConvertPage(test.nextLink)
			assert.ErrorIs(t, err, test.expectedErr)
			assert.Equal(t, PageResult{}, page)
		})
	}
}
