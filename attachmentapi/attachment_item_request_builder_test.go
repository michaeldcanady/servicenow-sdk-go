package attachmentapi

import (
	"context"
	"errors"
	"testing"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/mocking"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAttachmentItemRequestBuilderInternal(t *testing.T) {
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
			builder := NewAttachmentItemRequestBuilderInternal(tt.pathParameters, tt.requestAdapter)

			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentItemRequestBuilder{}, builder)
			assert.IsType(t, &internal.BaseRequestBuilder{}, builder.RequestBuilder)
			assert.Equal(t, tt.pathParameters, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestNewAttachmentItemRequestBuilder(t *testing.T) {
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
			builder := NewAttachmentItemRequestBuilder(tt.rawURL, tt.requestAdapter)

			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentItemRequestBuilder{}, builder)
			assert.IsType(t, &internal.BaseRequestBuilder{}, builder.RequestBuilder)
			assert.Equal(t, map[string]string{internal.RawURLKey: tt.rawURL}, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

func TestAttachmentItemRequestBuilder_File(t *testing.T) {
	pathParams := map[string]string{}
	reqAdapter := mocking.NewMockRequestAdapter()

	tests := []struct {
		name         string
		isNilBuilder bool
		setupMocks   func() *mocking.MockRequestBuilder
	}{
		{
			name: "Successful",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetPathParameters").Return(pathParams)
				mockInternal.On("GetRequestAdapter").Return(reqAdapter)
				return mockInternal
			},
		},
		{
			name:         "Nil requestBuilder",
			isNilBuilder: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemRequestBuilder
			if !tt.isNilBuilder {
				builder = &AttachmentItemRequestBuilder{tt.setupMocks()}
			}
			fileBuilder := builder.File()

			if tt.isNilBuilder {
				assert.Nil(t, fileBuilder)
			} else {
				assert.NotNil(t, fileBuilder)
				assert.Equal(t, pathParams, fileBuilder.GetPathParameters())
				assert.Equal(t, reqAdapter, fileBuilder.GetRequestAdapter())
			}
		})
	}
}

func TestAttachmentItemRequestBuilder_Get(t *testing.T) {
	mockPathParameters := map[string]string{}
	mockURLTemplate := ""
	mockParsable := internal.NewBaseServiceNowItemResponse[*Attachment](CreateAttachmentFromDiscriminatorValue)

	tests := []struct {
		name         string
		isNilBuilder bool
		setupMocks   func() *mocking.MockRequestBuilder
		expectedRes  any
		expectedErr  error
	}{
		{
			name: "Successful",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("Send", context.Background(), mock.Anything, mock.Anything, mock.Anything).Return(mockParsable, nil)
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedRes: mockParsable,
			expectedErr: nil,
		},
		{
			name: "Send Error",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("Send", context.Background(), mock.Anything, mock.Anything, mock.Anything).Return((*mocking.MockParsable)(nil), errors.New("send error"))
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedErr: errors.New("send error"),
		},
		{
			name: "Wrong type",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("Send", context.Background(), mock.Anything, mock.Anything, mock.Anything).Return(mocking.NewMockParsable(), nil)
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedErr: errors.New("res is not ServiceNowItemResponse[*Attachment]"),
		},
		{
			name: "Nil response",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("Send", context.Background(), mock.Anything, mock.Anything, mock.Anything).Return((*mocking.MockParsable)(nil), nil)
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedErr: errors.New("response is nil"),
		},
		{
			name:        "Nil inner model",
			setupMocks:  func() *mocking.MockRequestBuilder { return &mocking.MockRequestBuilder{} }, // Mocking is not used here but to satisfy interface
			expectedErr: nil,
		},
		{
			name:         "Nil model",
			isNilBuilder: true,
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemRequestBuilder
			if tt.name == "Nil inner model" {
				builder = &AttachmentItemRequestBuilder{nil}
			} else if !tt.isNilBuilder {
				builder = &AttachmentItemRequestBuilder{tt.setupMocks()}
			}

			result, err := builder.Get(context.Background(), nil)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				if !tt.isNilBuilder && tt.name != "Nil inner model" {
					assert.Equal(t, tt.expectedRes, result)
				} else {
					assert.Nil(t, result)
				}
			}
		})
	}
}

