package core

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewPageIterator(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	parsableFactory := mocking.NewMockParsableFactory().Factory

	validRes := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	validRes.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	validRes.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	validRes.On("GetNextLink").Return(nil, nil)
	validRes.On("GetPreviousLink").Return(nil, nil)
	validRes.On("GetFirstLink").Return(nil, nil)
	validRes.On("GetLastLink").Return(nil, nil)

	tests := []struct {
		name        string
		res         ServiceNowCollectionResponse[*mocking.MockParsable]
		reqAdapter  abstractions.RequestAdapter
		constructor serialization.ParsableFactory
		wantErr     bool
		errMsg      string
		errIs       error
	}{
		{
			name:        "Valid initialization",
			res:         validRes,
			reqAdapter:  reqAdapter,
			constructor: parsableFactory,
			wantErr:     false,
		},
		{
			name:        "Nil reqAdapter",
			res:         validRes,
			reqAdapter:  nil,
			constructor: parsableFactory,
			wantErr:     true,
			errMsg:      "requestAdapter cannot be nil",
			errIs:       snerrors.ErrNilRequestAdapter,
		},
		{
			name:        "Nil response",
			res:         nil,
			reqAdapter:  reqAdapter,
			constructor: parsableFactory,
			wantErr:     true,
			errMsg:      "response cannot be nil",
			errIs:       snerrors.ErrNilResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator, err := NewPageIterator(tt.res, tt.reqAdapter, tt.constructor)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, iterator)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
				if tt.errIs != nil {
					require.ErrorIs(t, err, tt.errIs)
				}
			} else {
				require.NoError(t, err)
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
			res.On("GetNextLink").Return(nil, nil)
			res.On("GetPreviousLink").Return(nil, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := mocking.NewMockRequestAdapter()
			constructor := mocking.NewMockParsableFactory().Factory

			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, constructor)
			require.NoError(t, err)

			err = iterator.Iterate(context.Background(), tt.reverse, func(_ *mocking.MockParsable) bool {
				return true
			})
			assert.NoError(t, err)
		})
	}
}

func TestPageIterator_NextItem(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], *mocking.MockRequestAdapter)
		wantErr   bool
		errIs     error
	}{
		{
			name: "Standard next item",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], _ *mocking.MockRequestAdapter) {
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
			},
			wantErr: false,
		},
		{
			name: "No more items",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], _ *mocking.MockRequestAdapter) {
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
			},
			wantErr: true,
			errIs:   ErrNoMoreItems,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			reqAdapter := mocking.NewMockRequestAdapter()
			tt.mockSetup(res, reqAdapter)

			constructor := mocking.NewMockParsableFactory().Factory
			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, constructor)
			require.NoError(t, err)

			item, err := iterator.NextItem(context.Background())
			if tt.wantErr {
				require.Error(t, err)
				if tt.errIs != nil {
					require.ErrorIs(t, err, tt.errIs)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, item)
			}
		})
	}
}

func TestPageIterator_PreviousItem(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], *mocking.MockRequestAdapter)
		setupIt   func(*PageIterator[*mocking.MockParsable])
		wantErr   bool
		errIs     error
	}{
		{
			name: "Standard previous item",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], _ *mocking.MockRequestAdapter) {
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
			},
			setupIt: func(pi *PageIterator[*mocking.MockParsable]) {
				pi.pauseIndex = 1
			},
			wantErr: false,
		},
		{
			name: "No previous items",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], _ *mocking.MockRequestAdapter) {
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
			},
			setupIt: func(pi *PageIterator[*mocking.MockParsable]) {
				pi.pauseIndex = 0
			},
			wantErr: true,
			errIs:   ErrNoMoreItems,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			reqAdapter := mocking.NewMockRequestAdapter()
			tt.mockSetup(res, reqAdapter)

			constructor := mocking.NewMockParsableFactory().Factory
			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, constructor)
			require.NoError(t, err)

			if tt.setupIt != nil {
				tt.setupIt(iterator)
			}

			item, err := iterator.PreviousItem(context.Background())
			if tt.wantErr {
				require.Error(t, err)
				if tt.errIs != nil {
					require.ErrorIs(t, err, tt.errIs)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, item)
			}
		})
	}
}

