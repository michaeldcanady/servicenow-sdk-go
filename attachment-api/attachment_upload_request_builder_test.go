package attachmentapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
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

// TODO: (TestAttachmentUploadRequestBuilder_Post) Add tests
func TestAttachmentUploadRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
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
					Content:            []byte("content"),
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				requestConfiguration := &AttachmentUploadRequestBuilderPostRequestConfiguration{}

				builder := &AttachmentUploadRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), mockContent, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
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
				assert.Equal(t, expected, requestInformation)
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
				assert.Equal(t, expected, requestInformation)
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
