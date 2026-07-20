//go:build snippets

package snippets

// [START sg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go"
	aggregationapi "github.com/michaeldcanady/servicenow-sdk-go/aggregationapi"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

// [END sg_imports]

func _() {
	// [START sg_main]
	// Step 1: Authenticate and initialize the client
	cred := credentials.NewBasicProvider("xSDK_USERNAMEx", "xSDK_PASSWORDx")

	client, err := servicenow.NewServiceNowServiceClient(
		servicenow.WithAuthenticationProvider(cred),
		servicenow.WithInstance("xSDK_SN_INSTANCEx"),
	)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	// Step 2: Choose the aggregates and the records they run over
	wantCount := true
	statsQuery := "active=true"
	config := &aggregationapi.StatsRequestBuilderGetRequestConfiguration{
		QueryParameters: &aggregationapi.StatsRequestBuilderGetQueryParameters{
			Count:     &wantCount,
			SumFields: []string{"reassignment_count"},
			Query:     &statsQuery,
		},
	}

	// Step 3: Request the aggregates — no records are transferred
	response, err := client.Now().Stats("incident").Get(context.Background(), config)
	if err != nil {
		log.Fatalf("stats request failed: %v", err)
	}

	// Step 4: Read the aggregate values
	result, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}
	stats, err := result.GetStats()
	if err != nil {
		log.Fatal(err)
	}
	count, err := stats.GetCount()
	if err != nil {
		log.Fatal(err)
	}

	// Aggregate values come back as strings, as ServiceNow returns them
	fmt.Printf("Active incidents: %s\n", *count)
	// [END sg_main]
}
