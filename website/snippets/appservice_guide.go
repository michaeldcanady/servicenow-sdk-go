// Copyright (c) 2021 Michael Canady
// SPDX-License-Identifier: MIT

//go:build snippets

package snippets

// [START apg_imports]
import (
	"context"
	"fmt"
	"log"

	servicenow "github.com/michaeldcanady/servicenow-sdk-go/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/credentials"
)

// [END apg_imports]

func _() {
	// [START apg_main]
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

	// Step 2: Look up an application service in the CSDM model
	found, err := client.Now().Cmdb().AppService().Csdm().FindService().Get(ctx, nil)
	if err != nil {
		log.Fatalf("find service failed: %v", err)
	}

	fmt.Printf("service: %+v\n", found)
	// [END apg_main]
}
