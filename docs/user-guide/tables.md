# Table Operations

The Table API allows you to perform CRUD (Create, Read, Update, Delete) operations on ServiceNow tables.

## Listing Records

To retrieve a collection of records from a table:

```go
import (
    "context"
    "fmt"
    "log"
    "github.com/michaeldcanady/servicenow-sdk-go"
    "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    // ... client initialization ...

    ctx := context.Background()
    
    // Get records from the 'incident' table
    response, err := client.Now2().TableV2("incident").Get(ctx, nil)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    for _, record := range response.GetValue() {
        fmt.Printf("Incident: %s
", record.Get("number"))
    }
}
```

## Creating a Record

To create a new record in a table:

```go
import (
    "context"
    "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    // ... client initialization ...

    ctx := context.Background()
    
    newIncident := tableapi.NewTableRecord()
    newIncident.Set("short_description", "System is down")
    newIncident.Set("priority", "1")

    response, err := client.Now2().TableV2("incident").Post(ctx, newIncident, nil)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Printf("Created incident with sys_id: %s
", response.GetResult().GetSysId())
}
```

## Updating a Record

To update an existing record, you first need its `sys_id`:

```go
sysId := "..." // your record sys_id

updateData := tableapi.NewTableRecord()
updateData.Set("short_description", "Updated description")

response, err := client.Now2().TableV2("incident").ById(sysId).Put(ctx, updateData, nil)
```

## Deleting a Record

```go
sysId := "..."
err := client.Now2().TableV2("incident").ById(sysId).Delete(ctx, nil)
```

## Querying and Filtering

You can use `TableRequestBuilder2GetRequestConfiguration` to add query parameters like `sysparm_query`.

```go
params := &tableapi.TableRequestBuilder2GetQueryParameters{
    Query: "priority=1^active=true",
    Limit: 10,
}

config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
    QueryParameters: params,
}

response, err := client.Now2().TableV2("incident").Get(ctx, config)
```

!!! tip "Fluent Query Builder (Preview)"
    For complex queries, you can use the [Fluent Query Builder](query-builder.md) to construct query strings in a type-safe manner. Note that this requires the `preview.query` build flag.
