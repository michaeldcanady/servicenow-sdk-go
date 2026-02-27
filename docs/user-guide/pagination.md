# Pagination

When working with large datasets in ServiceNow, responses are often split into
multiple pages. The SDK provides specialized iterators to help you traverse
these results efficiently.

## Understand pagination

ServiceNow uses limit and offset parameters or link headers to manage
pagination. The SDK abstracts this complexity by providing an iterator that
can automatically fetch subsequent pages as you process records.

The SDK provides specialized iterators for common APIs:
- `TablePageIterator`: For table record collections.
- `AttachmentPageIterator`: For attachment collections.

## Use the page iterator

### Basic iteration

To iterate over multiple pages of records, use the `NewDefaultTablePageIterator`
(for standard table records) or `NewTablePageIterator` (for custom types).

```go
{% include-markdown '../snippets/pagination.go' start='// [START pagination_table_basic]' end='// [START pagination_table_basic]' comments=false trailing-newlines=false dedent=true %}
```

### Iterate item-by-item

If you need more control than a callback, you can use `NextItem` and `HasNext`
to pull items one at a time.

```go
{% include-markdown '../snippets/pagination.go' start='// [START pagination_item_by_item]' end='// [START pagination_item_by_item]' comments=false trailing-newlines=false dedent=true %}
```

### Attachment iteration

Iterating over attachments is similar, using the `AttachmentPageIterator`.

```go
{% include-markdown '../snippets/pagination.go' start='// [START pagination_attachment]' end='// [START pagination_attachment]' comments=false trailing-newlines=false dedent=true %}
```

### Reverse iteration

You can iterate backwards through pages by passing `true` to the `reverse`
parameter of the `Iterate` method. This uses the `previous` link headers
provided by the ServiceNow API.

```go
err = iterator.Iterate(ctx, true, func(record *tableapi.TableRecord) bool {
    // Process records in reverse order
    return true
})
```

## Advanced features

### Manual page navigation

You can navigate between pages manually using methods like `Next()`,
`Previous()`, `First()`, and `Last()`.

```go
{% include-markdown '../snippets/pagination.go' start='// [START pagination_table_manual]' end='// [START pagination_table_manual]' comments=false trailing-newlines=false dedent=true %}
```

### State management

The iterator maintains state, allowing you to reset or restart iteration.

```go
{% include-markdown '../snippets/pagination.go' start='// [START pagination_state_management]' end='// [START pagination_state_management]' comments=false trailing-newlines=false dedent=true %}
```

- `Reset()`: Returns the iterator to the first page and first item.
- `ResetPage()`: Restarts iteration of the current page.

## Next steps

- **[Table Operations](tables.md):** Learn more about the initial requests
  that produce paginated results.
- **[Querying](query-builder.md):** Learn how to use queries to limit the number
  of records returned.
