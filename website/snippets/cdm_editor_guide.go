//go:build snippets

package snippets

// [START ceg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END ceg_imports]

func _() {
	// [START ceg_main]
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

	editor := client.Cdm().Editor()

	// Step 2: List configuration data nodes
	nodes, err := editor.Nodes().Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to list nodes: %v", err)
	}
	fmt.Printf("nodes: %+v\n", nodes)

	// Step 3: Validate the configuration data
	validation, err := editor.Validation().Get(ctx, nil)
	if err != nil {
		log.Fatalf("validation failed: %v", err)
	}
	fmt.Printf("validation: %+v\n", validation)
	// [END ceg_main]
}
