package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAttachmentFileRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		requestAdapter abstractions.RequestAdapter
	}{
		{
			name:           "Valid parameters",
			pathParameters: map[string]string{"dummy": "value"},
			requestAdapter: mocking.NewMockRequestAdapter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentFileRequestBuilderInternal(tt.pathParameters, tt.requestAdapter)
			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentFileRequestBuilder{}, builder)
			assert.Equal(t, tt.pathParameters, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestNewAttachmentFileRequestBuilder(t *testing.T) {
	tests := []struct {
		name           string
		rawURL         string
		requestAdapter abstractions.RequestAdapter
		expectedParams map[string]string
	}{
		{
			name:           "Valid url",
			rawURL:         "https://example.com",
			requestAdapter: mocking.NewMockRequestAdapter(),
			expectedParams: map[string]string{internal.RawURLKey: "https://example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentFileRequestBuilder(tt.rawURL, tt.requestAdapter)
			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentFileRequestBuilder{}, builder)
			assert.Equal(t, tt.expectedParams, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestAttachmentFileRequestBuilder_Post(t *testing.T) {
	defaultMedia := &Media{data: []byte("test-data"), contentType: "application/json"}
	defaultParams := &AttachmentFileRequestBuilderPostQueryParameters{
		TableSysID: internal.ToPointer("sys123"),
		TableName:  internal.ToPointer("incident"),
		FileName:   internal.ToPointer("test.txt"),
	}

	tests := []struct {
		name                 string
		nilBuilder           bool
		media                *Media
		requestConfig        *AttachmentFileRequestBuilderPostRequestConfiguration
		nilRequestAdapter    bool
		sendError            error
		sendResponse         serialization.Parsable
		serializationErr     error
		expectedErr          error
		expectedErrorMessage string
	}{
		{
			name: "Successful upload",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:        defaultMedia,
			sendResponse: core.NewBaseServiceNowItemResponse[*File](CreateFileFromDiscriminatorValue),
		},
		{
			name:       "Nil builder",
			nilBuilder: true,
			media:      defaultMedia,
		},
		{
			name:                 "Missing requestConfiguration",
			requestConfig:        nil,
			media:                defaultMedia,
			expectedErrorMessage: "requestConfiguration is nil",
		},
		{
			name: "Missing query parameters",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: nil,
			},
			media:                defaultMedia,
			expectedErrorMessage: "requestConfiguration.QueryParameters is nil",
		},
		{
			name: "Missing TableSysID",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
					TableName: internal.ToPointer("incident"),
					FileName:  internal.ToPointer("test.txt"),
				},
			},
			media:       defaultMedia,
			expectedErr: errors.New("requestConfiguration.QueryParameters.TableSysID is nil or empty"),
		},
		{
			name: "Missing TableName",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
					TableSysID: internal.ToPointer("sys123"),
					FileName:   internal.ToPointer("test.txt"),
				},
			},
			media:       defaultMedia,
			expectedErr: errors.New("requestConfiguration.QueryParameters.TableName is nil or empty"),
		},
		{
			name: "Missing FileName",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
					TableSysID: internal.ToPointer("sys123"),
					TableName:  internal.ToPointer("incident"),
				},
			},
			media:       defaultMedia,
			expectedErr: errors.New("requestConfiguration.QueryParameters.FileName is nil or empty"),
		},
		{
			name: "Nil media",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:       nil,
			expectedErr: errors.New("media is nil"),
		},
		{
			name: "Empty content type",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:       &Media{data: []byte("test"), contentType: ""},
			expectedErr: errors.New("media.contentType is nil or empty"),
		},
		{
			name: "Empty media data",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:       &Media{data: []byte{}, contentType: "application/json"},
			expectedErr: errors.New("media.data is nil or empty"),
		},
		{
			name: "Nil request adapter",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:             defaultMedia,
			nilRequestAdapter: true,
			expectedErr:       snerrors.ErrNilRequestAdapter,
		},
		{
			name: "Adapter send error",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:       defaultMedia,
			sendError:   errors.New("send failed"),
			expectedErr: errors.New("send failed"),
		},
		{
			name: "Nil response from adapter",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:        defaultMedia,
			sendResponse: nil,
			expectedErr:  snerrors.ErrNilResponse,
		},
		{
			name: "Wrong response type from adapter",
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: defaultParams,
			},
			media:        defaultMedia,
			sendResponse: mocking.NewMockParsable(),
			expectedErr:  errors.New("resp is not ServiceNowItemResponse[*File]"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentFileRequestBuilder
			if !tt.nilBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetURLTemplate").Return("")
				mockInternalRequestBuilder.On("GetPathParameters").Return(map[string]string{})

				if tt.nilRequestAdapter {
					mockInternalRequestBuilder.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
				} else {
					mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)

					writer := mocking.NewMockSerializationWriter()
					writer.On("WriteObjectValue", "", tt.media, mock.Anything).Return(nil)
					writer.On("Close").Return(nil)
					var data []byte
					if tt.media != nil {
						data = tt.media.data
					}
					writer.On("GetSerializedContent").Return(data, nil)

					factory := mocking.NewMockSerializationWriterFactory()
					factory.On("GetSerializationWriter", "application/json").Return(writer, nil)
					mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

					mockRequestAdapter.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.sendResponse, tt.sendError)
				}
				builder = &AttachmentFileRequestBuilder{mockInternalRequestBuilder}
			}

			res, err := builder.Post(context.Background(), tt.media, tt.requestConfig)

			if tt.nilBuilder {
				assert.Nil(t, res)
				assert.NoError(t, err)
				return
			}

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
				assert.Nil(t, res)
			} else if tt.expectedErrorMessage != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMessage)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.sendResponse, res)
			}
		})
	}
}

