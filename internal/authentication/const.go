package authentication

import "github.com/yosida95/uritemplate/v3"

var (
	authURLTemplate  = uritemplate.MustNew("{+baseurl}/oauth_auth.do")
	tokenURLTemplate = uritemplate.MustNew("{+baseurl}/oauth_token.do") //nolint:gosec
)
