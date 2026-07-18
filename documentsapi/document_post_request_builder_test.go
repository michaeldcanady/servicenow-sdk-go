package documentsapi

import (
	"context"
	"errors"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDocumentPostRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		nilBuilder  bool
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				resp := core.NewBaseServiceNowItemResponse[Document](CreateDocumentFromDiscriminatorValue)
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
			},
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("post failed"))
			},
			expectedErr: errors.New("post failed"),
		},
		{
			name: "Nil response",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
		},
		{
			name:       "Nil builder",
			nilBuilder: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *documentPostRequestBuilder

			if tt.nilBuilder {
				builder = nil
			} else {
				adapter := &mocking.MockRequestAdapter{}
				tt.setupMock(adapter)
				builder = newDocumentPostRequestBuilder(adapter, attachURLTemplate, map[string]string{"baseurl": "https://example.com"})
			}

			resp, err := builder.post(context.Background(), &documentPostRequestConfiguration{})

			if tt.nilBuilder {
				assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
				assert.Nil(t, resp)
				return
			}

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, resp)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestDocumentPostRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name            string
		requestConfig   *documentPostRequestConfiguration
		expectedHeaders map[string][]string
	}{
		{
			name:          "Nil config",
			requestConfig: nil,
			expectedHeaders: map[string][]string{
				"accept": {"application/json"},
			},
		},
		{
			name: "With custom headers",
			requestConfig: &documentPostRequestConfiguration{
				Headers: func() *abstractions.RequestHeaders {
					h := abstractions.NewRequestHeaders()
					h.Add("custom-header", "custom-value")
					return h
				}(),
			},
			expectedHeaders: map[string][]string{
				"accept":        {"application/json"},
				"custom-header": {"custom-value"},
			},
		},
		{
			name: "With request options",
			requestConfig: &documentPostRequestConfiguration{
				Options: func() []abstractions.RequestOption {
					opt := mocking.NewMockRequestOption()
					opt.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})
					return []abstractions.RequestOption{opt}
				}(),
			},
			expectedHeaders: map[string][]string{
				"accept": {"application/json"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			builder := newDocumentPostRequestBuilder(adapter, attachURLTemplate, map[string]string{"baseurl": "https://example.com"})

			reqInfo, err := builder.toPostRequestInformation(context.Background(), tt.requestConfig)

			assert.NoError(t, err)
			assert.NotNil(t, reqInfo)
			assert.Equal(t, abstractions.POST, reqInfo.Method)
			for k, v := range tt.expectedHeaders {
				assert.Equal(t, v, reqInfo.Headers.Get(k))
			}
		})
	}

	t.Run("With data", func(t *testing.T) {
		adapter := &mocking.MockRequestAdapter{}
		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(nil)
		writer.On("Close").Return(nil)
		writer.On("GetSerializedContent").Return([]byte("{}"), nil)

		factory := mocking.NewMockSerializationWriterFactory()
		factory.On("GetSerializationWriter", "application/json").Return(writer, nil)
		adapter.On("GetSerializationWriterFactory").Return(factory, nil)

		builder := newDocumentPostRequestBuilder(adapter, attachURLTemplate, map[string]string{"baseurl": "https://example.com"})

		reqInfo, err := builder.toPostRequestInformation(context.Background(), &documentPostRequestConfiguration{
			Data: NewDocument(),
		})

		assert.NoError(t, err)
		assert.NotNil(t, reqInfo)
	})

	t.Run("Nil builder", func(t *testing.T) {
		var builder *documentPostRequestBuilder
		reqInfo, err := builder.toPostRequestInformation(context.Background(), nil)
		assert.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
		assert.Nil(t, reqInfo)
	})
}
