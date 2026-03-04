package internal

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPageIterator(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}

func TestPageIterator_Iterate(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}
	item2 := &mocking.MockParsable{}

	res.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	nextLink := "https://example.com/next"
	res.On("GetNextLink").Return(&nextLink, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	res2 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res2.On("GetResult").Return([]*mocking.MockParsable{item2}, nil)
	res2.On("GetNextLink").Return(nil, nil)
	res2.On("GetPreviousLink").Return(nil, nil)
	res2.On("GetFirstLink").Return(nil, nil)
	res2.On("GetLastLink").Return(nil, nil)
	res2.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res2, nil)

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, func(serialization.ParseNode) (serialization.Parsable, error) { return nil, nil })

	var items []*mocking.MockParsable
	err := iterator.Iterate(context.Background(), false, func(item *mocking.MockParsable) bool {
		items = append(items, item)
		return true
	})

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, item1, items[0])
	assert.Equal(t, item2, items[1])
}

func TestPageIterator_HasNext(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	nextLink := "https://example.com/next"
	res.On("GetNextLink").Return(&nextLink, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	iterator, _ := NewPageIterator(res, reqAdapter, nil)

	assert.True(t, iterator.HasNext())

	res.On("GetNextLink").Unset()
	res.On("GetNextLink").Return(nil, nil)
	iterator, _ = NewPageIterator(res, reqAdapter, nil)
	assert.False(t, iterator.HasNext())
}

func TestPageIterator_NextItem(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}
	item2 := &mocking.MockParsable{}

	res.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	nextLink := "https://example.com/next"
	res.On("GetNextLink").Return(&nextLink, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	res2 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res2.On("GetResult").Return([]*mocking.MockParsable{item2}, nil)
	res2.On("GetNextLink").Return(nil, nil)
	res2.On("GetPreviousLink").Return(nil, nil)
	res2.On("GetFirstLink").Return(nil, nil)
	res2.On("GetLastLink").Return(nil, nil)
	res2.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res2, nil)

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, func(serialization.ParseNode) (serialization.Parsable, error) { return nil, nil })

	fetchedItem1, err := iterator.NextItem(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, item1, fetchedItem1)

	fetchedItem2, err := iterator.NextItem(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, item2, fetchedItem2)

	_, err = iterator.NextItem(context.Background())
	assert.ErrorIs(t, err, ErrNoMoreItems)
}

func TestPageIterator_Iterate_Reverse(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}
	item2 := &mocking.MockParsable{}

	// Page 2 (current)
	res.On("GetResult").Return([]*mocking.MockParsable{item2}, nil)
	res.On("GetNextLink").Return(nil, nil)
	prevLink := "https://example.com/prev"
	res.On("GetPreviousLink").Return(&prevLink, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	// Page 1 (previous)
	res1 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res1.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	res1.On("GetNextLink").Return(nil, nil)
	res1.On("GetPreviousLink").Return(nil, nil)
	res1.On("GetFirstLink").Return(nil, nil)
	res1.On("GetLastLink").Return(nil, nil)
	res1.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res1, nil)

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, func(serialization.ParseNode) (serialization.Parsable, error) { return nil, nil })

	assert.True(t, iterator.HasPrevious())

	var items []*mocking.MockParsable
	err := iterator.Iterate(context.Background(), true, func(item *mocking.MockParsable) bool {
		items = append(items, item)
		return true
	})

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, item2, items[0])
	assert.Equal(t, item1, items[1])
}

