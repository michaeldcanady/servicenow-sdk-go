package core

import "errors"

var (
	ErrNilResponse     = errors.New("Response is nil")
	ErrNilResponseBody = errors.New("Response body is nil")
)