func TestAttachmentFileRequestBuilder_ToPostRequestInformation(t *testing.T) {
	defaultMedia := &Media{data: []byte("I am content"), contentType: "application/json"}

	tests := []struct {
		name                    string
		nilBuilder              bool
		nilRequestBuilder       bool
		media                   *Media
		requestConfig           *AttachmentFileRequestBuilderPostRequestConfiguration
		nilRequestAdapter       bool
		expectedErr             bool
		expectedHeaders         map[string][]string
		expectedQueryParameters map[string]any
		expectedOptionsKeys     []string
	}{
		{
			name:          "Successful minimal",
			media:         defaultMedia,
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{},
			expectedHeaders: map[string][]string{
				"accept":       {"application/json"},
				"content-type": {"application/json"},
			},
		},
		{
			name:       "Nil builder",
			nilBuilder: true,
			media:      defaultMedia,
		},
		{
			name:              "Nil request builder internally",
			nilRequestBuilder: true,
			media:             defaultMedia,
		},
		{
			name:              "Nil request adapter",
			media:             defaultMedia,
			requestConfig:     &AttachmentFileRequestBuilderPostRequestConfiguration{},
			nilRequestAdapter: true,
			expectedErr:       true,
		},
		{
			name:  "With custom headers",
			media: defaultMedia,
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				Headers: func() *abstractions.RequestHeaders {
					h := abstractions.NewRequestHeaders()
					h.Add("custom-header", "custom-value")
					return h
				}(),
			},
			expectedHeaders: map[string][]string{
				"accept":        {"application/json"},
				"content-type":  {"application/json"},
				"custom-header": {"custom-value"},
			},
		},
		{
			name:  "With query parameters",
			media: defaultMedia,
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
					FileName: internal.ToPointer("test.txt"),
				},
			},
			expectedQueryParameters: map[string]any{
				"file_name": []any{"test.txt"},
			},
		},
		{
			name:  "With request options",
			media: defaultMedia,
			requestConfig: &AttachmentFileRequestBuilderPostRequestConfiguration{
				Options: func() []abstractions.RequestOption {
					opt := mocking.NewMockRequestOption()
					opt.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})
					return []abstractions.RequestOption{opt}
				}(),
			},
			expectedOptionsKeys: []string{"key"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentFileRequestBuilder

			if !tt.nilBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetURLTemplate").Return("")
				mockInternalRequestBuilder.On("GetPathParameters").Return(map[string]string{})

				if tt.nilRequestAdapter {
					mockInternalRequestBuilder.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
				} else {
					mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)

					writer := mocking.NewMockSerializationWriter()
					writer.On("WriteObjectValue", "", tt.media, mock.Anything).Return(nil)
					writer.On("Close").Return(nil)
					var data []byte
					if tt.media != nil {
						data = tt.media.data
					}
					writer.On("GetSerializedContent").Return(data, nil)

					factory := mocking.NewMockSerializationWriterFactory()
					factory.On("GetSerializationWriter", "application/json").Return(writer, nil)
					mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				}

				if tt.nilRequestBuilder {
					builder = &AttachmentFileRequestBuilder{nil}
				} else {
					builder = &AttachmentFileRequestBuilder{mockInternalRequestBuilder}
				}
			}

			reqInfo, err := builder.ToPostRequestInformation(context.Background(), tt.media, tt.requestConfig)

			if tt.nilBuilder || tt.nilRequestBuilder {
				assert.Nil(t, reqInfo)
				assert.NoError(t, err)
				return
			}

			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, reqInfo)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, reqInfo)
				assert.Equal(t, abstractions.POST, reqInfo.Method)
				assert.Equal(t, tt.media.data, reqInfo.Content)

				// Verify headers
				if tt.expectedHeaders != nil {
					for k, v := range tt.expectedHeaders {
						assert.Equal(t, v, reqInfo.Headers.Get(k))
					}
				}

				// Verify query parameters
				if tt.expectedQueryParameters != nil {
					assert.Equal(t, tt.expectedQueryParameters, reqInfo.QueryParametersAny)
				}

				// Verify options
				if len(tt.expectedOptionsKeys) > 0 {
					options := reqInfo.GetRequestOptions()
					assert.Equal(t, len(tt.expectedOptionsKeys), len(options))
					for i, opt := range options {
						assert.Equal(t, tt.expectedOptionsKeys[i], opt.GetKey().Key)
					}
				}
			}
		})
	}
}
