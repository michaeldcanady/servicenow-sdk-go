package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAttachmentItemFileRequestBuilderInternal(t *testing.T) {
	tests := []struct {
		name           string
		pathParameters map[string]string
		requestAdapter abstractions.RequestAdapter
	}{
		{
			name:           "Successful",
			pathParameters: map[string]string{},
			requestAdapter: mocking.NewMockRequestAdapter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentItemFileRequestBuilderInternal(tt.pathParameters, tt.requestAdapter)

			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentItemFileRequestBuilder{}, builder)
			assert.IsType(t, &internal.BaseRequestBuilder{}, builder.RequestBuilder)
			assert.Equal(t, tt.pathParameters, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestNewAttachmentItemFileRequestBuilder(t *testing.T) {
	tests := []struct {
		name           string
		rawURL         string
		requestAdapter abstractions.RequestAdapter
	}{
		{
			name:           "Successful",
			rawURL:         "",
			requestAdapter: mocking.NewMockRequestAdapter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlParams := map[string]string{internal.RawURLKey: tt.rawURL}

			builder := NewAttachmentItemFileRequestBuilder(tt.rawURL, tt.requestAdapter)

			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentItemFileRequestBuilder{}, builder)
			assert.IsType(t, &internal.BaseRequestBuilder{}, builder.RequestBuilder)
			assert.Equal(t, urlParams, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestAttachmentItemFileRequestBuilder_Get(t *testing.T) {
	// Reusable mocks
	mockPathParameters := map[string]string{}
	mockURLTemplate := ""
	mockContent := []byte("testing")

	tests := []struct {
		name                 string
		isNilBuilder         bool
		requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration
		setupMocks           func(t *testing.T) *mocking.MockRequestBuilder
		expectedContent      []byte
		expectedErr          error
	}{
		{
			name:                 "Successful",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			setupMocks: func(t *testing.T) *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendPrimitive", context.Background(), mock.Anything, "[]byte", mock.Anything).Return(mockContent, nil).Run(func(args mock.Arguments) {
					requestInformation := args.Get(1).(*abstractions.RequestInformation)
					opts := requestInformation.GetRequestOptions()
					for _, opt := range opts {
						if typedOpt, ok := opt.(*nethttplibrary.HeadersInspectionOptions); ok {
							respHeaders := abstractions.NewResponseHeaders()
							respHeaders.Add("X-Attachment-Metadata", "{}")
							typedOpt.ResponseHeaders = respHeaders
							break
						}
					}
				})

				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedContent: mockContent,
			expectedErr:     nil,
		},
		{
			name:                 "Send error",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			setupMocks: func(t *testing.T) *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendPrimitive", context.Background(), mock.Anything, "[]byte", mock.Anything).Return((*mocking.MockParsable)(nil), errors.New("send error"))

				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedContent: nil,
			expectedErr:     errors.New("send error"),
		},
		{
			name:                 "Wrong type",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			setupMocks: func(t *testing.T) *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendPrimitive", context.Background(), mock.Anything, "[]byte", mock.Anything).Return(mocking.NewMockParsable(), nil)

				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedContent: nil,
			expectedErr:     errors.New("resp is not []byte"),
		},
		{
			name:                 "Nil request adapter",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			setupMocks: func(t *testing.T) *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedContent: nil,
			expectedErr:     errors.New("requestAdapter is nil"),
		},
		{
			name:                 "Nil model",
			isNilBuilder:         true,
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			expectedContent:      nil,
			expectedErr:          nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemFileRequestBuilder
			if !tt.isNilBuilder {
				mockInternal := tt.setupMocks(t)
				builder = &AttachmentItemFileRequestBuilder{mockInternal}
			}

			result, err := builder.Get(context.Background(), tt.requestConfiguration)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				if tt.expectedContent != nil {
					resultContent, _ := result.GetContent()
					assert.Equal(t, tt.expectedContent, resultContent)
				} else {
					assert.Nil(t, result)
				}
			}
		})
	}
}

func TestAttachmentItemFileRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name                 string
		isNilBuilder         bool
		requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration
		setupMocks           func() *mocking.MockRequestBuilder
		expectedErr          error
		validate             func(t *testing.T, info *abstractions.RequestInformation)
	}{
		{
			name:                 "Successful - minimal",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			setupMocks: func() *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mocking.NewMockRequestAdapter())
				mockInternal.On("GetURLTemplate").Return("")
				mockInternal.On("GetPathParameters").Return(map[string]string{})
				return mockInternal
			},
			expectedErr: nil,
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Equal(t, abstractions.GET, info.Method)
			},
		},
		{
			name:                 "Successful - with headers",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{
				Headers: func() *abstractions.RequestHeaders {
					h := abstractions.NewRequestHeaders()
					h.Add("test", "test1")
					return h
				}(),
			},
			setupMocks: func() *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mocking.NewMockRequestAdapter())
				mockInternal.On("GetURLTemplate").Return("")
				mockInternal.On("GetPathParameters").Return(map[string]string{})
				return mockInternal
			},
			expectedErr: nil,
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				val := info.Headers.Get("test")
				assert.Equal(t, []string{"test1"}, val)
			},
		},
		{
			name:                 "Successful - with options",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{
				Options: func() []abstractions.RequestOption {
					mockOption := mocking.NewMockRequestOption()
					mockOption.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})
					return []abstractions.RequestOption{mockOption}
				}(),
			},
			setupMocks: func() *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mocking.NewMockRequestAdapter())
				mockInternal.On("GetURLTemplate").Return("")
				mockInternal.On("GetPathParameters").Return(map[string]string{})
				return mockInternal
			},
			expectedErr: nil,
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Len(t, info.GetRequestOptions(), 1)
			},
		},
		{
			name:                 "Nil model",
			isNilBuilder:         true,
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			expectedErr:          nil,
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Nil(t, info)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemFileRequestBuilder
			if !tt.isNilBuilder {
				mockInternal := tt.setupMocks()
				builder = &AttachmentItemFileRequestBuilder{mockInternal}
			}

			info, err := builder.ToGetRequestInformation(context.Background(), tt.requestConfiguration)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, info)
				}
			}
		})
	}
}
