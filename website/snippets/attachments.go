//go:build snippets

package snippets

// [START attachment_imports]
import (
	"context"
	"fmt"
	"log"
	"os"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	attachmentapi "github.com/michaeldcanady/servicenow-sdk-go/attachmentapi"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// [END attachment_imports]

func _() {
	var _ credentials.AccessToken

	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()

	// [START attachment_get_item]
	getItemConfig := &attachmentapi.AttachmentItemRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	getItemResponse, err := client.Now().Attachment().ByID("xSDK_SN_TABLE_SYS_IDx").Get(context.Background(), getItemConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_get_item]

	// [START attachment_std_get_item]
	// Step 3: Define raw URL
	getItemRawURL := "xSDK_SN_URLx/api/now/attachment/xSDK_SN_TABLE_SYS_IDx"

	// Step 4: Configure request
	getItemStdConfig := &attachmentapi.AttachmentItemRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	getItemBuilder := attachmentapi.NewAttachmentItemRequestBuilder(getItemRawURL, client.GetRequestAdapter())

	getItemStdResponse, err := getItemBuilder.Get(context.Background(), getItemStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_get_item]

	// [START attachment_list]
	listConfig := &attachmentapi.AttachmentRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	listResponse, err := client.Now().Attachment().Get(context.Background(), listConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_list]

	// [START attachment_std_list]
	// Step 3: Define raw URL
	listRawURL := "xSDK_SN_URLx/api/now/attachment"

	// Step 4: Configure request
	listStdConfig := &attachmentapi.AttachmentRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	listBuilder := attachmentapi.NewAttachmentRequestBuilder(listRawURL, client.GetRequestAdapter())

	listStdResponse, err := listBuilder.Get(context.Background(), listStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_list]

	// [START attachment_delete]
	// Step 3: Configure request
	deleteConfig := &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{
		// Optional configurations
	}

	err = client.Now().Attachment().ByID("xSDK_SN_TABLE_SYS_IDx").Delete(context.Background(), deleteConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_delete]

	// [START attachment_std_delete]
	// Step 3: Define raw URL
	rawURL := "xSDK_SN_URLx/api/now/attachment/xSDK_SN_TABLE_SYS_IDx"

	// Step 4: Configure request
	deleteConfig = &attachmentapi.AttachmentItemRequestBuilderDeleteRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	collectionRequestBuilder := attachmentapi.NewAttachmentItemRequestBuilder(rawURL, client.GetRequestAdapter())

	if err := collectionRequestBuilder.Delete(context.Background(), deleteConfig); err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_delete]

	// [START attachment_file_create]
	// Build media type
	dataContentType := "text/plain"
	data := []byte("this is example data")
	media := attachmentapi.NewMedia(dataContentType, data)

	createFileTableSysId := "xSDK_SN_TABLE_SYS_IDx"
	createFileTableName := "xSDK_SN_TABLEx"
	createFileFileName := "example.txt"

	createFileConfig := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableSysID: &createFileTableSysId, // required
			TableName:  &createFileTableName,  // required
			FileName:   &createFileFileName,   // required
		},
		// Optional configurations
	}

	createFileResponse, err := client.Now().Attachment().File().Post(context.Background(), media, createFileConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_file_create]

	// [START attachment_std_file_create]
	// Step 3: Define raw URL
	fileCreateRawURL := "xSDK_SN_URLx/api/now/attachment/file"

	// Step 4: Build media type
	fileCreateDataContentType := "text/plain"
	fileCreateData := []byte("this is example data")
	fileCreateMedia := attachmentapi.NewMedia(fileCreateDataContentType, fileCreateData)

	fileCreateTableSysId := "xSDK_SN_TABLE_SYS_IDx"
	fileCreateTableName := "xSDK_SN_TABLEx"
	fileCreateFileName := "example.txt"

	// Step 5: Configure request
	fileCreateStdConfig := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
			TableSysID: &fileCreateTableSysId, // required
			TableName:  &fileCreateTableName,  // required
			FileName:   &fileCreateFileName,   // required
		},
		// Optional configurations
	}

	// Step 6: Build request
	fileCreateBuilder := attachmentapi.NewAttachmentFileRequestBuilder(fileCreateRawURL, client.GetRequestAdapter())

	fileCreateStdResponse, err := fileCreateBuilder.Post(context.Background(), fileCreateMedia, fileCreateStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_file_create]

	// [START attachment_upload_create]
	var body abstractions.MultipartBody

	uploadConfig := &attachmentapi.AttachmentUploadRequestBuilderPostRequestConfiguration{
		// Optional configurations
	}

	uploadResponse, err := client.Now().Attachment().Upload().Post(context.Background(), body, uploadConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_upload_create]

	// [START attachment_std_upload_create]
	// Step 3: Define raw URL
	uploadRawURL := "xSDK_SN_URLx/api/now/attachment/upload"

	// Step 4: Build multipart body
	var uploadStdBody abstractions.MultipartBody

	// Step 5: Configure request
	uploadStdConfig := &attachmentapi.AttachmentUploadRequestBuilderPostRequestConfiguration{
		// Optional configurations
	}

	// Step 6: Build request
	uploadBuilder := attachmentapi.NewAttachmentUploadRequestBuilder(uploadRawURL, client.GetRequestAdapter())

	uploadStdResponse, err := uploadBuilder.Post(context.Background(), uploadStdBody, uploadStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_upload_create]

	// [START attachment_file_get]
	fileGetConfig := &attachmentapi.AttachmentItemFileRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	fileGetResponse, err := client.Now().Attachment().ByID("xSDK_SN_TABLE_SYS_IDx").File().Get(context.Background(), fileGetConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_file_get]

	// [START attachment_std_file_get]
	// Step 3: Define raw URL
	fileGetRawURL := "xSDK_SN_URLx/api/now/attachment/xSDK_SN_TABLE_SYS_IDx/file"

	// Step 4: Configure request
	fileGetStdConfig := &attachmentapi.AttachmentItemFileRequestBuilderGetRequestConfiguration{
		// Optional configurations
	}

	// Step 5: Build request
	fileGetBuilder := attachmentapi.NewAttachmentItemFileRequestBuilder(fileGetRawURL, client.GetRequestAdapter())

	fileGetStdResponse, err := fileGetBuilder.Get(context.Background(), fileGetStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END attachment_std_file_get]

	// [START attachment_list_guide]
	// List all attachments
	listGuideResponse, err := client.Now().Attachment().Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	listGuideResults, _ := listGuideResponse.GetResult()
	for _, attachment := range listGuideResults {
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
	defer func() { _ = file.Close() }()

	// Upload attachment for an incident
	guideTableName := "xSDK_SN_TABLEx"
	guideTableSysId := "xSDK_SN_TABLE_SYS_IDx"
	guideFileName := "file.txt"

	params := &attachmentapi.AttachmentFileRequestBuilderPostQueryParameters{
		TableName:  &guideTableName,
		TableSysID: &guideTableSysId,
		FileName:   &guideFileName,
	}

	config := &attachmentapi.AttachmentFileRequestBuilderPostRequestConfiguration{
		QueryParameters: params,
	}

	// Assuming 'file' can be used as media content
	createGuideResponse, err := client.Now().Attachment().File().Post(ctx, nil, config) // Placeholder
	if err != nil {
		log.Fatal(err)
	}

	result, err := createGuideResponse.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.GetSysID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created attachment with sys_id: %s\n", *id)
	// [END attachment_create_guide]

	// [START attachment_download_guide]
	sysIdToDownload := "xSDK_SN_TABLE_SYS_IDx"

	config2 := &attachmentapi.AttachmentItemFileRequestBuilderGetRequestConfiguration{}

	downloadFile, err := client.Now().Attachment().ByID(sysIdToDownload).File().Get(ctx, config2)
	if err != nil {
		log.Fatal(err)
	}

	content, _ := downloadFile.GetContent()
	// content is a []byte containing the file data
	// [END attachment_download_guide]

	_ = getItemResponse
	_ = getItemStdResponse
	_ = listResponse
	_ = listStdResponse
	_ = createFileResponse
	_ = fileCreateStdResponse
	_ = uploadResponse
	_ = uploadStdResponse
	_ = fileGetResponse
	_ = fileGetStdResponse
	_ = listGuideResults
	_ = createGuideResponse
	_ = id
	_ = content
}
