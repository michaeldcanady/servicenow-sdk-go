package cmdbinstanceapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/v2/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/v2/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCmdbItemRequestBuilder_Get(t *testing.T) {
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
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("get failed"))
			},
			expectedErr: errors.New("get failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)
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

func TestCmdbItemRequestBuilder_Put(t *testing.T) {
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
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update failed"))
			},
			expectedErr: errors.New("update failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)
			resp, err := builder.Put(context.Background(), nil, nil)

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

func TestCmdbItemRequestBuilder_Patch(t *testing.T) {
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
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("patch failed"))
			},
			expectedErr: errors.New("patch failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)
			resp, err := builder.Patch(context.Background(), nil, nil)

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

func TestCmdbItemRequestBuilder_NilReceiverGuards(t *testing.T) {
	builders := map[string]*CmdbItemRequestBuilder{
		"nil builder":              nil,
		"nil inner RequestBuilder": {},
	}
	for name, builder := range builders {
		t.Run(name, func(t *testing.T) {
			resp, err := builder.Get(context.Background(), nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)

			resp, err = builder.Put(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)

			resp, err = builder.Patch(context.Background(), nil, nil)
			require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
			assert.Nil(t, resp)
		})
	}
}

func TestCmdbItemRequestBuilder_NilRequestAdapterGuards(t *testing.T) {
	builder := NewCmdbItemRequestBuilderInternal(map[string]string{}, nil)

	resp, err := builder.Get(context.Background(), nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)

	resp, err = builder.Put(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)

	resp, err = builder.Patch(context.Background(), nil, nil)
	require.ErrorIs(t, err, snerrors.ErrNilRequestAdapter)
	assert.Nil(t, resp)
}

func TestCmdbItemRequestBuilder_Relation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)

	t.Run("Relation", func(t *testing.T) {
		relationBuilder := builder.Relation()
		assert.NotNil(t, relationBuilder)
	})
}

func newTestRequestOption() *mocking.MockRequestOption {
	option := mocking.NewMockRequestOption()
	option.On("GetKey").Return(abstractions.RequestOptionKey{Key: "testOption"})
	return option
}

func TestCmdbItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, &mocking.MockRequestAdapter{})

	t.Run("nil config", func(t *testing.T) {
		requestInfo, err := builder.ToGetRequestInformation(context.Background(), nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.GET, requestInfo.Method)
	})

	t.Run("config with headers and options", func(t *testing.T) {
		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &CmdbItemRequestBuilderGetRequestConfiguration{
			Headers: headers,
			Options: []abstractions.RequestOption{newTestRequestOption()},
		}

		requestInfo, err := builder.ToGetRequestInformation(context.Background(), config)
		require.NoError(t, err)
		assert.Equal(t, []string{"value"}, requestInfo.Headers.Get("X-Test"))
		assert.Len(t, requestInfo.GetRequestOptions(), 1)
	})
}

func TestCmdbItemRequestBuilder_ToPutRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)

	t.Run("nil config and nil body", func(t *testing.T) {
		requestInfo, err := builder.ToPutRequestInformation(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.PUT, requestInfo.Method)
	})

	t.Run("config with headers and options, and non-nil body", func(t *testing.T) {
		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &CmdbItemRequestBuilderPutRequestConfiguration{
			Headers: headers,
			Options: []abstractions.RequestOption{newTestRequestOption()},
		}

		requestInfo, err := builder.ToPutRequestInformation(context.Background(), NewCmdbInstance(), config)
		require.NoError(t, err)
		assert.Equal(t, []string{"value"}, requestInfo.Headers.Get("X-Test"))
		assert.Len(t, requestInfo.GetRequestOptions(), 1)
	})
}

func TestCmdbItemRequestBuilder_ToPatchRequestInformation(t *testing.T) {
	adapter := &mocking.MockRequestAdapter{}
	adapter.On("GetSerializationWriterFactory").Return(jsonserialization.NewJsonSerializationWriterFactory())
	builder := NewCmdbItemRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "className": "test", "sys_id": "123"}, adapter)

	t.Run("nil config and nil body", func(t *testing.T) {
		requestInfo, err := builder.ToPatchRequestInformation(context.Background(), nil, nil)
		require.NoError(t, err)
		assert.Equal(t, abstractions.PATCH, requestInfo.Method)
	})

	t.Run("config with headers and options, and non-nil body", func(t *testing.T) {
		headers := abstractions.NewRequestHeaders()
		headers.Add("X-Test", "value")

		config := &CmdbItemRequestBuilderPatchRequestConfiguration{
			Headers: headers,
			Options: []abstractions.RequestOption{newTestRequestOption()},
		}

		requestInfo, err := builder.ToPatchRequestInformation(context.Background(), NewCmdbInstance(), config)
		require.NoError(t, err)
		assert.Equal(t, []string{"value"}, requestInfo.Headers.Get("X-Test"))
		assert.Len(t, requestInfo.GetRequestOptions(), 1)
	})
}
