package internal

import (
	"github.com/RecoLabs/servicenow-sdk-go/internal/core"
)

var (
	ErrEmptyURI                = core.ErrEmptyURI
	ErrNilPathParameters       = core.ErrNilPathParameters
	ErrNilQueryParamters       = core.ErrNilQueryParamters
	ErrMissingBasePathParam    = core.ErrMissingBasePathParam
	ErrMissingBasePathTemplate = core.ErrMissingBasePathTemplate
	ErrInvalidHeaderType       = core.ErrInvalidHeaderType
	ErrEmptyRawURL             = core.ErrEmptyRawURL
	ErrMissingSchema           = core.ErrMissingSchema
	ErrNilResponse             = core.ErrNilResponse
	ErrNilResponseBody         = core.ErrNilResponseBody
	ErrNilSource               = core.ErrNilSource
	ErrNilCredential           = core.ErrNilCredential

	//Page Iterator
	ErrNilClient         = core.ErrNilClient
	ErrNilResult         = core.ErrNilResult
	ErrWrongResponseType = core.ErrWrongResponseType
	ErrParsing           = core.ErrParsing
	ErrNilCallback       = core.ErrNilCallback

	//Authorization Provider
	ErrNilRequest = core.ErrNilRequest
)
