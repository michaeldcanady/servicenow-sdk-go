//go:build snippets

package snippets

// [START pgg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/core"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	tableapi "github.com/michaeldcanady/servicenow-sdk-go/tableapi"
)

// [END pgg_imports]

func _() {
	// [START pgg_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	ctx := context.Background()

	// Step 2: Execute the first list request
	response, err := client.Now().Table("xSDK_SN_TABLEx").Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to list records: %v", err)
	}

	// Step 3: Create an iterator from the response
	iterator, err := core.NewPageIterator(response, client.GetRequestAdapter(), tableapi.CreateTableRecordFromDiscriminatorValue)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Iterate every record on every page — the iterator follows
	// the response's pagination links for you
	total := 0
	err = iterator.Iterate(ctx, false, func(record *tableapi.TableRecord) bool {
		total++
		return true // false stops the iteration early
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Visited %d records\n", total)
	// [END pgg_main]
}
