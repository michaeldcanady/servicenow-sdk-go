package attachmentapi

import (
	"context"
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	nethttplibrary "github.com/microsoft/kiota-http-go"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
)

const (
	attachmentItemFileURLTemplate = "{+baseurl}/api/now/v1/attachment{/sys_id}/file"
)

// AttachmentItemFileRequestBuilder provides operations to manage Service-Now attachments.
type AttachmentItemFileRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewAttachmentItemFileRequestBuilderInternal instantiates a new AttachmentItemFileRequestBuilder with custom parsable for table entries.
func NewAttachmentItemFileRequestBuilderInternal(
	pathParameters map[string]string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemFileRequestBuilder {
	m := &AttachmentItemFileRequestBuilder{
		BaseRequestBuilder: *abstractions.NewBaseRequestBuilder(requestAdapter, attachmentItemFileURLTemplate, pathParameters),
	}
	return m
}

// NewAttachmentItemFileRequestBuilder instantiates a new AttachmentItemFileRequestBuilder with custom parsable for table entries.
func NewAttachmentItemFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[rawURLKey] = rawURL
	return NewAttachmentItemFileRequestBuilderInternal(urlParams, requestAdapter)
}

func (rB *AttachmentItemFileRequestBuilder) Get(ctx context.Context, requestConfiguration *TableAttachmentItemFileRequestBuilderGetRequestConfiguration) (FileWithContentable, error) {
	if internal.IsNil(rB) {
		return nil, nil
	}

	if internal.IsNil(requestConfiguration) {
		requestConfiguration = &TableAttachmentItemFileRequestBuilderGetRequestConfiguration{}
	}

	opts := nethttplibrary.NewHeadersInspectionOptions()
	opts.InspectResponseHeaders = true

	requestConfiguration.Options = append(requestConfiguration.Options, opts)

	requestInfo, err := rB.toGetRequestInformation(ctx, nil, requestConfiguration)
	if err != nil {
		return nil, err
	}

	// TODO: add error factory
	errorMapping := abstractions.ErrorMappings{}

	resp, err := rB.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "byte", errorMapping)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	var data = make([]byte, 0, len(resp))
	for _, elem := range resp {
		byteElem, ok := elem.(byte)
		if !ok {
			return nil, errors.New("elem is not byte")
		}
		data = append(data, byteElem)
	}

	metadata := opts.ResponseHeaders.Get("X-Attachment-Metadata")[0]

	var node serialization.ParseNode

	node, err = jsonserialization.NewJsonParseNode([]byte(metadata))
	if err != nil {
		return nil, err
	}

	file, err := node.GetObjectValue(CreateFileWithContentFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	typedFile, ok := file.(FileWithContentable)
	if !ok {
		return nil, errors.New("file is not FileWithContentable")
	}

	if err := typedFile.setContent(data); err != nil {
		return nil, err
	}

	return typedFile, nil
}

func (rB *AttachmentItemFileRequestBuilder) toGetRequestInformation(_ context.Context, _ any, requestConfiguration *TableAttachmentItemFileRequestBuilderGetRequestConfiguration) (*abstractions.RequestInformation, error) { //nolint:unparam
	if internal.IsNil(rB) {
		return nil, nil
	}

	requestInfo := intHttp.NewRequestInformationWithMethodAndURLTemplateAndPathParameters(abstractions.GET, rB.UrlTemplate, rB.PathParameters)
	if !internal.IsNil(requestConfiguration) {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "*/*")

	return &requestInfo.RequestInformation, nil
}
