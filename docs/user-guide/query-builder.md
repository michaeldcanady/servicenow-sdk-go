# Querying

The SDK provides a powerful query builder to construct complex ServiceNow
Encoded Queries in a type-safe and readable way. This helps filter table
records precisely and minimize unnecessary data transfer.

> **Note:** This is a preview feature currently under active development.

## Understand encoded queries

ServiceNow uses a specific syntax for filtering records, known as Encoded
Queries. These queries are strings consisting of field names, operators, and
values.

### Key concepts

- **Operators:** Used to compare field values, for example, `=`, `!=`, `LIKE`,
  `STARTSWITH`.
- **Logical Operators:** Used to combine multiple conditions (`^` for and
  [conjunction], `^OR` for or [disjunction]).
- **Encoded String:** The final output of the query builder that's passed to
  the Table API.

## Build a query

To build a query, use the `query2` package. It provides methods for different
field types and logical operations.

```go
{% include-markdown 'snippets/query.go' start='// [START query_basic]' end='// [END query_basic]' comments=false trailing-newlines=false dedent=true %}
```

## Use a query with the Table API

Once you've built a query string, you can use it in the request
configuration for table operations.

```go
{% include-markdown 'snippets/query.go' start='// [START query_table_api]' end='// [END query_table_api]' comments=false trailing-newlines=false dedent=true %}
```

## Next steps

- **[Table Operations](tables.md):** Learn more about the operations where
  you'll use these queries.
- **[Pagination](pagination.md):** Learn how to handle results when your
  query returns many records.
