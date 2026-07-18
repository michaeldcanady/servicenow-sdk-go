//go:build snippets

package snippets

// [START ag_delete_aging_imports]
import (
	"context"
	"fmt"
	"log"
	"time"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachmentapi"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/michaeldcanady/servicenow-sdk-go/query"
)

// [END ag_delete_aging_imports]

func _() {
	// [START ag_delete_aging]
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

	// Step 3: Determine the cutoff date
	retentionDays := 30
	cutoff := time.Now().AddDate(0, 0, -retentionDays)

	// Step 4: List attachments for a specific record
	encodedQuery := query.And(
		query.String("table_sys_id").Is("xSDK_SN_TABLE_SYS_IDx"),
		query.String("table_name").Is("xSDK_SN_TABLEx"),
		query.Date("sys_created_on").Before(query.NewDateTimeValue(cutoff)),
	).String()

	config := &attachmentapi.AttachmentRequestBuilderGetRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentRequestBuilderGetQueryParameters{
			SysparmQuery: &encodedQuery,
		},
	}

	resp, err := client.Now().Attachment().Get(context.Background(), config)
	if err != nil {
		log.Fatalf("unable to send/receive request: %v", err)
	}

	attachments, err := resp.GetResult()
	if err != nil {
		log.Fatalf("unable to retrieve result(s): %v", err)
	}

	if len(attachments) == 0 {
		log.Println("No attachments found.")
		return
	}

	// Step 5: Delete attachments older than the cutoff
	for _, attachment := range attachments {
		createdOn, err := attachment.GetSysCreatedOn()
		if err != nil {
			log.Printf("unable to retrieve creation date: %v", err)
			continue
		}

		sysID, err := attachment.GetSysID()
		if err != nil {
			log.Printf("unable to retrieve sys_id: %v", err)
			continue
		}

		deleteConfig := &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{}

		if err := client.
			Now().
			Attachment().
			ByID(*sysID).
			Delete(context.Background(), deleteConfig); err != nil {

			log.Printf("failed to delete attachment %s: %v", *sysID, err)
			continue
		}

		fmt.Printf("Deleted attachment %s (created %s)\n", *sysID, createdOn.String())
	}
	// [END ag_delete_aging]
}
