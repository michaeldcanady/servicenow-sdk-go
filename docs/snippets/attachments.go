package snippets

// [START attachment_imports]
import (
	"context"
	"fmt"
	"log"
	"os"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachment-api"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// [END attachment_imports]

func _() {
	var _ credentials.AccessToken

	var client *servicenowsdkgo.ServiceNowClient
	ctx := context.Background()

	// [START attachment_get_item]
	get_item_config := &attachmentapi.AttachmentItemRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	get_item_response, err := client.Now2().Attachment2().ByID("xSDK_SN_TABLE_SYS_IDx").Get(context.Background(), get_item_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_get_item]

	// [START attachment_list]
	list_config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
		// Optional configurations
	}

	list_response, err := client.Now2().Attachment2().Get(context.Background(), list_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_list]

	// [START attachment_delete]
	// Step 3: Configure request
	deleteConfig := &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{
		// Optional configurations
	}

	err = client.Now2().Attachment2().ByID("xSDK_SN_TABLE_SYS_IDx").Delete(context.Background(), deleteConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_delete]

	// [START attachment_std_delete]
	// Step 3: Define raw URL
	rawURL := "https://xSDK_SN_URLx/api/now/v1/attachment/xSDK_SN_TABLE_SYS_IDx"

	// Step 4: Configure request
	deleteConfig = &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	collectionRequestBuilder := attachmentapi.NewAttachmentItemRequestBuilder2(rawURL, client.RequestAdapter)

	if err := collectionRequestBuilder.Delete(context.Background(), deleteConfig); err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_delete]

	// [START attachment_file_create]
	// Build media type
	dataContentType := "text/plain"
	data := []byte("this is example data")
	media := attachmentapi.NewMedia(dataContentType, data)

	create_file_config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableSysID: "xSDK_SN_TABLE_SYS_IDx", // required
			TableName:  "xSDK_SN_TABLEx",        // required
			FileName:   "example.txt",           // required
		},
		// Optional configurations
	}

	create_file_response, err := client.Now2().Attachment2().File().Post(context.Background(), media, create_file_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_file_create]

	// [START attachment_upload_create]
	// body := // TODO: how to make multipart body?
	var body abstractions.MultipartBody

	upload_config := &attachmentapi.AttachmentUploadRequestBuilderPostRequestConfiguration{
		// Optional configurations
	}

	upload_response, err := client.Now2().Attachment2().Upload().Post(context.Background(), body, upload_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_upload_create]

	// [START attachment_list_guide]
	// List all attachments
	list_guide_response, err := client.Now2().Attachment2().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	list_guide_results, _ := list_guide_response.GetResult()
	for _, attachment := range list_guide_results {
		name, err := attachment.GetFileName()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Attachment: %s\n", *name)
	}
	// [END attachment_list_guide]

	// [START attachment_create_guide]
	file, err := os.Open("path/to/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Upload attachment for an incident
	params := &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
		TableName:  "xSDK_SN_TABLEx",
		TableSysID: "xSDK_SN_TABLE_SYS_IDx",
		FileName:   "file.txt",
	}

	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: params,
	}

	// Assuming 'file' can be used as media content
	create_guide_response, err := client.Now2().Attachment2().File().Post(ctx, nil, config) // Placeholder
	if err != nil {
		log.Fatal(err)
	}

	id, err := create_guide_response.GetSysID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created attachment with sys_id: %s\n", *id)
	// [END attachment_create_guide]

	// [START attachment_download_guide]
	sysIdToDownload := "xSDK_SN_TABLE_SYS_IDx"

	config2 := &attachmentapi.AttachmentItemFileRequestBuilderGetRequestConfiguration{}

	download_file, err := client.Now2().Attachment2().ByID(sysIdToDownload).File().Get(ctx, config2)
	if err != nil {
		log.Fatal(err)
	}

	content, _ := download_file.GetContent()
	// content is a []byte containing the file data
	// [END attachment_download_guide]

	_ = get_item_response
	_ = list_response
	_ = create_file_response
	_ = upload_response
	_ = list_guide_results
	_ = content
}
