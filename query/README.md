# Query

Contains how to do querying

```golang
import (
    "time"

    "github.com/michaeldcanady/service-now-sdk-go/query"
)

func Main() {
    // Converts your built query to a string
    queryString := query.Query(
        query.And(
            query.Is("field", "value"),
            query.IsNot("field", 10),
            ...
        ),
        query.Or(
            query.And(
                query.IsDifferent("field", "value"),
                query.On("field", time.Now())
            ),
            ...
        )
    )
}
```
