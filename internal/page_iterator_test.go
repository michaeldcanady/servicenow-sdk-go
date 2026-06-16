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
	tests := []struct {
		name        string
		res         ServiceNowCollectionResponse[*mocking.MockParsable]
		reqAdapter  abstractions.RequestAdapter
		constructor serialization.ParsableFactory
		expectedErr bool
	}{
		{
			name: "Valid initialization",
			res: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
				return res
			}(),
			reqAdapter:  mocking.NewMockRequestAdapter(),
			constructor: mocking.NewMockParsableFactory().Factory,
			expectedErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			iterator, err := NewPageIterator(test.res, test.reqAdapter, test.constructor)
			if test.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, iterator)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, iterator)
			}
		})
	}
}

func TestPageIterator_Iterate(t *testing.T) {
	tests := []struct {
		name    string
		reverse bool
	}{
		{
			name:    "Standard forward iterate",
			reverse: false,
		},
		{
			name:    "Reverse iterate",
			reverse: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
			res.On("GetNextLink").Return(nil, nil)
			res.On("GetPreviousLink").Return(nil, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := mocking.NewMockRequestAdapter()
			constructor := mocking.NewMockParsableFactory().Factory

			iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, constructor)

			err := iterator.Iterate(context.Background(), test.reverse, func(item *mocking.MockParsable) bool {
				return true
			})
			assert.NoError(t, err)
		})
	}
}

func TestPageIterator_NextItem(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Standard next item",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
			res.On("GetNextLink").Return(nil, nil)
			res.On("GetPreviousLink").Return(nil, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := mocking.NewMockRequestAdapter()
			constructor := mocking.NewMockParsableFactory().Factory

			iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, constructor)

			item, err := iterator.NextItem(context.Background())
			assert.NoError(t, err)
			assert.NotNil(t, item)
		})
	}
}

func TestNewPageIterator_Errors(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()

	// Nil reqAdapter
	iterator, err := NewPageIterator[*mocking.MockParsable](nil, nil, nil)
	assert.Error(t, err)
	assert.Nil(t, iterator)

	// Nil response
	iterator, err = NewPageIterator[*mocking.MockParsable](nil, reqAdapter, nil)
	assert.Error(t, err)
	assert.Nil(t, iterator)
}

func TestPageIterator_Navigation_EmptyLinks(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

	resp, err := iterator.Next(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, resp)

	resp, err = iterator.Previous(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, resp)

	resp, err = iterator.First(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, resp)

	resp, err = iterator.Last(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, resp)
}

func TestNewPageIterator_WithOptions(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	headers := abstractions.NewRequestHeaders()
	headers.Add("test", "header")
	options := []abstractions.RequestOption{}

	iterator, err := NewPageIterator[*mocking.MockParsable](
		res,
		reqAdapter,
		nil,
		WithHeaders[*mocking.MockParsable](headers),
		WithRequestOptions[*mocking.MockParsable](options...),
	)

	assert.NoError(t, err)
	assert.NotNil(t, iterator)
	assert.Equal(t, headers, iterator.headers)
}

func TestPageIterator_Options(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

	iterator.SetHeaders(abstractions.NewRequestHeaders())
	iterator.AddRequestOptions([]abstractions.RequestOption{}...)
}

func TestPageIterator_Navigation(t *testing.T) {
	tests := []struct {
		name     string
		navFunc  func(context.Context, *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error)
		mockLink string
	}{
		{"Next", func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
			return pi.Next(ctx)
		}, "next"},
		{"Previous", func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
			return pi.Previous(ctx)
		}, "prev"},
		{"First", func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
			return pi.First(ctx)
		}, "first"},
		{"Last", func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
			return pi.Last(ctx)
		}, "last"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{}, nil)

			link := "https://example.com/" + tt.mockLink
			res.On("GetNextLink").Return(&link, nil)
			res.On("GetPreviousLink").Return(&link, nil)
			res.On("GetFirstLink").Return(&link, nil)
			res.On("GetLastLink").Return(&link, nil)

			reqAdapter := &mocking.MockRequestAdapter{}
			reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)

			iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)

			resp, err := tt.navFunc(context.Background(), iterator)
			assert.NoError(t, err)
			assert.NotNil(t, resp)
		})
	}
}

func TestPageIterator_Iterate(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable(), mocking.NewMockParsable()}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

	count := 0
	err := iterator.Iterate(context.Background(), func(item *mocking.MockParsable) bool {
		count++
		return true
	})

	assert.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestPageIterator_Reset(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Reset iterator",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
			res.On("GetNextLink").Return(nil, nil)
			res.On("GetPreviousLink").Return(nil, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := mocking.NewMockRequestAdapter()

			iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

			_, _ = iterator.NextItem(context.Background())
			assert.Equal(t, 1, iterator.pauseIndex)

			iterator.Reset()
			assert.Equal(t, 0, iterator.pauseIndex)
		})
	}
}

func TestPageIterator_ResetPage(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Reset page",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
			res.On("GetNextLink").Return(nil, nil)
			res.On("GetPreviousLink").Return(nil, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := mocking.NewMockRequestAdapter()

			iterator, _ := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)

			_, _ = iterator.NextItem(context.Background())
			assert.Equal(t, 1, iterator.pauseIndex)

			iterator.ResetPage()
			assert.Equal(t, 0, iterator.pauseIndex)
		})
	}
}
