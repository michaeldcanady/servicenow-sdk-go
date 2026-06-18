package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAttachmentUploadRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		requestAdapter abstractions.RequestAdapter
	}{
		{
			name:           "Successful",
			pathParameters: map[string]string{},
			requestAdapter: mocking.NewMockRequestAdapter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentUploadRequestBuilderInternal(tt.pathParameters, tt.requestAdapter)

			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentUploadRequestBuilder{}, builder)
			assert.IsType(t, &core.BaseRequestBuilder{}, builder.RequestBuilder)
			assert.Equal(t, tt.pathParameters, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestNewAttachmentUploadRequestBuilder(t *testing.T) {
	tests := []struct {
		name           string
		rawURL         string
		requestAdapter abstractions.RequestAdapter
	}{
		{
			name:           "Successful",
			rawURL:         "",
			requestAdapter: mocking.NewMockRequestAdapter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentUploadRequestBuilder(tt.rawURL, tt.requestAdapter)

			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentUploadRequestBuilder{}, builder)
			assert.IsType(t, &core.BaseRequestBuilder{}, builder.RequestBuilder)
			assert.Equal(t, map[string]string{internal.RawURLKey: tt.rawURL}, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestAttachmentUploadRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody)
		expectedErr bool
		nilRes      bool
	}{
		{
			name: "Successful",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
				_ = body.AddOrReplacePart("table_name", "text/plain", &s)
				_ = body.AddOrReplacePart("table_sys_id", "text/plain", &s)
				_ = body.AddOrReplacePart("uploadFile", "text/plain", &s)

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", mock.Anything).Return(mockSerializationWriter, nil)

				ra.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&File{}, nil)
			},
			expectedErr: false,
		},
		{
			name:        "Missing Content-Type",
			setup:       func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {},
			expectedErr: true,
		},
		{
			name: "Missing table_name",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
			},
			expectedErr: true,
		},
		{
			name: "Missing table_sys_id",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
				_ = body.AddOrReplacePart("table_name", "text/plain", &s)
			},
			expectedErr: true,
		},
		{
			name: "Missing uploadFile",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
				_ = body.AddOrReplacePart("table_name", "text/plain", &s)
				_ = body.AddOrReplacePart("table_sys_id", "text/plain", &s)
			},
			expectedErr: true,
		},
		{
			name: "Send Error",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
				_ = body.AddOrReplacePart("table_name", "text/plain", &s)
				_ = body.AddOrReplacePart("table_sys_id", "text/plain", &s)
				_ = body.AddOrReplacePart("uploadFile", "text/plain", &s)

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", mock.Anything).Return(mockSerializationWriter, nil)

				ra.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("send error"))
			},
			expectedErr: true,
		},
		{
			name: "Wrong Type Error",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
				_ = body.AddOrReplacePart("table_name", "text/plain", &s)
				_ = body.AddOrReplacePart("table_sys_id", "text/plain", &s)
				_ = body.AddOrReplacePart("uploadFile", "text/plain", &s)

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", mock.Anything).Return(mockSerializationWriter, nil)

				ra.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&Attachment{}, nil)
			},
			expectedErr: true,
		},
		{
			name: "Nil Result",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				s := "test"
				_ = body.AddOrReplacePart("Content-Type", "text/plain", &s)
				_ = body.AddOrReplacePart("table_name", "text/plain", &s)
				_ = body.AddOrReplacePart("table_sys_id", "text/plain", &s)
				_ = body.AddOrReplacePart("uploadFile", "text/plain", &s)

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", mock.Anything).Return(mockSerializationWriter, nil)

				ra.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
			expectedErr: false,
			nilRes:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ra := mocking.NewMockRequestAdapter()
			body := abstractions.NewMultipartBody()
			tt.setup(ra, body)

			rb := NewAttachmentUploadRequestBuilder("url", ra)
			res, err := rb.Post(context.Background(), body, nil)

			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				if tt.nilRes {
					assert.Nil(t, res)
				} else {
					assert.NotNil(t, res)
				}
			}
		})
	}
}

func TestAttachmentUploadRequestBuilder_ToPostRequestInformation(t *testing.T) {
	setupMockSerialization := func(mockContent abstractions.MultipartBody) *mocking.MockSerializationWriterFactory {
		mockSerializationWriter := mocking.NewMockSerializationWriter()
		mockSerializationWriter.On("WriteObjectValue", "", mockContent, ([]serialization.Parsable)(nil)).Return(nil)
		mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
		mockSerializationWriter.On("Close").Return(nil)

		mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
		mockSerializationWriterFactory.On("GetSerializationWriter", "application/json").Return(mockSerializationWriter, nil)
		return mockSerializationWriterFactory
	}

	tests := []struct {
		name                 string
		isNilBuilder         bool
		requestConfiguration *AttachmentUploadRequestBuilderPostRequestConfiguration
		headers              func() *abstractions.RequestHeaders
		options              func() []abstractions.RequestOption
		expectedMethod       abstractions.HttpMethod
	}{
		{
			name:                 "Successful - minimal",
			requestConfiguration: &AttachmentUploadRequestBuilderPostRequestConfiguration{},
			expectedMethod:       abstractions.POST,
		},
		{
			name: "Successful - with headers",
			requestConfiguration: &AttachmentUploadRequestBuilderPostRequestConfiguration{
				Headers: func() *abstractions.RequestHeaders {
					h := abstractions.NewRequestHeaders()
					h.Add("test", "test1")
					return h
				}(),
			},
			expectedMethod: abstractions.POST,
		},
		{
			name: "Successful - with options",
			requestConfiguration: &AttachmentUploadRequestBuilderPostRequestConfiguration{
				Options: func() []abstractions.RequestOption {
					mockOption := mocking.NewMockRequestOption()
					mockOption.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})
					return []abstractions.RequestOption{mockOption}
				}(),
			},
			expectedMethod: abstractions.POST,
		},
		{
			name:                 "Nil model",
			isNilBuilder:         true,
			requestConfiguration: &AttachmentUploadRequestBuilderPostRequestConfiguration{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContent := abstractions.NewMultipartBody()
			mockRequestAdapter := mocking.NewMockRequestAdapter()
			mockRequestAdapter.On("GetSerializationWriterFactory").Return(setupMockSerialization(mockContent))

			mockInternal := mocking.NewMockRequestBuilder()
			mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
			mockInternal.On("GetURLTemplate").Return("")
			mockInternal.On("GetPathParameters").Return(map[string]string{})

			var builder *AttachmentUploadRequestBuilder
			if !tt.isNilBuilder {
				builder = &AttachmentUploadRequestBuilder{mockInternal}
			}

			info, err := builder.ToPostRequestInformation(context.Background(), mockContent, tt.requestConfiguration)

			if tt.isNilBuilder {
				assert.Nil(t, info)
				assert.Nil(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedMethod, info.Method)
				if tt.requestConfiguration.Headers != nil {
					val := info.Headers.Get("test")
					assert.Equal(t, []string{"test1"}, val)
				}
				if tt.requestConfiguration.Options != nil {
					assert.Len(t, info.GetRequestOptions(), 1)
				}
			}
		})
	}
}
