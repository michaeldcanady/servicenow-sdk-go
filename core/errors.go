package core

import "github.com/michaeldcanady/servicenow-sdk-go/internal/core"

var (
	ErrEmptyURI                = core.ErrEmptyURI
	ErrNilPathParameters       = core.ErrNilPathParameters
	ErrNilQueryParamters       = core.ErrNilQueryParamters
	ErrMissingBasePathParam    = core.ErrMissingBasePathParam
	ErrMissingBasePathTemplate = core.ErrMissingBasePathTemplate
	ErrInvalidHeaderType       = core.ErrInvalidHeaderType
	ErrEmptyRawUrl             = core.ErrEmptyRawURL //nolint:stylecheck
	ErrMissingSchema           = core.ErrMissingSchema
	ErrNilResponse             = core.ErrNilResponse
	ErrNilResponseBody         = core.ErrNilResponseBody
	ErrNilSource               = core.ErrNilSource
)
