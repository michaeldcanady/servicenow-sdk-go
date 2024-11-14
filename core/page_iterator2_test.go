package core

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/assert"
)

func constructPersonCollection(response *http.Response) (CollectionResponse[person], error) {
	resp := &personCollectionResponse{}

	err := internal.ParseResponse(response, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TestNewPageIterator2(t *testing.T) {
	tests := []test[*PageIterator2[person]]{
		{
			title:       "Valid",
			input:       []interface{}{sharedCurrentPage, sharedClient},
			expected:    sharedPageIterator2,
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title:       "Missing Client",
			input:       []interface{}{sharedCurrentPage, (*mockClient)(nil)},
			expected:    (*PageIterator2[person])(nil),
			expectedErr: ErrNilClient,
		},
		{
			title:       "Missing Current Page",
			input:       []interface{}{(*personCollectionResponse)(nil), sharedClient},
			expected:    (*PageIterator2[person])(nil),
			expectedErr: ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			intInput := tt.input.([]interface{})
			pageIterator, err := NewPageIterator2(intInput[0].(*personCollectionResponse), intInput[1].(*mockClient), constructPersonCollection)

			assert.ErrorIs(t, err, tt.expectedErr)

			if !internal.IsNil(pageIterator) {
				assert.Equal(t, tt.expected.client, pageIterator.client)
				assert.Equal(t, tt.expected.client, pageIterator.client)
				assert.Equal(t, tt.expected.currentPage, pageIterator.currentPage)
				assert.Equal(t, tt.expected.pauseIndex, pageIterator.pauseIndex)
			}
		})
	}
}

func TestPageIterator2_Iterate(t *testing.T) {
	var count int

	//nolint:dupl
	tests := []test[int]{
		{
			title:       "Missing Callback",
			input:       (func(*person) bool)(nil),
			expected:    0,
			shouldErr:   true,
			expectedErr: ErrNilCallback,
		},
		{
			title: "Callback false",
			input: func(person *person) bool {
				return false
			},
			expected:    0,
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Missing Response",
			setup: func() {
				sharedPageIterator2.currentPage.NextPageLink = fakeLinkNilResp
			},
			cleanup: func() {
				sharedPageIterator2.currentPage.NextPageLink = ""
			},
			input:       func(person *person) bool { return true },
			expected:    0,
			shouldErr:   true,
			expectedErr: ErrNilResponse,
		},
		{
			title: "Missing Next Link",
			input: func(person *person) bool {
				count += 1

				return true
			},
			expected:    len(sharedCurrentPage.Result),
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Single Page",
			setup: func() {
				sharedPageIterator2.currentPage.NextPageLink = fakeNextLink
			},
			input: func(person *person) bool {
				count += 1

				return true
			},
			expected:    len(sharedCurrentPage.Result) + len(sharedCurrentPage.Result),
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			count = 0
			sharedPageIterator2.pauseIndex = 0

			if tt.setup != nil {
				tt.setup()
			}

			callback := tt.input.(func(*person) bool)

			err := sharedPageIterator2.Iterate(callback, false)

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, count)

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

//nolint:dupl
func TestPageIterator2_enumerate(t *testing.T) {
	var count int

	tests := []test[[]interface{}]{
		{
			title: "All",
			input: func(item *person) bool {
				count += 1

				return true
			},
			expected:    []interface{}{len(sharedCurrentPage.Result), true},
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Once",
			input: func(item *person) bool {
				count += 1

				return false
			},
			expected:    []interface{}{1, false},
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title: "Missing Page items",
			setup: func() {
				sharedPageIterator2.currentPage.Result = nil
			},
			input:       func(pageItem *person) bool { return true },
			expected:    []interface{}{0, false},
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			count = 0                          // Reset the count between runs
			sharedPageIterator2.pauseIndex = 0 // Reset the pause index between runs

			if tt.setup != nil {
				tt.setup()
			}

			keepIterating := sharedPageIterator2.enumerate(tt.input.(func(item *person) bool))

			assert.ErrorIs(t, nil, tt.expectedErr)
			assert.Equal(t, tt.expected, []interface{}{count, keepIterating})
		})
	}
}

func TestPageIterator2_nextPage(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title:       "Normal",
			input:       false,
			expected:    PageResult[person]{},
			expectedErr: nil,
		},
		{
			title:       "Reversed",
			input:       true,
			expected:    PageResult[person]{},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			page, err := sharedPageIterator2.nextPage(tt.input.(bool))

			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, page)
		})
	}
}

func TestPageIterator2_Next(t *testing.T) {
	tests := []test[PageResult[person]]{
		// TODO: Needs "valid" test
		{
			title:       "Missing Next Link",
			expected:    PageResult[person]{},
			expectedErr: ErrEmptyURI,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			pageResult, err := sharedPageIterator2.Next()

			assert.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator2_Last(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title: "Valid",
			setup: func() {
				sharedPageIterator2.currentPage.LastPageLink = fakeLastLink
			},
			input:       nil,
			expected:    sharedPageResult,
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			pageResult, err := sharedPageIterator2.Last()

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator2_First(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title: "Valid",
			setup: func() {
				sharedPageIterator2.currentPage.FirstPageLink = fakeLastLink
			},
			input:       nil,
			expected:    sharedPageResult,
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			pageResult, err := sharedPageIterator2.First()

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator2_Previous(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title: "Valid",
			setup: func() {
				sharedPageIterator2.currentPage.PreviousPageLink = fakeLastLink
			},
			input:       nil,
			expected:    sharedPageResult,
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			pageResult, err := sharedPageIterator2.Previous()

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator2_fetchAndConvertPage(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title:       "Valid",
			input:       fakeLastLink,
			expected:    sharedPageResult,
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title:       "Missing URI",
			input:       "",
			expected:    PageResult[person]{},
			expectedErr: ErrEmptyURI,
		},
		{
			title:       "Nil Response",
			input:       fakeLinkNilResp,
			expected:    PageResult[person]{},
			shouldErr:   false,
			expectedErr: ErrNilResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			page, err := sharedPageIterator2.fetchAndConvertPage(tt.input.(string))

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, page)
		})
	}
}

func TestPageIterator2_fetchPage(t *testing.T) {
	tests := []test[CollectionResponse[person]]{
		{
			title:       "Valid",
			input:       fakeLastLink,
			expected:    sharedCurrentPage,
			expectedErr: nil,
		},
		{
			title:       "Missing URI",
			input:       "",
			expected:    nil,
			expectedErr: ErrEmptyURI,
		},
		{
			title:       "Nil Response",
			input:       fakeLinkNilResp,
			expected:    nil,
			expectedErr: ErrNilResponse,
		},
		{
			title:    "Bad Request URI",
			input:    "https://www.badrequesturi.com#fragment",
			expected: nil,
			expectedErr: &url.Error{
				Op:  "parse",
				URL: "https://www.badrequesturi.com#fragment",
				Err: url.InvalidHostError("#"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			collection, err := sharedPageIterator2.fetchPage(tt.input.(string))

			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, collection)
		})
	}
}
