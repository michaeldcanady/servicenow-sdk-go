//go:build snippets

package snippets

// [START tgc_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/tableapi"
)

// [END tgc_imports]

func _() {
	// [START tgc_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 2: Build the record — only fields you set are sent
	newRecord := tableapi.NewTableRecord()
	if err := newRecord.SetValue("short_description", "System is down"); err != nil {
		log.Fatal(err)
	}
	if err := newRecord.SetValue("priority", "1"); err != nil {
		log.Fatal(err)
	}

	// Step 3: Create it
	response, err := client.Now().Table("xSDK_SN_TABLEx").Post(context.Background(), newRecord, nil)
	if err != nil {
		log.Fatalf("unable to create record: %v", err)
	}

	// Step 4: Read the new record's sys_id from the response
	created, err := response.GetResult()
	if err != nil {
		log.Fatalf("unable to read result: %v", err)
	}

	sysID, err := created.GetSysID()
	if err != nil {
		log.Fatalf("unable to read sys_id: %v", err)
	}

	fmt.Printf("Created record with sys_id: %s\n", *sysID)
	// [END tgc_main]
}
