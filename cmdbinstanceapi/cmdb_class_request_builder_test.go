package cmdbinstanceapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCmdbClassRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := core.NewBaseServiceNowCollectionResponse[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("query failed"))
			},
			expectedErr: errors.New("query failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCmdbClassRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test"}, adapter)
			resp, err := builder.Get(context.Background(), nil)

			if tt.expectedErr != nil {
				require.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestCmdbClassRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := core.NewBaseServiceNowItemResponse[CmdbInstance](CreateCmdbInstanceFromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("creation failed"))
			},
			expectedErr: errors.New("creation failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCmdbClassRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test"}, adapter)
			resp, err := builder.Post(context.Background(), nil, nil)

			if tt.expectedErr != nil {
				require.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

func TestCmdbClassRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*CmdbClassRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)

			itemResp, err := builder.Post(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, itemResp)
		})
	}
}

func TestCmdbClassRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewCmdbClassRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)

	itemResp, err := builder.Post(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, itemResp)
}

func TestCmdbClassRequestBuilder_ByID(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCmdbClassRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test"}, adapter)

	t.Run("ByID", func(t *testing.T) {
		itemBuilder := builder.ByID("sys_id_123")
		assert.NotNil(t, itemBuilder)
		assert.Equal(t, "sys_id_123", itemBuilder.GetPathParameters()["sys_id"])
	})
}

func TestCmdbClassRequestBuilder_ToGetRequestInformation(t *testing.T) {
	builder := NewCmdbClassRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test"}, &mocking.MockRequestAdapter{})

	t.Run("nil config", func(t *testing.T) {
		requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.GET, requestInfo.Method)
	})

	t.Run("config with headers, options and query parameters", func(t *testing.T) {
		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")
		query := "name=foo"
		limit := int32(10)
		offset := int32(0)

		option := mocking.NewMockRequestOption()
		option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "testOption"})

		config := &CmdbClassRequestBuilderGetRequestConfiguration{
			Headers: headers,
			Options: []abstractions.RequestOption{option},
			QueryParameters: &CmdbClassRequestBuilderGetQueryParameters{
				Query:  internal.ToPointer(query),
				Limit:  internal.ToPointer(limit),
				Offset: internal.ToPointer(offset),
			},
		}

		requestInfo, err := builder.ToGetRequestInformation(context.Background(), config)
		require.NoError(t, err)
		assert.Equal(t, []string{"value"}, requestInfo.Headers.Get("X-Test"))
		assert.Len(t, requestInfo.GetRequestOptions(), 1)
	})
}

func TestCmdbClassRequestBuilder_ToPostRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewCmdbClassRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test"}, adapter)

	t.Run("nil config and nil body", func(t *testing.T) {
		requestInfo, err := builder.ToPostRequestInformation(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.POST, requestInfo.Method)
	})

	t.Run("config with headers and options, and non-nil body", func(t *testing.T) {
		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		option := mocking.NewMockRequestOption()
		option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "testOption"})

		config := &CmdbClassRequestBuilderPostRequestConfiguration{
			Headers: headers,
			Options: []abstractions.RequestOption{option},
		}

		requestInfo, err := builder.ToPostRequestInformation(context.Background(), NewCmdbInstance(), config)
		require.NoError(t, err)
		assert.Equal(t, []string{"value"}, requestInfo.Headers.Get("X-Test"))
		assert.Len(t, requestInfo.GetRequestOptions(), 1)
	})
}
