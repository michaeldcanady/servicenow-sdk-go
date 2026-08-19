# CDM Editor API

The CDM Editor API provides endpoints for managing CDM editor nodes and validation.

## \[GET\] /sn_cdm/editor/v1/nodes

Retrieves editor nodes with optional filtering.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Editor().Nodes().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /sn_cdm/editor/v1/nodes

Creates a new editor node.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/cdmeditorapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cdmeditorapi.NodeCreateRequest{}
	response, err := client.Now().Cdm().Editor().Nodes().Post(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[PUT\] /sn_cdm/editor/v1/nodes/{node_sys_id}

Updates the specified editor node.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
	"github.com/michaeldcanady/servicenow-sdk-go/cdmeditorapi"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	body := &cdmeditorapi.NodeUpdateRequest{}
	response, err := client.Now().Cdm().Editor().Nodes().
		ByID("node_sys_id").Put(context.Background(), body, nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[DELETE\] /sn_cdm/editor/v1/nodes/{node_sys_id}

Deletes the specified editor node.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	err := client.Now().Cdm().Editor().Nodes().
		ByID("node_sys_id").Delete(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
```

## \[GET\] /sn_cdm/editor/v1/validation

Validates a CDM model.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Editor().Validation().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
