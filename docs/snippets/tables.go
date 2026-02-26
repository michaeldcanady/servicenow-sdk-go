package snippets

// [START table_imports]
import (
	"context"
	"fmt"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

// [END table_imports]

func _() {
	var _ credentials.AccessToken

	var client *servicenowsdkgo.ServiceNowClient
	var requestBuilder *tableapi.TableItemRequestBuilder2[*tableapi.TableRecord]
	var collectionRequestBuilder *tableapi.TableRequestBuilder2[*tableapi.TableRecord]
	ctx := context.Background()

	// [START table_standard_setup]
	// Step 3: Define raw URL
	rawURL := "https://xSDK_SN_URLx/api/now/v1/table/xSDK_SN_TABLEx"

	// Step 4: Build request
	requestBuilder = tableapi.NewDefaultTableItemRequestBuilder2(rawURL, client.RequestAdapter)
	// [END table_standard_setup]

	// [START table_collection_standard_setup]
	// Step 3: Define raw URL
	rawURL = "https://xSDK_SN_URLx/api/now/v1/table/xSDK_SN_TABLEx"

	// Step 4: Build request
	collectionRequestBuilder = tableapi.NewDefaultTableRequestBuilder2(rawURL, client.RequestAdapter)
	// [END table_collection_standard_setup]

	// [START table_get_fluent]
	// Step 3: Configure request
	getConfig := &tableapi.TableItemRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Execute request
	getResponse, err := client.Now2().TableV2("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Get(context.Background(), getConfig)
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

	// [START table_list_fluent]
	// Step 3: Configure request
	listConfig := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Execute request
	list_response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Get(context.Background(), listConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_list_fluent]

	// [START table_list_standard]
	// Step 5: Configure request
	listStdConfig := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2GetQueryParameters{
			// Optional configurations
		},
	}

	// Step 6: Execute request
	list_std_response, err := collectionRequestBuilder.Get(context.Background(), listStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_list_standard]

	// [START table_create_fluent]
	// Step 3: Configure request
	createConfig := &tableapi.TableRequestBuilder2PostRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilder2PostQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Build request body
	createData := tableapi.NewTableRecord()
	createData.SetValue("short_description", "example incident")
	createData.SetValue("description", "incident created by servicenow-sdk-go")

	// Step 5: Execute request
	create_response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Post(context.Background(), createData, createConfig)
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
	createStdData.SetValue("short_description", "example incident")
	createStdData.SetValue("description", "incident created by servicenow-sdk-go")

	// Step 7: Execute request
	create_std_response, err := collectionRequestBuilder.Post(context.Background(), createStdData, createStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_create_standard]

	// [START table_update_fluent]
	// Step 3: Configure request
	update_config := &tableapi.TableItemRequestBuilder2PutRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2PutQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Build request body
	updateData := tableapi.NewTableRecord()
	updateData.SetValue("short_description", "updated incident")

	// Step 5: Execute request
	updateResponse, err := client.Now2().TableV2("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Put(context.Background(), updateData, update_config)
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
	updateStdData.SetValue("short_description", "updated incident")

	// Step 7: Execute request
	updateStdResponse, err := requestBuilder.Put(context.Background(), updateStdData, updateStdConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END table_update_standard]

	// [START table_delete_fluent]
	// Step 3: Configure request
	delete_config := &tableapi.TableItemRequestBuilder2DeleteRequestConfiguration{
		QueryParameters: &tableapi.TableItemRequestBuilder2DeleteQueryParameters{
			// Optional configurations
		},
	}

	// Step 4: Execute request
	err = client.Now2().TableV2("xSDK_SN_TABLEx").ById("xSDK_SN_TABLE_SYS_IDx").Delete(context.Background(), delete_config)
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

	// [START table_list_guide]
	// Get records from the 'incident' table
	listGuideResponse, err := client.Now2().TableV2("xSDK_SN_TABLEx").Get(ctx, nil)
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
	newIncident.SetValue("short_description", "System is down")
	newIncident.SetValue("priority", "1")

	create_guide_response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Post(ctx, newIncident, nil)
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
	updateGuideData.SetValue("short_description", "Updated description")

	update_guide_response, err := client.Now2().TableV2("xSDK_SN_TABLEx").ById(sysIdStr).Put(ctx, updateGuideData, nil)
	// [END table_update_guide]

	// [START table_delete_guide]
	sysIdToDelete := "xSDK_SN_TABLE_SYS_IDx"
	err = client.Now2().TableV2("xSDK_SN_TABLEx").ById(sysIdToDelete).Delete(ctx, nil)
	// [END table_delete_guide]

	// [START table_query_guide]
	params := &tableapi.TableRequestBuilder2GetQueryParameters{
		Query: "priority=1^active=true",
		Limit: 10,
	}

	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: params,
	}

	query_guide_response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Get(ctx, config)
	// [END table_query_guide]

	_ = getResponse
	_ = getStdRecord
	_ = collectionRequestBuilder
	_ = list_response
	_ = list_std_response
	_ = create_response
	_ = create_std_response
	_ = updateResponse
	_ = updateStdResponse
	_ = update_guide_response
	_ = query_guide_response
}
