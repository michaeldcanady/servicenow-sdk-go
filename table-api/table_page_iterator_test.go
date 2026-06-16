package tableapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewTablePageIterator(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*TableRecord{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()

	iterator, err := NewTablePageIterator(res, reqAdapter, CreateTableRecordFromDiscriminatorValue)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}

func TestNewDefaultTablePageIterator(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*TableRecord]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*TableRecord{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()

	iterator, err := NewDefaultTablePageIterator(res, reqAdapter)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}
