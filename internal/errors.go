package internal

import "errors"

var (
	ErrEmptyURI                = errors.New("uri cannot be empty")
	ErrNilPathParameters       = errors.New("uri template parameters cannot be nil")
	ErrNilQueryParamters       = errors.New("uri query parameters cannot be nil")
	ErrMissingBasePathParam    = errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built")
	ErrMissingBasePathTemplate = errors.New("template must contain a placeholder for \"{+baseurl}\" for the URL to be built")
	ErrInvalidHeaderType       = errors.New("headers must be a pointer or an http.Header")
	ErrEmptyRawURL             = errors.New("empty raw URL")
	ErrMissingSchema           = errors.New("URL is missing schema")
	ErrNilResponse             = errors.New("Response is nil")
	ErrNilResponseBody         = errors.New("response body is nil")
	ErrNilSource               = errors.New("source is nil")
	ErrNilCredential           = errors.New("credential is nil")

	//Page Iterator
	ErrNilClient         = errors.New("client can't be nil")
	ErrNilResult         = errors.New("result property missing in response object")
	ErrWrongResponseType = errors.New("incorrect Response Type")
	ErrParsing           = errors.New("parsing nextLink url failed")
	ErrNilCallback       = errors.New("callback can't be nil")

	//Authorization Provider
	ErrNilRequest = errors.New("request can't be nil")
)
