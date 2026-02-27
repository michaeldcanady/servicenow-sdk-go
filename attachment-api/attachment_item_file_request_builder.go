package attachmentapi

import (
	"context"
	"errors"

	model "github.com/michaeldcanady/servicenow-sdk-go/internal/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/kiota"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
	kiota.RequestBuilder
}

// newAttachmentItemFileRequestBuilderInternal instantiates a new AttachmentItemFileRequestBuilder with the provided requestBuilder
func newAttachmentItemFileRequestBuilderInternal(requestBuilder kiota.RequestBuilder) *AttachmentItemFileRequestBuilder {
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
		kiota.NewBaseRequestBuilder(requestAdapter, attachmentItemFileURLTemplate, pathParameters),
	)
}

// NewAttachmentItemFileRequestBuilder instantiates a new AttachmentItemFileRequestBuilder with custom parsable for table entries.
func NewAttachmentItemFileRequestBuilder(
	rawURL string,
	requestAdapter abstractions.RequestAdapter,
) *AttachmentItemFileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams[utils.RawURLKey] = rawURL
	return NewAttachmentItemFileRequestBuilderInternal(urlParams, requestAdapter)
}

// Get returns file with content using provided parameters
func (rB *AttachmentItemFileRequestBuilder) Get(ctx context.Context, requestConfiguration *AttachmentItemFileRequestBuilderGetRequestConfiguration) (*FileWithContentModel, error) {
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	if utils.IsNil(requestConfiguration) {
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
		"XXX": model.CreateServiceNowErrorFromDiscriminatorValue,
	}

	requestAdapter := rB.GetRequestAdapter()
	if utils.IsNil(requestAdapter) {
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
	if utils.IsNil(rB) || utils.IsNil(rB.RequestBuilder) {
		return nil, nil
	}

	requestInfo := abstractions.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(abstractions.GET, rB.GetURLTemplate(), rB.GetPathParameters())
	kiotaRequestInfo := &kiota.KiotaRequestInformation{RequestInformation: requestInfo}
	if !utils.IsNil(requestConfiguration) {
		if headers := requestConfiguration.Headers; !utils.IsNil(headers) {
			kiotaRequestInfo.Headers.AddAll(headers)
		}
		if options := requestConfiguration.Options; !utils.IsNil(options) {
			kiotaRequestInfo.AddRequestOptions(options)
		}
	}
	requestInfo.Headers.TryAdd("Accept", "*/*")
	return kiotaRequestInfo.RequestInformation, nil
}
