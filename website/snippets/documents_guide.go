//go:build snippets

package snippets

// [START dg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END dg_imports]

func _() {
	// [START dg_main]
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

	documents := client.Now().Documents()

	// Step 2: Explore documents and folders
	results, err := documents.Explore().Get(ctx, nil)
	if err != nil {
		log.Fatalf("explore failed: %v", err)
	}
	fmt.Printf("explore returned: %+v\n", results)

	// Step 3: List a document's versions
	versions, err := documents.Versions("xSDK_SN_TABLE_SYS_IDx").Get(ctx, nil)
	if err != nil {
		log.Fatalf("versions failed: %v", err)
	}
	fmt.Printf("versions: %+v\n", versions)
	// [END dg_main]
}
