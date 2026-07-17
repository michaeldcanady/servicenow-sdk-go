// Package oauth2 provides a spec-compliant OAuth2 client for various grant types.
package oauth2

import "github.com/michaeldcanady/servicenow-sdk-go/internal/conversion"

// AuthMethod represents the method used by the client to authenticate with the authorization server.
type AuthMethod int

const (
	// AuthMethodUnknown represents an unknown or unsupported authentication method.
	AuthMethodUnknown AuthMethod = iota - 1
	// AuthMethodClientSecretPost sends the client credentials in the HTTP request body.
	AuthMethodClientSecretPost
	// AuthMethodClientSecretBasic sends the client credentials using HTTP Basic authentication.
	AuthMethodClientSecretBasic
	// AuthMethodNone represents a public client with no client secret.
	AuthMethodNone
)

var authMethodStrings = map[AuthMethod]string{
	AuthMethodClientSecretPost:  "client_secret_post",
	AuthMethodClientSecretBasic: "client_secret_basic",
	AuthMethodNone:              "none",
	AuthMethodUnknown:           "unknown",
}

// String returns the string representation of the AuthMethod.
func (e AuthMethod) String() string {
	return conversion.EnumString(authMethodStrings, e, authMethodStrings[AuthMethodUnknown])
}