func TestPageIterator_Next(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], *mocking.MockRequestAdapter)
		navFunc   func(context.Context, *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error)
		wantRes   bool
	}{
		{
			name: "Next with link",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/next"
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(&link, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)

				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Next(ctx)
			},
			wantRes: true,
		},
		{
			name: "Next without link",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], _ *mocking.MockRequestAdapter) {
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Next(ctx)
			},
			wantRes: false,
		},
		{
			name: "Previous with link",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/prev"
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(&link, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)

				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Previous(ctx)
			},
			wantRes: true,
		},
		{
			name: "First with link",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/first"
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(&link, nil)
				res.On("GetLastLink").Return(nil, nil)

				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.First(ctx)
			},
			wantRes: true,
		},
		{
			name: "Last with link",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/last"
				res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(&link, nil)

				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Last(ctx)
			},
			wantRes: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			reqAdapter := &mocking.MockRequestAdapter{}
			tt.mockSetup(res, reqAdapter)

			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
			require.NoError(t, err)

			resp, err := tt.navFunc(context.Background(), iterator)
			require.NoError(t, err)
			if tt.wantRes {
				assert.NotEqual(t, PageResult[*mocking.MockParsable]{}, resp)
			} else {
				assert.Equal(t, PageResult[*mocking.MockParsable]{}, resp)
			}
		})
	}
}

func TestNewPageIterator_OptionError(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	optErr := errors.New("bad option")
	failingOption := func(*PageIterator[*mocking.MockParsable]) error {
		return optErr
	}

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil, failingOption)
	require.Error(t, err)
	assert.Equal(t, optErr, err)
	assert.Nil(t, iterator)
}

func TestPageIterator_Iterate_Errors(t *testing.T) {
	tests := []struct {
		name    string
		reverse bool
	}{
		{name: "forward propagates non-ErrNoMoreItems error", reverse: false},
		{name: "reverse propagates non-ErrNoMoreItems error", reverse: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := "https://example.com/page"
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
			res.On("GetNextLink").Return(&link, nil)
			res.On("GetPreviousLink").Return(&link, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := &mocking.MockRequestAdapter{}
			reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
				Return(nil, errors.New("fetch failed"))

			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
			require.NoError(t, err)

			called := false
			err = iterator.Iterate(context.Background(), tt.reverse, func(_ *mocking.MockParsable) bool {
				called = true
				return true
			})
			require.Error(t, err)
			assert.False(t, called)
		})
	}
}

func TestPageIterator_Iterate_CallbackStopsEarly(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable(), mocking.NewMockParsable()}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	count := 0
	err = iterator.Iterate(context.Background(), false, func(_ *mocking.MockParsable) bool {
		count++
		return false
	})
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestPageIterator_NextItem_PauseIndexNegative(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	iterator.pauseIndex = -1
	item, err := iterator.NextItem(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, 1, iterator.pauseIndex)
}

func TestPageIterator_NextItem_FetchError(t *testing.T) {
	link := "https://example.com/next"
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(&link, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := &mocking.MockRequestAdapter{}
	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("fetch failed"))

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	item, err := iterator.NextItem(context.Background())
	require.Error(t, err)
	assert.Equal(t, (*mocking.MockParsable)(nil), item)
}

func TestPageIterator_PreviousItem_PauseIndexOverflow(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	iterator.pauseIndex = 5
	item, err := iterator.PreviousItem(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, 0, iterator.pauseIndex)
}

func TestPageIterator_PreviousItem_FetchError(t *testing.T) {
	link := "https://example.com/prev"
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(&link, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := &mocking.MockRequestAdapter{}
	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil, errors.New("fetch failed"))

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	item, err := iterator.PreviousItem(context.Background())
	require.Error(t, err)
	assert.Equal(t, (*mocking.MockParsable)(nil), item)
}

func TestPageIterator_Previous_NoLink(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	page, err := iterator.Previous(context.Background())
	require.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, page)
}

