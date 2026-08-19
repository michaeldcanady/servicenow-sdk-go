# Documents API

The Documents API provides endpoints for managing ServiceNow documents, including exploration, creation, versioning, content retrieval, and synchronization.

## \[GET\] /now/v1/documents/explore

Explores documents with optional filtering by table, folder, or record.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Documents().Explore().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[POST\] /now/v1/documents/create

Creates a new document.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Documents().Create().Post(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/v1/documents/{document_sys_id}/content

Retrieves the content of a document.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Documents().Content("document_sys_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /now/v1/documents/versions/{document_sys_id}

Retrieves versions of a document.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Documents().Versions("document_sys_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[DELETE\] /now/v1/documents/delete

Deletes a document.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	err := client.Now().Documents().Delete().Delete(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
```
