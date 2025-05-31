package attachmentapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachmentItemRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}
				requestAdapter := mocking.NewMockRequestAdapter()

				builder := NewAttachmentItemRequestBuilderInternal(pathParameters, requestAdapter)

				assert.IsType(t, &AttachmentItemRequestBuilder{}, builder)
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

func TestNewAttachmentItemRequestBuilder(t *testing.T) {
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

				builder := NewAttachmentItemRequestBuilder(rawURL, requestAdapter)

				assert.IsType(t, &AttachmentItemRequestBuilder{}, builder)
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

func TestAttachmentItemRequestBuilder_File(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}
				requestAdapter := mocking.NewMockRequestAdapter()

				internalRequestBuilder := mocking.NewMockRequestBuilder()
				internalRequestBuilder.On("GetPathParameters").Return(pathParameters)
				internalRequestBuilder.On("GetRequestAdapter").Return(requestAdapter)

				builder := &AttachmentItemRequestBuilder{internalRequestBuilder}

				fileBuilder := builder.File()

				assert.Equal(t, &AttachmentItemFileRequestBuilder{
					&newInternal.BaseRequestBuilder{
						BaseRequestBuilder: abstractions.BaseRequestBuilder{
							PathParameters: pathParameters,
							RequestAdapter: requestAdapter,
							UrlTemplate:    attachmentItemFileURLTemplate,
						},
					},
				}, fileBuilder)
			},
		},
		{
			name: "Nil requestBuilder",
			test: func(t *testing.T) {
				builder := (*AttachmentItemRequestBuilder)(nil)
				fileBuilder := builder.File()

				assert.Nil(t, fileBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentItemRequestBuilder_Get) Add tests
func TestAttachmentItemRequestBuilder_Get(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentItemRequestBuilder_Delete) Add tests
func TestAttachmentItemRequestBuilder_Delete(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachmentItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
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
				mockHeaders.Add("Accept", "application/json")
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

				requestConfiguration := &AttachmentItemRequestBuilderGetRequestConfiguration{}

				builder := &AttachmentItemRequestBuilder{mockInternalRequestBuilder}

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
				mockHeaders.Add("Accept", "application/json")
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

				requestConfiguration := &AttachmentItemRequestBuilderGetRequestConfiguration{
					Headers: headers,
				}

				builder := &AttachmentItemRequestBuilder{mockInternalRequestBuilder}

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
				mockHeaders.Add("Accept", "application/json")
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

				requestConfiguration := &AttachmentItemRequestBuilderGetRequestConfiguration{
					Options: []abstractions.RequestOption{mockRequestOption},
				}

				builder := &AttachmentItemRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				requestConfiguration := &AttachmentItemRequestBuilderGetRequestConfiguration{}

				builder := (*AttachmentItemRequestBuilder)(nil)

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

func TestAttachmentItemRequestBuilder_ToDeleteRequestInformation(t *testing.T) {
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
				mockHeaders.Add("Accept", "application/json")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.DELETE,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{})

				requestConfiguration := &AttachmentItemRequestBuilderDeleteRequestConfiguration{}

				builder := &AttachmentItemRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToDeleteRequestInformation(context.Background(), requestConfiguration)

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
				mockHeaders.Add("Accept", "application/json")
				mockHeaders.Add("test", "test1")
				mockURLTemplate := ""

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.DELETE,
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

				requestConfiguration := &AttachmentItemRequestBuilderDeleteRequestConfiguration{
					Headers: headers,
				}

				builder := &AttachmentItemRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToDeleteRequestInformation(context.Background(), requestConfiguration)

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
				mockHeaders.Add("Accept", "application/json")
				mockURLTemplate := ""

				mockRequestOption := mocking.NewMockRequestOption()
				mockRequestOption.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})

				mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
				mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternalRequestBuilder.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternalRequestBuilder.On("GetPathParameters").Return(mockPathParameters)

				expected := &abstractions.RequestInformation{
					Method:             abstractions.DELETE,
					UrlTemplate:        mockURLTemplate,
					PathParameters:     mockPathParameters,
					PathParametersAny:  mockPathParametersAny,
					QueryParameters:    mockQueryParameters,
					QueryParametersAny: mockQueryParametersAny,
					Headers:            mockHeaders,
				}

				expected.AddRequestOptions([]abstractions.RequestOption{mockRequestOption})

				requestConfiguration := &AttachmentItemRequestBuilderDeleteRequestConfiguration{
					Options: []abstractions.RequestOption{mockRequestOption},
				}

				builder := &AttachmentItemRequestBuilder{mockInternalRequestBuilder}

				requestInformation, err := builder.ToDeleteRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				requestConfiguration := &AttachmentItemRequestBuilderDeleteRequestConfiguration{}

				builder := (*AttachmentItemRequestBuilder)(nil)

				requestInformation, err := builder.ToDeleteRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, requestInformation)
				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}
