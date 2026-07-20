//go:build snippets

package snippets

// [START tgq_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/tableapi"
)

// [END tgq_imports]

func _() {
	// [START tgq_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 2: Configure the query — an encoded query filters server-side
	tableQuery := "active=true^priority=1"
	limit := int32(10)
	config := &tableapi.TableRequestBuilderGetRequestConfiguration{
		QueryParameters: &tableapi.TableRequestBuilderGetQueryParameters{
			Query: &tableQuery,
			Limit: &limit,
		},
	}

	// Step 3: Fetch the matching records
	response, err := client.Now().Table("xSDK_SN_TABLEx").Get(context.Background(), config)
	if err != nil {
		log.Fatalf("unable to list records: %v", err)
	}

	records, err := response.GetResult()
	if err != nil {
		log.Fatalf("unable to read results: %v", err)
	}

	// Step 4: Read a field from each record
	for _, record := range records {
		element, err := record.Get("number")
		if err != nil {
			log.Fatal(err)
		}
		value, err := element.GetValue()
		if err != nil {
			log.Fatal(err)
		}
		number, err := value.GetStringValue()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Record: %s\n", *number)
	}
	// [END tgq_main]
}
