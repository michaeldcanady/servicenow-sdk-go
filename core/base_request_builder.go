package core

import (
	"errors"
	"strings"

	snerrors "github.com/michaeldcanady/servicenow-sdk-go/errors"
	"github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// RequestBuilder
type RequestBuilder interface {
	GetPathParameters() map[string]string
	SetPathParameters(map[string]string) error
	GetRequestAdapter() abstractions.RequestAdapter
	SetRequestAdapter(abstractions.RequestAdapter) error
	GetURLTemplate() string
	SetURLTemplate(string) error
}

// BaseRequestBuilder is the base every *api request builder in this SDK embeds.
type BaseRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

// NewBaseRequestBuilder instantiates a new BaseRequestBuilder.
func NewBaseRequestBuilder(
	requestAdapter abstractions.RequestAdapter,
	urlTemplate string,
	pathParameters map[string]string,
) *BaseRequestBuilder {
	if pathParameters == nil {
		pathParameters = make(map[string]string)
	}

	return &BaseRequestBuilder{
		abstractions.BaseRequestBuilder{
			PathParameters: pathParameters,
			UrlTemplate:    urlTemplate,
			RequestAdapter: requestAdapter,
		},
	}
}

// GetPathParameters returns the request builder's URL path parameters.
func (rB *BaseRequestBuilder) GetPathParameters() map[string]string {
	if conversion.IsNil(rB) {
		return nil
	}

	return rB.PathParameters
}

// SetPathParameters sets the request builder's URL path parameters.
func (rB *BaseRequestBuilder) SetPathParameters(pathParameters map[string]string) error {
	if conversion.IsNil(rB) {
		return snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(pathParameters) {
		return snerrors.ErrNilPathParameters
	}

	rB.PathParameters = pathParameters
	return nil
}

// GetRequestAdapter returns the request builder's RequestAdapter.
func (rB *BaseRequestBuilder) GetRequestAdapter() abstractions.RequestAdapter {
	if conversion.IsNil(rB) {
		return nil
	}

	return rB.RequestAdapter
}

// SetRequestAdapter sets the request builder's RequestAdapter.
func (rB *BaseRequestBuilder) SetRequestAdapter(requestAdapter abstractions.RequestAdapter) error {
	if conversion.IsNil(rB) {
		return snerrors.ErrNilRequestBuilder
	}

	if conversion.IsNil(requestAdapter) {
		return snerrors.ErrNilRequestAdapter
	}

	rB.RequestAdapter = requestAdapter
	return nil
}

// GetURLTemplate returns the request builder's URL template.
func (rB *BaseRequestBuilder) GetURLTemplate() string {
	if conversion.IsNil(rB) {
		return ""
	}

	return rB.UrlTemplate
}

// SetURLTemplate sets the request builder's URL template.
func (rB *BaseRequestBuilder) SetURLTemplate(urlTemplate string) error {
	if conversion.IsNil(rB) {
		return snerrors.ErrNilRequestBuilder
	}

	urlTemplate = strings.TrimSpace(urlTemplate)
	if urlTemplate == "" {
		return errors.New("urlTemplate is empty")
	}

	rB.UrlTemplate = urlTemplate
	return nil
}
