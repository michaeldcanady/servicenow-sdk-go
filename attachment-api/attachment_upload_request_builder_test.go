package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAttachmentUploadRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}
				requestAdapter := mocking.NewMockRequestAdapter()

				builder := NewAttachmentUploadRequestBuilderInternal(pathParameters, requestAdapter)

				assert.IsType(t, &AttachmentUploadRequestBuilder{}, builder)
				assert.IsType(t, &newInternal.BaseRequestBuilder{}, builder.RequestBuilder)
				assert.Equal(t, pathParameters, builder.GetPathParameters())
				assert.Equal(t, requestAdapter, builder.GetRequestAdapter())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNewAttachmentUploadRequestBuilder(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				rawURL := ""
				requestAdapter := mocking.NewMockRequestAdapter()

				urlParams := map[string]string{newInternal.RawURLKey: rawURL}

				builder := NewAttachmentUploadRequestBuilder(rawURL, requestAdapter)

				assert.IsType(t, &AttachmentUploadRequestBuilder{}, builder)
				assert.IsType(t, &newInternal.BaseRequestBuilder{}, builder.RequestBuilder)
				assert.Equal(t, urlParams, builder.GetPathParameters())
				assert.Equal(t, requestAdapter, builder.GetRequestAdapter())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachmentUploadRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody)
		expectedErr bool
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
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&FileModel{}, nil)
			},
			expectedErr: false,
		},
		{
			name: "Missing Content-Type",
			setup: func(ra *mocking.MockRequestAdapter, body abstractions.MultipartBody) {
				// empty body
			},
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
				ra.On("Send", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&Attachment2Model{}, nil)
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
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ra := mocking.NewMockRequestAdapter()
			body := abstractions.NewMultipartBody()
			test.setup(ra, body)

			rb := NewAttachmentUploadRequestBuilder("url", ra)
			res, err := rb.Post(context.Background(), body, nil)

			if test.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, res)
			} else {
				assert.NoError(t, err)
				if test.name == "Nil Result" {
					assert.Nil(t, res)
				} else {
					assert.NotNil(t, res)
				}
			}
		})
	}
}

func TestAttachmentUploadRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful - minimal",
			test: func(t *testing.T) {
				mockContent := abstractions.NewMultipartBody()

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", "", mockContent, ([]serialization.Parsable)(nil)).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", "application/json").Return(mockSerializationWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)

				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "application/json")
				// TODO: has to include a boundary value, which is changing
				mockHeaders.Add("content-type", "application/json")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				requestConfiguration := &AttachmentUploadRequestBuilderPostRequestConfiguration{}

				builder := &AttachmentUploadRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), mockContent, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected.Method, requestInformation.Method)
				assert.Equal(t, expected.UrlTemplate, requestInformation.UrlTemplate)
				assert.Equal(t, expected.PathParameters, requestInformation.PathParameters)
				assert.Equal(t, expected.PathParametersAny, requestInformation.PathParametersAny)
				assert.Equal(t, expected.QueryParameters, requestInformation.QueryParameters)
				assert.Equal(t, expected.QueryParametersAny, requestInformation.QueryParametersAny)
			},
		},
		{
			name: "Successful - with headers",
			test: func(t *testing.T) {
				mockContent := abstractions.NewMultipartBody()

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", "", mockContent, ([]serialization.Parsable)(nil)).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", "application/json").Return(mockSerializationWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)

				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "application/json")
				// TODO: has to include a boundary value, which is changing
				mockHeaders.Add("content-type", "application/json")
				mockHeaders.Add("test", "test1")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            []byte("content"),
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				headers := &abstractions.RequestHeaders{}
				headers.Add("test", "test1")

				requestConfiguration := &AttachmentUploadRequestBuilderPostRequestConfiguration{
					Headers: headers,
				}

				builder := &AttachmentUploadRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), mockContent, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected.Method, requestInformation.Method)
				assert.Equal(t, expected.UrlTemplate, requestInformation.UrlTemplate)
				assert.Equal(t, expected.PathParameters, requestInformation.PathParameters)
				assert.Equal(t, expected.PathParametersAny, requestInformation.PathParametersAny)
				assert.Equal(t, expected.QueryParameters, requestInformation.QueryParameters)
				assert.Equal(t, expected.QueryParametersAny, requestInformation.QueryParametersAny)
			},
		},
		{
			name: "Successful - with options",
			test: func(t *testing.T) {
				mockContent := abstractions.NewMultipartBody()

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", "", mockContent, ([]serialization.Parsable)(nil)).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", "application/json").Return(mockSerializationWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)

				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "application/json")
				// TODO: has to include a boundary value, which is changing
				mockHeaders.Add("content-type", "application/json")
				mockURLTemplate := ""

				mockRequestOption := mocking.NewMockRequestOption()
				mockRequestOption.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            []byte("content"),
				}

				expected.AddRequestOptions([]abstractions.RequestOption{mockRequestOption})

				requestConfiguration := &AttachmentUploadRequestBuilderPostRequestConfiguration{
					Options: []abstractions.RequestOption{mockRequestOption},
				}

				builder := &AttachmentUploadRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), mockContent, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected.Method, requestInformation.Method)
				assert.Equal(t, expected.UrlTemplate, requestInformation.UrlTemplate)
				assert.Equal(t, expected.PathParameters, requestInformation.PathParameters)
				assert.Equal(t, expected.PathParametersAny, requestInformation.PathParametersAny)
				assert.Equal(t, expected.QueryParameters, requestInformation.QueryParameters)
				assert.Equal(t, expected.QueryParametersAny, requestInformation.QueryParametersAny)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				mockContent := abstractions.NewMultipartBody()

				mockSerializationWriter := mocking.NewMockSerializationWriter()
				mockSerializationWriter.On("WriteObjectValue", "", mockContent, ([]serialization.Parsable)(nil)).Return(nil)
				mockSerializationWriter.On("GetSerializedContent").Return([]byte("content"), nil)
				mockSerializationWriter.On("Close").Return(nil)

				mockSerializationWriterFactory := mocking.NewMockSerializationWriterFactory()
				mockSerializationWriterFactory.On("GetSerializationWriter", "application/json").Return(mockSerializationWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(mockSerializationWriterFactory)

				requestConfiguration := &AttachmentUploadRequestBuilderPostRequestConfiguration{}

				builder := (*AttachmentUploadRequestBuilder)(nil)

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), mockContent, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
