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

	// [START attachment_std_get_item]
	// Step 3: Define raw URL
	get_item_rawURL := "https://xSDK_SN_URLx/api/now/v1/attachment/xSDK_SN_TABLE_SYS_IDx"

	// Step 4: Configure request
	get_item_std_config := &attachmentapi.AttachmentItemRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	get_item_builder := attachmentapi.NewAttachmentItemRequestBuilder(get_item_rawURL, client.RequestAdapter)

	get_item_std_response, err := get_item_builder.Get(context.Background(), get_item_std_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_get_item]

	// [START attachment_list]
	list_config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
		// Optional configurations
	}

	list_response, err := client.Now2().Attachment2().Get(context.Background(), list_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_list]

	// [START attachment_std_list]
	// Step 3: Define raw URL
	list_rawURL := "https://xSDK_SN_URLx/api/now/v1/attachment"

	// Step 4: Configure request
	list_std_config := &attachmentapi.AttachmentRequestBuilder2GetRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	list_builder := attachmentapi.NewAttachmentRequestBuilder2(list_rawURL, client.RequestAdapter)

	list_std_response, err := list_builder.Get(context.Background(), list_std_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_list]

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

	create_file_table_sys_id := "xSDK_SN_TABLE_SYS_IDx"
	create_file_table_name := "xSDK_SN_TABLEx"
	create_file_file_name := "example.txt"

	create_file_config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableSysID: &create_file_table_sys_id, // required
			TableName:  &create_file_table_name,   // required
			FileName:   &create_file_file_name,    // required
		},
		// Optional configurations
	}

	create_file_response, err := client.Now2().Attachment2().File().Post(context.Background(), media, create_file_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_file_create]

	// [START attachment_std_file_create]
	// Step 3: Define raw URL
	file_create_rawURL := "https://xSDK_SN_URLx/api/now/v1/attachment/file"

	// Step 4: Build media type
	file_create_dataContentType := "text/plain"
	file_create_data := []byte("this is example data")
	file_create_media := attachmentapi.NewMedia(file_create_dataContentType, file_create_data)

	file_create_table_sys_id := "xSDK_SN_TABLE_SYS_IDx"
	file_create_table_name := "xSDK_SN_TABLEx"
	file_create_file_name := "example.txt"

	// Step 5: Configure request
	file_create_std_config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableSysID: &file_create_table_sys_id, // required
			TableName:  &file_create_table_name,   // required
			FileName:   &file_create_file_name,    // required
		},
		// Optional configurations
	}

	// Step 6: Build request
	file_create_builder := attachmentapi.NewAttachmentFileRequestBuilder(file_create_rawURL, client.RequestAdapter)

	file_create_std_response, err := file_create_builder.Post(context.Background(), file_create_media, file_create_std_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_file_create]

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

	// [START attachment_std_upload_create]
	// Step 3: Define raw URL
	upload_rawURL := "https://xSDK_SN_URLx/api/now/v1/attachment/upload"

	// Step 4: Build multipart body
	var upload_std_body abstractions.MultipartBody

	// Step 5: Configure request
	upload_std_config := &attachmentapi.AttachmentUploadRequestBuilderPostRequestConfiguration{
		// Optional configurations
	}

	// Step 6: Build request
	upload_builder := attachmentapi.NewAttachmentUploadRequestBuilder(upload_rawURL, client.RequestAdapter)

	upload_std_response, err := upload_builder.Post(context.Background(), upload_std_body, upload_std_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_upload_create]

	// [START attachment_file_get]
	file_get_config := &attachmentapi.AttachmentItemFileRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	file_get_response, err := client.Now2().Attachment2().ByID("xSDK_SN_TABLE_SYS_IDx").File().Get(context.Background(), file_get_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_file_get]

	// [START attachment_std_file_get]
	// Step 3: Define raw URL
	file_get_rawURL := "https://xSDK_SN_URLx/api/now/v1/attachment/xSDK_SN_TABLE_SYS_IDx/file"

	// Step 4: Configure request
	file_get_std_config := &attachmentapi.AttachmentItemFileRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	file_get_builder := attachmentapi.NewAttachmentItemFileRequestBuilder(file_get_rawURL, client.RequestAdapter)

	file_get_std_response, err := file_get_builder.Get(context.Background(), file_get_std_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_file_get]

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
	guide_table_name := "xSDK_SN_TABLEx"
	guide_table_sys_id := "xSDK_SN_TABLE_SYS_IDx"
	guide_file_name := "file.txt"

	params := &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
		TableName:  &guide_table_name,
		TableSysID: &guide_table_sys_id,
		FileName:   &guide_file_name,
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
	_ = get_item_std_response
	_ = list_response
	_ = list_std_response
	_ = create_file_response
	_ = file_create_std_response
	_ = upload_response
	_ = upload_std_response
	_ = file_get_response
	_ = file_get_std_response
	_ = list_guide_results
	_ = create_guide_response
	_ = id
	_ = content
}