func TestPageIterator_First_NoLink(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	page, err := iterator.First(context.Background())
	require.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, page)
}

func TestPageIterator_Last_NoLink(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	page, err := iterator.Last(context.Background())
	require.NoError(t, err)
	assert.Equal(t, PageResult[*mocking.MockParsable]{}, page)
}

func TestPageIterator_First_Last_Errors(t *testing.T) {
	tests := []struct {
		name      string
		navFunc   func(context.Context, *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error)
		mockSetup func(*mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], *mocking.MockRequestAdapter)
	}{
		{
			name: "First Send error",
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.First(ctx)
			},
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/first"
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(&link, nil)
				res.On("GetLastLink").Return(nil, nil)
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send error"))
			},
		},
		{
			name: "First convertToPage error",
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.First(ctx)
			},
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/first"
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(&link, nil)
				res.On("GetLastLink").Return(nil, nil)
				badRes := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				badRes.On("GetResult").Return([]*mocking.MockParsable{}, errors.New("get result error"))
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(badRes, nil)
			},
		},
		{
			name: "Last Send error",
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Last(ctx)
			},
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/last"
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(&link, nil)
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send error"))
			},
		},
		{
			name: "Last convertToPage error",
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Last(ctx)
			},
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/last"
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(nil, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(&link, nil)
				badRes := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				badRes.On("GetResult").Return([]*mocking.MockParsable{}, errors.New("get result error"))
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(badRes, nil)
			},
		},
		{
			name: "Previous Send error",
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Previous(ctx)
			},
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/prev"
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(&link, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send error"))
			},
		},
		{
			name: "Previous convertToPage error",
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Previous(ctx)
			},
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/prev"
				res.On("GetNextLink").Return(nil, nil)
				res.On("GetPreviousLink").Return(&link, nil)
				res.On("GetFirstLink").Return(nil, nil)
				res.On("GetLastLink").Return(nil, nil)
				badRes := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				badRes.On("GetResult").Return([]*mocking.MockParsable{}, errors.New("get result error"))
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(badRes, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{}, nil)

			reqAdapter := &mocking.MockRequestAdapter{}
			tt.mockSetup(res, reqAdapter)

			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
			require.NoError(t, err)

			_, err = tt.navFunc(context.Background(), iterator)
			assert.Error(t, err)
		})
	}
}

func TestPageIterator_FetchPage_URLParseError(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	badLink := "://not a valid url"
	_, err = iterator.fetchPage(context.Background(), &badLink)
	require.Error(t, err)
	assert.Equal(t, "parsing nextLink url failed", err.Error())
}

func TestPageIterator_FetchPage_NilLink(t *testing.T) {
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := mocking.NewMockRequestAdapter()
	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	page, err := iterator.fetchPage(context.Background(), nil)
	require.NoError(t, err)
	assert.Nil(t, page)
}

