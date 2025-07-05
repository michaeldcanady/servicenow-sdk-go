package batchapi

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

func TestNewBatchRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				var params = map[string]string{"key": "value"}
				var mockAdapter = mocking.NewMockRequestAdapter()

				builder := NewBatchRequestBuilderInternal(params, mockAdapter)

				assert.NotNil(t, builder)
				assert.IsType(t, &BatchRequestBuilder{}, builder)
				assert.Equal(t, mockAdapter, builder.RequestBuilder.GetRequestAdapter())
				assert.Equal(t, params, builder.RequestBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestNewBatchRequestBuilder2(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				const rawURL = ""
				var mockAdapter = mocking.NewMockRequestAdapter()

				builder := NewBatchRequestBuilder(rawURL, mockAdapter)

				assert.NotNil(t, builder)
				assert.IsType(t, &BatchRequestBuilder{}, builder)
				assert.Equal(t, mockAdapter, builder.RequestBuilder.GetRequestAdapter())
				assert.Equal(t, map[string]string{newInternal.RawURLKey: rawURL}, builder.RequestBuilder.GetPathParameters())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				request := newMockBatchRequest()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockContent := &BatchResponseModel{}
				mockData := []byte{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            mockData,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})
				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				mockWriter.On("Close").Return(nil)
				mockWriter.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(mockWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.IsType(abstractions.ErrorMappings{})).Return(mockContent, nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), request, nil)

				assert.Equal(t, mockContent, result)
				assert.Nil(t, err)
				mockRequestAdapter.AssertExpectations(t)
				mockInternalRequestBuilder.AssertExpectations(t)
				mockWriter.AssertExpectations(t)
			},
		},
		{
			name: "Send error",
			test: func(t *testing.T) {
				request := newMockBatchRequest()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockData := []byte{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            mockData,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})
				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				mockWriter.On("Close").Return(nil)
				mockWriter.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(mockWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.IsType(abstractions.ErrorMappings{})).Return(nil, errors.New("send error"))

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), request, nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("send error"), err)
				mockRequestAdapter.AssertExpectations(t)
				mockInternalRequestBuilder.AssertExpectations(t)
				mockWriter.AssertExpectations(t)
			},
		},
		{
			name: "wrong response type",
			test: func(t *testing.T) {
				request := newMockBatchRequest()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockContent := mocking.NewMockParsable()
				mockData := []byte{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            mockData,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})
				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				mockWriter.On("Close").Return(nil)
				mockWriter.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(mockWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.IsType(abstractions.ErrorMappings{})).Return(mockContent, nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), request, nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("resp is not *BatchResponse"), err)
				mockRequestAdapter.AssertExpectations(t)
				mockInternalRequestBuilder.AssertExpectations(t)
				mockWriter.AssertExpectations(t)
			},
		},
		{
			name: "Nil response",
			test: func(t *testing.T) {
				request := newMockBatchRequest()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockData := []byte{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            mockData,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})
				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				mockWriter.On("Close").Return(nil)
				mockWriter.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(mockWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)
				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), mock.IsType(abstractions.ErrorMappings{})).Return(nil, nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), request, nil)

				assert.Nil(t, result)
				assert.Nil(t, err)
				mockRequestAdapter.AssertExpectations(t)
				mockInternalRequestBuilder.AssertExpectations(t)
				mockWriter.AssertExpectations(t)
			},
		},
		{
			name: "requestInformation conversion error",
			test: func(t *testing.T) {
				request := newMockBatchRequest()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockData := []byte{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            mockData,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})
				mockWriter := mocking.NewMockSerializationWriter()
				mockWriter.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				mockWriter.On("Close").Return(nil)
				mockWriter.On("GetSerializedContent").Return([]byte{}, errors.New("get content error"))

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(mockWriter, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), request, nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("get content error"), err)
				mockRequestAdapter.AssertExpectations(t)
				mockInternalRequestBuilder.AssertExpectations(t)
				mockWriter.AssertExpectations(t)
			},
		},
		{
			name: "Nil body",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockData := []byte{}

				expected := &abstractions.RequestInformation{
					Method:             abstractions.POST,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
					Content:            mockData,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{})

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Post(context.Background(), nil, nil)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("body can't be nil"), err)
			},
		},
		{
			name: "Nil inner requestBuilder",
			test: func(t *testing.T) {
				request := newMockBatchRequest()

				builder := &BatchRequestBuilder{nil}

				requestInformation, err := builder.Post(context.Background(), request, nil)

				assert.Nil(t, err)
				assert.Nil(t, requestInformation)
			},
		},
		{
			name: "Nil requestBuilder",
			test: func(t *testing.T) {
				request := newMockBatchRequest()

				builder := (*BatchRequestBuilder)(nil)

				requestInformation, err := builder.Post(context.Background(), request, nil)

				assert.Nil(t, err)
				assert.Nil(t, requestInformation)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestBatchRequestBuilder_toPostRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful - minimal",
			test: func(t *testing.T) {
				mockData := []byte("I am content")

				request := newMockBatchRequest()

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockPathParameters := map[string]string{"baseurl": "https://mock.service-now.com"}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockContent := mockData
				mockURLTemplate := "{+baseurl}/api/now/v1/batch"

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

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.toPostRequestInformation(context.Background(), request, nil)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
				writer.AssertExpectations(t)
			},
		},
		{
			name: "Successful with headers",
			test: func(t *testing.T) {
				mockData := []byte("I am content")

				request := newMockBatchRequest()

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockPathParameters := map[string]string{"baseurl": "https://mock.service-now.com"}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockHeaders.Add("random", "value")
				mockContent := mockData
				mockURLTemplate := "{+baseurl}/api/now/v1/batch"

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

				headers := abstractions.NewRequestHeaders()
				headers.Add("random", "value")

				config := &BatchRequestBuilderPostRequestConfiguration{
					Headers: headers,
				}

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.toPostRequestInformation(context.Background(), request, config)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
				writer.AssertExpectations(t)
			},
		},
		{
			name: "Successful with options",
			test: func(t *testing.T) {
				mockData := []byte("I am content")
				mockRequestOption := mocking.NewMockRequestOption()
				mockRequestOption.On("GetKey").Return(abstractions.RequestOptionKey{Key: "optionKey"})

				request := newMockBatchRequest()

				writer := mocking.NewMockSerializationWriter()
				writer.On("WriteObjectValue", "", request, mock.AnythingOfType("[]serialization.Parsable")).Return(nil)
				writer.On("Close").Return(nil)
				writer.On("GetSerializedContent").Return(mockData, nil)

				factory := mocking.NewMockSerializationWriterFactory()
				factory.On("GetSerializationWriter", "application/json").Return(writer, nil)

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("GetSerializationWriterFactory").Return(factory, nil)

				mockPathParameters := map[string]string{"baseurl": "https://mock.service-now.com"}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockContent := mockData
				mockURLTemplate := "{+baseurl}/api/now/v1/batch"

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

				options := []abstractions.RequestOption{mockRequestOption}

				config := &BatchRequestBuilderPostRequestConfiguration{
					Options: options,
				}

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.toPostRequestInformation(context.Background(), request, config)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
				writer.AssertExpectations(t)
			},
		},
		{
			name: "SetContentFromParsableError",
			test: func(t *testing.T) {
				request := newMockBatchRequest()

				mockPathParameters := map[string]string{"baseurl": "https://mock.service-now.com"}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
				mockHeaders.Add("content-type", "application/json")
				mockURLTemplate := "{+baseurl}/api/now/v1/batch"

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(nil)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				builder := &BatchRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.toPostRequestInformation(context.Background(), request, nil)

				assert.Equal(t, errors.New("requestAdapter cannot be nil"), err)
				assert.Nil(t, requestInformation)
			},
		},
		{
			name: "Nil inner requestBuilder",
			test: func(t *testing.T) {
				request := newMockBatchRequest()

				builder := &BatchRequestBuilder{nil}

				requestInformation, err := builder.toPostRequestInformation(context.Background(), request, nil)

				assert.Nil(t, err)
				assert.Nil(t, requestInformation)
			},
		},
		{
			name: "Nil requestBuilder",
			test: func(t *testing.T) {
				request := newMockBatchRequest()

				builder := (*BatchRequestBuilder)(nil)

				requestInformation, err := builder.toPostRequestInformation(context.Background(), request, nil)

				assert.Nil(t, err)
				assert.Nil(t, requestInformation)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
