package snippets

// [START auth_imports]
import (
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	// [START credentials_import]
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	// [END credentials_import]
)

// [END auth_imports]

// [START auth_interface]
type Credential interface {
	GetAuthentication() (string, error)
}

// [END auth_interface]

func _() {
	// [START auth_basic_admin]
	credAdmin := credentials.NewUsernamePasswordCredential("admin", "password")
	// [END auth_basic_admin]

	// [START client_init_panic]
	clientPanic, err := servicenowsdkgo.NewServiceNowClient2(credAdmin, "your-instance")
	if err != nil {
		panic(err)
	}
	// [END client_init_panic]

	// [START auth_basic]
	cred := credentials.NewUsernamePasswordCredential("xSDK_USERNAMEx", "xSDK_PASSWORDx")
	// [END auth_basic]

	// [START auth_token]
	// You'll need your Client ID, Client Secret, and the Base URL of your instance.
	credToken, err := credentials.NewTokenCredential(
		"your-client-id",
		"your-client-secret",
		"https://your-instance.service-now.com",
		nil, // Use default prompt for username/password or provide your own
	)
	if err != nil {
		log.Fatal(err)
	}
	// [END auth_token]

	// [START client_init]
	client, err := servicenowsdkgo.NewServiceNowClient2(cred, "xSDK_SN_INSTANCEx")
	if err != nil {
		log.Fatal(err)
	}
	// [END client_init]

	_ = clientPanic
	_ = client
	_ = credToken
}
