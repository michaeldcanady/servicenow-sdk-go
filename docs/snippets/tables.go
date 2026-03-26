package snippets

// [START table_imports]
import (
	"context"
	"fmt"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

// [END table_imports]

func _() {
	tableBasicSetup()
	tableGetSnippets()
	tableListSnippets()
	tableCreateSnippets()
	tableUpdateSnippets()
	tableDeleteSnippets()
	tableGuideSnippets()
}

func tableBasicSetup() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var requestBuilder *tableapi.TableItemRequestBuilder2[*tableapi.TableRecord]
	var collectionRequestBuilder *tableapi.TableRequestBuilder2[*tableapi.TableRecord]

	// [START table_standard_setup]
	// Step 3: Define raw URL
	rawURL := "https://xSDK_SN_URLx/api/now/v1/table/xSDK_SN_TABLEx"

	// Step 4: Build request
	requestBuilder = tableapi.NewDefaultTableItemRequestBuilder2(rawURL, client.GetRequestAdapter())
	// [END table_standard_setup]

	// [START table_collection_standard_setup]
	// Step 3: Define raw URL
	rawURL = "https://xSDK_SN_URLx/api/now/v1/table/xSDK_SN_TABLEx"

	// Step 4: Build request
	collectionRequestBuilder = tableapi.NewDefaultTableRequestBuilder2(rawURL, client.GetRequestAdapter())
	// [END table_collection_standard_setup]
	_ = requestBuilder
	_ = collectionRequestBuilder
}

func tableGetSnippets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var requestBuilder *tableapi.TableItemRequestBuilder2[*tableapi.TableRecord]

	// [START table_get_fluent]
	// Step 3: Configure request
	getConfig := &tableapi.TableItemRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Execute request
	getResponse, err := client.Now().Table("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Get(context.Background(), getConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_get_fluent]

	// [START table_get_standard]
	// Step 5: Configure request
	getStdConfig := &tableapi.TableItemRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Execute request
	getStdRecord, err := requestBuilder.Get(context.Background(), getStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_get_standard]
	_ = getResponse
	_ = getStdRecord
}

func tableListSnippets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var collectionRequestBuilder *tableapi.TableRequestBuilder2[*tableapi.TableRecord]

	// [START table_list_fluent]
	// Step 3: Configure request
	listConfig := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Execute request
	list_response, err := client.Now().Table("xSDK_SN_TABLEx").Get(context.Background(), listConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_list_fluent]

	// [START table_list_standard]
	// Step 5: Configure request
	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Execute request
	list_std_response, err := collectionRequestBuilder.Get(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_list_standard]
	_ = list_response
	_ = list_std_response
}

func tableCreateSnippets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var collectionRequestBuilder *tableapi.TableRequestBuilder2[*tableapi.TableRecord]

	// [START table_create_fluent]
	// Step 3: Configure request
	config := &tableapi.TableRequestBuilder2PostRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2PostQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Build request body
	createData := tableapi.NewTableRecord()
	if err := createData.SetValue("short_description", "example incident"); err != nil {
		log.Fatal(err)
	}
	if err := createData.SetValue("description", "incident created by servicenow-sdk-go"); err != nil {
		log.Fatal(err)
	}

	// Step 5: Execute request
	create_response, err := client.Now().Table("xSDK_SN_TABLEx").Post(context.Background(), createData, config)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_create_fluent]

	// [START table_create_standard]
	// Step 5: Configure request
	createStdConfig := &tableapi.TableRequestBuilder2PostRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2PostQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Build request body
	createStdData := tableapi.NewTableRecord()
	if err := createStdData.SetValue("short_description", "example incident"); err != nil {
		log.Fatal(err)
	}
	if err := createStdData.SetValue("description", "incident created by servicenow-sdk-go"); err != nil {
		log.Fatal(err)
	}

	// Step 7: Execute request
	create_std_response, err := collectionRequestBuilder.Post(context.Background(), createStdData, createStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_create_standard]
	_ = create_response
	_ = create_std_response
}

