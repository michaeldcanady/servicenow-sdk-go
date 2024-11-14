package core

import (
	"errors"

	"github.com/RecoLabs/servicenow-sdk-go/internal"
)

var (
	ErrEmptyURI                = internal.ErrEmptyURI
	ErrNilPathParameters       = internal.ErrNilPathParameters
	ErrNilQueryParamters       = internal.ErrNilQueryParamters
	ErrMissingBasePathParam    = internal.ErrMissingBasePathParam
	ErrMissingBasePathTemplate = internal.ErrMissingBasePathTemplate
	ErrInvalidHeaderType       = internal.ErrInvalidHeaderType
	ErrEmptyRawUrl             = internal.ErrEmptyRawURL //nolint:stylecheck
	ErrMissingSchema           = internal.ErrMissingSchema
	ErrNilResponse             = internal.ErrNilResponse
	ErrNilResponseBody         = internal.ErrNilResponseBody
	ErrNilSource               = internal.ErrNilSource

	//Page Iterator
	ErrNilClient          = internal.ErrNilClient
	ErrNilResult          = internal.ErrNilResult
	ErrWrongResponseType  = internal.ErrWrongResponseType
	ErrParsing            = internal.ErrParsing
	ErrNilCallback        = internal.ErrNilCallback
	ErrNilConstructorFunc = errors.New("constructorFunc can't be nil")
)
