# Fluent Query Builder (Preview)

!!! danger "Preview Feature"
    This feature is currently in **preview** and requires the `preview.query` build flag. 
    To use it, build your application with: `-tags preview.query`

The Fluent Query Builder provides a type-safe and readable way to construct complex ServiceNow query strings (`sysparm_query`).

## Setup

Ensure you have the `preview.query` build tag enabled in your environment or build command.

## Basic Usage

Instead of writing raw query strings like `"priority=1^active=true"`, you can use the fluent API:

```go
import (
    "fmt"
    "github.com/michaeldcanady/servicenow-sdk-go/query"
)

func main() {
    q := query.NewQuery().
        StringField("short_description").Contains("System").
        And().
        NumericField("priority").Equals(1).
        String()

    fmt.Println(q) // Output: short_descriptionLIKESystem^priority=1
}
```

## Integration with Table API

You can use the built query string with the Table API's query parameters.

```go
import (
    "context"
    "github.com/michaeldcanady/servicenow-sdk-go/query"
    "github.com/michaeldcanady/servicenow-sdk-go/table-api"
)

func main() {
    // ... client initialization ...

    ctx := context.Background()

    // Build the query
    q := query.NewQuery().
        StringField("active").Equals("true").
        And().
        StringField("priority").In("1", "2").
        String()

    params := &tableapi.TableRequestBuilder2GetQueryParameters{
        Query: q,
    }

    config := &tableapi.TableRequestBuilder2GetRequestConfiguration{
        QueryParameters: params,
    }

    response, err := client.Now2().TableV2("incident").Get(ctx, config)
    // ... handle response ...
}
```

## Supported Operators

The builder supports most common ServiceNow operators:

### String Fields
- `Equals(val)`
- `NotEquals(val)`
- `Contains(val)`
- `DoesNotContain(val)`
- `StartsWith(val)`
- `EndsWith(val)`
- `IsEmptyString()`

### Numeric Fields
- `Equals(val)`
- `NotEquals(val)`
- `GreaterThan(val)`
- `LessThan(val)`
- `GreaterOrEqualTo(val)`
- `LessOrEqualTo(val)`

### Logical Operators
- `And()`
- `Or()`