func TestPageIterator_Iterate_EmptyPage(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}
	item2 := &mocking.MockParsable{}

	// Page 1
	res.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	nextLink2 := "https://example.com/page2"
	res.On("GetNextLink").Return(&nextLink2, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	// Page 2 (Empty)
	res2 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res2.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	nextLink3 := "https://example.com/page3"
	res2.On("GetNextLink").Return(&nextLink3, nil)
	res2.On("GetPreviousLink").Return(nil, nil)
	res2.On("GetFirstLink").Return(nil, nil)
	res2.On("GetLastLink").Return(nil, nil)
	res2.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	// Page 3
	res3 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res3.On("GetResult").Return([]*mocking.MockParsable{item2}, nil)
	res3.On("GetNextLink").Return(nil, nil)
	res3.On("GetPreviousLink").Return(nil, nil)
	res3.On("GetFirstLink").Return(nil, nil)
	res3.On("GetLastLink").Return(nil, nil)
	res3.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	reqAdapter.On("Send", mock.Anything, mock.MatchedBy(func(ri *abstractions.RequestInformation) bool {
		uri, _ := ri.GetUri()
		return uri.String() == "https://example.com/page2"
	}), mock.Anything, mock.Anything).Return(res2, nil)

	reqAdapter.On("Send", mock.Anything, mock.MatchedBy(func(ri *abstractions.RequestInformation) bool {
		uri, _ := ri.GetUri()
		return uri.String() == "https://example.com/page3"
	}), mock.Anything, mock.Anything).Return(res3, nil)

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, func(serialization.ParseNode) (serialization.Parsable, error) { return nil, nil })

	var items []*mocking.MockParsable
	err := iterator.Iterate(context.Background(), false, func(item *mocking.MockParsable) bool {
		items = append(items, item)
		return true
	})

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, item1, items[0])
	assert.Equal(t, item2, items[1])
}

func TestPageIterator_Options(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	headers := abstractions.NewRequestHeaders()
	headers.Add("Test", "Value")

	option := &mocking.MockRequestOption{}

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil,
		WithHeaders[*mocking.MockParsable](headers),
		WithRequestOptions[*mocking.MockParsable](option),
	)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
	assert.Equal(t, headers, iterator.headers)
	assert.Contains(t, iterator.reqOptions, (abstractions.RequestOption)(option))
}

func TestPageIterator_PreviousItem(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}
	item2 := &mocking.MockParsable{}

	// Page 2 (current)
	res.On("GetResult").Return([]*mocking.MockParsable{item2}, nil)
	res.On("GetNextLink").Return(nil, nil)
	prevLink := "https://example.com/prev"
	res.On("GetPreviousLink").Return(&prevLink, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	// Page 1 (previous)
	res1 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res1.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	res1.On("GetNextLink").Return(nil, nil)
	res1.On("GetPreviousLink").Return(nil, nil)
	res1.On("GetFirstLink").Return(nil, nil)
	res1.On("GetLastLink").Return(nil, nil)
	res1.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res1, nil)

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)
	// Start at the end of page 2
	iterator.pauseIndex = 1

	fetchedItem2, err := iterator.PreviousItem(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, item2, fetchedItem2)

	fetchedItem1, err := iterator.PreviousItem(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, item1, fetchedItem1)

	_, err = iterator.PreviousItem(context.Background())
	assert.ErrorIs(t, err, ErrNoMoreItems)
}

func TestPageIterator_Reset(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}
	item2 := &mocking.MockParsable{}

	res.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	nextLink := "https://example.com/next"
	res.On("GetNextLink").Return(&nextLink, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	res2 := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res2.On("GetResult").Return([]*mocking.MockParsable{item2}, nil)
	res2.On("GetNextLink").Return(nil, nil)
	res2.On("GetPreviousLink").Return(nil, nil)
	res2.On("GetFirstLink").Return(nil, nil)
	res2.On("GetLastLink").Return(nil, nil)
	res2.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res2, nil)

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, func(serialization.ParseNode) (serialization.Parsable, error) { return nil, nil })

	// Advance to second page
	_, _ = iterator.NextItem(context.Background())
	_, _ = iterator.NextItem(context.Background())

	assert.Equal(t, 1, iterator.pauseIndex)

	iterator.Reset()

	assert.Equal(t, 0, iterator.pauseIndex)
	assert.Equal(t, 1, len(iterator.currentPage.Result))
	assert.Equal(t, item1, iterator.currentPage.Result[0])
}

func TestPageIterator_ResetPage(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}

	item1 := &mocking.MockParsable{}

	res.On("GetResult").Return([]*mocking.MockParsable{item1}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)
	res.On("ParseHeaders", mock.AnythingOfType("*abstractions.ResponseHeaders")).Return()

	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

	_, _ = iterator.NextItem(context.Background())
	assert.Equal(t, 1, iterator.pauseIndex)

	iterator.ResetPage()
	assert.Equal(t, 0, iterator.pauseIndex)
}
