//go:build preview.query

package snippets

// [START query_basic_imports]
import (
	"context"
	"fmt"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/query2"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

// [END query_basic_imports]

func _() {
	var client *servicenowsdkgo.ServiceNowClient
	ctx := context.Background()

	// [START query_basic]
	q := query2.String("short_description").Contains("System").
		And(query2.Number("priority").Is(1)).String()

	fmt.Println(q) // Output: short_descriptionLIKESystem^priority=1
	// [END query_basic]

	// [START query_table_api]
	// Build the query
	q2 := query2.Boolean("active").Is(true).
		And(query2.String("priority").
			IsOneOf("1", "2")).String()

	params := &tableapi.TableRequestBuilder2GetQueryParameters{
		Query: q2,
	}

	config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
		QueryParameters: params,
	}

	response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Get(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	// [END query_table_api]

	_ = response
}
