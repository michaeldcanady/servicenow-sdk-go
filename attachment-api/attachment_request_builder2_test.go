package attachmentapi

import (
	"context"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
)

func TestNewAttachmentRequestBuilder2Internal(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{}
				requestAdapter := mocking.NewMockRequestAdapter()

				builder := NewAttachmentRequestBuilder2Internal(pathParameters, requestAdapter)

				assert.IsType(t, &AttachmentRequestBuilder2{}, builder)
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

func TestNewAttachmentRequestBuilder2(t *testing.T) {
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

				builder := NewAttachmentRequestBuilder2(rawURL, requestAdapter)

				assert.IsType(t, &AttachmentRequestBuilder2{}, builder)
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

func TestAttachmentRequestBuilder2_ByID(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{
		{
			name: "Successful",
			test: func(t *testing.T) {
				pathParameters := map[string]string{sysIDKey: "id"}
				requestAdapter := mocking.NewMockRequestAdapter()

				internalRequestBuilder := mocking.NewMockRequestBuilder()
				internalRequestBuilder.On("GetPathParameters").Return(pathParameters)
				internalRequestBuilder.On("GetRequestAdapter").Return(requestAdapter)

				builder := &AttachmentRequestBuilder2{internalRequestBuilder}

				itemBuilder := builder.ByID("id")

				assert.Equal(t, &AttachmentRequestBuilder2{
					&newInternal.BaseRequestBuilder{
						BaseRequestBuilder: abstractions.BaseRequestBuilder{
							PathParameters: pathParameters,
							RequestAdapter: requestAdapter,
							UrlTemplate:    attachmentFileURLTemplate,
						},
					},
				}, itemBuilder)
			},
		},
		{
			name: "Nil requestBuilder",
			test: func(t *testing.T) {
				builder := (*AttachmentRequestBuilder2)(nil)
				itemBuilder := builder.ByID("id")

				assert.Nil(t, itemBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachmentRequestBuilder2_File(t *testing.T) {
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

				builder := &AttachmentRequestBuilder2{internalRequestBuilder}

				itemBuilder := builder.File()

				assert.Equal(t, &AttachmentFileRequestBuilder{
					&newInternal.BaseRequestBuilder{
						BaseRequestBuilder: abstractions.BaseRequestBuilder{
							PathParameters: pathParameters,
							RequestAdapter: requestAdapter,
							UrlTemplate:    attachmentFileURLTemplate,
						},
					},
				}, itemBuilder)
			},
		},
		{
			name: "Nil requestBuilder",
			test: func(t *testing.T) {
				builder := (*AttachmentRequestBuilder2)(nil)
				itemBuilder := builder.File()

				assert.Nil(t, itemBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachmentRequestBuilder2_Upload(t *testing.T) {
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

				builder := &AttachmentRequestBuilder2{internalRequestBuilder}

				itemBuilder := builder.Upload()

				assert.Equal(t, &AttachmentUploadRequestBuilder{
					&newInternal.BaseRequestBuilder{
						BaseRequestBuilder: abstractions.BaseRequestBuilder{
							PathParameters: pathParameters,
							RequestAdapter: requestAdapter,
							UrlTemplate:    attachmentURLTemplate,
						},
					},
				}, itemBuilder)
			},
		},
		{
			name: "Nil requestBuilder",
			test: func(t *testing.T) {
				builder := (*AttachmentRequestBuilder2)(nil)
				itemBuilder := builder.Upload()

				assert.Nil(t, itemBuilder)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

// TODO: (TestAttachmentRequestBuilder2_Get) Add tests
func TestAttachmentRequestBuilder2_Get(t *testing.T) {
	tests := []struct {
		name string
		test func(*testing.T)
	}{}

	for _, test := range tests {
		t.Run(test.name, test.test)
	}
}

func TestAttachmentRequestBuilder2_ToGetRequestInformation(t *testing.T) {
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

				requestConfiguration := &AttachmentRequestBuilder2GetRequestConfiguration{}

				builder := &AttachmentRequestBuilder2{mockInternalRequestBuilder}

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

				requestConfiguration := &AttachmentRequestBuilder2GetRequestConfiguration{
					Headers: headers,
				}

				builder := &AttachmentRequestBuilder2{mockInternalRequestBuilder}

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

				requestConfiguration := &AttachmentRequestBuilder2GetRequestConfiguration{
					Options: []abstractions.RequestOption{mockRequestOption},
				}

				builder := &AttachmentRequestBuilder2{mockInternalRequestBuilder}

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Successful - with query parameters",
			test: func(t *testing.T) {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockPathParameters := map[string]string{}
				mockPathParametersAny := map[string]any{}
				mockQueryParameters := map[string]string{"sysparm_limit": "1"}
				mockQueryParametersAny := map[string]any{"sysparm_limit": []interface{}{"1"}}
				mockHeaders := &abstractions.RequestHeaders{}
				mockHeaders.Add("accept", "application/json")
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

				queryParameters := &AttachmentRequestBuilder2GetQueryParameters{
					SysparmLimit: newInternal.ToPointer(1),
				}

				requestConfiguration := &AttachmentRequestBuilder2GetRequestConfiguration{
					QueryParameters: queryParameters,
				}

				builder := &AttachmentRequestBuilder2{mockInternalRequestBuilder}

				requestInformation, err := builder.ToGetRequestInformation(context.Background(), requestConfiguration)

				assert.Nil(t, err)
				assert.Equal(t, expected, requestInformation)
			},
		},
		{
			name: "Nil model",
			test: func(t *testing.T) {
				requestConfiguration := &AttachmentRequestBuilder2GetRequestConfiguration{}

				builder := (*AttachmentRequestBuilder2)(nil)

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
