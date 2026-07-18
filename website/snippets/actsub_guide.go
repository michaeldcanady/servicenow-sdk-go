//go:build snippets

package snippets

// [START asg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END asg_imports]

func _() {
	// [START asg_main]
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

	actSub := client.Now().ActSub()

	// Step 2: List the activities available to the current user
	activities, err := actSub.Activities().Get(ctx, nil)
	if err != nil {
		log.Fatalf("unable to list activities: %v", err)
	}

	list, err := activities.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d activities\n", len(list))

	// Step 3: Check whether the current user is subscribed to an object
	subscribed, err := actSub.Subscriptions().
		ByObjectId("xSDK_SN_TABLE_SYS_IDx").
		IsSubscribed().
		Get(ctx, nil)
	if err != nil {
		log.Fatalf("subscription check failed: %v", err)
	}
	fmt.Printf("subscription state: %+v\n", subscribed)
	// [END asg_main]
}
