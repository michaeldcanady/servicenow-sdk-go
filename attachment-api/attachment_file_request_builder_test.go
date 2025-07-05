package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAttachmentFileRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}
				requestAdapter := mocking.NewMockRequestAdapter()

				builder := NewAttachmentFileRequestBuilderInternal(pathParameters, requestAdapter)

				assert.IsType(t, &AttachmentFileRequestBuilder{}, builder)
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

func TestNewAttachmentFileRequestBuilder(t *testing.T) {
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

				builder := NewAttachmentFileRequestBuilder(rawURL, requestAdapter)

				assert.IsType(t, &AttachmentFileRequestBuilder{}, builder)
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

func TestAttachmentFileRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{"file_name": "fileName", "table_name": "tableName", "table_sys_id": "sysId"}
				mockQueryParametersAny := map[string]any{"file_name": []interface{}{"fileName"}, "table_name": []interface{}{"tableName"}, "table_sys_id": []interface{}{"sysId"}}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("Content-Type", "application/json")
				mockContent := []byte("testing")
				mockURLTemplate := ""
				mockParsable := &FileModel{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Content:            mockContent,
					Headers:            mockHeaders,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})

				contentType := "application/json"
				mockData := mockContent
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.AnythingOfType("abstractions.ErrorMappings")).Return(mockParsable, nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
						FileName:   newInternal.ToPointer("fileName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Equal(t, mockParsable, result)
				assert.Nil(t, err)
			},
		},
		{
			name: "Send error",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{"file_name": "fileName", "table_name": "tableName", "table_sys_id": "sysId"}
				mockQueryParametersAny := map[string]any{"file_name": []interface{}{"fileName"}, "table_name": []interface{}{"tableName"}, "table_sys_id": []interface{}{"sysId"}}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("Content-Type", "application/json")
				mockContent := []byte("testing")
				mockURLTemplate := ""

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Content:            mockContent,
					Headers:            mockHeaders,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})

				contentType := "application/json"
				mockData := mockContent
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.AnythingOfType("abstractions.ErrorMappings")).Return((*mocking.MockParsable)(nil), errors.New("send error"))

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
						FileName:   newInternal.ToPointer("fileName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("send error"), err)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{"file_name": "fileName", "table_name": "tableName", "table_sys_id": "sysId"}
				mockQueryParametersAny := map[string]any{"file_name": []interface{}{"fileName"}, "table_name": []interface{}{"tableName"}, "table_sys_id": []interface{}{"sysId"}}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("Content-Type", "application/json")
				mockContent := []byte("testing")
				mockURLTemplate := ""

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Content:            mockContent,
					Headers:            mockHeaders,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})

				contentType := "application/json"
				mockData := mockContent
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.AnythingOfType("abstractions.ErrorMappings")).Return(mocking.NewMockParsable(), nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
						FileName:   newInternal.ToPointer("fileName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("resp is not *FileModel"), err)
			},
		},
		{
			name: "Nil request adapter",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("Content-Type", "text/plain")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				contentType := "text/plain"
				mockData := []byte("testing")
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
						FileName:   newInternal.ToPointer("fileName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				result, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("requestAdapter is nil"), err)
			},
		},
		{
			name: "Missing data",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := "text/plain"
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
						FileName:   newInternal.ToPointer("fileName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("data is empty"), err)
			},
		},
		{
			name: "Missing content type",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
						FileName:   newInternal.ToPointer("fileName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("contentType can't be empty"), err)
			},
		},
		{
			name: "Missing FileName query parameter",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
						TableName:  newInternal.ToPointer("tableName"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("requestConfiguration.QueryParameters.FileName can't be empty"), err)
			},
		},
		{
			name: "Missing TableName query parameter",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: newInternal.ToPointer("sysId"),
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("requestConfiguration.QueryParameters.TableName can't be empty"), err)
			},
		},
		{
			name: "Missing TableSysID query parameter",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: &AttachmentFileRequestBuilderPostQueryParameters{
						TableSysID: nil,
					},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("requestConfiguration.QueryParameters.TableSysID can't be empty"), err)
			},
		},
		{
			name: "Missing query parameters",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("requestConfiguration or requestConfiguration.QueryParameters can't be empty"), err)
			},
		},
		{
			name: "Missing request configuration",
			test: func(t *testing.T) {
				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				contentType := ""
				mockData := []byte{}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, nil)

				assert.Nil(t, requestInformation)
				assert.Equal(t, errors.New("requestConfiguration or requestConfiguration.QueryParameters can't be empty"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{}

				builder := (*AttachmentFileRequestBuilder)(nil)

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.Post(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachmentFileRequestBuilder_ToPostRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful - minimal",
			test: func(t *testing.T) {
				contentType := "application/json"
				mockData := []byte("I am content")
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockContent := mockData
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
					Content:            mockContent,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), media, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Successful - with headers",
			test: func(t *testing.T) {
				contentType := "application/json"
				mockData := []byte("I am content")
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockHeaders.Add("test", "test1")
				mockContent := mockData
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
					Content:            mockContent,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				headers := &abstractions.RequestHeaders{}
				headers.Add("test", "test1")

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					Headers: headers,
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), media, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Successful - with options",
			test: func(t *testing.T) {
				contentType := "application/json"
				mockData := []byte("I am content")
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockContent := mockData
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
					Content:            mockContent,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{mockRequestOption})

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					Options: []abstractions.RequestOption{mockRequestOption},
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), media, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Successful - with query parameters",
			test: func(t *testing.T) {
				contentType := "application/json"
				mockData := []byte("I am content")
				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", media, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{"file_name": "file name"}
				mockQueryParametersAny := map[string]any{"file_name": []interface{}{"file name"}}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockContent := mockData
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
					Content:            mockContent,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				queryParameters := &AttachmentFileRequestBuilderPostQueryParameters{
					FileName: newInternal.ToPointer("file name"),
				}

				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{
					QueryParameters: queryParameters,
				}

				builder := &AttachmentFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), media, requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				contentType := ""
				mockData := []byte{}
				requestConfiguration := &AttachmentFileRequestBuilderPostRequestConfiguration{}

				builder := (*AttachmentFileRequestBuilder)(nil)

				media := &Media{
					data:        mockData,
					contentType: contentType,
				}

				requestInformation, err := builder.ToPostRequestInformation(context.Background(), media, requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
