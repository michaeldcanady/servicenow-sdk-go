package internal

import (
	intHttp "github.com/michaeldcanady/servicenow-sdk-go/internal/http"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

func ConfigureRequestInformation[T any](request *intHttp.KiotaRequestInformation, config *abstractions.RequestConfiguration[T]) {
	if request == nil {
		return
	}
	if config == nil {
		return
	}
	if params := config.QueryParameters; !IsNil(params) {
		request.AddQueryParameters(*params)
	}
	if headers := config.Headers; !IsNil(headers) {
		request.Headers.AddAll(headers)
	}
	if options := config.Options; !IsNil(options) {
		request.AddRequestOptions(options)
	}
}
