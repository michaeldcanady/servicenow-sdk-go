package kiota

import (
	"errors"
	"strings"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/utils"
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
	if utils.IsNil(rB) {
		return nil
	}

	return rB.PathParameters
}

func (rB *BaseRequestBuilder) SetPathParameters(pathParameters map[string]string) error {
	if utils.IsNil(rB) {
		return nil
	}

	if utils.IsNil(pathParameters) {
		return errors.New("pathParameters is nil")
	}

	rB.PathParameters = pathParameters
	return nil
}

func (rB *BaseRequestBuilder) GetRequestAdapter() abstractions.RequestAdapter {
	if utils.IsNil(rB) {
		return nil
	}

	return rB.RequestAdapter
}

func (rB *BaseRequestBuilder) SetRequestAdapter(requestAdapter abstractions.RequestAdapter) error {
	if utils.IsNil(rB) {
		return nil
	}

	if utils.IsNil(requestAdapter) {
		return errors.New("requestAdapter is nil")
	}

	rB.RequestAdapter = requestAdapter
	return nil
}

func (rB *BaseRequestBuilder) GetURLTemplate() string {
	if utils.IsNil(rB) {
		return ""
	}

	return rB.UrlTemplate
}

func (rB *BaseRequestBuilder) SetURLTemplate(urlTemplate string) error {
	if utils.IsNil(rB) {
		return nil
	}

	urlTemplate = strings.TrimSpace(urlTemplate)
	if urlTemplate == "" {
		return errors.New("urlTemplate is empty")
	}

	rB.UrlTemplate = urlTemplate
	return nil
}
