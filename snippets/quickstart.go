package snippets

import (
	"context"
	"fmt"
	"log"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/credentials"
)

func _() {
	// [START quickstart]
	cred := credentials.NewUsernamePasswordCredential("xSDK_USERNAMEx", "xSDK_PASSWORDx")
	client, err := servicenowsdkgo.NewServiceNowClient2(cred, "xSDK_SN_INSTANCEx")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	response, err := client.Now2().TableV2("xSDK_SN_TABLEx").Get(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	results, err := response.GetResult()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range results {
		num, _ := record.Get("number")
		val, _ := num.GetValue()
		strVal, _ := val.GetStringValue()
		fmt.Printf("Incident: %s", *strVal)
	}
	// [END quickstart]
}