func tableUpdateSnippets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var requestBuilder *tableapi.TableItemRequestBuilder2[*tableapi.TableRecord]

	// [START table_update_fluent]
	// Step 3: Configure request
	update_config := &tableapi.TableItemRequestBuilder2PutRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2PutQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Build request body
	updateData := tableapi.NewTableRecord()
	if err := updateData.SetValue("short_description", "updated incident"); err != nil {
		log.Fatal(err)
	}

	// Step 5: Execute request
	updateResponse, err := client.Now().Table("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Put(context.Background(), updateData, update_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_update_fluent]

	// [START table_update_standard]
	// Step 5: Configure request
	updateStdConfig := &tableapi.TableItemRequestBuilder2PutRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2PutQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Build request body
	updateStdData := tableapi.NewTableRecord()
	if err := updateStdData.SetValue("short_description", "updated incident"); err != nil {
		log.Fatal(err)
	}

	// Step 7: Execute request
	updateStdResponse, err := requestBuilder.Put(context.Background(), updateStdData, updateStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_update_standard]
	_ = updateResponse
	_ = updateStdResponse
}

func tableDeleteSnippets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var requestBuilder *tableapi.TableItemRequestBuilder2[*tableapi.TableRecord]

	// [START table_delete_fluent]
	// Step 3: Configure request
	delete_config := &tableapi.TableItemRequestBuilder2DeleteRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2DeleteQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Execute request
	err := client.Now().Table("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Delete(context.Background(), delete_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_delete_fluent]

	// [START table_delete_standard]
	// Step 5: Configure request
	delete_std_config := &tableapi.TableItemRequestBuilder2DeleteRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2DeleteQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Execute request
	err = requestBuilder.Delete(context.Background(), delete_std_config)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_delete_standard]
}

func tableGuideSnippets() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()

	// [START table_list_guide]
	// Get records from the 'incident' table
	listGuideResponse, err := client.Now().Table("xSDK_SN_TABLEx").Get(ctx, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	results, err := listGuideResponse.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range results {
		num, err := record.Get("number")
		if err != nil {
			log.Fatal(err)
		}

		val, err := num.GetValue()
		if err != nil {
			log.Fatal(err)
		}
		strVal, err := val.GetStringValue()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Incident: %s\n", *strVal)
	}
	// [END table_list_guide]

	// [START table_create_guide]
	newIncident := tableapi.NewTableRecord()
	if err := newIncident.SetValue("short_description", "System is down"); err != nil {
		log.Fatal(err)
	}
	if err := newIncident.SetValue("priority", "1"); err != nil {
		log.Fatal(err)
	}

	create_guide_response, err := client.Now().Table("xSDK_SN_TABLEx").Post(ctx, newIncident, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result, _ := create_guide_response.GetResult()

	sysId, _ := result.GetSysID()
	fmt.Printf("Created incident with sys_id: %s\n", *sysId)
	// [END table_create_guide]

	// [START table_update_guide]
	sysIdStr := "xSDK_SN_TABLE_SYS_IDx" // your record sys_id

	updateGuideData := tableapi.NewTableRecord()
	if err := updateGuideData.SetValue("short_description", "Updated description"); err != nil {
		log.Fatal(err)
	}

	update_guide_response, err := client.Now().Table("xSDK_SN_TABLEx").ById(sysIdStr).Put(ctx, updateGuideData, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_update_guide]

	// [START table_delete_guide]
	sysIdToDelete := "xSDK_SN_TABLE_SYS_IDx"
	err = client.Now().Table("xSDK_SN_TABLEx").ById(sysIdToDelete).Delete(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_delete_guide]

	// [START table_query_guide]
	params := &tableapi.TableRequestBuilder2GetQueryParameters{
		Query: "priority=1^active=true",
		Limit: 10,
	}

	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: params,
	}

	query_guide_response, err := client.Now().Table("xSDK_SN_TABLEx").Get(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_query_guide]
	_ = update_guide_response
	_ = query_guide_response
}
