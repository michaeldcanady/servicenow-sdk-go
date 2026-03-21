package tests

import (
	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/microsoft/kiota-abstractions-go/serialization"
)

type sharedTestContext struct {
	instance string
	provider authentication.AuthenticationProvider
	client   *servicenowsdkgo.ServiceNowServiceClient
	resp     serialization.Parsable
	err      error
}
