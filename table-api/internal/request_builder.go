package internal

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{unreleased}.
//
// RequestBuilder ...
type RequestBuilder interface {
	SendGet2(*core.RequestConfiguration) error
	SendDelete2(*core.RequestConfiguration) error
	SendPut2(*core.RequestConfiguration) error
	SendPost3(*core.RequestConfiguration) error
	ToHeadRequestInformation() (*core.RequestInformation, error)
}
