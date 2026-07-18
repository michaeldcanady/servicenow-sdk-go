//go:build snippets

package snippets

// [START tgu_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/tableapi"
)

// [END tgu_imports]

func _() {
	// [START tgu_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 2: Build the update — set only the fields to change;
	// everything else on the record is left untouched
	update := tableapi.NewTableRecord()
	if err := update.SetValue("short_description", "Updated description"); err != nil {
		log.Fatal(err)
	}

	// Step 3: Apply it to the record by sys_id
	sysID := "xSDK_SN_TABLE_SYS_IDx"
	response, err := client.Now().Table("xSDK_SN_TABLEx").ByID(sysID).Put(context.Background(), update, nil)
	if err != nil {
		log.Fatalf("unable to update record: %v", err)
	}

	if _, err := response.GetResult(); err != nil {
		log.Fatalf("unable to read result: %v", err)
	}

	fmt.Printf("Updated record %s\n", sysID)
	// [END tgu_main]
}
