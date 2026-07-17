package attachmentapi

import (
	"context"
	"errors"
	"testing"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
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
			name:           "Valid parameters",
			pathParameters: map[string]string{"sys_id": "test-sys-id"},
			requestAdapter: mocking.NewMockRequestAdapter(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentItemFileRequestBuilderInternal(tt.pathParameters, tt.requestAdapter)
			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentItemFileRequestBuilder{}, builder)
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
		expectedParams map[string]string
	}{
		{
			name:           "Valid URL",
			rawURL:         "https://example.com",
			requestAdapter: mocking.NewMockRequestAdapter(),
			expectedParams: map[string]string{internal.RawURLKey: "https://example.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewAttachmentItemFileRequestBuilder(tt.rawURL, tt.requestAdapter)
			assert.NotNil(t, builder)
			assert.IsType(t, &AttachmentItemFileRequestBuilder{}, builder)
			assert.Equal(t, tt.expectedParams, builder.GetPathParameters())
			assert.Equal(t, tt.requestAdapter, builder.GetRequestAdapter())
		})
	}
}

type attachmentItemFileGetTestCase struct {
	name                 string
	nilBuilder           bool
	requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration
	nilRequestAdapter    bool
	sendError            error
	sendResponse         any
	metadataHeaderValue  string
	expectedErr          error
	expectedErrorMessage string
}

// newAttachmentItemFileGetTestBuilder builds the *AttachmentItemFileRequestBuilder
// described by tt, wiring up the mock request adapter's SendPrimitive call - including
// injecting tt.metadataHeaderValue as a response header when present - needed for a Get call.
func newAttachmentItemFileGetTestBuilder(tt attachmentItemFileGetTestCase) *AttachmentItemFileRequestBuilder {
	if tt.nilBuilder {
		return nil
	}

	mockRequestAdapter := mocking.NewMockRequestAdapter()
	mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
	mockInternalRequestBuilder.On("GetURLTemplate").Return("")
	mockInternalRequestBuilder.On("GetPathParameters").Return(map[string]string{})

	if tt.nilRequestAdapter {
		mockInternalRequestBuilder.On("GetRequestAdapter").Return((*mocking.MockRequestAdapter)(nil))
		return &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
	}

	mockInternalRequestBuilder.On("GetRequestAdapter").Return(mockRequestAdapter)
	mockRequestAdapter.On("SendPrimitive", mock.Anything, mock.Anything, "[]byte", mock.Anything).Return(tt.sendResponse, tt.sendError).Run(func(args mock.Arguments) {
		if tt.metadataHeaderValue == "" {
			return
		}
		requestInformation := args.Get(1).(*abstractions.RequestInformation)
		for _, opt := range requestInformation.GetRequestOptions() {
			typedOpt, ok := opt.(*nethttplibrary.HeadersInspectionOptions)
			if !ok {
				continue
			}
			respHeaders := abstractions.NewResponseHeaders()
			respHeaders.Add(attachmentMetadataHeader, tt.metadataHeaderValue)
			typedOpt.ResponseHeaders = respHeaders
			break
		}
	})

	return &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
}

// assertAttachmentItemFileGetResult verifies the Get result described by tt.
func assertAttachmentItemFileGetResult(t *testing.T, tt attachmentItemFileGetTestCase, mockContent []byte, result *FileWithContent, err error) {
	t.Helper()

	if tt.nilBuilder {
		assert.Nil(t, result)
		assert.NoError(t, err)
		return
	}

	if tt.expectedErr != nil {
		assert.Error(t, err)
		assert.Equal(t, tt.expectedErr, err)
		assert.Nil(t, result)
		return
	}

	if tt.expectedErrorMessage != "" {
		assert.Error(t, err)
		assert.Contains(t, err.Error(), tt.expectedErrorMessage)
		assert.Nil(t, result)
		return
	}

	assert.NoError(t, err)
	if tt.sendResponse == nil {
		assert.Nil(t, result)
		return
	}
	assert.NotNil(t, result)
	content, _ := result.GetContent()
	assert.Equal(t, mockContent, content)
}

