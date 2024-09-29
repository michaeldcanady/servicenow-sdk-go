package internal

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type RequestBuilder interface {
	SendPost3(config *core.RequestConfiguration) error
}
