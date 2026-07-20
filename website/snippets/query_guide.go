//go:build snippets

package snippets

// [START qg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	"github.com/michaeldcanady/servicenow-sdk-go/query"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/tableapi"
)

// [END qg_imports]

func _() {
	// [START qg_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 2: Compose the query from typed conditions
	q := query.Boolean("active").Is(true).
		And(query.String("priority").IsOneOf("1", "2")).
		String() // renders the encoded query: active=true^priorityIN1,2

	// Step 3: Use it in the request configuration
	config := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Query: &q,
		},
	}

	response, err := client.Now().Table("xSDK_SN_TABLEx").Get(context.Background(), config)
	if err != nil {
		log.Fatalf("unable to list records: %v", err)
	}

	records, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %d records\n", len(records))
	// [END qg_main]
}
