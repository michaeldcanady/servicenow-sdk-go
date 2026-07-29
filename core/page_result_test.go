package core

import (
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertToPage(t *testing.T) {
	link := "https://example.com"

	tests := []struct {
		name     string
		response func() ServiceNowCollectionResponse[*mocking.MockParsable]
		wantErr  error
		wantPage bool
	}{
		{
			name: "nil response",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				return nil
			},
			wantErr: snerrors.ErrNilResponse,
		},
		{
			name: "GetResult error",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetResult").Return([]*mocking.MockParsable{}, errors.New("get result error"))
				return res
			},
			wantErr: errors.New("get result error"),
		},
		{
			name: "GetNextLink error",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(nil, errors.New("get next link error"))
				return res
			},
			wantErr: errors.New("get next link error"),
		},
		{
			name: "GetPreviousLink error",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(&link, nil)
				res.On("GetPreviousLink").Return(nil, errors.New("get previous link error"))
				return res
			},
			wantErr: errors.New("get previous link error"),
		},
		{
			name: "GetFirstLink error",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(&link, nil)
				res.On("GetPreviousLink").Return(&link, nil)
				res.On("GetFirstLink").Return(nil, errors.New("get first link error"))
				return res
			},
			wantErr: errors.New("get first link error"),
		},
		{
			name: "GetLastLink error",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetResult").Return([]*mocking.MockParsable{}, nil)
				res.On("GetNextLink").Return(&link, nil)
				res.On("GetPreviousLink").Return(&link, nil)
				res.On("GetFirstLink").Return(&link, nil)
				res.On("GetLastLink").Return(nil, errors.New("get last link error"))
				return res
			},
			wantErr: errors.New("get last link error"),
		},
		{
			name: "Successful",
			response: func() ServiceNowCollectionResponse[*mocking.MockParsable] {
				res := &mocking.MockServiceNowCollectionResponse[*mocking.MockParsable]{}
				res.On("GetResult").Return([]*mocking.MockParsable{mocking.NewMockParsable()}, nil)
				res.On("GetNextLink").Return(&link, nil)
				res.On("GetPreviousLink").Return(&link, nil)
				res.On("GetFirstLink").Return(&link, nil)
				res.On("GetLastLink").Return(&link, nil)
				return res
			},
			wantPage: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page, err := convertToPage(tt.response())
			if tt.wantErr != nil {
				require.Error(t, err)
				if errors.Is(tt.wantErr, snerrors.ErrNilResponse) {
					require.ErrorIs(t, err, snerrors.ErrNilResponse)
				} else {
					require.Equal(t, tt.wantErr, err)
				}
				return
			}
			require.NoError(t, err)
			if tt.wantPage {
				assert.Len(t, page.Result, 1)
				assert.Equal(t, &link, page.NextLink)
				assert.Equal(t, &link, page.PrevLink)
				assert.Equal(t, &link, page.FirstLink)
				assert.Equal(t, &link, page.LastLink)
			}
		})
	}
}
