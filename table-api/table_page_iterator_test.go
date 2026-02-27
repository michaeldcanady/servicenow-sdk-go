package tableapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTablePageIterator(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}

	res.On("GetResult").Return([]*TableRecord{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, err := NewTablePageIterator[*TableRecord](res, reqAdapter, CreateTableRecordFromDiscriminatorValue)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}

func TestNewDefaultTablePageIterator(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}

	res.On("GetResult").Return([]*TableRecord{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, err := NewDefaultTablePageIterator(res, reqAdapter)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}

func TestTablePageIterator_Iterate(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}

	record1 := &TableRecord{}
	record2 := &TableRecord{}

	res.On("GetResult").Return([]*TableRecord{record1}, nil)
	nextLink := "https://example.com/next"
	res.On("GetNextLink").Return(&nextLink, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	res2 := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}
	res2.On("GetResult").Return([]*TableRecord{record2}, nil)
	res2.On("GetNextLink").Return(nil, nil)
	res2.On("GetPreviousLink").Return(nil, nil)
	res2.On("GetFirstLink").Return(nil, nil)
	res2.On("GetLastLink").Return(nil, nil)

	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res2, nil)

	iterator, _ := NewTablePageIterator[*TableRecord](res, reqAdapter, CreateTableRecordFromDiscriminatorValue)

	var items []*TableRecord
	err := iterator.Iterate(context.Background(), false, func(item *TableRecord) bool {
		items = append(items, item)
		return true
	})

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, record1, items[0])
	assert.Equal(t, record2, items[1])
}

func TestTablePageIterator_NextItem(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}

	record1 := &TableRecord{}

	res.On("GetResult").Return([]*TableRecord{record1}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, _ := NewTablePageIterator[*TableRecord](res, reqAdapter, CreateTableRecordFromDiscriminatorValue)

	item, err := iterator.NextItem(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, record1, item)

	_, err = iterator.NextItem(context.Background())
	assert.ErrorIs(t, err, newInternal.ErrNoMoreItems)
}
