package documentsapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDocumentGetRequestBuilder_Get(t *testing.T) {
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
				m.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("get failed"))
			},
			expectedErr: errors.New("get failed"),
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
			var builder *documentGetRequestBuilder

			if tt.nilBuilder {
				builder = nil
			} else {
				adapter := &mocking.MockRequestAdapter{}
				tt.setupMock(adapter)
				builder = newDocumentGetRequestBuilder(adapter, versionStateURLTemplate, map[string]string{"baseurl": "https://example.com"})
			}

			resp, err := builder.get(context.Background(), &documentGetRequestConfiguration{})

			if tt.nilBuilder {
				assert.NoError(t, err)
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

func TestDocumentGetRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name            string
		requestConfig   *documentGetRequestConfiguration
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
			requestConfig: &documentGetRequestConfiguration{
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
			requestConfig: &documentGetRequestConfiguration{
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
			builder := newDocumentGetRequestBuilder(adapter, versionStateURLTemplate, map[string]string{"baseurl": "https://example.com"})

			reqInfo, err := builder.toGetRequestInformation(context.Background(), tt.requestConfig)

			assert.NoError(t, err)
			assert.NotNil(t, reqInfo)
			assert.Equal(t, abstractions.GET, reqInfo.Method)
			for k, v := range tt.expectedHeaders {
				assert.Equal(t, v, reqInfo.Headers.Get(k))
			}
		})
	}

	t.Run("Nil builder", func(t *testing.T) {
		var builder *documentGetRequestBuilder
		reqInfo, err := builder.toGetRequestInformation(context.Background(), nil)
		assert.NoError(t, err)
		assert.Nil(t, reqInfo)
	})
}
