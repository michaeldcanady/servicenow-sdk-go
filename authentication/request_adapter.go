package authentication

import (
	"github.com/microsoft/kiota-abstractions-go/authentication"
	nethttplibrary "github.com/microsoft/kiota-http-go"
)

func newRequestAdapter() *nethttplibrary.NetHttpRequestAdapter {

	requestAdapter, _ := nethttplibrary.NewNetHttpRequestAdapter(&authentication.AnonymousAuthenticationProvider{})

	return requestAdapter
}
