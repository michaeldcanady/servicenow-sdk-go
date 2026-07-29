package documentsapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestActionRequestBuilder_Document(t *testing.T) {
	builder := NewActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, &mocking.MockRequestAdapter{})

	child := builder.Document("doc-id")

	require.NotNil(t, child)
	assert.Equal(t, "doc-id", child.GetPathParameters()["documentSysId"])
	assert.Equal(t, actionURLTemplate, child.GetURLTemplate())
}

func TestDocumentActionRequestBuilder_Version(t *testing.T) {
	builder := NewDocumentActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com", "documentSysId": "doc-id"}, &mocking.MockRequestAdapter{})

	child := builder.Version("version-id")

	require.NotNil(t, child)
	assert.Equal(t, "version-id", child.GetPathParameters()["versionSysId"])
	assert.Equal(t, actionURLTemplate, child.GetURLTemplate())
}

func TestVersionActionRequestBuilder_Patch_NilBuilder(t *testing.T) {
	var builder *VersionActionRequestBuilder

	err := builder.Patch(context.Background(), nil)

	require.ErrorIs(t, err, snerrors.ErrNilRequestBuilder)
}

func TestVersionActionRequestBuilder_Patch(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(*mocking.MockRequestAdapter)
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			setupMock: func(m *mocking.MockRequestAdapter) {
				m.On("SendNoContent", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("patch failed"))
			},
			expectedErr: errors.New("patch failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapter := &mocking.MockRequestAdapter{}
			tt.setupMock(adapter)

			builder := NewVersionActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)
			err := builder.Patch(context.Background(), nil)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestVersionActionRequestBuilder_ToPatchRequestInformation(t *testing.T) {
	tests := []struct {
		name            string
		requestConfig   *VersionActionRequestBuilderPatchRequestConfiguration
		expectedHeaders map[string][]string
	}{
		{
			name:          "nil config",
			requestConfig: nil,
			expectedHeaders: map[string][]string{
				"accept": {"application/json"},
			},
		},
		{
			name: "with custom headers",
			requestConfig: &VersionActionRequestBuilderPatchRequestConfiguration{
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
			name: "with request options",
			requestConfig: &VersionActionRequestBuilderPatchRequestConfiguration{
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
			builder := NewVersionActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

			reqInfo, err := builder.ToPatchRequestInformation(context.Background(), tt.requestConfig)

			require.NoError(t, err)
			assert.NotNil(t, reqInfo)
			assert.Equal(t, abstractions.PATCH, reqInfo.Method)
			for k, v := range tt.expectedHeaders {
				assert.Equal(t, v, reqInfo.Headers.Get(k))
			}
		})
	}

	t.Run("with data", func(t *testing.T) {
		adapter := &mocking.MockRequestAdapter{}
		writer := mocking.NewMockSerializationWriter()
		writer.On("WriteObjectValue", "", mock.Anything, mock.Anything).Return(nil)
		writer.On("Close").Return(nil)
		writer.On("GetSerializedContent").Return([]byte("{}"), nil)

		factory := mocking.NewMockSerializationWriterFactory()
		factory.On("GetSerializationWriter", "application/json").Return(writer, nil)
		adapter.On("GetSerializationWriterFactory").Return(factory, nil)

		builder := NewVersionActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, adapter)

		reqInfo, err := builder.ToPatchRequestInformation(context.Background(), &VersionActionRequestBuilderPatchRequestConfiguration{
			Data: NewDocument(),
		})

		require.NoError(t, err)
		assert.NotNil(t, reqInfo)
	})

	t.Run("with data serialization error", func(t *testing.T) {
		builder := NewVersionActionRequestBuilderInternal(map[string]string{"baseurl": "https://example.com"}, nil)

		reqInfo, err := builder.ToPatchRequestInformation(context.Background(), &VersionActionRequestBuilderPatchRequestConfiguration{
			Data: NewDocument(),
		})

		require.Error(t, err)
		assert.Nil(t, reqInfo)
	})
}
