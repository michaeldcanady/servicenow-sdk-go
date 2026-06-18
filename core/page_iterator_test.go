package core

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
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
			errMsg:      "reqAdapter can't be nil",
		},
		{
			name:        "Nil response",
			res:         nil,
			reqAdapter:  reqAdapter,
			constructor: parsableFactory,
			wantErr:     true,
			errMsg:      "response cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator, err := NewPageIterator(tt.res, tt.reqAdapter, tt.constructor)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, iterator)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
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

			err = iterator.Iterate(context.Background(), tt.reverse, func(item *mocking.MockParsable) bool {
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
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
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
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
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
				assert.Error(t, err)
				if tt.errIs != nil {
					assert.ErrorIs(t, err, tt.errIs)
				}
			} else {
				assert.NoError(t, err)
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
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
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
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
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
				assert.Error(t, err)
				if tt.errIs != nil {
					assert.ErrorIs(t, err, tt.errIs)
				}
			} else {
				assert.NoError(t, err)
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
			mockSetup: func(res *mocking.MockServiceNowCollectionResponse[*mocking.MockParsable], req *mocking.MockRequestAdapter) {
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
			assert.NoError(t, err)
			if tt.wantRes {
				assert.NotEqual(t, PageResult[*mocking.MockParsable]{}, resp)
			} else {
				assert.Equal(t, PageResult[*mocking.MockParsable]{}, resp)
			}
		})
	}
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
				assert.NoError(t, err)
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
