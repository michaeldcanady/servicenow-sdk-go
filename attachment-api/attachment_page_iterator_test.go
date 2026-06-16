package attachmentapi

import (
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachmentPageIterator(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*Attachment]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*Attachment{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()

	iterator, err := NewAttachmentPageIterator(res, reqAdapter)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
}
