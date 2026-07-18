//go:build snippets

package snippets

// [START pmg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END pmg_imports]

func _() {
	// [START pmg_main]
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

	mappings := client.Cdm().Policies().Mappings()

	// Step 2: Create a policy mapping (parameters go in the request configuration)
	created, err := mappings.Post(ctx, nil)
	if err != nil {
		log.Fatalf("unable to create mapping: %v", err)
	}

	mapping, err := created.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created mapping: %+v\n", mapping)

	// Step 3: Delete a mapping when it's no longer needed
	if err := mappings.Delete(ctx, nil); err != nil {
		log.Fatalf("unable to delete mapping: %v", err)
	}
	// [END pmg_main]
}
