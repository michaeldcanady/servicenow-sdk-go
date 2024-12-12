package authentication

import (
	"github.com/michaeldcanady/servicenow-sdk-go/internal/authentication"
	kauthentication "github.com/microsoft/kiota-abstractions-go/authentication"
	"github.com/pkg/browser"
)

func NewInteractiveBrowserCredential(clientID, clientSecret string, port int, scopes []string) *kauthentication.BaseBearerTokenAuthenticationProvider {
	return kauthentication.NewBaseBearerTokenAuthenticationProvider(authentication.NewAuthorizationCodeTokenProvider(clientID, clientSecret, port, scopes, browser.OpenURL))
}
