//go:build snippets

package snippets

// [START tgd_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END tgd_imports]

func _() {
	// [START tgd_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 2: Delete the record by sys_id — this is permanent
	sysID := "xSDK_SN_TABLE_SYS_IDx"
	if err := client.Now().Table("xSDK_SN_TABLEx").ByID(sysID).Delete(context.Background(), nil); err != nil {
		log.Fatalf("unable to delete record: %v", err)
	}

	fmt.Printf("Deleted record %s\n", sysID)
	// [END tgd_main]
}
