package internal

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
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
