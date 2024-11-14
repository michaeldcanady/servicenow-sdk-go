package internal

import "github.com/RecoLabs/servicenow-sdk-go/core"

type RequestBuilder interface {
	SendPost3(config *core.RequestConfiguration) error
}
