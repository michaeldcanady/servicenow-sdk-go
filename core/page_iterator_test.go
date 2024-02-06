package core

import (
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
	sharedPageIterator = &BasePageIterator[person, personCollectionResponse]{
		currentPage: PageResult[person]{
			Result:           sharedCurrentPage.Result,
			NextPageLink:     "",
			PreviousPageLink: "",
			LastPageLink:     "",
			FirstPageLink:    "",
		},
		client: sharedClient,
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
	tests := []test[*BasePageIterator[person, personCollectionResponse]]{
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
			expected:    (*BasePageIterator[person, personCollectionResponse])(nil),
			expectedErr: ErrNilClient,
		},
		{
			title:       "Missing Current Page",
			input:       []interface{}{(*personCollectionResponse)(nil), sharedClient},
			expected:    (*BasePageIterator[person, personCollectionResponse])(nil),
			expectedErr: ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			intInput := tt.input.([]interface{})
			pageIterator, err := NewPageIterator[person, personCollectionResponse](intInput[0].(*personCollectionResponse), intInput[1].(*mockClient))

			assert.ErrorIs(t, err, tt.expectedErr)

			assert.Equal(t, tt.expected, pageIterator)
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

			assert.Equal(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, page)
		})
	}
}
