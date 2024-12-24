package core

import (
	"errors"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

var (
	// Deprecated: deprecated since v{unreleased}.
	ErrEmptyURI = internal.ErrEmptyURI
	// Deprecated: deprecated since v{unreleased}.
	ErrNilPathParameters = internal.ErrNilPathParameters
	// Deprecated: deprecated since v{unreleased}.
	ErrNilQueryParamters = internal.ErrNilQueryParamters
	// Deprecated: deprecated since v{unreleased}.
	ErrMissingBasePathParam = internal.ErrMissingBasePathParam
	// Deprecated: deprecated since v{unreleased}.
	ErrMissingBasePathTemplate = internal.ErrMissingBasePathTemplate
	// Deprecated: deprecated since v{unreleased}.
	ErrInvalidHeaderType = internal.ErrInvalidHeaderType
	// Deprecated: deprecated since v{unreleased}.
	ErrEmptyRawUrl = internal.ErrEmptyRawURL //nolint:stylecheck
	// Deprecated: deprecated since v{unreleased}.
	ErrMissingSchema = internal.ErrMissingSchema
	// Deprecated: deprecated since v{unreleased}.
	ErrNilResponse = internal.ErrNilResponse
	// Deprecated: deprecated since v{unreleased}.
	ErrNilResponseBody = internal.ErrNilResponseBody
	// Deprecated: deprecated since v{unreleased}.
	ErrNilSource = internal.ErrNilSource

	// Deprecated: deprecated since v{unreleased}.
	ErrNilClient = internal.ErrNilClient
	// Deprecated: deprecated since v{unreleased}.
	ErrNilResult = internal.ErrNilResult
	// Deprecated: deprecated since v{unreleased}.
	ErrWrongResponseType = internal.ErrWrongResponseType
	// Deprecated: deprecated since v{unreleased}.
	ErrParsing = internal.ErrParsing
	// Deprecated: deprecated since v{unreleased}.
	ErrNilCallback = internal.ErrNilCallback
	// Deprecated: deprecated since v{unreleased}.
	ErrNilConstructorFunc = errors.New("constructorFunc can't be nil")
)
