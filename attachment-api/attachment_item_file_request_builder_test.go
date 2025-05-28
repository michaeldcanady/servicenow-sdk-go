package attachmentapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachmentItemFileRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}
				requestAdapter := mocking.NewMockRequestAdapter()

				builder := NewAttachmentItemFileRequestBuilderInternal(pathParameters, requestAdapter)

				assert.IsType(t, &AttachmentItemFileRequestBuilder{}, builder)
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

func TestNewAttachmentItemFileRequestBuilder(t *testing.T) {
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

				builder := NewAttachmentItemFileRequestBuilder(rawURL, requestAdapter)

				assert.IsType(t, &AttachmentItemFileRequestBuilder{}, builder)
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

// TODO: (TestAttachmentItemFileRequestBuilder_Get) Add tests
//func TestAttachmentItemFileRequestBuilder_Get(t *testing.T) {
//	tests := []struct {
//		name string
//		test func(*testing.T)
//	}{
//		{
//			name: "Successful",
//			test: func(t *testing.T) {
//				mockPathParameters := map[string]string{}
//				mockPathParametersAny := map[string]any{}
//				mockHeaders := &abstractions.RequestHeaders{}
//				mockHeaders.Add("Accept", "*/*")
//				mockQueryParameters := map[string]string{}
//				mockQueryParametersAny := map[string]any{}
//				mockURLTemplate := ""
//				mockContent := []byte("testing")
//				mockParsable := &FileWithContentModel{}
//
//				opts := nethttplibrary.NewHeadersInspectionOptions()
//				opts.InspectResponseHeaders = true
//
//				expected := &abstractions.RequestInformation{
//					Method:             abstractions.GET,
//					UrlTemplate:        mockURLTemplate,
//					PathParameters:     mockPathParameters,
//					PathParametersAny:  mockPathParametersAny,
//					QueryParameters:    mockQueryParameters,
//					QueryParametersAny: mockQueryParametersAny,
//					Headers:            mockHeaders,
//				}
//				expected.AddRequestOptions([]abstractions.RequestOption{opts})
//
//				mockRequestAdapter := mocking.NewMockRequestAdapter()
//				// TODO: how to have the inspection option update after SendPrimitive
//				mockRequestAdapter.On("SendPrimitive", context.Background(), expected, "[]byte", abstractions.ErrorMappings{}).Return(mockContent, nil)
//
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
//				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
//				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)
//
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				result, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Equal(t, mockParsable, result)
//				assert.Nil(t, err)
//			},
//		},
//		{
//			name: "Send error",
//			test: func(t *testing.T) {
//				mockPathParameters := map[string]string{}
//				mockPathParametersAny := map[string]any{}
//				mockHeaders := &abstractions.RequestHeaders{}
//				mockHeaders.Add("Accept", "*/*")
//				mockURLTemplate := ""
//
//				expected := &abstractions.RequestInformation{
//					Method:            abstractions.GET,
//					UrlTemplate:       mockURLTemplate,
//					PathParameters:    mockPathParameters,
//					PathParametersAny: mockPathParametersAny,
//					Headers:           mockHeaders,
//				}
//				expected.AddRequestOptions([]abstractions.RequestOption{})
//
//				mockRequestAdapter := mocking.NewMockRequestAdapter()
//				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), abstractions.ErrorMappings{}).Return((*mocking.MockParsable)(nil), errors.New("send error"))
//
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
//				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
//				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)
//
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				result, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Nil(t, result)
//				assert.Equal(t, errors.New("send error"), err)
//			},
//		},
//		{
//			name: "Wrong type",
//			test: func(t *testing.T) {
//				mockPathParameters := map[string]string{}
//				mockPathParametersAny := map[string]any{}
//				mockHeaders := &abstractions.RequestHeaders{}
//				mockHeaders.Add("Accept", "*/*")
//				mockURLTemplate := ""
//
//				expected := &abstractions.RequestInformation{
//					Method:            abstractions.POST,
//					UrlTemplate:       mockURLTemplate,
//					PathParameters:    mockPathParameters,
//					PathParametersAny: mockPathParametersAny,
//					Headers:           mockHeaders,
//				}
//				expected.AddRequestOptions([]abstractions.RequestOption{})
//
//				mockRequestAdapter := mocking.NewMockRequestAdapter()
//				mockRequestAdapter.On("Send", context.Background(), expected, mock.AnythingOfType("serialization.ParsableFactory"), abstractions.ErrorMappings{}).Return(mocking.NewMockParsable(), nil)
//
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
//				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
//				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)
//
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				result, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Nil(t, result)
//				assert.Equal(t, errors.New("resp is not *FileModel"), err)
//			},
//		},
//		{
//			name: "Nil request adapter",
//			test: func(t *testing.T) {
//				mockPathParameters := map[string]string{}
//				mockHeaders := &abstractions.RequestHeaders{}
//				mockHeaders.Add("Accept", "*/*")
//				mockURLTemplate := ""
//
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//				mockInternalRequestBuilder.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
//				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
//				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)
//
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				result, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Nil(t, result)
//				assert.Equal(t, errors.New("requestAdapter is nil"), err)
//			},
//		},
//		{
//			name: "Missing data",
//			test: func(t *testing.T) {
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				requestInformation, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Nil(t, requestInformation)
//				assert.Equal(t, errors.New("data is empty"), err)
//			},
//		},
//		{
//			name: "Missing content type",
//			test: func(t *testing.T) {
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				requestInformation, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Nil(t, requestInformation)
//				assert.Equal(t, errors.New("contentType can't be empty"), err)
//			},
//		},
//		{
//			name: "Missing request configuration",
//			test: func(t *testing.T) {
//				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
//
//				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
//
//				requestInformation, err := builder.Get(context.Background(), nil)
//
//				assert.Nil(t, requestInformation)
//				assert.Equal(t, errors.New("requestConfiguration or requestConfiguration.QueryParameters can't be empty"), err)
//			},
//		},
//		{
//			name: "Nil model",
//			test: func(t *testing.T) {
//				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
//
//				builder := (*AttachmentItemFileRequestBuilder)(nil)
//
//				requestInformation, err := builder.Get(context.Background(), requestConfiguration)
//
//				assert.Nil(t, requestInformation)
//				assert.Nil(t, err)
//			},
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, test.test)
//	}
//}

func TestAttachmentItemFileRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful - minimal",
			test: func(t *testing.T) {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.GET,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Successful - with headers",
			test: func(t *testing.T) {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockHeaders.Add("test", "test1")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.GET,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				headers := &abstractions.RequestHeaders{}
				headers.Add("test", "test1")

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{
					Headers: headers,
				}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Successful - with options",
			test: func(t *testing.T) {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockURLTemplate := ""

				mockRequestOption := mocking.NewMockRequestOption()
				mockRequestOption.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.GET,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{mockRequestOption})

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{
					Options: []abstractions.RequestOption{mockRequestOption},
				}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := (*AttachmentItemFileRequestBuilder)(nil)

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
