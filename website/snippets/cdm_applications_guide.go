//go:build snippets

package snippets

// [START cag_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END cag_imports]

func _() {
	// [START cag_main]
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

	applications := client.Cdm().Applications()

	// Step 2: Check the status of a configuration upload
	status, err := applications.UploadStatus().ByID("xSDK_SN_TABLE_SYS_IDx").Get(ctx, nil)
	if err != nil {
		log.Fatalf("upload status failed: %v", err)
	}
	fmt.Printf("upload status: %+v\n", status)

	// Step 3: List a deployable's exports
	exports, err := applications.Deployables().Exports().Get(ctx, nil)
	if err != nil {
		log.Fatalf("exports failed: %v", err)
	}
	fmt.Printf("exports: %+v\n", exports)
	// [END cag_main]
}
