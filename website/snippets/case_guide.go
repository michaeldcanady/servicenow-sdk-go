//go:build snippets

package snippets

// [START csg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END csg_imports]

func _() {
	// [START csg_main]
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

	cases := client.CustomerService().Case()

	// Step 2: List cases
	response, err := cases.Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to list cases: %v", err)
	}

	// Step 3: Read fields from each case
	list, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range list {
		number, err := c.GetNumber()
		if err != nil {
			log.Fatal(err)
		}
		short, err := c.GetShortDescription()
		if err != nil {
			log.Fatal(err)
		}
		if number != nil && short != nil {
			fmt.Printf("%s — %s\n", *number, *short)
		}
	}
	// [END csg_main]
}
