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

// TODO: add tests
func TestBatchRequestBuilder_Post(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: add tests
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