func TestPageIterator_FetchPage_NoExistingHeaderOption(t *testing.T) {
	link := "https://example.com/next"
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(&link, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := &mocking.MockRequestAdapter{}
	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	// Strip the HeadersInspectionOptions that NewPageIterator seeds by default,
	// forcing fetchPage to construct and append a new one.
	iterator.reqOptions = nil

	_, err = iterator.fetchPage(context.Background(), &link)
	require.NoError(t, err)
	assert.Len(t, iterator.reqOptions, 1)
}

func TestPageIterator_FetchPage_ExistingHeaderOption(t *testing.T) {
	link := "https://example.com/next"
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(&link, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	reqAdapter := &mocking.MockRequestAdapter{}
	reqAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(res, nil)

	iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
	require.NoError(t, err)

	// iterator already carries a HeadersInspectionOptions from NewPageIterator;
	// calling fetchPage a second time should reuse it instead of appending another.
	initialOptCount := len(iterator.reqOptions)
	_, err = iterator.fetchPage(context.Background(), &link)
	require.NoError(t, err)
	assert.Len(t, iterator.reqOptions, initialOptCount)
}

func TestPageIterator_Options(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	headers := abstractions.NewRequestHeaders()
	headers.Add("test", "header")
	options := []abstractions.RequestOption{
		&mocking.MockRequestOption{},
	}

	tests := []struct {
		name string
		run  func(*testing.T)
	}{
		{
			name: "NewPageIterator with options",
			run: func(t *testing.T) {
				iterator, err := NewPageIterator[*mocking.MockParsable](
					res,
					reqAdapter,
					nil,
					WithHeaders[*mocking.MockParsable](headers),
					WithRequestOptions[*mocking.MockParsable](options...),
				)
				require.NoError(t, err)
				assert.NotNil(t, iterator)
				assert.Equal(t, headers, iterator.headers)
				assert.Contains(t, iterator.reqOptions, options[0])
			},
		},
		{
			name: "SetHeaders and AddRequestOptions",
			run: func(t *testing.T) {
				iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)
				require.NoError(t, err)

				newHeaders := abstractions.NewRequestHeaders()
				newHeaders.Add("new", "header")
				iterator.SetHeaders(newHeaders)
				assert.Equal(t, newHeaders, iterator.headers)

				iterator.AddRequestOptions(options...)
				assert.Contains(t, iterator.reqOptions, options[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

func TestPageIterator_Reset(t *testing.T) {
	reqAdapter := mocking.NewMockRequestAdapter()
	res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
	res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
	res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
	res.On("GetNextLink").Return(nil, nil)
	res.On("GetPreviousLink").Return(nil, nil)
	res.On("GetFirstLink").Return(nil, nil)
	res.On("GetLastLink").Return(nil, nil)

	tests := []struct {
		name string
		run  func(*testing.T, *PageIterator[*mocking.MockParsable])
	}{
		{
			name: "Reset",
			run: func(t *testing.T, iterator *PageIterator[*mocking.MockParsable]) {
				_, _ = iterator.NextItem(context.Background())
				assert.Equal(t, 1, iterator.pauseIndex)

				iterator.Reset()
				assert.Equal(t, 0, iterator.pauseIndex)
			},
		},
		{
			name: "ResetPage",
			run: func(t *testing.T, iterator *PageIterator[*mocking.MockParsable]) {
				_, _ = iterator.NextItem(context.Background())
				assert.Equal(t, 1, iterator.pauseIndex)

				iterator.ResetPage()
				assert.Equal(t, 0, iterator.pauseIndex)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, nil)
			require.NoError(t, err)
			tt.run(t, iterator)
		})
	}
}

func TestPageIterator_Navigation_Errors(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(*mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], *mocking.MockRequestAdapter)
		navFunc   func(context.Context, *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error)
	}{
		{
			name: "Send error",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/next"
				res.On("GetNextLink").Return(&link, nil)
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send error"))
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Next(ctx)
			},
		},
		{
			name: "Wrong response type",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/next"
				res.On("GetNextLink").Return(&link, nil)
				// Return a mock that is not ServiceNowCollectionResponse
				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&mocking.MockParsable{}, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Next(ctx)
			},
		},

		{
			name: "convertToPage error",
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
				link := "https://example.com/next"
				res.On("GetNextLink").Return(&link, nil)

				// Return a response that will cause convertToPage to fail
				badRes := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				// Return an empty slice instead of nil to avoid panic in mock
				badRes.On("GetResult").Return([]*mocking.MockParsable{}, errors.New("get result error"))

				req.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(badRes, nil)
			},
			navFunc: func(ctx context.Context, pi *PageIterator[*mocking.MockParsable]) (PageResult[*mocking.MockParsable], error) {
				return pi.Next(ctx)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
			res.On("GetBackingStore").Return(mocking.NewMockBackingStore())
			res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
			res.On("GetPreviousLink").Return(nil, nil)
			res.On("GetFirstLink").Return(nil, nil)
			res.On("GetLastLink").Return(nil, nil)

			reqAdapter := &mocking.MockRequestAdapter{}
			tt.mockSetup(res, reqAdapter)

			iterator, err := NewPageIterator[*mocking.MockParsable](res, reqAdapter, mocking.NewMockParsableFactory().Factory)
			require.NoError(t, err)

			_, err = tt.navFunc(context.Background(), iterator)
			assert.Error(t, err)
		})
	}
}
