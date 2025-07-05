package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestAttachmentItemFileRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockContent := []byte("testing")
				mockParsable := &FileWithContentModel{
					NewFile(),
				}

				err := mockParsable.SetContent(mockContent)
				assert.Nil(t, err)

				opts := nethttplibrary.NewHeadersInspectionOptions()
				opts.InspectResponseHeaders = true

				expected := &abstractions.RequestInformation{
					Method:             abstractions.GET,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{opts})

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendPrimitive", context.Background(), expected, "[]byte", mock.IsType(abstractions.ErrorMappings{})).Return(mockContent, nil).Run(func(args mock.Arguments) {
					value := args.Get(1)
					requestInformation, _ := value.(*abstractions.RequestInformation)
					opts := requestInformation.GetRequestOptions()
					for _, opt := range opts {
						if typedOpt, ok := opt.(*nethttplibrary.HeadersInspectionOptions); ok {
							respHeaders := abstractions.NewResponseHeaders()
							respHeaders.Add("X-Attachment-Metadata", "{}")
							typedOpt.ResponseHeaders = respHeaders
							break
						}
					}
					requestInformation.AddRequestOptions(opts)
				})

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Get(context.Background(), requestConfiguration)

				// TODO: make more robust without testing functions
				expectedContent, _ := mockParsable.GetContent()
				resultContent, _ := result.GetContent()
				assert.Equal(t, expectedContent, resultContent)
				assert.Nil(t, err)
			},
		},
		{
			name: "Send error",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockContent := []byte("testing")
				mockParsable := &FileWithContentModel{
					NewFile(),
				}
				err := mockParsable.SetContent(mockContent)
				assert.Nil(t, err)

				opts := nethttplibrary.NewHeadersInspectionOptions()
				opts.InspectResponseHeaders = true

				expected := &abstractions.RequestInformation{
					Method:             abstractions.GET,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{opts})

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendPrimitive", context.Background(), expected, "[]byte", mock.IsType(abstractions.ErrorMappings{})).Return((*mocking.MockParsable)(nil), errors.New("send error")).Run(func(args mock.Arguments) {
					value := args.Get(1)
					requestInformation, _ := value.(*abstractions.RequestInformation)
					opts := requestInformation.GetRequestOptions()
					for _, opt := range opts {
						if typedOpt, ok := opt.(*nethttplibrary.HeadersInspectionOptions); ok {
							respHeaders := abstractions.NewResponseHeaders()
							respHeaders.Add("X-Attachment-Metadata", "{}")
							typedOpt.ResponseHeaders = respHeaders
							break
						}
					}
					requestInformation.AddRequestOptions(opts)
				})

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Get(context.Background(), requestConfiguration)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("send error"), err)
			},
		},
		{
			name: "Wrong type",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockQueryParameters := map[string]string{}
				mockQueryParametersAny := map[string]any{}
				mockURLTemplate := ""
				mockContent := []byte("testing")
				mockParsable := &FileWithContentModel{
					NewFile(),
				}
				err := mockParsable.SetContent(mockContent)
				assert.Nil(t, err)

				opts := nethttplibrary.NewHeadersInspectionOptions()
				opts.InspectResponseHeaders = true

				expected := &abstractions.RequestInformation{
					Method:             abstractions.GET,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}
				expected.AddRequestOptions([]abstractions.RequestOption{opts})

				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendPrimitive", context.Background(), expected, "[]byte", mock.IsType(abstractions.ErrorMappings{})).Return(mocking.NewMockParsable(), nil)

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Get(context.Background(), requestConfiguration)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("resp is not []byte"), err)
			},
		},
		{
			name: "Nil request adapter",
			test: func(t *testing.T) {
				mockPathParameters := map[string]string{}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("Accept", "*/*")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}

				result, err := builder.Get(context.Background(), requestConfiguration)

				assert.Nil(t, result)
				assert.Equal(t, errors.New("requestAdapter is nil"), err)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				requestConfiguration := &AttachmentItemFileRequestBuilderGetRequestConfiguration{}

				builder := (*AttachmentItemFileRequestBuilder)(nil)

				requestInformation, err := builder.Get(context.Background(), requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

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
