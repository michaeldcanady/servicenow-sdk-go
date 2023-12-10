package tableapi

import (
	"errors"
)

var (
	ErrNilClient         = errors.New("client can't be nil")
	ErrNilResponse       = errors.New("response can't be nil")
	ErrNilResult         = errors.New("result property missing in response object")
	ErrWrongResponseType = errors.New("incorrect Response Type")
	ErrParsing           = errors.New("parsing nextLink url failed")
)
