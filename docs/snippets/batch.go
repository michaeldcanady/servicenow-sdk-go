package snippets

// [START batch_imports]
import (
	"context"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	batchapi "github.com/michaeldcanady/servicenow-sdk-go/batch-api"
	abstractions "github.com/microsoft/kiota-abstractions-go"
)

// [END batch_imports]

// [START batch_helper]
// batchRequests A helper function to combine provided request information into a single `BatchRequest`.
func batchRequests(excludeResponseHeaders bool, requests ...*abstractions.RequestInformation) (*batchapi.BatchRequestModel, error) {
	body := batchapi.NewBatchRequestModel()
	for _, request := range requests {
		restRequest, err := batchapi.CreateRestRequestFromRequestInformation(request, excludeResponseHeaders)
		if err != nil {
			return nil, err
		}
		if err := body.AddRequest(restRequest); err != nil {
			return nil, err
		}
	}
	return body, nil
}

// [END batch_helper]

func _() {
	var client *servicenowsdkgo.ServiceNowServiceClient
	var requests []*abstractions.RequestInformation

	// [START batch_create]
	// Step 3. Build requests, using ToXXXRequestInformation method
	body, err := batchRequests(true, requests...)
	if err != nil {
		log.Fatal(err)
	}

	// Step 4: Configure request
	postConfig := &batchapi.BatchRequestBuilderPostRequestConfiguration{}

	// Step 5: Execute request
	batch_response, err := client.Now().Batch().Post(context.Background(), body, postConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END batch_create]

	// [START batch_std_create]
	// Step 3: Define raw URL
	rawURL := "https://xSDK_SN_URLx/api/now/v1/batch"

	// Step 4. Build requests, using ToXXXRequestInformation method
	body, err = batchRequests(true, requests...)
	if err != nil {
		log.Fatal(err)
	}

	// Step 5: Configure request
	postConfig = &batchapi.BatchRequestBuilderPostRequestConfiguration{}

	// Step 6: Build request
	requestBuilder := batchapi.NewBatchRequestBuilder(rawURL, client.GetRequestAdapter())

	// Step 7: Execute request
	resp, err := requestBuilder.Post(context.Background(), body, postConfig)
	if err != nil {
		log.Fatal(err)
	}
	// [END batch_std_create]

	// Need to be "used" to be valid
	_ = batch_response
	_ = resp
}
