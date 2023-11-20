package core

import "errors"

var (
	ErrEmptyUri                = errors.New("uri cannot be empty")
	ErrNilPathParameters       = errors.New("uri template parameters cannot be nil")
	ErrNilQueryParamters       = errors.New("uri query parameters cannot be nil")
	ErrMissingBasePathParam    = errors.New("pathParameters must contain a value for \"baseurl\" for the URL to be built")
	ErrMissingBasePathTemplate = errors.New("template must contain a placeholder for \"{+baseurl}\" for the URL to be built")
	ErrInvalidHeaderType       = errors.New("headers must be a pointer or an http.Header")
	ErrEmptyRawUrl             = errors.New("empty raw URL")
	ErrMissingSchema           = errors.New("URL is missing schema")
	ErrNilResponse             = errors.New("Response is nil")
	ErrNilSource               = errors.New("source is nil")
)
