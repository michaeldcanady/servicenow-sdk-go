package credential

import (
	"fmt"
	"net/http"

	"github.com/michaeldcanady/servicenow-sdk-go/internal/oauth2"
)

const (
	snOauthRedirectURITemplate = "/oauth_redirect.do"
	snOauthTokenURITemplate    = "/oauth_token.do"
)

func serviceNowOauthHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func newServiceNowOauthServer(address string, opts ...oauth2.ServerOption) (oauth2.Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serviceNowOauthHandlerFunc)

	return oauth2.NewOauthServer(address, mux, opts...)
}
