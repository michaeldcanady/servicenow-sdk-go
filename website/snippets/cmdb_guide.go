//go:build snippets

package snippets

// [START cig_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END cig_imports]

func _() {
	// [START cig_main]
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

	// Step 2: Query the CIs of a class
	response, err := client.Now().Cmdb().Instance().ByClass("cmdb_ci_linux_server").Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to query CIs: %v", err)
	}

	// Step 3: Read fields from each CI
	cis, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	for _, ci := range cis {
		name, err := ci.GetName()
		if err != nil {
			log.Fatal(err)
		}
		if name != nil {
			fmt.Printf("CI: %s\n", *name)
		}
	}
	// [END cig_main]
}
