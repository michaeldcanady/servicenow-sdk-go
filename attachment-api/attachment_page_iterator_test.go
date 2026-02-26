package attachmentapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachmentPageIterator(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[Attachment]{}

	res.On("GetResult").Return([]Attachment{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, err := NewAttachmentPageIterator(res, reqAdapter, CreateAttachmentFromDiscriminatorValue)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}

func TestAttachmentPageIterator_Iterate(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[Attachment]{}

	attachment1 := &Attachment2Model{}
	attachment2 := &Attachment2Model{}

	res.On("GetResult").Return([]Attachment{attachment1, attachment2}, nil) // Mock single page with 2 items for simplicity in specialized test
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, _ := NewAttachmentPageIterator(res, reqAdapter, CreateAttachmentFromDiscriminatorValue)

	var items []Attachment
	err := iterator.Iterate(context.Background(), false, func(item Attachment) bool {
		items = append(items, item)
		return true
	})

	assert.NoError(t, err)
	assert.Len(t, items, 2)
	assert.Equal(t, attachment1, items[0])
	assert.Equal(t, attachment2, items[1])
}

func TestAttachmentPageIterator_NextItem(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[Attachment]{}

	attachment1 := &Attachment2Model{}

	res.On("GetResult").Return([]Attachment{attachment1}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	iterator, _ := NewAttachmentPageIterator(res, reqAdapter, CreateAttachmentFromDiscriminatorValue)

	item, err := iterator.NextItem(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, attachment1, item)

	_, err = iterator.NextItem(context.Background())
	assert.ErrorIs(t, err, newInternal.ErrNoMoreItems)
}
