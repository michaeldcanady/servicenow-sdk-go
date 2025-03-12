package internal

import (
	"errors"
	"strings"

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
	if IsNil(rB) {
		return nil
	}

	return rB.PathParameters
}

func (rB *BaseRequestBuilder) SetPathParameters(pathParameters map[string]string) error {
	if IsNil(rB) {
		return nil
	}

	if IsNil(pathParameters) {
		return errors.New("pathParameters is nil")
	}

	rB.PathParameters = pathParameters
	return nil
}

func (rB *BaseRequestBuilder) GetRequestAdapter() abstractions.RequestAdapter {
	if IsNil(rB) {
		return nil
	}

	return rB.RequestAdapter
}

func (rB *BaseRequestBuilder) SetRequestAdapter(requestAdapter abstractions.RequestAdapter) error {
	if IsNil(rB) {
		return nil
	}

	if IsNil(requestAdapter) {
		return errors.New("requestAdapter is nil")
	}

	rB.RequestAdapter = requestAdapter
	return nil
}

func (rB *BaseRequestBuilder) GetURLTemplate() string {
	if IsNil(rB) {
		return ""
	}

	return rB.UrlTemplate
}

func (rB *BaseRequestBuilder) SetURLTemplate(urlTemplate string) error {
	if IsNil(rB) {
		return nil
	}

	urlTemplate = strings.TrimSpace(urlTemplate)
	if urlTemplate == "" {
		return errors.New("urlTemplate is empty")
	}

	rB.UrlTemplate = urlTemplate
	return nil
}
