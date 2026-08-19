# CDM Changeset API

The CDM Changeset API provides endpoints for managing Continuous Delivery Model changesets, including listing, deleting, activity tracking, commit status, and impacted deployables/shared components.

## \[GET\] /sn_cdm/changesets

Retrieves multiple changesets.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Changesets().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[DELETE\] /sn_cdm/changesets

Deletes changesets matching the specified query.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	err := client.Now().Cdm().Changesets().Delete(context.Background(), nil)
	if err != nil {
		panic(err)
	}
}
```

## \[GET\] /sn_cdm/changesets/activity

Retrieves activity for changesets.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Changesets().Activity().Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```

## \[GET\] /sn_cdm/changesets/commit-status/{commit_id}

Retrieves the commit status for a specific changeset.

```golang
package main

import (
	"context"

	servicenowsdkgo "github.com/michaeldcanady/servicenow-sdk-go"
)

func main() {
	var client *servicenowsdkgo.ServiceNowServiceClient

	response, err := client.Now().Cdm().Changesets().CommitStatus().
		ByID("commit_id").Get(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	_ = response
}
```
