//go:build snippets

package snippets

// [START ag_download_imports]
import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachmentapi"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/michaeldcanady/servicenow-sdk-go/query"
)

// [END ag_download_imports]

func _() {
	// [START ag_download]
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

	// Step 3: List attachments for a specific record
	encodedQuery := query.And(
		query.String("table_sys_id").Is("xSDK_SN_TABLE_SYS_IDx"),
		query.String("table_name").Is("xSDK_SN_TABLEx"),
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

	// Step 4: Retrieve the attachment results
	attachments, err := resp.GetResult()
	if err != nil {
		log.Fatalf("unable to retrieve result(s): %v", err)
	}

	if len(attachments) == 0 {
		log.Println("No attachments found.")
		return
	}

	// Step 5: Download each attachment
	for _, attachment := range attachments {
		sysID, err := attachment.GetSysID()
		if err != nil {
			log.Printf("unable to retrieve attachment sys_id: %v", err)
			continue
		}

		fileName, err := attachment.GetFileName()
		if err != nil {
			log.Printf("unable to retrieve attachment name: %v", err)
			continue
		}

		// Download the file content
		fileResp, err := client.
			Now().
			Attachment().
			ByID(*sysID).
			File().
			Get(context.Background(), nil)

		if err != nil {
			log.Printf("unable to download attachment %s: %v", *sysID, err)
			continue
		}

		content, err := fileResp.GetContent()
		if err != nil {
			log.Printf("unable to read attachment content: %v", err)
			continue
		}

		// Write the file to disk
		outputPath := filepath.Join("./downloads", *fileName)
		if err := os.WriteFile(outputPath, content, 0644); err != nil {
			log.Printf("failed to write file %s: %v", outputPath, err)
			continue
		}

		fmt.Printf("Downloaded attachment: %s → %s\n", *fileName, outputPath)
	}
	// [END ag_download]
}
