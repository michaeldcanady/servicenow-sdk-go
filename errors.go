package servicenowsdkgo

import "errors"

var (
	ErrNilRequestInfo = errors.New("requestInfo cannot be nil")
	ErrNilContext     = errors.New("ctx cannot be nil")
)
