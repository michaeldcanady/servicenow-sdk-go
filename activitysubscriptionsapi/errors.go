package activitysubscriptionsapi

import "errors"

var (
	ErrNilContextQueryParameter         = errors.New("context query parameters cannot be nil")
	ErrNilOrEmptyQueryParameters        = errors.New("query parameters cannot be nil or empty")
	ErrNilContextInstanceQueryParameter = errors.New("context instance query parameters cannot be nil")
)
