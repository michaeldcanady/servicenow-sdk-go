package internal

import "github.com/michaeldcanady/servicenow-sdk-go/core"

// Deprecated: deprecated since v{unreleased}.
//
// RequestBuilder ...
type RequestBuilder interface {
	SendGet2(*core.RequestConfiguration) error                   //nolint: staticcheck
	SendDelete2(*core.RequestConfiguration) error                //nolint: staticcheck
	SendPut2(*core.RequestConfiguration) error                   //nolint: staticcheck
	SendPost3(*core.RequestConfiguration) error                  //nolint: staticcheck
	ToHeadRequestInformation() (*core.RequestInformation, error) //nolint: staticcheck
}
