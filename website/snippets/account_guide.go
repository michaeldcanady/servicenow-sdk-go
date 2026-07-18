//go:build snippets

package snippets

// [START acg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END acg_imports]

func _() {
	// [START acg_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	ctx := context.Background()

	// Step 2: List customer accounts
	response, err := client.Now().Account().Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to list accounts: %v", err)
	}

	// Step 3: Read fields from each account
	accounts, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	for _, account := range accounts {
		name, err := account.GetName()
		if err != nil {
			log.Fatal(err)
		}
		if name != nil {
			fmt.Printf("Account: %s\n", *name)
		}
	}
	// [END acg_main]
}
