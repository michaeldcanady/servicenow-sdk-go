package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	newInternal "github.com/michaeldcanady/servicenow-sdk-go/internal/new"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
)

const (
	// attachmentItemFileURLTemplate the url template for Service-Now's attachment item file endpoint
	attachmentItemFileURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}/file"
)

// AttachmentItemFileRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemFileRequestBuilder struct {
	newInternal.RequestBuilder
}

// NewV1CompatibleAttachmentUploadRequestBuilder instantiates a new AttachmentUploadRequestBuilder.
func NewV1CompatibleAttachmentItemFileRequestBuilder(
	pathParameters map[string]string,
	client core.Client,
) *AttachmentItemFileRequestBuilder {
	authProvider := core.NewAPIV1ClientAdapter(client)
	adapter, _ := nethttplibrary.NewNetHttpRequestAdapter(authProvider)

	return newAttachmentItemFileRequestBuilderInternal(
		newInternal.NewBaseRequestBuilder(adapter, attachmentURLTemplate, pathParameters),
	)
}

// newAttachmentItemFileRequestBuilderInternal instantiates a new AttachmentItemFileRequestBuilder with the provided requestBuilder
func newAttachmentItemFileRequestBuilderInternal(requestBuilder newInternal.RequestBuilder) *AttachmentItemFileRequestBuilder {
	m := &AttachmentItemFileRequestBuilder{
		requestBuilder,
	}
	return m
}

// NewAttachmentItemFileRequestBuilderInternal instantiates a new AttachmentItemFileRequestBuilder with custom parsable for table entries.
func NewAttachmentItemFileRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemFileRequestBuilder {
	return newAttachmentItemFileRequestBuilderInternal(
		newInternal.NewBaseRequestBuilder(requestAdapter, attachmentItemFileURLTemplate, pathParameters),
	)
}

// NewAttachmentItemFileRequestBuilder instantiates a new AttachmentItemFileRequestBuilder with custom parsable for table entries.
func NewAttachmentItemFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[newInternal.RawURLKey] = rawURL
	return NewAttachmentItemFileRequestBuilderInternal(urlParams, requestAdapter)
}

// Get returns file with content using provided parameters
func (rB *AttachmentItemFileRequestBuilder) Get(ctx context.Context, requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration) (*FileWithContentModel, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &AttachmentItemFileRequestBuilderGetRequestConfiguration{}
	}

	opts := nethttplibrary.NewHeadersInspectionOptions()
	opts.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, opts)

	requestInfo, err := rB.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}

	errorMapping := abstractions.ErrorMappings{
		"XXX": newInternal.CreateServiceNowErrorFromDiscriminatorValue,
	}

	requestAdapter := rB.GetRequestAdapter()
	if internal.IsNil(requestAdapter) {
		return nil, errors.New("requestAdapter is nil")
	}

	resp, err := rB.GetRequestAdapter().SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	typedResp, ok := resp.([]byte)
	if !ok {
		return nil, errors.New("resp is not []byte")
	}

	metadata := opts.ResponseHeaders.Get("x-attachment-metadata")[0]

	var node serialization.ParseNode

	node, err = jsonserialization.NewJsonParseNode([]byte(metadata))
	if err != nil {
		return nil, err
	}

	file, err := node.GetObjectValue(CreateFileWithContentFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	typedFile, ok := file.(*FileWithContentModel)
	if !ok {
		return nil, errors.New("file is not *FileWithContentModel")
	}

	if err := typedFile.SetContent(typedResp); err != nil {
		return nil, err
	}

	return typedFile, nil
}

// ToGetRequestInformation converts request configurations to Get request information.
func (rB *AttachmentItemFileRequestBuilder) ToGetRequestInformation(_ context.Context, requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) {
	if internal.IsNil(rB) || internal.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &newInternal.KiotaRequestInformation{RequestInformation: requestInfo}
	if !internal.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !internal.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !internal.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd("Accept", "*/*")
	return kiotaRequestInfo.RequestInformation, nil
}
