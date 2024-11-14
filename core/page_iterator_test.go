package core

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fakeLastLink    = "https://fake-link.com?last"
	fakeNextLink    = "https://fake-link.com?next_with_response"
	fakeLinkNilResp = "https://fake-link.com?nil_resp"
)

var (
	sharedClient      = &mockClient{}
	sharedCurrentPage = &personCollectionResponse{
		Result: []*person{
			{
				Name: "bob",
				Age:  25,
			},
			{
				Name: "steve",
				Age:  38,
			},
			{
				Name: "jerry",
				Age:  50,
			},
			{
				Name: "tom",
				Age:  18,
			},
			{
				Name: "mary",
				Age:  22,
			},
			{
				Name: "bill",
				Age:  30,
			},
		},
	}
	sharedPageIterator = &PageIterator[person, personCollectionResponse]{
		currentPage: PageResult[person]{
			Result:           sharedCurrentPage.Result,
			NextPageLink:     "",
			PreviousPageLink: "",
			LastPageLink:     "",
			FirstPageLink:    "",
		},
		client:     sharedClient,
		pauseIndex: 0,
		ctx:        context.Background(),
	}
	sharedPageIterator2 = &PageIterator2[person]{
		currentPage: PageResult[person]{
			Result:           sharedCurrentPage.Result,
			NextPageLink:     "",
			PreviousPageLink: "",
			LastPageLink:     "",
			FirstPageLink:    "",
		},
		client:          sharedClient,
		pauseIndex:      0,
		constructorFunc: constructPersonCollection,
	}
	sharedPageResult = PageResult[person]{
		Result:           sharedCurrentPage.Result,
		NextPageLink:     "",
		PreviousPageLink: "",
		LastPageLink:     "",
		FirstPageLink:    "",
	}
)

func sharedCurrentPageToJSON() []byte {
	data, _ := json.Marshal(sharedCurrentPage)

	return data
}

func TestNewPageIterator(t *testing.T) {
	tests := []test[*PageIterator[person, personCollectionResponse]]{
		{
			title:       "Valid",
			input:       []interface{}{sharedCurrentPage, sharedClient},
			expected:    sharedPageIterator,
			shouldErr:   false,
			expectedErr: nil,
		},
		{
			title:       "Missing Client",
			input:       []interface{}{sharedCurrentPage, (*mockClient)(nil)},
			expected:    (*PageIterator[person, personCollectionResponse])(nil),
			expectedErr: ErrNilClient,
		},
		{
			title:       "Missing Current Page",
			input:       []interface{}{(*personCollectionResponse)(nil), sharedClient},
			expected:    (*PageIterator[person, personCollectionResponse])(nil),
			expectedErr: ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			intInput := tt.input.([]interface{})
			pageIterator, err := NewPageIterator[person, personCollectionResponse](context.Background(), intInput[0].(*personCollectionResponse), intInput[1].(*mockClient))

			assert.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, tt.expected, pageIterator)
		})
	}
}

func TestPageIterator_Iterate(t *testing.T) {
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
				sharedPageIterator.currentPage.NextPageLink = fakeLinkNilResp
			},
			cleanup: func() {
				sharedPageIterator.currentPage.NextPageLink = ""
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
				sharedPageIterator.currentPage.NextPageLink = fakeNextLink
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
			sharedPageIterator.pauseIndex = 0

			if tt.setup != nil {
				tt.setup()
			}

			err := sharedPageIterator.Iterate(tt.input.(func(*person) bool))

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, count)

			if tt.cleanup != nil {
				tt.cleanup()
			}
		})
	}
}

//nolint:dupl
func TestPageIterator_enumerate(t *testing.T) {
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
				sharedPageIterator.currentPage.Result = nil
			},
			input:       func(pageItem *person) bool { return true },
			expected:    []interface{}{0, false},
			shouldErr:   false,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			count = 0                         // Reset the count between runs
			sharedPageIterator.pauseIndex = 0 // Reset the pause index between runs

			if tt.setup != nil {
				tt.setup()
			}

			keepIterating := sharedPageIterator.enumerate(tt.input.(func(item *person) bool))

			assert.ErrorIs(t, nil, tt.expectedErr)
			assert.Equal(t, tt.expected, []interface{}{count, keepIterating})
		})
	}
}

func TestPageIterator_Next(t *testing.T) {
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
			pageResult, err := sharedPageIterator.Next()

			assert.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator_Last(t *testing.T) {
	tests := []test[PageResult[person]]{
		{
			title: "Valid",
			setup: func() {
				sharedPageIterator.currentPage.LastPageLink = fakeLastLink
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

			pageResult, err := sharedPageIterator.Last()

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, pageResult)
		})
	}
}

func TestPageIterator_fetchAndConvertPage(t *testing.T) {
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
			page, err := sharedPageIterator.fetchAndConvertPage(tt.input.(string))

			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, page)
		})
	}
}
