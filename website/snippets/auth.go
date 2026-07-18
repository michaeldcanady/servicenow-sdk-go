//go:build snippets

package snippets

// [START auth_imports]
import (
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END auth_imports]

import "github.com/microsoft/kiota-abstractions-go/authentication"

// [START auth_interface]
type Credential interface {
	GetAuthentication() (string, error)
}

// [END auth_interface]

func _() {
	// [START auth_basic_admin]
	credAdmin := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")
	// [END auth_basic_admin]

	// [START client_init_panic]
	clientPanic, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(credAdmin),
		servicenowsdkgo.WithURL("xSDK_SN_URLx"),
	)
	if err != nil {
		panic(err)
	}
	// [END client_init_panic]

	// [START auth_basic]
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")
	// [END auth_basic]

	// [START auth_token]
	// You'll need your Client ID.
	credToken, err := credentials.NewPublicAuthorizationCodeProvider(
		"your-client-id",
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END auth_token]

	// [START client_init]
	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithURL("xSDK_SN_URLx"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END client_init]

	_ = clientPanic
	_ = client
	_ = credToken
}

func authBasicFull() {
	// [START auth_basic_full]
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_basic_full]
}

func authROPC() {
	// [START auth_ropc]
	cred, err := credentials.NewROPCProvider(
		"{clientID}",
		"{clientSecret}",
		"xSDK_USERNAMEx",
		"xSDK_PASSWORDx",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_ropc]
}

func authClientCredentials() {
	// [START auth_client_credentials]
	cred, err := credentials.NewClientCredentialsProvider(
		"{clientID}",
		"{clientSecret}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_client_credentials]
}

func authCodePrivate() {
	// [START auth_code_private]
	cred, err := credentials.NewPrivateAuthorizationCodeProvider(
		"{clientID}",
		"{clientSecret}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_code_private]
}

func authCodePublic() {
	// [START auth_code_public]
	cred, err := credentials.NewPublicAuthorizationCodeProvider(
		"{clientID}",
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_code_public]
}

func authJWT() {
	// tokenProvider must generate signed JWT assertions; it is user-provided
	// and must satisfy kiota's authentication.AccessTokenProvider.
	var tokenProvider authentication.AccessTokenProvider

	// [START auth_jwt]
	cred, err := credentials.NewJWTProvider(
		"{clientID}",
		"{clientSecret}",
		tokenProvider,
		credentials.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client, err := servicenowsdkgo.NewServiceNowServiceClient(
		servicenowsdkgo.WithAuthenticationProvider(cred),
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Client is now authenticated and ready to use
	_ = client
	// [END auth_jwt]
}

func _() {
	authBasicFull()
	authROPC()
	authClientCredentials()
	authCodePrivate()
	authCodePublic()
	authJWT()
}