func TestAttachmentItemRequestBuilder_Delete(t *testing.T) {
	mockPathParameters := map[string]string{}
	mockURLTemplate := ""

	tests := []struct {
		name         string
		isNilBuilder bool
		setupMocks   func() *mocking.MockRequestBuilder
		expectedErr  error
	}{
		{
			name: "Successful",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendNoContent", context.Background(), mock.Anything, mock.Anything).Return(nil)
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedErr: nil,
		},
		{
			name: "SendNoContext Error",
			setupMocks: func() *mocking.MockRequestBuilder {
				mockRequestAdapter := mocking.NewMockRequestAdapter()
				mockRequestAdapter.On("SendNoContent", context.Background(), mock.Anything, mock.Anything).Return(errors.New("send error"))
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mockRequestAdapter)
				mockInternal.On("GetURLTemplate").Return(mockURLTemplate)
				mockInternal.On("GetPathParameters").Return(mockPathParameters)
				return mockInternal
			},
			expectedErr: errors.New("send error"),
		},
		{
			name:        "Nil inner model",
			setupMocks:  func() *mocking.MockRequestBuilder { return &mocking.MockRequestBuilder{} },
			expectedErr: nil,
		},
		{
			name:         "Nil model",
			isNilBuilder: true,
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemRequestBuilder
			if tt.name == "Nil inner model" {
				builder = &AttachmentItemRequestBuilder{nil}
			} else if !tt.isNilBuilder {
				builder = &AttachmentItemRequestBuilder{tt.setupMocks()}
			}

			err := builder.Delete(context.Background(), nil)

			if tt.expectedErr != nil {
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAttachmentItemRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []struct {
		name                 string
		isNilBuilder         bool
		requestConfiguration *AttachmentItemRequestBuilderGetRequestConfiguration
		setupMocks           func() *mocking.MockRequestBuilder
		validate             func(t *testing.T, info *abstractions.RequestInformation)
	}{
		{
			name:                 "Successful - minimal",
			requestConfiguration: &AttachmentItemRequestBuilderGetRequestConfiguration{},
			setupMocks: func() *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mocking.NewMockRequestAdapter())
				mockInternal.On("GetURLTemplate").Return("")
				mockInternal.On("GetPathParameters").Return(map[string]string{})
				return mockInternal
			},
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Equal(t, abstractions.GET, info.Method)
			},
		},
		{
			name:                 "Successful - with headers",
			requestConfiguration: &AttachmentItemRequestBuilderGetRequestConfiguration{
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
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				val := info.Headers.Get("test")
				assert.Equal(t, []string{"test1"}, val)
			},
		},
		{
			name:                 "Successful - with options",
			requestConfiguration: &AttachmentItemRequestBuilderGetRequestConfiguration{
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
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Len(t, info.GetRequestOptions(), 1)
			},
		},
		{
			name:                 "Nil model",
			isNilBuilder:         true,
			requestConfiguration: &AttachmentItemRequestBuilderGetRequestConfiguration{},
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Nil(t, info)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemRequestBuilder
			if !tt.isNilBuilder {
				builder = &AttachmentItemRequestBuilder{tt.setupMocks()}
			}

			info := builder.ToGetRequestInformation(context.Background(), tt.requestConfiguration)

			if tt.validate != nil {
				tt.validate(t, info)
			}
		})
	}
}

func TestAttachmentItemRequestBuilder_ToDeleteRequestInformation(t *testing.T) {
	tests := []struct {
		name                 string
		isNilBuilder         bool
		requestConfiguration *AttachmentItemRequestBuilderDeleteRequestConfiguration
		setupMocks           func() *mocking.MockRequestBuilder
		validate             func(t *testing.T, info *abstractions.RequestInformation)
	}{
		{
			name:                 "Successful - minimal",
			requestConfiguration: &AttachmentItemRequestBuilderDeleteRequestConfiguration{},
			setupMocks: func() *mocking.MockRequestBuilder {
				mockInternal := mocking.NewMockRequestBuilder()
				mockInternal.On("GetRequestAdapter").Return(mocking.NewMockRequestAdapter())
				mockInternal.On("GetURLTemplate").Return("")
				mockInternal.On("GetPathParameters").Return(map[string]string{})
				return mockInternal
			},
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Equal(t, abstractions.DELETE, info.Method)
			},
		},
		{
			name:                 "Successful - with headers",
			requestConfiguration: &AttachmentItemRequestBuilderDeleteRequestConfiguration{
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
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				val := info.Headers.Get("test")
				assert.Equal(t, []string{"test1"}, val)
			},
		},
		{
			name:                 "Successful - with options",
			requestConfiguration: &AttachmentItemRequestBuilderDeleteRequestConfiguration{
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
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Len(t, info.GetRequestOptions(), 1)
			},
		},
		{
			name:                 "Nil model",
			isNilBuilder:         true,
			requestConfiguration: &AttachmentItemRequestBuilderDeleteRequestConfiguration{},
			validate: func(t *testing.T, info *abstractions.RequestInformation) {
				assert.Nil(t, info)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var builder *AttachmentItemRequestBuilder
			if !tt.isNilBuilder {
				builder = &AttachmentItemRequestBuilder{tt.setupMocks()}
			}

			info := builder.ToDeleteRequestInformation(context.Background(), tt.requestConfiguration)

			if tt.validate != nil {
				tt.validate(t, info)
			}
		})
	}
}
