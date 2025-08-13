# Query building

The Service-Now API query syntax can be difficult to approach if you're not familiar or are not getting queries using the UI.

Below are the two proposed versions of the Query building API.

=== "V1"

    This design is leveraging the functional patterns for building the query.
    For this design keep in mind when you start there is an implicit *and* between the initial condition meaning the below query translates to: `example>7^example2>7^ORexample3=random`m.

    if you want to start with an *or* you'll need to use `OrGroup`.

    ```golang
    import (
        "github.com/michaeldcanady/query"
    )

    func main() {
        builder := query.NewQueryBuilder()

        builder.
            AddFilter("example", query.GreaterThanCondition(7)).
            // you can group conditions using OrGroup/AndGroup
            OrGroup(func(q *query.QueryBuilder) {
                q.AddFilter("example2", query.GreaterThanCondition(7))
                q.AddFilter("example3", query.IsCondition("random"))
            })
    }
    ```

=== "V2"

    This design is leveraging the well-known `Builder Pattern` for building the query.

    ```golang
    import (
        "github.com/michaeldcanady/query1"
    )

    func main() {
        builder := query.NewQuery().NumericField("example").GreaterThan(7).
            Or().NumericField("example2").GreaterThan(7).
            Or().StringField("example3").Is("random")
    }
    ```
