package internal

import "github.com/RecoLabs/servicenow-sdk-go/core"

type RequestBuilder interface {
	SendGet2(*core.RequestConfiguration) error
	SendDelete2(*core.RequestConfiguration) error
	SendPut2(*core.RequestConfiguration) error
	SendPost3(*core.RequestConfiguration) error
	ToHeadRequestInformation() (*core.RequestInformation, error)
}
