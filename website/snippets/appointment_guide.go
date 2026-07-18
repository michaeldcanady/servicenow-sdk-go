//go:build snippets

package snippets

// [START abg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END abg_imports]

func _() {
	// [START abg_main]
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

	// Step 2: Read the booking configuration for your service
	response, err := client.AppointmentBooking().Configuration().Get(ctx, nil)
	if err != nil {
		log.Fatalf("configuration request failed: %v", err)
	}

	// Step 3: Read settings out of the result
	result, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	active, err := result.GetActive()
	if err != nil {
		log.Fatal(err)
	}

	if active != nil {
		fmt.Printf("appointment booking active: %t\n", *active)
	}
	// [END abg_main]
}