func TestAttachmentItemFileRequestBuilder_Get(t *testing.T) {
	mockContent := []byte("testing-content")

	tests := []attachmentItemFileGetTestCase{
		{
			name:                 "Successful with metadata",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			sendResponse:         mockContent,
			metadataHeaderValue:  `{"file_name":"test.txt"}`,
		},
		{
			name:                 "Send error",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			sendError:            errors.New("send error"),
			expectedErr:          errors.New("send error"),
		},
		{
			name:                 "Wrong response type",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			sendResponse:         "not-bytes-slice",
			expectedErr:          errors.New("resp is not []byte"),
		},
		{
			name:                 "Nil request adapter",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			nilRequestAdapter:    true,
			expectedErr:          snerrors.ErrNilRequestAdapter,
		},
		{
			name:                 "Nil response",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			sendResponse:         nil,
		},
		{
			name:                 "Nil builder",
			nilBuilder:           true,
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := newAttachmentItemFileGetTestBuilder(tt)
			result, err := builder.Get(context.Background(), tt.requestConfiguration)
			assertAttachmentItemFileGetResult(t, tt, mockContent, result, err)
		})
	}
}

type attachmentItemFileToGetRequestInfoTestCase struct {
	name                 string
	nilBuilder           bool
	nilRequestBuilder    bool
	requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration
	expectedErr          bool
	expectedHeaders      map[string][]string
	expectedOptionsKeys  []string
}

// newAttachmentItemFileToGetRequestInfoTestBuilder builds the
// *AttachmentItemFileRequestBuilder (or nil variants) described by tt.
func newAttachmentItemFileToGetRequestInfoTestBuilder(tt attachmentItemFileToGetRequestInfoTestCase) *AttachmentItemFileRequestBuilder {
	if tt.nilBuilder {
		return nil
	}

	mockInternalRequestBuilder := mocking.NewMockRequestBuilder()
	mockInternalRequestBuilder.On("GetURLTemplate").Return("")
	mockInternalRequestBuilder.On("GetPathParameters").Return(map[string]string{})

	if tt.nilRequestBuilder {
		return &AttachmentItemFileRequestBuilder{nil}
	}
	return &AttachmentItemFileRequestBuilder{mockInternalRequestBuilder}
}

// assertAttachmentItemFileToGetRequestInformation verifies the ToGetRequestInformation
// result described by tt.
func assertAttachmentItemFileToGetRequestInformation(t *testing.T, tt attachmentItemFileToGetRequestInfoTestCase, reqInfo *abstractions.RequestInformation, err error) {
	t.Helper()

	if tt.nilBuilder || tt.nilRequestBuilder {
		assert.Nil(t, reqInfo)
		assert.NoError(t, err)
		return
	}

	if tt.expectedErr {
		assert.Error(t, err)
		assert.Nil(t, reqInfo)
		return
	}

	assert.NoError(t, err)
	assert.NotNil(t, reqInfo)
	assert.Equal(t, abstractions.GET, reqInfo.Method)

	if tt.expectedHeaders != nil {
		for k, v := range tt.expectedHeaders {
			assert.Equal(t, v, reqInfo.Headers.Get(k))
		}
	}

	if len(tt.expectedOptionsKeys) > 0 {
		options := reqInfo.GetRequestOptions()
		assert.Equal(t, len(tt.expectedOptionsKeys), len(options))
		for i, opt := range options {
			assert.Equal(t, tt.expectedOptionsKeys[i], opt.GetKey().Key)
		}
	}
}

func TestAttachmentItemFileRequestBuilder_ToGetRequestInformation(t *testing.T) {
	tests := []attachmentItemFileToGetRequestInfoTestCase{
		{
			name:                 "Successful minimal",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{},
			expectedHeaders: map[string][]string{
				"accept": {"*/*"},
			},
		},
		{
			name:       "Nil builder",
			nilBuilder: true,
		},
		{
			name:              "Nil request builder internally",
			nilRequestBuilder: true,
		},
		{
			name: "With custom headers",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{
				Headers: func() *abstractions.RequestHeaders {
					h := abstractions.NewRequestHeaders()
					h.Add("custom-header", "custom-value")
					return h
				}(),
			},
			expectedHeaders: map[string][]string{
				"accept":        {"*/*"},
				"custom-header": {"custom-value"},
			},
		},
		{
			name: "With request options",
			requestConfiguration: &AttachmentItemFileRequestBuilderGetRequestConfiguration{
				Options: func() []abstractions.RequestOption {
					opt := mocking.NewMockRequestOption()
					opt.On("GetKey").Return(abstractions.RequestOptionKey{Key: "key"})
					return []abstractions.RequestOption{opt}
				}(),
			},
			expectedHeaders: map[string][]string{
				"accept": {"*/*"},
			},
			expectedOptionsKeys: []string{"key"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := newAttachmentItemFileToGetRequestInfoTestBuilder(tt)
			reqInfo, err := builder.ToGetRequestInformation(context.Background(), tt.requestConfiguration)
			assertAttachmentItemFileToGetRequestInformation(t, tt, reqInfo, err)
		})
	}
}
