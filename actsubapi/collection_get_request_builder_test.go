package actsubapi

import (
	"context"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCollectionGetRequestBuilder_Get(t *testing.T) {
	type testCase struct {
		name       string
		nilBuilder bool
		nilInner   bool
		config     *abstractions.RequestConfiguration[abstractions.DefaultQueryParameters]
		mockRes    interface{}
		mockErr    error
		expectErr  bool
	}

	tests := []testCase{
		{
			name:      "Success",
			mockRes:   core.NewBaseServiceNowCollectionResponse[*ActivitySubscriptionModel](CreateActivitySubscriptionModelFromDiscriminatorValue),
			expectErr: false,
		},
		{
			name:      "Error",
			mockErr:   assert.AnError,
			expectErr: true,
		},
		{
			name: "NilResponse",
		},
		{
			name:       "Nil builder",
			nilBuilder: true,
		},
		{
			name:     "Nil inner request builder",
			nilInner: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var builder *collectionGetRequestBuilder

			switch {
			case tc.nilBuilder:
				builder = nil
			case tc.nilInner:
				builder = &collectionGetRequestBuilder{nil}
			default:
				adapter := &mocking.MockRequestAdapter{}
				adapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.mockRes, tc.mockErr)
				builder = newCollectionGetRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter, activitiesURLTemplate)
			}

			resp, err := builder.Get(context.Background(), tc.config)

			switch {
			case tc.nilBuilder, tc.nilInner:
				assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
				assert.Nil(t, resp)
			case tc.expectErr:
				assert.Error(t, err)
				assert.Nil(t, resp)
			default:
				assert.NoError(t, err)
				if tc.mockRes != nil {
					assert.Equal(t, tc.mockRes, resp)
				} else {
					assert.Nil(t, resp)
				}
			}
		})
	}
}

func TestCollectionGetRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name       string
		nilBuilder bool
		nilInner   bool
	}{
		{name: "Success"},
		{name: "Nil builder", nilBuilder: true},
		{name: "Nil inner request builder", nilInner: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *collectionGetRequestBuilder

			switch {
			case tt.nilBuilder:
				builder = nil
			case tt.nilInner:
				builder = &collectionGetRequestBuilder{nil}
			default:
				adapter := &mocking.MockRequestAdapter{}
				builder = newCollectionGetRequestBuilder(map[string]string{"baseurl": "https://example.com"}, adapter, activitiesURLTemplate)
			}

			reqInfo, err := builder.ToGetRequestInformation(context.Background(), nil)

			if tt.nilBuilder || tt.nilInner {
				assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
				assert.Nil(t, reqInfo)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, reqInfo)
			assert.Equal(t, abstractions.GET, reqInfo.Method)
			assert.Equal(t, activitiesURLTemplate, reqInfo.UrlTemplate)
			assert.Equal(t, "https://example.com", reqInfo.PathParameters["baseurl"])
		})
	}
}
