//go:build snippets

package snippets

import (
	"context"
	"fmt"
	"log"
	"time"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/tableapi"
)

func _() {
	conceptsChaining()
	conceptsConfiguration()
	conceptsEnvelope()
	conceptsBackedModels()
	conceptsContext()
}

func conceptsChaining() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START concepts_chaining]
	// Each call narrows the URL. Nothing is sent until a verb method runs.
	//
	//   client.Now()                          → /api/now
	//         .Table("incident")              → /api/now/table/incident
	//         .ByID("xSDK_SN_TABLE_SYS_IDx")  → /api/now/table/incident/{sys_id}
	builder := client.Now().Table("incident").ByID("xSDK_SN_TABLE_SYS_IDx")

	// The verb method (Get/Post/Put/Delete) executes the request.
	response, err := builder.Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END concepts_chaining]
	_ = response

	// [START concepts_standard]
	// Standard modality: construct the request builder yourself from a raw URL.
	rawURL := "xSDK_SN_URLx/api/now/v1/table/xSDK_SN_TABLEx"
	requestBuilder := tableapi.NewDefaultTableItemRequestBuilder(rawURL, client.GetRequestAdapter())

	stdRecord, err := requestBuilder.Get(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END concepts_standard]
	_ = stdRecord
}

func conceptsConfiguration() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START concepts_configuration]
	// Every verb has its own RequestConfiguration and QueryParameters types.
	// Pass nil when you have nothing to configure.
	config := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Query: "active=true^priority=1",
			Limit: 10,
		},
	}

	response, err := client.Now().Table("xSDK_SN_TABLEx").Get(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	// [END concepts_configuration]
	_ = response
}

func conceptsEnvelope() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	ctx := context.Background()

	// [START concepts_envelope]
	// Collection endpoints return a collection envelope…
	listResponse, err := client.Now().Table("xSDK_SN_TABLEx").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// …and the records live behind GetResult().
	records, err := listResponse.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("got %d records\n", len(records))

	// Item endpoints work the same way, with a single record inside.
	itemResponse, err := client.Now().Table("xSDK_SN_TABLEx").ByID("xSDK_SN_TABLE_SYS_IDx").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	record, err := itemResponse.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	// [END concepts_envelope]
	_ = record
}

func conceptsBackedModels() {
	var record *tableapi.TableRecord

	// [START concepts_backed_read]
	// 1. Get the field. A missing field is an error, not a zero value.
	element, err := record.Get("number")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Unwrap the element. Fields carry a value, a display value, and a link.
	value, err := element.GetValue()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Convert to the type you expect. The result is a pointer: nil means
	//    the instance sent no value, distinct from "" or 0.
	number, err := value.GetStringValue()
	if err != nil {
		log.Fatal(err)
	}

	if number != nil {
		fmt.Printf("number: %s\n", *number)
	}
	// [END concepts_backed_read]

	// [START concepts_backed_write]
	newRecord := tableapi.NewTableRecord()
	if err := newRecord.SetValue("short_description", "created by servicenow-sdk-go"); err != nil {
		log.Fatal(err)
	}
	if err := newRecord.SetValue("priority", "1"); err != nil {
		log.Fatal(err)
	}
	// [END concepts_backed_write]
}

func conceptsContext() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	// [START concepts_context]
	// Every verb method takes a context.Context; deadlines and cancellation
	// propagate to the underlying HTTP call.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := client.Now().Table("xSDK_SN_TABLEx").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [END concepts_context]
	_ = response
}
