//go:build snippets

package snippets

// [START bg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// [END bg_imports]

func _() {
	// [START bg_main]
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

	// Step 2: Describe the requests to combine — ToXRequestInformation
	// builds a request without sending it
	var requests []*abstractions.RequestInformation

	incidents, err := client.Now().Table("incident").ToGetRequestInformation(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	users, err := client.Now().Table("sys_user").ToGetRequestInformation(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	requests = append(requests, incidents, users)

	// Step 3: Combine them into one batch body
	body, err := batchRequests(true, requests...)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Send the batch in a single HTTP call
	response, err := client.Now().Batch().Post(ctx, body, nil)
	if err != nil {
		log.Fatalf("batch request failed: %v", err)
	}

	// Step 5: Inspect the per-request results
	serviced, err := response.GetServicedRequests()
	if err != nil {
		log.Fatal(err)
	}
	unserviced, err := response.GetUnservicedRequests()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d requests serviced, %d unserviced\n", len(serviced), len(unserviced))
	// [END bg_main]
}
