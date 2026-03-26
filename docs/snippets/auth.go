package snippets

// [START auth_imports]
import (
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END auth_imports]

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
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
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
		servicenowsdkgo.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END client_init]

	// Need to be "used" to be valid
	_ = clientPanic
	_ = client
	_ = credToken
}
