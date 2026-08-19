# CDM Applications API

The CDM Applications API provides endpoints for managing Continuous Delivery Model applications, including deployables, exports, uploads, shared components, and shared libraries.

## \[GET\] /sn_cdm/applications/deployables/exports

Retrieves export records for deployables.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Applications().Deployables().Exports().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /sn_cdm/applications/deployables/exports/{export_id}/content

Downloads the content of an export.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Applications().Deployables().Exports().
		ByID("export_id").Content().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /sn_cdm/applications/uploads/components

Uploads components to an application.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
	"github.com/michaeldcanady/servicenow-sdk-go/v2/cdmapplicationsapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cdmapplicationsapi.ComponentUploadRequest{}
	response, err := client.Now().Cdm().Applications().Uploads().Components().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[PUT\] /sn_cdm/applications/deployables

Updates deployable records.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go/v2/v2"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Applications().Deployables().Put(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
