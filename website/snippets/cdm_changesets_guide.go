//go:build snippets

package snippets

// [START ccg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END ccg_imports]

func _() {
	// [START ccg_main]
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

	changesets := client.Cdm().Changesets()

	// Step 2: List changesets
	response, err := changesets.Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to list changesets: %v", err)
	}

	list, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d changesets\n", len(list))

	// Step 3: Check the status of a commit
	status, err := changesets.CommitStatus().ByID("xSDK_SN_TABLE_SYS_IDx").Get(ctx, nil)
	if err != nil {
		log.Fatalf("commit status failed: %v", err)
	}
	fmt.Printf("commit status: %+v\n", status)
	// [END ccg_main]
}
