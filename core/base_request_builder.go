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

type BaseRequestBuilder struct {
	abstractions.BaseRequestBuilder
}

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

func (rB *BaseRequestBuilder) GetPathParameters() map[string]string {
	if conversion.IsNil(rB) {
		return nil
	}

	return rB.PathParameters
}

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

func (rB *BaseRequestBuilder) GetRequestAdapter() abstractions.RequestAdapter {
	if conversion.IsNil(rB) {
		return nil
	}

	return rB.RequestAdapter
}

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

func (rB *BaseRequestBuilder) GetURLTemplate() string {
	if conversion.IsNil(rB) {
		return ""
	}

	return rB.UrlTemplate
}

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
