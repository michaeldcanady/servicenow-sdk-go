package tests

import (
	"os"
	"strings"

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

	// Table API specific
	tableName   string
	sysID       string
	requestBody serialization.Parsable
	executeReq  func() (serialization.Parsable, error)
}

func isHeadless() bool {
	switch strings.ToLower(os.Getenv("SN_HEADLESS")) {
	case "true", "1":
		return true
	default:
		return false
	}
}

func isOffline() bool {
	switch strings.ToLower(os.Getenv("SN_OFFLINE")) {
	case "true", "1":
		return true
	default:
		return false
	}
}
