package internal

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
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
