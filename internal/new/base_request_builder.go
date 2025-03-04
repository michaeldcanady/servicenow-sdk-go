package internal

import abstractions "github.com/microsoft/kiota-abstractions-go"

type BaseRequestBuilder interface {
	GetPathParameters() map[string]string
	SetPathParameters(map[string]string)
	GetRequestAdapter() abstractions.RequestAdapter
	SetRequestAdapter(abstractions.RequestAdapter)
	GetURLTemplate() string
	SetURLTemplate(string)
}

type RequestBuilder struct {
	abstractions.BaseRequestBuilder
}

func NewRequestBuilder(requestAdapter abstractions.RequestAdapter, urlTemplate string, pathParameters map[string]string) *RequestBuilder {
	return &RequestBuilder{
		*abstractions.NewBaseRequestBuilder(requestAdapter, urlTemplate, pathParameters),
	}
}

func (rB *RequestBuilder) GetPathParameters() map[string]string {
	return rB.PathParameters
}

func (rB *RequestBuilder) SetPathParameters(pathParameters map[string]string) {
	rB.PathParameters = pathParameters
}

func (rB *RequestBuilder) GetRequestAdapter() abstractions.RequestAdapter {
	return rB.RequestAdapter
}

func (rB *RequestBuilder) SetRequestAdapter(requestAdapter abstractions.RequestAdapter) {
	rB.RequestAdapter = requestAdapter
}

func (rB *RequestBuilder) GetURLTemplate() string {
	return rB.UrlTemplate
}

func (rB *RequestBuilder) SetURLTemplate(urlTemplate string) {
	rB.UrlTemplate = urlTemplate
}
