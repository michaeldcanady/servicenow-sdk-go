//go:build snippets

package snippets

// [START ag_recent_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachmentapi"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END ag_recent_imports]

func _() {
	// [START ag_recent]
	// Step 1: Authenticate with your ServiceNow instance
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	clientOpts := []servicenow.ServiceNowServiceClientOption{
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	}

	// Step 2: Initialize the ServiceNow client
	client, err := servicenow.NewServiceNowServiceClient(clientOpts...)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 3: List attachments (filter using query params as needed)
	limit := int32(1)
	encodedQuery := "ORDERBYDESCsys_created_on"

	config := &attachmentapi.AttachmentRequestBuilderGetRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentRequestBuilderGetQueryParameters{
			SysparmLimit: &limit,
			SysparmQuery: &encodedQuery,
		},
	}

	resp, err := client.Now().Attachment().Get(
		context.Background(),
		config,
	)
	if err != nil {
		log.Fatalf("unable to send/receive request: %v", err)
	}

	// Step 4: Retrieve the attachment results
	attachments, err := resp.GetResult()
	if err != nil {
		log.Fatalf("unable to retrieve result(s): %v", err)
	}

	if len(attachments) == 0 {
		log.Println("No attachments found.")
		return
	}

	// Step 5: Select the most recent attachment (already sorted by ORDERBYDESC)
	latest := attachments[0]

	fileName, err := latest.GetFileName()
	if err != nil {
		log.Fatalf("unable to retrieve attachment's name: %v", err)
	}

	sysID, err := latest.GetSysID()
	if err != nil {
		log.Fatalf("unable to retrieve attachment's sys_id: %v", err)
	}

	fmt.Printf("Most recent attachment: %s (sys_id: %s)\n", *fileName, *sysID)
	// [END ag_recent]
}
