//go:build snippets

package snippets

// [START ag_upload_imports]
import (
	"context"
	"fmt"
	"log"
	"os"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachmentapi"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END ag_upload_imports]

func _() {
	// [START ag_upload]
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

	// Step 3: Read the file from disk
	filePath := "./example.txt"
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// Step 4: Wrap the file content in a Media object
	mediaType := "text/plain" // adjust based on your file type
	media := attachmentapi.NewMedia(mediaType, fileBytes)

	// Step 5: Provide the required attachment parameters
	tableName := "xSDK_SN_TABLEx"
	tableSysID := "xSDK_SN_TABLE_SYS_IDx"
	fileName := "example.txt"

	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableName:  &tableName,
			TableSysID: &tableSysID,
			FileName:   &fileName,
		},
	}

	// Step 6: Upload the file
	resp, err := client.Now().Attachment().File().Post(
		context.Background(),
		media,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	uploaded, err := resp.GetResult()
	if err != nil {
		log.Fatalf("unable to retrieve uploaded attachment: %v", err)
	}

	sysID, err := uploaded.GetSysID()
	if err != nil {
		log.Fatalf("unable to retrieve attachment sys_id: %v", err)
	}

	fmt.Printf("Uploaded attachment with sys_id: %s\n", *sysID)
	// [END ag_upload]
}
